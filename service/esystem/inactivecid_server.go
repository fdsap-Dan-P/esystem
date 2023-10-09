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

func (server *DumpServer) CreateDumpInActiveCID(stream pb.DumpService_CreateDumpInActiveCIDServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpInActiveCID:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		inActiveCID := req.GetInActiveCID()

		log.Printf("received InActiveCID request: id = %s", inActiveCID)

		inActiveCIDReq := model.InActiveCID{
			ModCtr:        inActiveCID.ModCtr,
			BrCode:        inActiveCID.BrCode,
			CID:           inActiveCID.CID,
			InActive:      inActiveCID.InActive,
			DateStart:     inActiveCID.DateStart.AsTime(),
			DateEnd:       util.NullProto2Time(inActiveCID.DateEnd),
			UserId:        inActiveCID.UserId,
			DeactivatedBy: util.NullProto2String(inActiveCID.DeactivatedBy),
		}
		err = server.dump.CreateInActiveCID(context.Background(), inActiveCIDReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create InActiveCID: %v", err))
		}

		res := &pb.CreateDumpInActiveCIDResponse{
			ModCtr: []int64{inActiveCID.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateInActiveCID(context.Background(), stream.)
		// log.Printf("CreateDumpInActiveCID: %v", req)
	}
	return nil
}
