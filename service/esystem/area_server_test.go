package service_test

// import (
// 	"context"
// 	"io"
// 	"log"
// 	"testing"
// )

// func TestCreateDumpArea(t *testing.T) {
// 	log.Printf("Test Start %v", 1)

// 	stream, err := client.RouteChat(context.Background())
// 	waitc := make(chan struct{})
// 	go func() {
// 		for {
// 			in, err := stream.Recv()
// 			if err == io.EOF {
// 				// read done.
// 				close(waitc)
// 				return
// 			}
// 			if err != nil {
// 				log.Fatalf("Failed to receive a note : %v", err)
// 			}
// 			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
// 		}
// 	}()
// 	for _, note := range notes {
// 		if err := stream.Send(note); err != nil {
// 			log.Fatalf("Failed to send a note: %v", err)
// 		}
// 	}
// 	stream.CloseSend()
// 	<-waitc

// 	// t.Parallel()

// 	// areaNoBrCode := sample.NewDumpArea()
// 	// areaNoBrCode.BrCode = ""

// 	// areaInvalidBrBo := sample.NewDumpArea()
// 	// areaInvalidBrBo.BrCode = "23213"

// 	// areaDuplicateID := sample.NewDumpArea()

// 	// err := testDumpStore.CreateDumpArea(context.Background(),areaDuplicateID)
// 	// require.Nil(t, err)

// 	// testCases := []struct {
// 	// 	name  string
// 	// 	area  *pb.DumpArea
// 	// 	store *service.DumpServer
// 	// 	code  codes.Code
// 	// }{
// 	// 	{
// 	// 		name:  "success_with_id",
// 	// 		area:  sample.NewDumpArea(),
// 	// 		store: service.NewDumpServer(testDB),
// 	// 		code:  codes.OK,
// 	// 	},
// 	// 	{
// 	// 		name:  "success_no_id",
// 	// 		area:  areaDuplicateID,
// 	// 		store: service.NewDumpServer(testDB),
// 	// 		code:  codes.OK,
// 	// 	},
// 	// 	{
// 	// 		name:  "failure_invalid_id",
// 	// 		area:  areaInvalidBrBo,
// 	// 		store: service.NewDumpServer(testDB),
// 	// 		code:  codes.InvalidArgument,
// 	// 	},
// 	// 	{
// 	// 		name:  "failure_duplicate_id",
// 	// 		area:  areaDuplicateID,
// 	// 		store: service.NewDumpServer(testDB),
// 	// 		code:  codes.AlreadyExists,
// 	// 	},
// 	// }

// 	// for i := range testCases {
// 	// 	tc := testCases[i]

// 	// 	t.Run(tc.name, func(t *testing.T) {
// 	// 		t.Parallel()

// 	// 		req := &pb.CreateareaRequest{
// 	// 			area: tc.area,
// 	// 		}

// 	// 		server := service.NewareaServer(tc.store, nil, nil)
// 	// 		res, err := server.Createarea(context.Background(), req)
// 	// 		if tc.code == codes.OK {
// 	// 			require.NoError(t, err)
// 	// 			require.NotNil(t, res)
// 	// 			require.NotEmpty(t, res.Id)
// 	// 			if len(tc.area.Id) > 0 {
// 	// 				require.Equal(t, tc.area.Id, res.Id)
// 	// 			}
// 	// 		} else {
// 	// 			require.Error(t, err)
// 	// 			require.Nil(t, res)
// 	// 			st, ok := status.FromError(err)
// 	// 			require.True(t, ok)
// 	// 			require.Equal(t, tc.code, st.Code())
// 	// 		}
// 	// 	})
// 	// }

// 	// require.True(t, false)
// }
