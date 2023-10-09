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

func (dumpClient *DumpClient) CreateDumpSaMaster(saMasterList []local.SaMasterInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpSaMaster(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for saMaster: %v", err)
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
	for _, saMaster := range saMasterList {
		req := &pb.CreateDumpSaMasterRequest{
			SaMaster: &pb.DumpSaMaster{
				ModCtr:     saMaster.ModCtr,
				BrCode:     saMaster.BrCode,
				ModAction:  saMaster.ModAction,
				Acc:        saMaster.Acc,
				CID:        saMaster.CID,
				Type:       saMaster.Type,
				Balance:    util.NullDecimal2Proto(saMaster.Balance),
				DoLastTrn:  util.NullTime2Proto(saMaster.DoLastTrn),
				DoStatus:   util.NullTime2Proto(saMaster.DoStatus),
				Dopen:      util.NullTime2Proto(saMaster.Dopen),
				DoMaturity: util.NullTime2Proto(saMaster.DoMaturity),
				Status:     util.NullString2Proto(saMaster.Status),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send saMaster stream request: %v - %v", err, stream.RecvMsg(nil))
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
