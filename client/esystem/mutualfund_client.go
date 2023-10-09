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

func (dumpClient *DumpClient) CreateDumpMutualFund(mutualFundList []local.MutualFundInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpMutualFund(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for mutualFund: %v", err)
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
	for _, mutualFund := range mutualFundList {
		req := &pb.CreateDumpMutualFundRequest{
			MutualFund: &pb.DumpMutualFund{
				ModCtr:    mutualFund.ModCtr,
				BrCode:    mutualFund.BrCode,
				ModAction: mutualFund.ModAction,
				CID:       mutualFund.CID,
				OrNo:      util.NullInt642Proto(mutualFund.OrNo),
				TrnDate:   timestamppb.New(mutualFund.TrnDate),
				TrnType:   util.NullString2Proto(mutualFund.TrnType),
				TrnAmt:    util.Decimal2Proto(mutualFund.TrnAmt),
				UserName:  util.NullString2Proto(mutualFund.UserName),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send mutualFund stream request: %v - %v", err, stream.RecvMsg(nil))
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
