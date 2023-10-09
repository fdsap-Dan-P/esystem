package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	client "simplebank/client"
	kplus "simplebank/client/kplus"

	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	enableTLS = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	// serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
	// serverAddress = flag.String("address", "localhost:52898", "The server address in the format of host:port")

	serverAddress = flag.String("address", "localhost:52898", "The server address in the format of host:port")
)

// serverAddress = flag.String("address", "localhost:52898", "The server address in the format of host:port")
// serverAddress = flag.String("address", "34.143.228.170:8080", "The server address in the format of host:port")
// serverAddress      = flag.String("address", "esystemci.fortress-asya.com:8080", "The server address in the format of host:port")

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

// ResetFunc resets the context timeout timer
type ResetFunc func()

func WithTimeoutReset(parent context.Context, d time.Duration) (context.Context, context.CancelFunc, ResetFunc) {
	ctx, cancel0 := context.WithCancel(parent)
	timer := time.AfterFunc(d, cancel0)
	cancel := func() {
		cancel0()
		timer.Stop()
	}
	reset := func() { timer.Reset(d) }
	return ctx, cancel, reset
}

func main() {
	// serverAddress := flag.String("address", "", "the server address")
	// serverAddress = "34.143.228.170:8080"
	log.Printf("serverAddress: %s", *serverAddress)
	// flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	var dialOption grpc.DialOption
	// var transportOption credentials.TransportCredentials
	// transportOption := grpc.WithTransportCredentials
	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		dialOption = grpc.WithTransportCredentials(tlsCredentials)
		if err != nil {
			log.Fatal("cannot dial server: ", err)
		}
		// credentials.NewClientTLSFromCert(tlsCredentials, "")
	} else {
		dialOption = grpc.WithTransportCredentials(insecure.NewCredentials())
		// grpc.WithInsecure()
		// transportOption = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	}

	// log.Printf("Client Main : %v", "grpc.Dial(*serverAddress, transportOption")

	cc1, err := grpc.Dial(*serverAddress, dialOption)
	log.Println("Disbled TLS")
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	// log.Printf("Client Main : %v", "client.NewAuthClient(cc1, username, password)")
	authClient := client.NewAuthClient(cc1, username, password)
	log.Printf("Client Main : authClient:%v-> %v", authClient, "client.NewAuthInterceptor(")
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	// log.Printf("Client Main : %v", "grpc.Dial(")
	cc2, err := grpc.Dial(
		*serverAddress,
		dialOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	// log.Printf("Client Main : %v", 1)

	// // modified := client.
	// log.Printf("Client Main : %v", "client.NewLaptopClient(cc2)")
	// laptopClient := client.NewLaptopClient(cc2)
	// testRateLaptop(laptopClient)

	client := kplus.NewKPlusClient(cc2)

	ctx, cancel, reset := WithTimeoutReset(context.Background(), time.Minute*5)
	defer cancel()

	cus, cuserr := client.SearchCustomerCID(ctx, 6710)
	if cuserr != nil {
		log.Printf("SearchCustomerCID ERROR: %v", cuserr)
	}
	log.Printf("SearchCustomerCID Main : %v", cus)
	reset()

	saveList, saveListerr := client.SearchCustomerCID(ctx, 143)
	if saveListerr != nil {
		log.Printf("saveSearchCustomerCIDList ERROR: %v", saveListerr)
	}
	log.Printf("SearchCustomerCID Data : %v", saveList)
	reset()

}

// for mod : range
