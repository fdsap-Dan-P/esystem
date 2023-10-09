package service

import (
	"database/sql"
	"log"

	pb "simplebank/pb"

	kplus "simplebank/db/datastore/kplus"
)

// var _ pb.KplusServiceServer = (*docServer.KplusServer)(nil)
var _ pb.KPlusServiceServer = (*KPlusServer)(nil)

// KplusServer is the server that provides Area services

type KPlusServer struct {
	db    *sql.DB
	kplus *kplus.QueriesKPlus
	// imageStore  ds.ImageStore
	// ratingStore ds.RatingStore
	pb.UnimplementedKPlusServiceServer
}

// NewKplusServer returns a new KplusServer
// func NewKplusServer(AreaStore dsAcc.StoreAccount, imageStore ds.ImageStore, ratingStore ds.RatingStore) *KplusServer {

func NewKplusServer(db *sql.DB) *KPlusServer {
	// a := pb.UnimplementedAreaServiceServer{}
	log.Println("NewKplusServer..........")
	return &KPlusServer{
		db:    db,
		kplus: kplus.New(db),
	}
}
