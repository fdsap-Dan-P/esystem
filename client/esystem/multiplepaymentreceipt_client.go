package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	local "simplebank/db/datastore/esystemlocal"
	pb "simplebank/pb"
	"simplebank/util"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (dumpClient *DumpClient) CreateDumpMultiplePaymentReceipt(multiplePaymentReceiptList []local.MultiplePaymentReceipt) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpMultiplePaymentReceipt(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for multiplePaymentReceipt: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			QueriesLocal.UpdateModifiedTableUploaded(context.Background(), res.ModCtr, true)
			log.Printf("received response: %v", res)
		}
	}()

	// send requests
	for _, multiplePaymentReceipt := range multiplePaymentReceiptList {
		req := &pb.CreateDumpMultiplePaymentReceiptRequest{
			MultiplePaymentReceipt: &pb.DumpMultiplePaymentReceipt{
				ModCtr:   multiplePaymentReceipt.ModCtr,
				BrCode:   multiplePaymentReceipt.BrCode,
				TrnDate:  timestamppb.New(multiplePaymentReceipt.TrnDate),
				OrNo:     multiplePaymentReceipt.OrNo,
				CID:      multiplePaymentReceipt.CID,
				PrNo:     multiplePaymentReceipt.PrNo,
				UserName: multiplePaymentReceipt.UserName,
				TermId:   multiplePaymentReceipt.TermId,
				AmtPaid:  util.Decimal2Proto(multiplePaymentReceipt.AmtPaid),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send multiplePaymentReceipt stream request: %v - %v", err, stream.RecvMsg(nil))
		}
		log.Print("sent request: ", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}
