package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	client "simplebank/client"
	laptop "simplebank/client/laptop"
	"simplebank/util"

	// pb "simplebank/pb"
	local "simplebank/db/datastore/esystemlocal"
	sample "simplebank/sample"

	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	enableTLS          = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddress      = flag.String("address", "localhost:52898", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func testCreateLaptop(laptopClient *laptop.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

// func testSearchLaptop(laptopClient *client.LaptopClient) {
// 	for i := 0; i < 10; i++ {
// 		laptopClient.CreateLaptop(sample.NewLaptop())
// 	}

// 	filter := &pb.Filter{
// 		MaxPriceUsd: 3000,
// 		MinCpuCores: 4,
// 		MinCpuGhz:   2.5,
// 		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
// 	}

// 	laptopClient.SearchLaptop(filter)
// }

func testUploadImage(laptopClient *laptop.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func testRateLaptop(laptopClient *laptop.LaptopClient) {
	n := 3
	laptopIDs := make([]int64, n)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

const (
	username        = "olive.mercado0609@gmail.com"
	password        = "1234"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/simplebank.LaptopService/"

	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

// func get

var QueriesLocal *local.QueriesLocal
var DB *sql.DB

func OpenLocalDB() {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DB, err = sql.Open("mssql", config.DBeSystemLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	QueriesLocal = local.New(DB)

}

func main() {
	// serverAddress := flag.String("address", "", "the server address")
	log.Printf("serverAddress: %v", serverAddress)
	// flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithInsecure()

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	log.Printf("Client Main : %v", "grpc.Dial(*serverAddress, transportOption")
	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	log.Printf("Client Main : %v", "client.NewAuthClient(cc1, username, password)")
	authClient := client.NewAuthClient(cc1, username, password)
	log.Printf("Client Main : authClient:%v-> %v", authClient, "client.NewAuthInterceptor(")
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	log.Printf("Client Main : %v", "grpc.Dial(")
	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	// modified := client.
	log.Printf("Client Main : %v", "client.NewLaptopClient(cc2)")
	laptopClient := laptop.NewLaptopClient(cc2)
	testRateLaptop(laptopClient)

}

// for mod : range
