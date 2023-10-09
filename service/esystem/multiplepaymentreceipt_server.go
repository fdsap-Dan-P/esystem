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

func (server *DumpServer) CreateDumpMultiplePaymentReceipt(stream pb.DumpService_CreateDumpMultiplePaymentReceiptServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpMultiplePaymentReceipt:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		multiplePaymentReceipt := req.GetMultiplePaymentReceipt()

		log.Printf("received MultiplePaymentReceipt request: id = %s", multiplePaymentReceipt)

		multiplePaymentReceiptReq := model.MultiplePaymentReceipt{
			ModCtr:   multiplePaymentReceipt.ModCtr,
			BrCode:   multiplePaymentReceipt.BrCode,
			TrnDate:  multiplePaymentReceipt.TrnDate.AsTime(),
			OrNo:     multiplePaymentReceipt.OrNo,
			CID:      multiplePaymentReceipt.CID,
			PrNo:     multiplePaymentReceipt.PrNo,
			UserName: multiplePaymentReceipt.UserName,
			TermId:   multiplePaymentReceipt.TermId,
			AmtPaid:  util.Proto2Decimal(multiplePaymentReceipt.AmtPaid),
		}
		err = server.dump.CreateMultiplePaymentReceipt(context.Background(), multiplePaymentReceiptReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create MultiplePaymentReceipt: %v", err))
		}

		res := &pb.CreateDumpMultiplePaymentReceiptResponse{
			ModCtr: []int64{multiplePaymentReceipt.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateMultiplePaymentReceipt(context.Background(), stream.)
		// log.Printf("CreateDumpMultiplePaymentReceipt: %v", req)
	}
	return nil
}
