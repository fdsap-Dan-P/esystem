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
)

func (dumpClient *DumpClient) CreateDumpJnlDetails(jnlDetailsList []local.JnlDetailsInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpJnlDetails(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for jnlDetails: %v", err)
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
	for _, jnlDetails := range jnlDetailsList {
		req := &pb.CreateDumpJnlDetailsRequest{
			JnlDetails: &pb.DumpJnlDetails{
				ModCtr:    jnlDetails.ModCtr,
				BrCode:    jnlDetails.BrCode,
				ModAction: jnlDetails.ModAction,
				Acc:       jnlDetails.Acc,
				Trn:       jnlDetails.Trn,
				Series:    util.NullInt642Proto(jnlDetails.Series),
				Debit:     util.NullDecimal2Proto(jnlDetails.Debit),
				Credit:    util.NullDecimal2Proto(jnlDetails.Credit),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send jnlDetails stream request: %v - %v", err, stream.RecvMsg(nil))
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
