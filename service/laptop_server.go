package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "simplebank/pb"
	"simplebank/util"

	dsAcc "simplebank/db/datastore/account"
	dsRef "simplebank/db/datastore/reference"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxImageSize = 1 << 20

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	db  *sql.DB
	ref *dsRef.QueriesReference
	inv dsAcc.StoreAccount //ds.LaptopStore
	// imageStore  ds.ImageStore
	// ratingStore ds.RatingStore
	pb.UnimplementedLaptopServiceServer
}

// NewLaptopServer returns a new LaptopServer
// func NewLaptopServer(laptopStore dsAcc.StoreAccount, imageStore ds.ImageStore, ratingStore ds.RatingStore) *LaptopServer {

func NewLaptopServer(db *sql.DB) *LaptopServer {
	// a := pb.UnimplementedLaptopServiceServer{}
	return &LaptopServer{
		db:  db,
		ref: dsRef.New(db),
		inv: dsAcc.NewStoreAccount(db)}
}

func (server *LaptopServer) mustEmbedUnimplementedLaptopServiceServer() {
}

func SeachItem(items []dsRef.ReferenceInfo, item string) int64 {
	for i := range items {
		if items[i].Title == item {
			return int64(i)
		}
	}
	return 0
}

/*
UpdateInventoryItem(*InventoryItemRequest, InventoryItemService_UpdateInventoryItemServer) error
GetInventoryItem(*GetInventoryItemRequest, InventoryItemService_GetInventoryItemServer) error
GetInventoryItembyUuid(*GetInventoryItemRequestbyUuid, InventoryItemService_GetInventoryItembyUuidServer) error
GetListInventoryItembyGeneric(*GetListInventoryItemRequestbyGeneric, InventoryItemService_GetListInventoryItembyGenericServer) error
GetListInventoryItembyBrand(*GetListInventoryItemRequestbyBrand, InventoryItemService_GetListInventoryItembyBrandServer) error
SearchInventoryItem(*SearchInventoryItemRequest, InventoryItemService_SearchInventoryItemServer) error

type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error
	UploadImage(LaptopService_UploadImageServer) error
	RateLaptop(LaptopService_RateLaptopServer) error
	mustEmbedUnimplementedLaptopServiceServer()
}

*/

// func (server *LaptopServer) GetInventoryItem(
// 	in *pb.GetInventoryItemRequest,
// 	stream pb.InventoryItemService_GetInventoryItemServer) error {

// 	log.Printf("fetch response for id : %d", in.Id)

