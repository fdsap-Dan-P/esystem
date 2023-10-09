package main

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/grpc"

	_ "github.com/denisenkom/go-mssqldb"

	pb "your-package-name/protobuf" // Replace with your protobuf package

	_ "github.com/lib/pq"
)

// Configurations
const (
	postgresConnectionString = "your-postgres-connection-string"
	grpcServerAddress        = "localhost:50051" // Replace with your gRPC server address
)

func main() {
	// Connect to Postgres database
	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer db.Close()

	// Initialize gRPC connection
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewFileTransferClient(conn)

	// Create a channel to receive table names
	tables := make(chan string, 10) // Adjust buffer size as per your requirements

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Fetch table names from Postgres
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		log.Fatalf("Failed to fetch table names: %v", err)
	}
	defer rows.Close()

	// Start goroutines for processing tables
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Printf("Failed to scan table name: %v", err)
			continue
		}

		// Add table name to the channel
		tables <- tableName

		// Increment wait group counter
		wg.Add(1)

		// Start goroutine to process table
		go func(table string) {
			defer wg.Done()

			// Process table and generate CSV
			if err := processTable(db, table); err != nil {
				log.Printf("Failed to process table %s: %v", table, err)
				return
			}

			// Compress CSV to ZIP
			zipFile := fmt.Sprintf("%s.zip", table)
			if err := compressToZIP(table, zipFile); err != nil {
				log.Printf("Failed to compress table %s: %v", table, err)
				return
			}

			// Send ZIP file via gRPC
			if err := sendFile(client, zipFile); err != nil {
				log.Printf("Failed to send file %s: %v", zipFile, err)
				return
			}

			log.Printf("Processed table %s", table)
		}(tableName)
	}

	// Close the tables channel
	close(tables)

	// Wait for all goroutines to finish
	wg.Wait()

	log.Println("All tables processed")
}

func processTable(db *sql.DB, table string) error {
	// Create a CSV file for the table
	csvFile, err := os.Create(fmt.Sprintf("%s.csv", table))
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer csvFile.Close()

	// Query table data from Postgres
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		return fmt.Errorf("failed to query table data: %v", err)
	}
	defer rows.Close()

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Fetch column names from the query result
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to fetch column names: %v", err)
	}

	// Write column names to CSV
	if err := csvWriter.Write(columns); err != nil {
		return fmt.Errorf("failed to write column names: %v", err)
	}

	// Prepare a slice for storing row values
	rowValues := make([]interface{}, len(columns))
	rowPointers := make([]interface{}, len(columns))
	for i := range rowValues {
		rowPointers[i] = &rowValues[i]
	}

	// Fetch and write row data to CSV
	for rows.Next() {
		if err := rows.Scan(rowPointers...); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		// Convert row values to strings
		rowData := make([]string, len(columns))
		for i, rv := range rowValues {
			if rv == nil {
				rowData[i] = ""
			} else {
				rowData[i] = fmt.Sprintf("%v", rv)
			}
		}

		// Write row data to CSV
		if err := csvWriter.Write(rowData); err != nil {
			return fmt.Errorf("failed to write row data: %v", err)
		}
	}

	return nil
}

func compressToZIP(sourceFile, targetFile string) error {
	// Create a new ZIP file
	zipFile, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create ZIP file: %v", err)
	}
	defer zipFile.Close()

	// Create a ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Open the source file
	source, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer source.Close()

	// Get file information
	sourceInfo, err := source.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file information: %v", err)
	}

	// Create a new ZIP file entry
	writer, err := zipWriter.Create(sourceInfo.Name())
	if err != nil {
		return fmt.Errorf("failed to create ZIP file entry: %v", err)
	}

	// Copy the source file contents to the ZIP entry
	if _, err := io.Copy(writer, source); err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}

	return nil
}

func sendFile(client pb.FileTransferClient, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a buffer to store file contents
	fileContents := bytes.Buffer{}
	if _, err := io.Copy(&fileContents, file); err != nil {
		return fmt.Errorf("failed to read file contents: %v", err)
	}

	// Prepare the file transfer request
	request := &pb.FileTransferRequest{
		Filename: filepath.Base(filePath),
		Content:  fileContents.Bytes(),
	}

	// Send the file via gRPC
	_, err = client.SendFile(context.Background(), request)
	if err != nil {
		return fmt.Errorf("failed to send file via gRPC: %v", err)
	}

	return nil
}
