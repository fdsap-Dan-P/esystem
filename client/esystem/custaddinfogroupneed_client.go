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

func (dumpClient *DumpClient) CreateDumpCustAddInfoGroupNeed(custAddInfoGroupNeedList []local.CustAddInfoGroupNeedInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCustAddInfoGroupNeed(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for custAddInfoGroupNeed: %v", err)
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
	for _, custAddInfoGroupNeed := range custAddInfoGroupNeedList {
		req := &pb.CreateDumpCustAddInfoGroupNeedRequest{
			CustAddInfoGroupNeed: &pb.DumpCustAddInfoGroupNeed{
				ModCtr:      custAddInfoGroupNeed.ModCtr,
				BrCode:      custAddInfoGroupNeed.BrCode,
				ModAction:   custAddInfoGroupNeed.ModAction,
				InfoGroup:   custAddInfoGroupNeed.InfoGroup,
				InfoCode:    custAddInfoGroupNeed.InfoCode,
				InfoProcess: util.NullString2Proto(custAddInfoGroupNeed.InfoProcess),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send custAddInfoGroupNeed stream request: %v - %v", err, stream.RecvMsg(nil))
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
