package client

import (
	"context"
	"database/sql"
	"log"
	pb "simplebank/pb"
	"simplebank/util"
	"time"

	local "simplebank/db/datastore/esystemlocal"

	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc"
)

var QueriesLocal *local.QueriesLocal
var DBLocal *sql.DB

// DumpAreaClient is a client to call DumpArea service RPCs
type DumpClient struct {
	service             pb.DumpServiceClient
	document            pb.DocumentServiceClient
	ModTableOrderList   map[int16]modTableOrder
	ModTableLocalList   map[string]int64
	ModTableCentralList map[string]int64
}

func init() {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DBLocal, err = sql.Open("mssql", config.DBeSystemLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	QueriesLocal = local.New(DBLocal)
}

// NewLaptopClient returns a new laptop client
func NewDumpClient(cc *grpc.ClientConn) *DumpClient {
	service := pb.NewDumpServiceClient(cc)
	document := pb.NewDocumentServiceClient(cc)
	modTableOrderList := make(map[int16]modTableOrder)
	modTableLocalList := make(map[string]int64)
	modTableCentralList := make(map[string]int64)

	// DumpAreaClient is a client to call DumpArea service RPCs

	modTableOrderList[0] = modTableOrder{CentralTableName: "Area", LocalTableName: "Area", ModCtr: 0}
	modTableOrderList[1] = modTableOrder{CentralTableName: "Managers", LocalTableName: "Managers", ModCtr: 0}
	modTableOrderList[2] = modTableOrder{CentralTableName: "center", LocalTableName: "Center", ModCtr: 0}
	modTableOrderList[3] = modTableOrder{CentralTableName: "customer", LocalTableName: "Customer", ModCtr: 0}
	modTableOrderList[4] = modTableOrder{CentralTableName: "addresses", LocalTableName: "Addresses", ModCtr: 0}
	modTableOrderList[5] = modTableOrder{CentralTableName: "lnmaster", LocalTableName: "lnMaster", ModCtr: 0}
	modTableOrderList[6] = modTableOrder{CentralTableName: "samaster", LocalTableName: "saMaster", ModCtr: 0}
	modTableOrderList[7] = modTableOrder{CentralTableName: "trnmaster", LocalTableName: "trnMaster", ModCtr: 0}
	modTableOrderList[8] = modTableOrder{CentralTableName: "satrnmaster", LocalTableName: "satrnMaster", ModCtr: 0}
	modTableOrderList[9] = modTableOrder{CentralTableName: "loaninst", LocalTableName: "LoanInst", ModCtr: 0}
	modTableOrderList[10] = modTableOrder{CentralTableName: "lnchrgdata", LocalTableName: "lnChrgData", ModCtr: 0}
	modTableOrderList[11] = modTableOrder{CentralTableName: "custaddinfolist", LocalTableName: "CustAddInfoList", ModCtr: 0}
	modTableOrderList[12] = modTableOrder{CentralTableName: "custaddinfogroup", LocalTableName: "CustAddInfoGroup", ModCtr: 0}
	modTableOrderList[13] = modTableOrder{CentralTableName: "custaddinfogroupneed", LocalTableName: "CustAddInfoGroupNeed", ModCtr: 0}
	modTableOrderList[14] = modTableOrder{CentralTableName: "custaddinfo", LocalTableName: "CustAddInfo", ModCtr: 0}
	modTableOrderList[15] = modTableOrder{CentralTableName: "mutualfund", LocalTableName: "Mutual_Fund", ModCtr: 0}
	modTableOrderList[16] = modTableOrder{CentralTableName: "referencesdetails", LocalTableName: "ReferencesDetails", ModCtr: 0}
	modTableOrderList[17] = modTableOrder{CentralTableName: "centerworker", LocalTableName: "Center_Worker", ModCtr: 0}
	modTableOrderList[18] = modTableOrder{CentralTableName: "writeoff", LocalTableName: "Writeoff", ModCtr: 0}
	modTableOrderList[19] = modTableOrder{CentralTableName: "accounts", LocalTableName: "Accounts", ModCtr: 0}
	modTableOrderList[20] = modTableOrder{CentralTableName: "jnldetails", LocalTableName: "jnlDetails", ModCtr: 0}
	modTableOrderList[21] = modTableOrder{CentralTableName: "jnlheaders", LocalTableName: "jnlHeaders", ModCtr: 0}
	modTableOrderList[22] = modTableOrder{CentralTableName: "ledgerdetails", LocalTableName: "Ledger_Details", ModCtr: 0}

	modTableOrderList[22] = modTableOrder{CentralTableName: "MultiplePaymentReceipt", LocalTableName: "MultiplePaymentReceipt", ModCtr: 0}
	modTableOrderList[22] = modTableOrder{CentralTableName: "InActiveCID", LocalTableName: "InActiveCID", ModCtr: 0}
	modTableOrderList[22] = modTableOrder{CentralTableName: "ReactivateWriteoff", LocalTableName: "ReactivateWriteoff", ModCtr: 0}
	modTableOrderList[22] = modTableOrder{CentralTableName: "LnBeneficiary", LocalTableName: "LnBeneficiary", ModCtr: 0}

	return &DumpClient{
		service:             service,
		document:            document,
		ModTableLocalList:   modTableLocalList,
		ModTableCentralList: modTableCentralList}
}

// ResetFunc resets the context timeout timer
type ResetFunc func()

// WithTimeoutReset returns a child context which is canceled after the provided duration elapses.
// The returned ResetFunc may be called before the context is canceled to restart the timeout timer.
// Unlike context.WithTimeout, the returned context will not report the correct deadline.
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
