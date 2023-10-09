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

func (dumpClient *DumpClient) CreateDumpSaTrnMaster(saTrnMasterList []local.SaTrnMasterInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpSaTrnMaster(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for saTrnMaster: %v", err)
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
	for _, saTrnMaster := range saTrnMasterList {
		req := &pb.CreateDumpSaTrnMasterRequest{
			SaTrnMaster: &pb.DumpSaTrnMaster{
				ModCtr:      saTrnMaster.ModCtr,
				BrCode:      saTrnMaster.BrCode,
				ModAction:   saTrnMaster.ModAction,
				Acc:         saTrnMaster.Acc,
				TrnDate:     timestamppb.New(saTrnMaster.TrnDate),
				Trn:         saTrnMaster.Trn,
				TrnType:     util.NullInt642Proto(saTrnMaster.TrnType),
				OrNo:        util.NullInt642Proto(saTrnMaster.OrNo),
				TrnAmt:      util.NullDecimal2Proto(saTrnMaster.TrnAmt),
				RefNo:       util.NullString2Proto(saTrnMaster.RefNo),
				Particular:  saTrnMaster.Particular,
				TermId:      saTrnMaster.TermId,
				UserName:    saTrnMaster.UserName,
				PendApprove: saTrnMaster.PendApprove,
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send saTrnMaster stream request: %v - %v", err, stream.RecvMsg(nil))
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
