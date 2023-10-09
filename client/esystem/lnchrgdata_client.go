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

func (dumpClient *DumpClient) CreateDumpLnChrgData(lnChrgDataList []local.LnChrgDataInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpLnChrgData(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for lnChrgData: %v", err)
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
	for _, lnChrgData := range lnChrgDataList {
		req := &pb.CreateDumpLnChrgDataRequest{
			LnChrgData: &pb.DumpLnChrgData{
				ModCtr:    lnChrgData.ModCtr,
				BrCode:    lnChrgData.BrCode,
				ModAction: lnChrgData.ModAction,
				Acc:       lnChrgData.Acc,
				ChrgCode:  lnChrgData.ChrgCode,
				RefAcc:    util.NullString2Proto(lnChrgData.RefAcc),
				ChrAmnt:   util.Decimal2Proto(lnChrgData.ChrAmnt),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send lnChrgData stream request: %v - %v", err, stream.RecvMsg(nil))
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
