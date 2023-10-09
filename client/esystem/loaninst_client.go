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

func (dumpClient *DumpClient) CreateDumpLoanInst(loanInstList []local.LoanInstInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpLoanInst(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for loanInst: %v", err)
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
	for _, loanInst := range loanInstList {
		req := &pb.CreateDumpLoanInstRequest{
			LoanInst: &pb.DumpLoanInst{
				ModCtr:    loanInst.ModCtr,
				BrCode:    loanInst.BrCode,
				ModAction: loanInst.ModAction,
				Acc:       loanInst.Acc,
				Dnum:      loanInst.Dnum,
				DueDate:   timestamppb.New(loanInst.DueDate),
				InstFlag:  loanInst.InstFlag,
				DuePrin:   util.Decimal2Proto(loanInst.DuePrin),
				DueInt:    util.Decimal2Proto(loanInst.DueInt),
				UpInt:     util.Decimal2Proto(loanInst.UpInt),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send loanInst stream request: %v - %v", err, stream.RecvMsg(nil))
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
