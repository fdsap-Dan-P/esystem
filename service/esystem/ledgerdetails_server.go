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

func (server *DumpServer) CreateDumpLedgerDetails(stream pb.DumpService_CreateDumpLedgerDetailsServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpLedgerDetails:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		ledgerDetails := req.GetLedgerDetails()

		log.Printf("received LedgerDetails request: id = %s", ledgerDetails)

		ledgerDetailsReq := model.LedgerDetails{
			ModCtr:    ledgerDetails.ModCtr,
			BrCode:    ledgerDetails.BrCode,
			ModAction: ledgerDetails.ModAction,
			TrnDate:   ledgerDetails.TrnDate.AsTime(),
			Acc:       ledgerDetails.Acc,
			Balance:   util.Proto2Decimal(ledgerDetails.Balance),
		}
		err = server.dump.CreateLedgerDetails(context.Background(), ledgerDetailsReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create LedgerDetails: %v", err))
		}

		res := &pb.CreateDumpLedgerDetailsResponse{
			ModCtr: []int64{ledgerDetails.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateLedgerDetails(context.Background(), stream.)
		// log.Printf("CreateDumpLedgerDetails: %v", req)
	}
	return nil
}
