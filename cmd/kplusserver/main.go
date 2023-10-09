package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	dsUser "simplebank/db/datastore/user"

	pb "simplebank/pb"
	authService "simplebank/service/auth"
	docService "simplebank/service/document"
	eSysService "simplebank/service/esystem"
	kplusService "simplebank/service/kplus"

	"simplebank/util"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

const (
	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func main() {

	// Get Terminal Parameter Switches
	port := flag.Int("port", 8090, "the server port")
	// port := 52898
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	// serverType := flag.String("type", "grpc", "type of server (grpc/rest)")
	// endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

	// Connect to Database Server based on Config
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	db, err = sql.Open(config.DBDriver, postgresqlDbInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	userStore := dsUser.New(db)
	jwtManager := authService.NewJWTManager(secretKey, tokenDuration)
	kPlusServiceServer := kplusService.NewKPlusServiceServer(db, *userStore)
	eSysServiceServer := eSysService.NewDumpServer(db)
	docServiceServer := docService.NewDocumentServer(db, docService.New("."))
	authServiceServer := authService.NewAuthServiceServer(*userStore, jwtManager)

	// Create a gRPC server object
	interceptor := authService.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Println("cannot load TLS credentials: %w", err)
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	// grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(serverOptions...)

	// Attach the KPlusService service to the server
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer)
	pb.RegisterKPlusServiceServer(grpcServer, kPlusServiceServer)
	pb.RegisterDocumentServiceServer(grpcServer, docServiceServer)
	pb.RegisterDumpServiceServer(grpcServer, eSysServiceServer)

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// Register KPlusService
	err = pb.RegisterKPlusServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register KPlus:", err)
	}

	// Register DocumentService
	err = pb.RegisterDocumentServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register Document:", err)
	}

	// Register DumpService
	err = pb.RegisterDumpServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register Dump:", err)
	}

	// Register AuthService
	err = pb.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", *port),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/simplebank.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin", "Bookkeeper"},
		laptopServicePath + "UploadImage":  {"admin", "Bookkeeper"},
		laptopServicePath + "RateLaptop":   {"admin", "user", "Bookkeeper"},
	}
}
