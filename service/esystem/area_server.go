package service

import (
	"context"
	"io"

	pb "simplebank/pb"
	"simplebank/util"

	"log"

	model "simplebank/db/datastore/esystemlocal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.DumpServiceServer = (*DumpServer)(nil)

func (server *DumpServer) CreateDumpArea(stream pb.DumpService_CreateDumpAreaServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpArea:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		area := req.GetArea()

		log.Printf("received Area request: id = %s", area)

		areaReq := model.Area{
			ModCtr:    area.ModCtr,
			BrCode:    area.BrCode,
			ModAction: area.ModAction,
			AreaCode:  area.AreaCode,
			Area:      util.NullProto2String(area.Area),
		}
		err = server.dump.CreateArea(context.Background(), areaReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Area: %v", err))
		}

		res := &pb.CreateDumpAreaResponse{
			ModCtr: []int64{area.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateArea(context.Background(), stream.)
		// log.Printf("CreateDumpArea: %v", req)
	}
	return nil
}
