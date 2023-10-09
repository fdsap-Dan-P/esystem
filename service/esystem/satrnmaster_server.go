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

func (server *DumpServer) CreateDumpSaTrnMaster(stream pb.DumpService_CreateDumpSaTrnMasterServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpSaTrnMaster:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		saTrnMaster := req.GetSaTrnMaster()

		log.Printf("received SaTrnMaster request: id = %s", saTrnMaster)

		saTrnMasterReq := model.SaTrnMaster{
			ModCtr:      saTrnMaster.ModCtr,
			BrCode:      saTrnMaster.BrCode,
			ModAction:   saTrnMaster.ModAction,
			Acc:         saTrnMaster.Acc,
			TrnDate:     saTrnMaster.TrnDate.AsTime(),
			Trn:         saTrnMaster.Trn,
			TrnType:     util.NullProto2Int64(saTrnMaster.TrnType),
			OrNo:        util.NullProto2Int64(saTrnMaster.OrNo),
			TrnAmt:      util.NullProto2Decimal(saTrnMaster.TrnAmt),
			RefNo:       util.NullProto2String(saTrnMaster.RefNo),
			Particular:  saTrnMaster.Particular,
			TermId:      saTrnMaster.TermId,
			UserName:    saTrnMaster.UserName,
			PendApprove: saTrnMaster.PendApprove,
		}
		err = server.dump.CreateSaTrnMaster(context.Background(), saTrnMasterReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create SaTrnMaster: %v", err))
		}

		res := &pb.CreateDumpSaTrnMasterResponse{
			ModCtr: []int64{saTrnMaster.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateSaTrnMaster(context.Background(), stream.)
		// log.Printf("CreateDumpSaTrnMaster: %v", req)
	}
	return nil
}
