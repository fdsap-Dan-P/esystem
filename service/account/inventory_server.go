package service

import (
	"context"
	"database/sql"
	"log"

	pb "simplebank/pb"
	"simplebank/util"

	dsAcc "simplebank/db/datastore/account"
	dsRef "simplebank/db/datastore/reference"

	"google.golang.org/grpc"
)

const maxImageSize = 1 << 20

// LaptopServer is the server that provides laptop services
type InventoryServer struct {
	Db     *sql.DB
	Ref    *dsRef.QueriesReference
	Inv    dsAcc.StoreAccount //ds.LaptopStore
	Server grpc.Server
	// imageStore  ds.ImageStore
	// ratingStore ds.RatingStore
	pb.UnimplementedInventoryItemServiceServer
}

func NewInventoryServer(db *sql.DB) *InventoryServer {
	// a := pb.UnimplementedLaptopServiceServer{}
	return &InventoryServer{
		Db:     db,
		Ref:    dsRef.New(db),
		Inv:    dsAcc.NewStoreAccount(db),
		Server: *grpc.NewServer(),
	}
}

func (server *InventoryServer) GetInventoryItem(
	ctx context.Context, req *pb.GetInventoryItemRequest) (*pb.InventoryItemResponse, error) {

	invItem, err := server.Inv.GetInventoryItem(context.Background(), req.Id)
	if err != nil {
		log.Printf("send error %v", err)
		return nil, err
	}

	return &pb.InventoryItemResponse{
		InventoryItem: &pb.InventoryItem{
			Id:              invItem.Id,
			Uuid:            invItem.Uuid.String(),
			BarCode:         invItem.BarCode.String,
			ItemName:        invItem.ItemName,
			UniqueVariation: invItem.UniqueVariation,
			ParentId:        invItem.ParentId.Int64,
			GenericNameId:   util.NullInt642Proto(invItem.GenericNameId),
			BrandNameId:     util.NullInt642Proto(invItem.BrandNameId),
			MeasureId:       invItem.MeasureId,
			ImageId:         util.NullInt642Proto(invItem.ImageId),
			Remarks:         invItem.Remarks,
			OtherInfo:       util.NullString2Proto(invItem.OtherInfo),
		},
		// GenericName: pb.GenericName{},
		// Measure: pb.Measure{},
		// Image:              Document{},
		// InventorySpecsString: InventorySpecsString{},]
		// InventorySpecsNumber: InventorySpecsNumber{},
	}, nil
}

// log.Printf("fetch response for id : %d", in.Id)

//use wait group to allow process to be concurrent
// var wg sync.WaitGroup

// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	//time sleep to simulate server process time
// 	time.Sleep(time.Duration(1) * time.Second)

// 	invItem, err := server.Inv.GetInventoryItem(context.Background(), req.InventoryItem.Id)
// 	if err != nil {
// 		log.Printf("send error %v", err)
// 	}
// 	resp := pb.InventoryItemResponse{
// 		InventoryItem: &pb.InventoryItem{
// 			Id:              invItem.Id,
// 			Uuid:            invItem.Uuid.String(),
// 			BarCode:         invItem.BarCode.String,
// 			ItemName:        invItem.ItemName,
// 			UniqueVariation: invItem.UniqueVariation,
// 			ParentId:        invItem.ParentId.Int64,
// 			GenericNameId:   util.NullInt64Proto(invItem.GenericNameId),
// 			BrandNameId:     util.NullInt64Proto(invItem.BrandNameId),
// 			MeasureId:       invItem.MeasureId,
// 			ImageId:         util.NullInt64Proto(invItem.ImageId),
// 			Remarks:         invItem.Remarks,
// 			OtherInfo:       util.NullString2Proto(invItem.OtherInfo),
// 		},
// 		// GenericName: pb.GenericName{},
// 		// Measure: pb.Measure{},
// 		// Image:              Document{},
// 		// InventorySpecsString: InventorySpecsString{},]
// 		// InventorySpecsNumber: InventorySpecsNumber{},
// 		// InventorySpecsDate:   InventorySpecsDate{},
// 	}
// 	if err := srv.Send(&resp); err != nil {
// 		log.Printf("send error %v", err)
// 	}
// 	// log.Printf("finishing request number : %d", count)
// }()

// // for i := 0; i < 5; i++ {
// // 	wg.Add(1)
// // 	go func(count int64) {
// // 		defer wg.Done()

// // 		//time sleep to simulate server process time
// // 		time.Sleep(time.Duration(count) * time.Second)
// // 		resp := pb.InventoryItemResponse{
// // 		  }
// // 		if err := srv.Send(&resp); err != nil {
// // 			log.Printf("send error %v", err)
// // 		}
// // 		log.Printf("finishing request number : %d", count)
// // 	}(int64(i))
// // }

// wg.Wait()
// return nil

// var _ pb.InventoryItemServiceServer = (*InventoryServer)(nil)
