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

func (server *DumpServer) CreateDumpLnChrgData(stream pb.DumpService_CreateDumpLnChrgDataServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpLnChrgData:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		lnChrgData := req.GetLnChrgData()

		log.Printf("received LnChrgData request: id = %s", lnChrgData)

		lnChrgDataReq := model.LnChrgData{
			ModCtr:    lnChrgData.ModCtr,
			BrCode:    lnChrgData.BrCode,
			ModAction: lnChrgData.ModAction,
			Acc:       lnChrgData.Acc,
			ChrgCode:  lnChrgData.ChrgCode,
			RefAcc:    util.NullProto2String(lnChrgData.RefAcc),
			ChrAmnt:   util.Proto2Decimal(lnChrgData.ChrAmnt),
		}
		err = server.dump.CreateLnChrgData(context.Background(), lnChrgDataReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create LnChrgData: %v", err))
		}

		res := &pb.CreateDumpLnChrgDataResponse{
			ModCtr: []int64{lnChrgData.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateLnChrgData(context.Background(), stream.)
		// log.Printf("CreateDumpLnChrgData: %v", req)
	}
	return nil
}