// 	//use wait group to allow process to be concurrent
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(count int64) {
// 			defer wg.Done()

// 			//time sleep to simulate server process time
// 			time.Sleep(time.Duration(count) * time.Second)
// 			resp := pb.Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
// 			if err := srv.Send(&resp); err != nil {
// 				log.Printf("send error %v", err)
// 			}
// 			log.Printf("finishing request number : %d", count)
// 		}(int64(i))
// 	}

// 	wg.Wait()
// 	return nil

// 	return errors.New("te")
// }

func (server *LaptopServer) Laptop2InventoryItem(ctx context.Context, laptop *pb.Laptop) dsAcc.InventoryItemRequest {
	gen, _ := server.ref.GetReferenceInfobyTitle(ctx, "GenericName", 0, "Laptop")
	measure, _ := server.ref.GetReferenceInfobyTitle(ctx, "UnitMeasure", 0, "Unit")

	invReq := dsAcc.InventoryItemRequest{
		Id:            laptop.Id,
		Uuid:          util.ToUUID(laptop.Uuid),
		ItemName:      laptop.Name,
		GenericNameId: util.SetNullInt64(gen.Id),
		BrandNameId:   util.SetNullInt64(laptop.BrandId),
		MeasureId:     measure.Id,
		// SupplierId:,
		// Inventory:,
		// UnitPrice:,
		// BookValue:,
		// Unit:,
		// MeasureId:,
		// Discount:,
		// TaxRate:,
		// DateManufactured:,
		// DateExpired:,
		// Remarks:,
		// OtherInfo:,
	}
	// Get all Inventory Specs

	items, _ := server.ref.ListReference(ctx, dsRef.ListReferenceParams{
		RefType: "InventorySpecs",
		Limit:   0,
		Offset:  1,
	})

	// RAM
	itemId := SeachItem(items, "Random Access Memory")
	invReq.SpecsNumberList = append(invReq.SpecsNumberList,
		dsAcc.InventorySpecsNumberRequest{
			SpecsId: itemId,
			Value:   decimal.NewFromInt(int64(laptop.Ram.Value)),
			Value2:  decimal.NewFromFloat(1),
		})

	// GPUs
	invReq.InventoryItemList = append(invReq.InventoryItemList,
		dsAcc.InventoryItemRequest{
			ParentId:      sql.NullInt64{Valid: true, Int64: invReq.Id},
			ItemName:      laptop.Name,
			GenericNameId: sql.NullInt64{Valid: true, Int64: gen.Id},
			BrandNameId:   sql.NullInt64{Valid: true, Int64: laptop.BrandId},
			// SupplierId:,
			// Inventory:,
			// UnitPrice:,
			// BookValue:,
			// Unit:,
			// MeasureId:,
			// Discount:,
			// TaxRate:,
			// DateManufactured:,
			// DateExpired:,
			// Remarks:,
			// OtherInfo:,
		})

	return invReq
}

/*
func (server *LaptopServer) InventoryItem2Laptop(ctx context.Context, invReq model.InventoryItem) *pb.Laptop {

	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{

		Id:            laptop.Id,
		Uuid:          util.ToUUID(laptop.Uuid),
		ItemName:      laptop.Name,
		GenericNameId: util.SetNullInt64(gen.Id),
		BrandNameId:   util.SetNullInt64(laptop.BrandId),
		MeasureId:     measure.Id,

		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.StorageSize{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	laptop := pb.Laptop

	gen, _ := server.ref.GetReferenceInfobyTitle(ctx, "GenericName", 0, "Laptop")
	measure, _ := server.ref.GetReferenceInfobyTitle(ctx, "UnitMeasure", 0, "Unit")

	invReq := dsAcc.InventoryItemRequest{
		Id:            laptop.AccountId,
		ItemName:      laptop.Name,
		GenericNameId: sql.NullInt64{Valid: true, Int64: gen.Id},
		BrandNameId:   sql.NullInt64{Valid: true, Int64: laptop.BrandId},
		MeasureId:     measure.Id,
		// SupplierId:,
		// Inventory:,
		// UnitPrice:,
		// BookValue:,
		// Unit:,
		// MeasureId:,
		// Discount:,
		// TaxRate:,
		// DateManufactured:,
		// DateExpired:,
		// Remarks:,
		// OtherInfo:,
	}
	// Get all Inventory Specs

	items, _ := server.ref.ListReference(ctx, dsRef.ListReferenceParams{
		RefType: "InventorySpecs",
		Limit:   0,
		Offset:  1,
	})

	// RAM
	itemId := SeachItem(items, "Random Access Memory")
	invReq.SpecsNumberList = append(invReq.SpecsNumberList,
		dsAcc.InventorySpecsNumberRequest{
			SpecsId: itemId,
			Value:   decimal.NewFromInt(int64(laptop.Ram.Value)),
			Value2:  decimal.NewFromFloat(1),
		})

	// GPUs
	invReq.InventoryItemList = append(invReq.InventoryItemList,
		dsAcc.InventoryItemRequest{
			ParentId:      sql.NullInt64{Valid: true, Int64: invReq.Id},
			ItemName:      laptop.Name,
			GenericNameId: sql.NullInt64{Valid: true, Int64: gen.Id},
			BrandNameId:   sql.NullInt64{Valid: true, Int64: laptop.BrandId},
			// SupplierId:,
			// Inventory:,
			// UnitPrice:,
			// BookValue:,
			// Unit:,
			// MeasureId:,
			// Discount:,
			// TaxRate:,
			// DateManufactured:,
			// DateExpired:,
			// Remarks:,
			// OtherInfo:,
		})

	return invReq
}
*/

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %v %s", laptop.Id, laptop.Uuid)

	if len(laptop.Uuid) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Uuid)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Uuid = id.String()
	}

	// some heavy processing
	// time.Sleep(6 * time.Second)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to store
	invreq := server.Laptop2InventoryItem(ctx, laptop)
	inv, err := server.inv.CreateInventoryItemFull(ctx, invreq)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, util.ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %v", inv)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

