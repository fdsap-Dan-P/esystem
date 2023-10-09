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

func (server *DumpServer) CreateDumpReactivateWriteoff(stream pb.DumpService_CreateDumpReactivateWriteoffServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpReactivateWriteoff:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		reactivateWriteoff := req.GetReactivateWriteoff()

		log.Printf("received ReactivateWriteoff request: id = %s", reactivateWriteoff)

		reactivateWriteoffReq := model.ReactivateWriteoff{
			ModCtr:       reactivateWriteoff.ModCtr,
			BrCode:       reactivateWriteoff.BrCode,
			ID:           reactivateWriteoff.ID,
			CID:          reactivateWriteoff.CID,
			DeactivateBy: util.NullProto2String(reactivateWriteoff.DeactivateBy),
			ReactivateBy: util.NullProto2String(reactivateWriteoff.ReactivateBy),
			Status:       reactivateWriteoff.Status,
			StatusDate:   reactivateWriteoff.StatusDate.AsTime(),
		}
		err = server.dump.CreateReactivateWriteoff(context.Background(), reactivateWriteoffReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create ReactivateWriteoff: %v", err))
		}

		res := &pb.CreateDumpReactivateWriteoffResponse{

			ModCtr: []int64{reactivateWriteoff.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateReactivateWriteoff(context.Background(), stream.)
		// log.Printf("CreateDumpReactivateWriteoff: %v", req)
	}
	return nil
}
