package service

import (
	"database/sql"

	pb "simplebank/pb"

	document "simplebank/db/datastore/document"
	dump "simplebank/db/datastore/esystemdump"
)

// var _ pb.DocumentServiceServer = (*docServer.DocumentServer)(nil)
var _ pb.DumpServiceServer = (*DumpServer)(nil)

// DumpServer is the server that provides Area services
type DumpServer struct {
	db   *sql.DB
	dump *dump.QueriesDump
	// storage  docServer.Manager
	document *document.QueriesDocument
	// imageStore  ds.ImageStore
	// ratingStore ds.RatingStore
	*pb.UnimplementedDumpServiceServer
	*pb.UnimplementedDocumentServiceServer
}

// NewDumpServer returns a new DumpServer
// func NewDumpServer(AreaStore dsAcc.StoreAccount, imageStore ds.ImageStore, ratingStore ds.RatingStore) *DumpServer {

func NewDumpServer(db *sql.DB) *DumpServer {
	// a := pb.UnimplementedAreaServiceServer{}
	return &DumpServer{
		db:       db,
		dump:     dump.New(db),
		document: document.New(db),
	}
}
