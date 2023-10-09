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

func (dumpClient *DumpClient) CreateDumpTrnMaster(trnMasterList []local.TrnMasterInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpTrnMaster(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for trnMaster: %v", err)
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
	for _, trnMaster := range trnMasterList {
		req := &pb.CreateDumpTrnMasterRequest{
			TrnMaster: &pb.DumpTrnMaster{
				ModCtr:     trnMaster.ModCtr,
				BrCode:     trnMaster.BrCode,
				ModAction:  trnMaster.ModAction,
				Acc:        trnMaster.Acc,
				TrnDate:    timestamppb.New(trnMaster.TrnDate),
				Trn:        trnMaster.Trn,
				TrnType:    util.NullInt642Proto(trnMaster.TrnType),
				OrNo:       util.NullInt642Proto(trnMaster.OrNo),
				Prin:       util.Decimal2Proto(trnMaster.Prin),
				IntR:       util.Decimal2Proto(trnMaster.IntR),
				WaivedInt:  util.Decimal2Proto(trnMaster.WaivedInt),
				RefNo:      util.NullString2Proto(trnMaster.RefNo),
				UserName:   util.NullString2Proto(trnMaster.UserName),
				Particular: util.NullString2Proto(trnMaster.Particular),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send trnMaster stream request: %v - %v", err, stream.RecvMsg(nil))
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
