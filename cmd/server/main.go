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
	"strings"
	"time"

	dsUser "simplebank/db/datastore/user"

	// dsUsr "simplebank/db/datastore/user"
	// service "simplebank/service"
	// documentService "simplebank/service/document"
	// eSystemService "simplebank/service/esystem"
	// kpluservice "simplebank/service/kplus"

	pb "simplebank/pb"
	service "simplebank/service/auth"
	documentService "simplebank/service/document"
	eSystemService "simplebank/service/esystem"
	kpluservice "simplebank/service/kplus"

	// u "simplebank/user"
	"simplebank/util"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var db *sql.DB

// func getUsers(userStore dsuser.QuerierUser) error {
// 	err :=  createUser(userStore, "admin1", "secret", "admin")
// 	if err != nil {
// 		return err
// 	}
// 	return createUser(userStore, "user1", "secret", "user")
// }

// func createUser(userStore ds.UserStore, username, password, role string) error {
// 	user, err := u.NewUser(username, password, role)
// 	if err != nil {
// 		return err
// 	}
// 	return userStore.Save(user)
// }

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

const (
	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/simplebank.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin", "Bookkeeper"},
		laptopServicePath + "UploadImage":  {"admin", "Bookkeeper"},
		laptopServicePath + "RateLaptop":   {"admin", "user", "Bookkeeper"},
	}
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

func runGRPCServer(
	authServer pb.AuthServiceServer,
	dumpServer pb.DumpServiceServer,
	documentServer pb.DocumentServiceServer,
	kplusServer pb.KPlusServiceServer,
	jwtManager *service.JWTManager,
	enableTLS bool,
	listener net.Listener,
) error {
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	if enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}

	grpcServer := grpc.NewServer(serverOptions...)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterDumpServiceServer(grpcServer, dumpServer)
	pb.RegisterDocumentServiceServer(grpcServer, documentServer)
	pb.RegisterKPlusServiceServer(grpcServer, kplusServer)
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}

func runRESTServer(
	authServer pb.AuthServiceServer,
	dumpServer pb.DumpServiceServer,
	documentServer pb.DocumentServiceServer,
	kplusServer pb.KPlusServiceServer,
	jwtManager *service.JWTManager,
	enableTLS bool,
	listener net.Listener,
	grpcEndpoint string,
) error {
	// mux := runtime.NewServeMux()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(func(header string) (string, bool) {
		if header == "x-grpc-gateway-path" {
			return header, true
		}
		return "", false
	}))

	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// in-process handler
	err := pb.RegisterAuthServiceHandlerServer(ctx, mux, authServer)
	if err != nil {
		return err
	}

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}

	// in-process handler DumpServer
	err = pb.RegisterDumpServiceHandlerServer(ctx, mux, dumpServer)
	if err != nil {
		return err
	}
	err = pb.RegisterDumpServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}

	// in-process handler DocumentServer
	err = pb.RegisterDocumentServiceHandlerServer(ctx, mux, documentServer)
	if err != nil {
		return err
	}
	err = pb.RegisterDocumentServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}

	// in-process handler KPlusServer
	err = pb.RegisterKPlusServiceHandlerServer(ctx, mux, kplusServer)
	if err != nil {
		return err
	}
	err = pb.RegisterKPlusServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	if err != nil {
		return err
	}

	// sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swagger/")))
	// r.PathPrefix("/swaggerui/").Handler(sh)
	log.Println("Available REST paths:")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, item := range strings.Split(r.Header.Get("X-Endpoint-API-UserInfo"), ",") {
			if strings.HasPrefix(item, "x-grpc-gateway-path=") {
				fmt.Println(strings.TrimPrefix(item, "x-grpc-gateway-path="))
			}
		}
		mux.ServeHTTP(w, r)
	})

	log.Printf("Start REST server at %s, TLS = %t", listener.Addr().String(), enableTLS)

	if enableTLS {
		return http.ServeTLS(listener, mux, serverCertFile, serverKeyFile)
	}
	return http.Serve(listener, mux)
}

func main() {
	port := flag.Int("port", 52898, "the server port")
	// port := 52898
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	serverType := flag.String("type", "grpc", "type of server (grpc/rest)")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

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

	userStore := dsUser.New(db)
	// userStore := dsUsr.New(db)

	// err := seedUsers(userStore)
	// if err != nil {
	// 	log.Fatal("cannot seed users: ", err)
	// }

	log.Println("Connected to DB")

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServiceServer(*userStore, jwtManager)

	dumpServer := eSystemService.NewDumpServer(db)
	kplusServer := kpluservice.NewKPlusServiceServer(db, *userStore)
	documentServer := documentService.NewDocumentServer(db, documentService.New("."))
	// pb.DumpServiceServer

	address := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	// go func() {
	// 	log.Println("Starting gRPC server on port 50051")
	// 	if err := runGRPCServer(authServer, dumpServer, documentServer, kplusServer, jwtManager, *enableTLS, listener); err != nil {
	// 		log.Fatalf("failed to start gRPC server: %v", err)
	// 	}
	// }()

	// listener2, err2 := net.Listen("tcp", address)
	// err2 = runRESTServer(authServer, dumpServer, documentServer, kplusServer, jwtManager, *enableTLS, listener2, *endPoint)
	// if err2 != nil {
	// 	log.Fatal("cannot start server: ", err)
	// }

	if *serverType == "grpc" {
		err = runGRPCServer(authServer, dumpServer, documentServer, kplusServer, jwtManager, *enableTLS, listener)
	} else {
		err = runRESTServer(authServer, dumpServer, documentServer, kplusServer, jwtManager, *enableTLS, listener, *endPoint)
	}

}
