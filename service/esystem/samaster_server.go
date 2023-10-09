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

func (server *DumpServer) CreateDumpSaMaster(stream pb.DumpService_CreateDumpSaMasterServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpSaMaster:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		saMaster := req.GetSaMaster()

		log.Printf("received SaMaster request: id = %s", saMaster)

		saMasterReq := model.SaMaster{
			ModCtr:     saMaster.ModCtr,
			BrCode:     saMaster.BrCode,
			ModAction:  saMaster.ModAction,
			Acc:        saMaster.Acc,
			CID:        saMaster.CID,
			Type:       saMaster.Type,
			Balance:    util.NullProto2Decimal(saMaster.Balance),
			DoLastTrn:  util.NullProto2Time(saMaster.DoLastTrn),
			DoStatus:   util.NullProto2Time(saMaster.DoStatus),
			Dopen:      util.NullProto2Time(saMaster.Dopen),
			DoMaturity: util.NullProto2Time(saMaster.DoMaturity),
			Status:     util.NullProto2String(saMaster.Status),
		}
		err = server.dump.CreateSaMaster(context.Background(), saMasterReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create SaMaster: %v", err))
		}

		res := &pb.CreateDumpSaMasterResponse{
			ModCtr: []int64{saMaster.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateSaMaster(context.Background(), stream.)
		// log.Printf("CreateDumpSaMaster: %v", req)
	}
	return nil
}
