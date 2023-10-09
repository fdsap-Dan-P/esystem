package service_test

import (
	"context"
	"io"
	"net"
	"testing"

	pb "simplebank/pb"
	"simplebank/sample"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateDumpArea(t *testing.T) {

	t.Parallel()

	area := sample.NewDumpArea()
	// err := testDumpStore.CreateDumpArea(laptop)
	// require.NoError(t, err)

	serverAddress := startTestDumpAreaServer(t, testDumpStore)
	laptopClient := newTestDumpAreaClient(t, serverAddress)

	stream, err := laptopClient.CreateDumpArea(context.Background())
	require.NoError(t, err)

	n := 1
	for i := 0; i < n; i++ {
		req := &pb.CreateDumpAreaRequest{
			Area: area,
		}

		err := stream.Send(req)
		require.NoError(t, err)
	}

	err = stream.CloseSend()
	require.NoError(t, err)

	for idx := 0; ; idx++ {
		res, err := stream.Recv()
		if err == io.EOF {
			require.Equal(t, n, idx)
			return
		}

		require.NoError(t, err)
		require.Equal(t, area.ModCtr, res.GetModCtr())
	}
}

func startTestDumpAreaServer(t *testing.T, testDump pb.DumpServiceServer) string {

	grpcServer := grpc.NewServer()

	pb.RegisterDumpServiceServer(grpcServer, testDump)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestDumpAreaClient(t *testing.T, serverAddress string) pb.DumpServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewDumpServiceClient(conn)
}

// func TestClientCreateDumpArea(t *testing.T) {

// 	t.Parallel()

// 	AreaStore := ds.NewInMemoryLaptopStore()
// 	serverAddress := startTestLaptopServer(t, laptopStore, nil, nil)
// 	laptopClient := newTestLaptopClient(t, serverAddress)

// 	laptop := sample.NewLaptop()
// 	expectedID := laptop.Id
// 	req := &pb.CreateLaptopRequest{
// 		Laptop: laptop,
// 	}

// 	res, err := laptopClient.CreateLaptop(context.Background(), req)
// 	require.NoError(t, err)
// 	require.NotNil(t, res)
// 	require.Equal(t, expectedID, res.Id)

// 	// check that the laptop is saved to the store
// 	other, err := laptopStore.Find(res.Id)
// 	require.NoError(t, err)
// 	require.NotNil(t, other)

// 	// check that the saved laptop is the same as the one we send
// 	requireSameLaptop(t, laptop, other)
// }

// func TestClientSearchLaptop(t *testing.T) {
// 	t.Parallel()

// 	filter := &pb.Filter{
// 		MaxPriceUsd: 2000,
// 		MinCpuCores: 4,
// 		MinCpuGhz:   2.2,
// 		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
// 	}

// 	laptopStore := ds.NewInMemoryLaptopStore()
// 	expectedIDs := make(map[string]bool)

// 	for i := 0; i < 6; i++ {
// 		laptop := sample.NewLaptop()

// 		switch i {
// 		case 0:
// 			laptop.PriceUsd = 2500
// 		case 1:
// 			laptop.Cpu.NumberCores = 2
// 		case 2:
// 			laptop.Cpu.MinGhz = 2.0
// 		case 3:
// 			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
// 		case 4:
// 			laptop.PriceUsd = 1999
// 			laptop.Cpu.NumberCores = 4
// 			laptop.Cpu.MinGhz = 2.5
// 			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
// 			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
// 			expectedIDs[laptop.Id] = true
// 		case 5:
// 			laptop.PriceUsd = 2000
// 			laptop.Cpu.NumberCores = 6
// 			laptop.Cpu.MinGhz = 2.8
// 			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
// 			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
// 			expectedIDs[laptop.Id] = true
// 		}

// 		err := laptopStore.Save(laptop)
// 		require.NoError(t, err)
// 	}

// 	serverAddress := startTestLaptopServer(t, laptopStore, nil, nil)
// 	laptopClient := newTestLaptopClient(t, serverAddress)

// 	req := &pb.SearchLaptopRequest{Filter: filter}
// 	stream, err := laptopClient.SearchLaptop(context.Background(), req)
// 	require.NoError(t, err)

// 	found := 0
// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}

// 		require.NoError(t, err)
// 		require.Contains(t, expectedIDs, res.GetLaptop().GetId())

// 		found += 1
// 	}

// 	require.Equal(t, len(expectedIDs), found)
// }

// func TestClientRateLaptop(t *testing.T) {
// 	t.Parallel()

// 	laptopStore := ds.NewInMemoryLaptopStore()
// 	ratingStore := ds.NewInMemoryRatingStore()

// 	laptop := sample.NewLaptop()
// 	err := laptopStore.Save(laptop)
// 	require.NoError(t, err)

// 	serverAddress := startTestLaptopServer(t, laptopStore, nil, ratingStore)
// 	laptopClient := newTestLaptopClient(t, serverAddress)

// 	stream, err := laptopClient.RateLaptop(context.Background())
// 	require.NoError(t, err)

// 	scores := []float64{8, 7.5, 10}
// 	averages := []float64{8, 7.75, 8.5}

// 	n := len(scores)
// 	for i := 0; i < n; i++ {
// 		req := &pb.RateLaptopRequest{
// 			LaptopId: laptop.GetId(),
// 			Score:    scores[i],
// 		}

// 		err := stream.Send(req)
// 		require.NoError(t, err)
// 	}

// 	err = stream.CloseSend()
// 	require.NoError(t, err)

// 	for idx := 0; ; idx++ {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			require.Equal(t, n, idx)
// 			return
// 		}

// 		require.NoError(t, err)
// 		require.Equal(t, laptop.GetId(), res.GetLaptopId())
// 		require.Equal(t, uint32(idx+1), res.GetRatedCount())
// 		require.Equal(t, averages[idx], res.GetAverageScore())
// 	}
// }

// func startTestLaptopServer(t *testing.T, laptopStore ds.LaptopStore,
// 	imageStore ds.ImageStore, ratingStore ds.RatingStore) string {
// 	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

// 	listener, err := net.Listen("tcp", ":0") // random available port
// 	require.NoError(t, err)

// 	go grpcServer.Serve(listener)

// 	return listener.Addr().String()
// }

// func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
// 	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
// 	require.NoError(t, err)
// 	return pb.NewLaptopServiceClient(conn)
// }

// func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
// 	json1, err := serializer.ProtobufToJSON(laptop1)
// 	require.NoError(t, err)

// 	json2, err := serializer.ProtobufToJSON(laptop2)
// 	require.NoError(t, err)

// 	require.Equal(t, json1, json2)
// }
