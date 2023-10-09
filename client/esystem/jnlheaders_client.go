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

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func (dumpClient *DumpClient) CreateDumpJnlHeaders(jnlHeadersList []local.JnlHeadersInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpJnlHeaders(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for jnlHeaders: %v", err)
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
	for _, jnlHeaders := range jnlHeadersList {
		req := &pb.CreateDumpJnlHeadersRequest{
			JnlHeaders: &pb.DumpJnlHeaders{
				ModCtr:      jnlHeaders.ModCtr,
				BrCode:      jnlHeaders.BrCode,
				ModAction:   jnlHeaders.ModAction,
				Trn:         jnlHeaders.Trn,
				TrnDate:     timestamp.New(jnlHeaders.TrnDate),
				Particulars: jnlHeaders.Particulars,
				UserName:    util.NullString2Proto(jnlHeaders.UserName),
				Code:        jnlHeaders.Code,
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send jnlHeaders stream request: %v - %v", err, stream.RecvMsg(nil))
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
