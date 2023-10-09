package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	local "simplebank/db/datastore/esystemlocal"
	pb "simplebank/pb"
)

func (dumpClient *DumpClient) CreateDumpCustAddInfoList(custAddInfoListList []local.CustAddInfoListInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCustAddInfoList(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for custAddInfoList: %v", err)
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
	for _, custAddInfoList := range custAddInfoListList {
		req := &pb.CreateDumpCustAddInfoListRequest{
			CustAddInfoList: &pb.DumpCustAddInfoList{
				ModCtr:     custAddInfoList.ModCtr,
				BrCode:     custAddInfoList.BrCode,
				ModAction:  custAddInfoList.ModAction,
				InfoCode:   custAddInfoList.InfoCode,
				InfoOrder:  custAddInfoList.InfoOrder,
				Title:      custAddInfoList.Title,
				InfoType:   custAddInfoList.InfoType,
				InfoLen:    custAddInfoList.InfoLen,
				InfoFormat: custAddInfoList.InfoFormat,
				InputType:  custAddInfoList.InputType,
				InfoSource: custAddInfoList.InfoSource,
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send custAddInfoList stream request: %v - %v", err, stream.RecvMsg(nil))
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
