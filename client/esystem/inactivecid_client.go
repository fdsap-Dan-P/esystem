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

func (dumpClient *DumpClient) CreateDumpInActiveCID(inActiveCIDList []local.InActiveCID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpInActiveCID(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for inActiveCID: %v", err)
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
	for _, inActiveCID := range inActiveCIDList {
		req := &pb.CreateDumpInActiveCIDRequest{
			InActiveCID: &pb.DumpInActiveCID{
				ModCtr:        inActiveCID.ModCtr,
				BrCode:        inActiveCID.BrCode,
				CID:           inActiveCID.CID,
				InActive:      inActiveCID.InActive,
				DateStart:     timestamppb.New(inActiveCID.DateStart),
				DateEnd:       util.NullTime2Proto(inActiveCID.DateEnd),
				UserId:        inActiveCID.UserId,
				DeactivatedBy: util.NullString2Proto(inActiveCID.DeactivatedBy),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send inActiveCID stream request: %v - %v", err, stream.RecvMsg(nil))
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
