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

func (server *DumpServer) CreateDumpTrnMaster(stream pb.DumpService_CreateDumpTrnMasterServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpTrnMaster:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		trnMaster := req.GetTrnMaster()

		log.Printf("received TrnMaster request: id = %s", trnMaster)

		trnMasterReq := model.TrnMaster{
			ModCtr:    trnMaster.ModCtr,
			BrCode:    trnMaster.BrCode,
			ModAction: trnMaster.ModAction,
			Acc:       trnMaster.Acc,
			TrnDate:   trnMaster.TrnDate.AsTime(),
			Trn:       trnMaster.Trn,
			TrnType:   util.NullProto2Int64(trnMaster.TrnType),
			OrNo:      util.NullProto2Int64(trnMaster.OrNo),
			Prin:      util.Proto2Decimal(trnMaster.Prin),
			IntR:      util.Proto2Decimal(trnMaster.IntR),
			WaivedInt: util.Proto2Decimal(trnMaster.WaivedInt),
			RefNo:     util.NullProto2String(trnMaster.RefNo),
		}
		err = server.dump.CreateTrnMaster(context.Background(), trnMasterReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create TrnMaster: %v", err))
		}

		res := &pb.CreateDumpTrnMasterResponse{
			ModCtr: []int64{trnMaster.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateTrnMaster(context.Background(), stream.)
		// log.Printf("CreateDumpTrnMaster: %v", req)
	}
	return nil
}
