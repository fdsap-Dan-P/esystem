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

func (server *DumpServer) CreateDumpLnMaster(stream pb.DumpService_CreateDumpLnMasterServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpLnMaster:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		lnMaster := req.GetLnMaster()

		log.Printf("received LnMaster request: id = %s", lnMaster)

		lnMasterReq := model.LnMaster{
			ModCtr:      lnMaster.ModCtr,
			BrCode:      lnMaster.BrCode,
			ModAction:   lnMaster.ModAction,
			CID:         lnMaster.CID,
			Acc:         lnMaster.Acc,
			AcctType:    util.NullProto2Int64(lnMaster.AcctType),
			DisbDate:    util.NullProto2Time(lnMaster.DisbDate),
			Principal:   util.NullProto2Decimal(lnMaster.Principal),
			Interest:    util.NullProto2Decimal(lnMaster.Interest),
			NetProceed:  util.NullProto2Decimal(lnMaster.NetProceed),
			Gives:       util.NullProto2Int64(lnMaster.Gives),
			Frequency:   util.NullProto2Int64(lnMaster.Frequency),
			AnnumDiv:    util.NullProto2Int64(lnMaster.AnnumDiv),
			Prin:        util.NullProto2Decimal(lnMaster.Prin),
			IntR:        util.NullProto2Decimal(lnMaster.IntR),
			WaivedInt:   util.NullProto2Decimal(lnMaster.WaivedInt),
			WeeksPaid:   util.NullProto2Int64(lnMaster.WeeksPaid),
			DoMaturity:  util.NullProto2Time(lnMaster.DoMaturity),
			ConIntRate:  util.NullProto2Decimal(lnMaster.ConIntRate),
			Status:      util.NullProto2String(lnMaster.Status),
			Cycle:       util.NullProto2Int64(lnMaster.Cycle),
			LNGrpCode:   util.NullProto2Int64(lnMaster.LNGrpCode),
			Proff:       util.NullProto2Int64(lnMaster.Proff),
			FundSource:  util.NullProto2String(lnMaster.FundSource),
			DOSRI:       util.NullProto2Bool(lnMaster.DOSRI),
			LnCategory:  util.NullProto2Int64(lnMaster.LnCategory),
			OpenDate:    util.NullProto2Time(lnMaster.OpenDate),
			LastTrnDate: util.NullProto2Time(lnMaster.LastTrnDate),
			DisbBy:      util.NullProto2String(lnMaster.DisbBy),
		}
		err = server.dump.CreateLnMaster(context.Background(), lnMasterReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create LnMaster: %v", err))
		}

		res := &pb.CreateDumpLnMasterResponse{
			ModCtr: []int64{lnMaster.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateLnMaster(context.Background(), stream.)
		// log.Printf("CreateDumpLnMaster: %v", req)
	}
	return nil
}
