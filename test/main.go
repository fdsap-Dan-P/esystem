package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	server := "<server_address>"
	port := 1433
	database := "<database_name>"
	user := "<username>"
	password := "<password>"

	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=true", server, port, database, user, password)

	// Set the TLS configuration to use TLS 1.2
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
	}

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Set the custom TLS configuration
	err = db.QueryRowContext(context.WithValue(context.Background(), "tlsConfig", tlsConfig), "SELECT 1").Scan(new(int))
	if err != nil {
		fmt.Println("Error pinging the database:", err.Error())
		return
	}

	fmt.Println("Successfully connected to the database!")
}
