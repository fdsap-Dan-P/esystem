package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	local "simplebank/db/datastore/esystemlocal"
	pb "simplebank/pb"

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func (dumpClient *DumpClient) CreateDumpCustAddInfo(custAddInfoList []local.CustAddInfoInfo) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ctx, cancel, reset := WithTimeoutReset(context.Background(), time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCustAddInfo(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for custAddInfo: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			reset()
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

			// var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")

			// ctx, err = context.WithTimeout(ctx, time.Duration(*deadlineMs)*time.Millisecond)

			// if err != nil {
			// 	log.Print("no more responses")
			// 	waitResponse <- fmt.Errorf("Deadline related issue: %v", err)
			// }

			QueriesLocal.UpdateModifiedTableUploaded(context.Background(), res.ModCtr, true)
			log.Printf("received response: %v", res)
		}
	}()

	// send requests
	for _, custAddInfo := range custAddInfoList {
		req := &pb.CreateDumpCustAddInfoRequest{
			CustAddInfo: &pb.DumpCustAddInfo{
				ModCtr:    custAddInfo.ModCtr,
				BrCode:    custAddInfo.BrCode,
				ModAction: custAddInfo.ModAction,
				CID:       custAddInfo.CID,
				InfoDate:  timestamp.New(custAddInfo.InfoDate),
				InfoCode:  custAddInfo.InfoCode,
				Info:      custAddInfo.Info,
				InfoValue: custAddInfo.InfoValue,
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send custAddInfo stream request: %v - %v", err, stream.RecvMsg(nil))
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
