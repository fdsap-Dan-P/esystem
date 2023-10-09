package client

import (
	"database/sql"
	pb "simplebank/pb"

	local "simplebank/db/datastore/esystemlocal"

	_ "github.com/denisenkom/go-mssqldb"
	"google.golang.org/grpc"
)

var QueriesLocal *local.QueriesLocal
var DBLocal *sql.DB

// KPlusAreaClient is a client to call KPlusArea service RPCs
type KPlusClient struct {
	service pb.KPlusServiceClient
}

// NewLaptopClient returns a new laptop client
func NewKPlusClient(cc *grpc.ClientConn) *KPlusClient {
	service := pb.NewKPlusServiceClient(cc)

	// KPlusAreaClient is a client to call KPlusArea service RPCs

	return &KPlusClient{
		service: service,
	}
}
