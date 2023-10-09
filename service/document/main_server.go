package service

import (
	"database/sql"

	pb "simplebank/pb"

	doc "simplebank/db/datastore/document"
)

var _ pb.DocumentServiceServer = (*DocumentServer)(nil)

// DocumentServer is the server that provides Area services
type DocumentServer struct {
	db       *sql.DB
	storage  Manager
	document *doc.QueriesDocument
	*pb.UnimplementedDocumentServiceServer
}

// NewDocumentServer returns a new DocumentServer
// func NewDocumentServer(AreaStore dsAcc.StoreAccount, imageStore ds.ImageStore, ratingStore ds.RatingStore) *DocumentServer {

func NewDocumentServer(db *sql.DB, storage Manager) *DocumentServer {
	// a := pb.UnimplementedDocumentServiceServer{}
	return &DocumentServer{
		db:       db,
		document: doc.New(db),
		storage:  storage,
		// a:        a,
	}
}
