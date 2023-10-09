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

func (dumpClient *DumpClient) CreateDumpWriteoff(writeoffList []local.WriteoffInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpWriteoff(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for writeoff: %v", err)
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
	for _, writeoff := range writeoffList {
		req := &pb.CreateDumpWriteoffRequest{
			Writeoff: &pb.DumpWriteoff{
				ModCtr:     writeoff.ModCtr,
				BrCode:     writeoff.BrCode,
				ModAction:  writeoff.ModAction,
				Acc:        writeoff.Acc,
				DisbDate:   timestamppb.New(writeoff.DisbDate),
				Principal:  util.Decimal2Proto(writeoff.Principal),
				Interest:   util.Decimal2Proto(writeoff.Interest),
				BalPrin:    util.Decimal2Proto(writeoff.BalPrin),
				BalInt:     util.Decimal2Proto(writeoff.BalInt),
				TrnDate:    timestamppb.New(writeoff.TrnDate),
				AcctType:   writeoff.AcctType,
				Print:      util.NullString2Proto(writeoff.Print),
				PostedBy:   util.NullString2Proto(writeoff.PostedBy),
				VerifiedBy: util.NullString2Proto(writeoff.VerifiedBy),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send writeoff stream request: %v - %v", err, stream.RecvMsg(nil))
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
