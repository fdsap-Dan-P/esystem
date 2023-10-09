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

func (dumpClient *DumpClient) CreateDumpCenterWorker(centerWorkerList []local.CenterWorkerInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCenterWorker(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for centerWorker: %v", err)
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
	for _, centerWorker := range centerWorkerList {
		req := &pb.CreateDumpCenterWorkerRequest{
			CenterWorker: &pb.DumpCenterWorker{
				ModCtr:      centerWorker.ModCtr,
				BrCode:      centerWorker.BrCode,
				ModAction:   centerWorker.ModAction,
				AOID:        centerWorker.AOID,
				Lname:       util.NullString2Proto(centerWorker.Lname),
				FName:       util.NullString2Proto(centerWorker.FName),
				Mname:       util.NullString2Proto(centerWorker.Mname),
				PhoneNumber: util.NullString2Proto(centerWorker.PhoneNumber),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send centerWorker stream request: %v - %v", err, stream.RecvMsg(nil))
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
