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
	"time"

	client "simplebank/client"
	esystem "simplebank/client/esystem"
	"simplebank/util"

	// pb "simplebank/pb"
	local "simplebank/db/datastore/esystemlocal"

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
	// serverAddress = flag.String("address", "localhost:8080", "The server address in the format of host:port")

	// serverAddress = flag.String("address", "ces.fdsasya.com:8080", "The server address in the format of host:port")
	// serverAddress = flag.String("address", "34.143.228.170:8080", "The server address in the format of host:port")
)

// serverAddress = flag.String("address", "localhost:52898", "The server address in the format of host:port")
// serverAddress = flag.String("address", "34.143.228.170:8080", "The server address in the format of host:port")
// serverAddress      = flag.String("address", "esystemci.fortress-asya.com:8080", "The server address in the format of host:port")

const (
	username        = "olive.mercado0609@gmail.com"
	password        = "1234"
	refreshDuration = 10 * time.Minute
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

var (
	QueriesLocal   *local.QueriesLocal
	serverPath     string
	dockerPath     string
	dockerImgID    string
	serverAddress  string
	dBeSystemLocal string
	DB             *sql.DB
)

func initSetup() {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	serverPath = config.HomeFolder
	dockerPath = "/var/lib/postgresql/"
	dockerImgID = config.DockerSQLImgID
	serverAddress = config.ServerAddress
	dBeSystemLocal = config.DBeSystemLocal
}

func OpenLocalDB() {

	DB, err := sql.Open("mssql", dBeSystemLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	QueriesLocal = local.New(DB)

}

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

type forUpload struct {
	FilePath    string
	TargetTable string
	CSVFunc     func(context.Context, string) error
}

func main() {
	// serverAddress := flag.String("address", "", "the server address")
	// serverAddress = "34.143.228.170:8080"
	initSetup()
	log.Printf("serverAddress: %s", serverAddress)
	// flag.Parse()
	log.Printf("dial server %s, TLS = %t", serverAddress, *enableTLS)

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

	cc1, err := grpc.Dial(serverAddress, dialOption)
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
		serverAddress,
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
	OpenLocalDB()

	dumpClient := esystem.NewDumpClient(cc2)

	ctxQuery, cancelQuery, resetQuery := WithTimeoutReset(context.Background(), time.Minute*200)
	defer cancelQuery()

	// ctxUpload, cancelUpload, resetUpload := WithTimeoutReset(context.Background(), time.Minute*5)
	// defer cancelUpload()

	// log.Printf("Client Main : %v", 2)

	OrgParm, _ := QueriesLocal.GetOrgParms(ctxQuery)
	log.Printf("OrgParm: %v", OrgParm)
	resetQuery()
	log.Printf("Client Main : %v", OrgParm)
	dumpClient.CreateDumpBranchList(OrgParm)
	log.Printf("Client Main : %v", 4)

	brList, _ := dumpClient.GetDumpBranchList(OrgParm.BrCode)
	log.Printf("Client Main BrList: %v", brList)

	// dumpClient.GetModifiedTable(ctx, &pb.GetModifiedTableRequest{})
	// dumpClient.GetModifiedTable(ctx, &pb.GetModifiedTableRequest{})

	// localMod, _ := QueriesLocal.ListModifiedTable(ctx)
	// log.Printf("localMod: %v", localMod)
	// for _, modi := range localMod {
	// 	dumpClient.ModTableLocalList[modi.LocalTableName] = modi.ModCtr
	// }

	// localUser, _ := QueriesLocal.ListUsersList(ctx)
	// log.Printf("localMod: %v", localUser)
	// dumpClient.CreateDumpUsersList(localUser)

	// for _, user := range localUser {
	// }

	// localColSht, _ := QueriesLocal.GetColSht(ctx)
	// reset()
	// log.Printf("localColSht: %v", localColSht)
	// dumpClient.CreateDumpColSht(localColSht)

	// for _, colSht := range localColSht {
	// }

	// area, err := QueriesLocal.ListArea(ctx)
	// log.Printf("area: %v", err)
	// dumpClient.CreateDumpArea(area)

	// log.Printf("area: %v L:%v C:%v", area, dumpClient.ModTableCentralList["Area"], dumpClient.ModTableLocalList["Area"])

	// center, err := QueriesLocal.ListCenter(ctx)
	// util.LogErrorApp("CreateCSV:center", err)
	// dumpClient.CreateDumpCenter(center)

	// centerWorker, err := QueriesLocal.ListCenterWorker(ctx)
	// util.LogErrorApp("CreateCSV:CenterWorker", err)
	// dumpClient.CreateDumpCenterWorker(centerWorker)

	// custAddInfo, err := QueriesLocal.ListCustAddInfo(ctx)
	// util.LogErrorApp("CreateCSV:CustAddInfo", err)
	// dumpClient.CreateDumpCustAddInfo(custAddInfo)

	uploadFileRequest := esystem.UploadFileRequest{
		RefCode:     OrgParm.BrCode,
		Remarks:     "",
		ServerPath:  serverPath,
		DockerPath:  dockerPath,
		DockerImgID: dockerImgID}
	// /*

	forUploads := []forUpload{
		// /*
		{FilePath: "csv/Area.csv",
			TargetTable: "staging.Area",
			CSVFunc:     QueriesLocal.AreaCSV},
		{FilePath: "csv/Unit.csv",
			TargetTable: "staging.Unit",
			CSVFunc:     QueriesLocal.UnitCSV},
		{FilePath: "csv/Center.csv",
			TargetTable: "staging.Center",
			CSVFunc:     QueriesLocal.CenterCSV},
		{FilePath: "csv/Customer.csv",
			TargetTable: "staging.Customer",
			CSVFunc:     QueriesLocal.CustomerCSV},
		{FilePath: "csv/Addresses.csv",
			TargetTable: "staging.Addresses",
			CSVFunc:     QueriesLocal.AddressesCSV},
		{FilePath: "csv/LnMaster.csv",
			TargetTable: "staging.LnMaster",
			CSVFunc:     QueriesLocal.LnMasterCSV},
		{FilePath: "csv/SaMaster.csv",
			TargetTable: "staging.SaMaster",
			CSVFunc:     QueriesLocal.SaMasterCSV},
		{FilePath: "csv/TrnMaster.csv",
			TargetTable: "staging.TrnMaster",
			CSVFunc:     QueriesLocal.TrnMasterCSV},
		{FilePath: "csv/SaTrnMaster.csv",
			TargetTable: "staging.SaTrnMaster",
			CSVFunc:     QueriesLocal.SaTrnMasterCSV},
		{FilePath: "csv/LoanInst.csv",
			TargetTable: "staging.LoanInst",
			CSVFunc:     QueriesLocal.LoanInstCSV},
		{FilePath: "csv/LnChrgData.csv",
			TargetTable: "staging.LnChrgData",
			CSVFunc:     QueriesLocal.LnChrgDataCSV},
		{FilePath: "csv/CustAddInfoList.csv",
			TargetTable: "staging.CustAddInfoList",
			CSVFunc:     QueriesLocal.CustAddInfoListCSV},
		{FilePath: "csv/CustAddInfoGroup.csv",
			TargetTable: "staging.CustAddInfoGroup",
			CSVFunc:     QueriesLocal.CustAddInfoGroupCSV},
		{FilePath: "csv/CustAddInfoGroupNeed.csv",
			TargetTable: "staging.CustAddInfoGroupNeed",
			CSVFunc:     QueriesLocal.CustAddInfoGroupNeedCSV},
		{FilePath: "csv/CustAddInfo.csv",
			TargetTable: "staging.CustAddInfo",
			CSVFunc:     QueriesLocal.CustAddInfoCSV},
		{FilePath: "csv/MutualFund.csv",
			TargetTable: "staging.MutualFund",
			CSVFunc:     QueriesLocal.MutualFundCSV},
		{FilePath: "csv/ReferencesDetails.csv",
			TargetTable: "staging.ReferencesDetails",
			CSVFunc:     QueriesLocal.ReferencesDetailsCSV},
		{FilePath: "csv/CenterWorker.csv",
			TargetTable: "staging.CenterWorker",
			CSVFunc:     QueriesLocal.CenterWorkerCSV},
		{FilePath: "csv/Writeoff.csv",
			TargetTable: "staging.Writeoff",
			CSVFunc:     QueriesLocal.WriteoffCSV},
		{FilePath: "csv/Accounts.csv",
			TargetTable: "staging.Accounts",
			CSVFunc:     QueriesLocal.AccountsCSV},
		{FilePath: "csv/JnlDetails.csv",
			TargetTable: "staging.JnlDetails",
			CSVFunc:     QueriesLocal.JnlDetailsCSV},
		{FilePath: "csv/JnlHeaders.csv",
			TargetTable: "staging.JnlHeaders",
			CSVFunc:     QueriesLocal.JnlHeadersCSV},
		{FilePath: "csv/LedgerDetails.csv",
			TargetTable: "staging.LedgerDetails",
			CSVFunc:     QueriesLocal.LedgerDetailsCSV},
		{FilePath: "csv/UsersList.csv",
			TargetTable: "staging.UsersList",
			CSVFunc:     QueriesLocal.UsersListCSV},
		{FilePath: "csv/MultiplePaymentReceipt.csv",
			TargetTable: "staging.MultiplePaymentReceipt",
			CSVFunc:     QueriesLocal.MultiplePaymentReceiptCSV},
		{FilePath: "csv/ReactivateWriteoff.csv",
			TargetTable: "staging.ReactivateWriteoff",
			CSVFunc:     QueriesLocal.ReactivateWriteoffCSV},
		{FilePath: "csv/LnBeneficiary.csv",
			TargetTable: "staging.LnBeneficiary",
			CSVFunc:     QueriesLocal.LnBeneficiaryCSV},
		// */
		{FilePath: "csv/ColSht.csv",
			TargetTable: "staging.ColSht",
			CSVFunc:     QueriesLocal.ColShtCSV},
	}

	for _, fUs := range forUploads {
		log.Printf("START--> %s", fUs.TargetTable)
		uploadFileRequest.FilePath = fUs.FilePath
		uploadFileRequest.TargetTable = fUs.TargetTable
		createCSV(fUs)
		uploadCSV(dumpClient, uploadFileRequest)
		// ctxQuery, cancelQuery, _ = WithTimeoutReset(context.Background(), time.Minute*200)
		// defer cancelQuery()
		// err = fUs.CSVFunc(ctxQuery, uploadFileRequest.FilePath)
		// util.LogErrorApp("CreateCSV: "+uploadFileRequest.TargetTable, err)

		// ctxUpload, cancelUpload, _ = WithTimeoutReset(context.Background(), time.Minute*500)
		// defer cancelUpload()
		// dumpClient.UploadFile(ctxUpload, uploadFileRequest)
	}

	// // cmd := exec.Command("/bin/sh", "-c", "docker cp /Users/rhickmercado/Documents/Programming/go/src/simplebank/db/datastore/esystemlocal/csv/CustAddInfo.csv 263d578de57f:/var/lib/postgresql/CustAddInfo.csv")
	// // err = cmd.Run()
	// reset()

	// // if err != nil {
	// reset()
	// // 	log.Println(err)
	// }
	// require.NoError(t, err)

	// // for _, modOrd := range dumpClient.ModTableOrderList {
	// // 	// lastModCentral := dumpClient.ModTableCentralList[modOrd.CentralTableName].ModCtr
	// // 	// if modOrd.CentralTableName == "Area" {
	// // 	// 	dumpClient.CreateArea(ctx, &pb.NoParam{})
	// // 	// }
	// // 	log.Printf("test %v", modOrd)
	// // 	// dumpClient.ModTableLocalList := local.New()
	// // }

}

// for mod : range

func createCSV(f forUpload) {
	ctx, cancel, _ := WithTimeoutReset(context.Background(), time.Minute*200)
	defer cancel()
	err := f.CSVFunc(ctx, f.FilePath)
	util.LogErrorApp("CreateCSV: "+f.TargetTable, err)
}

func uploadCSV(client *esystem.DumpClient, f esystem.UploadFileRequest) {
	ctx, cancel, _ := WithTimeoutReset(context.Background(), time.Minute*500)
	defer cancel()
	client.UploadFile(ctx, f)
}
