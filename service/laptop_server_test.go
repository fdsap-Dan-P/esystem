package service

import (
	"context"
	"log"
	pb "simplebank/pb"
	"simplebank/sample"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLaptop2InventoryItem(t *testing.T) {
	log.Printf("Test Start %v", 1)

	laptop := sample.NewLaptop()

	brand, e := testLaptop.ref.GetReferenceInfobyTitle(context.Background(), "BrandName", 0, laptop.Brand)
	require.NoError(t, e)
	laptop.BrandId = brand.Id

	// log.Printf("TestLaptop2InventoryItem testLaptop:  %+v", testLaptop)
	// log.Printf("TestLaptop2InventoryItem laptop: %+v", laptop)
	invReq := testLaptop.Laptop2InventoryItem(context.Background(), laptop)
	// a := testLaptop.GetInventoryItem()
	log.Printf("Test Start %+v", invReq)

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	lap, err := testLaptop.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	log.Printf("Test Start lap Error: %+v ", err)
	log.Printf("Test Start lap: %+v ", lap)

	// require.True(t, false)
}

// func TestServerCreateLaptop(t *testing.T) {
// 	t.Parallel()

// 	laptopNoID := sample.NewLaptop()
// 	laptopNoID.Uuid = ""

// 	laptopInvalidID := sample.NewLaptop()
// 	laptopInvalidID.Uuid = "invalid-uuid"

// 	laptopDuplicateID := sample.NewLaptop()
// 	storeDuplicateID := ds.NewInMemoryLaptopStore()
// 	err := storeDuplicateID.Save(laptopDuplicateID)
// 	require.Nil(t, err)

// testCases := []struct {
// 	name   string
// 	laptop *pb.Laptop
// 	store  ds.LaptopStore
// 	code   codes.Code
// }{
// 	{
// 		name:   "success_with_id",
// 		laptop: sample.NewLaptop(),
// 		store:  ds.NewInMemoryLaptopStore(),
// 		code:   codes.OK,
// 	},
// 	{
// 		name:   "success_no_id",
// 		laptop: laptopNoID,
// 		store:  ds.NewInMemoryLaptopStore(),
// 		code:   codes.OK,
// 	},
// 	{
// 		name:   "failure_invalid_id",
// 		laptop: laptopInvalidID,
// 		store:  ds.NewInMemoryLaptopStore(),
// 		code:   codes.InvalidArgument,
// 	},
// 	{
// 		name:   "failure_duplicate_id",
// 		laptop: laptopDuplicateID,
// 		store:  storeDuplicateID,
// 		code:   codes.AlreadyExists,
// 	},
// }

// for i := range testCases {
// 	tc := testCases[i]

// 	t.Run(tc.name, func(t *testing.T) {
// 		t.Parallel()

// 		req := &pb.CreateLaptopRequest{
// 			Laptop: tc.laptop,
// 		}

// 		server := service.NewLaptopServer(tc.store, nil, nil)
// 		res, err := server.CreateLaptop(context.Background(), req)
// 		if tc.code == codes.OK {
// 			require.NoError(t, err)
// 			require.NotNil(t, res)
// 			require.NotEmpty(t, res.Id)
// 			if len(tc.laptop.Id) > 0 {
// 				require.Equal(t, tc.laptop.Id, res.Id)
// 			}
// 		} else {
// 			require.Error(t, err)
// 			require.Nil(t, res)
// 			st, ok := status.FromError(err)
// 			require.True(t, ok)
// 			require.Equal(t, tc.code, st.Code())
// 		}
// 	})
// }
// }
