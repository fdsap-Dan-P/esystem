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

func (server *DumpServer) CreateDumpMutualFund(stream pb.DumpService_CreateDumpMutualFundServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpMutualFund:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		mutualFund := req.GetMutualFund()

		log.Printf("received MutualFund request: id = %s", mutualFund)

		mutualFundReq := model.MutualFund{
			ModCtr:    mutualFund.ModCtr,
			BrCode:    mutualFund.BrCode,
			ModAction: mutualFund.ModAction,
			CID:       mutualFund.CID,
			OrNo:      util.NullProto2Int64(mutualFund.OrNo),
			TrnDate:   mutualFund.TrnDate.AsTime(),
			TrnType:   util.NullProto2String(mutualFund.TrnType),
			TrnAmt:    util.Proto2Decimal(mutualFund.TrnAmt),
			UserName:  util.NullProto2String(mutualFund.UserName),
		}
		err = server.dump.CreateMutualFund(context.Background(), mutualFundReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create MutualFund: %v", err))
		}

		res := &pb.CreateDumpMutualFundResponse{
			ModCtr: []int64{mutualFund.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateMutualFund(context.Background(), stream.)
		// log.Printf("CreateDumpMutualFund: %v", req)
	}
	return nil
}