// SearchLaptop is a server-streaming RPC to search for laptops
func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {
	filter := req.SearchRequest
	log.Printf("receive a search-laptop request with filter: %v", filter)

	// --- Remarks for correction
	// err := server.laptopStore.Search(
	// 	stream.Context(),
	// 	filter,
	// 	func(laptop *pb.Laptop) error {
	// 		res := &pb.SearchLaptopResponse{Laptop: laptop}
	// 		err := stream.Send(res)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		log.Printf("sent laptop with id: %s", laptop.GetId())
	// 		return nil
	// 	},
	// )

	// if err != nil {
	// 	return status.Errorf(codes.Internal, "unexpected error: %v", err)
	// }

	return nil
}

// // UploadImage is a client-streaming RPC to upload a laptop image
// func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
// 	req, err := stream.Recv()
// 	if err != nil {
// 		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
// 	}

// 	laptopID := req.GetInfo().GetLaptopId()
// 	imageType := req.GetInfo().GetImageType()
// 	log.Printf("receive an upload-image request for laptop %+v with image type %+v", laptopID, imageType)

// 	return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))

// 	// laptop, err := server.laptopStore.Find(laptopID)
// 	// if err != nil {
// 	// 	return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
// 	// }
// 	// if laptop == nil {
// 	// 	return logError(status.Errorf(codes.InvalidArgument, "laptop id %s doesn't exist", laptopID))
// 	// }

// 	imageData := bytes.Buffer{}
// 	imageSize := 0

// 	for {
// 		err := contextError(stream.Context())
// 		if err != nil {
// 			return err
// 		}

// 		log.Print("waiting to receive more data")

// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			log.Print("no more data")
// 			break
// 		}
// 		if err != nil {
// 			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
// 		}

// 		chunk := req.GetChunkData()
// 		size := len(chunk)

// 		log.Printf("received a chunk with size: %d", size)

// 		imageSize += size
// 		if imageSize > maxImageSize {
// 			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
// 		}

// 		// write slowly
// 		// time.Sleep(time.Second)

// 		_, err = imageData.Write(chunk)
// 		if err != nil {
// 			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
// 		}
// 	}

// 	// imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
// 	// if err != nil {
// 	// 	return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
// 	// }

// 	// res := &pb.UploadImageResponse{
// 	// 	Id:   imageID,
// 	// 	Size: uint32(imageSize),
// 	// }

// 	// err = stream.SendAndClose(res)
// 	// if err != nil {
// 	// 	return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
// 	// }

// 	// log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
// 	return nil
// }

// // RateLaptop is a bidirectional-streaming RPC that allows client to rate a stream of laptops
// // with a score, and returns a stream of average score for each of them
// func (server *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
// 	for {
// 		err := contextError(stream.Context())
// 		if err != nil {
// 			return err
// 		}

// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			log.Print("no more data")
// 			break
// 		}
// 		if err != nil {
// 			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
// 		}

// 		laptopID := req.GetLaptopId()
// 		score := req.GetScore()

// 		log.Printf("received a rate-laptop request: id = %+v, score = %.2f", laptopID, score)

// 		// 	found, err := server.laptopStore.Find(laptopID)
// 		// 	if err != nil {
// 		// 		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
// 		// 	}
// 		// 	if found == nil {
// 		// 		return logError(status.Errorf(codes.NotFound, "laptopID %s is not found", laptopID))
// 		// 	}

// 		// 	rating, err := server.ratingStore.Add(laptopID, score)
// 		// 	if err != nil {
// 		// 		return logError(status.Errorf(codes.Internal, "cannot add rating to the store: %v", err))
// 		// 	}

// 		// 	res := &pb.RateLaptopResponse{
// 		// 		LaptopId:     laptopID,
// 		// 		RatedCount:   rating.Count,
// 		// 		AverageScore: rating.Sum / float64(rating.Count),
// 		// 	}

// 		// 	err = stream.Send(res)
// 		// 	if err != nil {
// 		// 		return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
// 		// 	}
// 	}

// 	return nil
// }

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

var _ pb.LaptopServiceServer = (*LaptopServer)(nil)
