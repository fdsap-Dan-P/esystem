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

func (dumpClient *DumpClient) CreateDumpCustAddInfoGroup(custAddInfoGroupList []local.CustAddInfoGroupInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCustAddInfoGroup(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for custAddInfoGroup: %v", err)
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
	for _, custAddInfoGroup := range custAddInfoGroupList {
		req := &pb.CreateDumpCustAddInfoGroupRequest{
			CustAddInfoGroup: &pb.DumpCustAddInfoGroup{
				ModCtr:     custAddInfoGroup.ModCtr,
				BrCode:     custAddInfoGroup.BrCode,
				ModAction:  custAddInfoGroup.ModAction,
				InfoGroup:  custAddInfoGroup.InfoGroup,
				GroupTitle: util.NullString2Proto(custAddInfoGroup.GroupTitle),
				Remarks:    util.NullString2Proto(custAddInfoGroup.Remarks),
				ReqOnEntry: util.NullBool2Proto(custAddInfoGroup.ReqOnEntry),
				ReqOnExit:  util.NullBool2Proto(custAddInfoGroup.ReqOnExit),
				Link2Loan:  util.NullInt642Proto(custAddInfoGroup.Link2Loan),
				Link2Save:  util.NullInt642Proto(custAddInfoGroup.Link2Save),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send custAddInfoGroup stream request: %v - %v", err, stream.RecvMsg(nil))
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
