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

func (dumpClient *DumpClient) CreateDumpArea(areaList []local.AreaInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpArea(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for area: %v", err)
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
	for _, area := range areaList {
		req := &pb.CreateDumpAreaRequest{
			Area: &pb.DumpArea{
				ModCtr:    area.ModCtr,
				BrCode:    area.BrCode,
				ModAction: area.ModAction,
				AreaCode:  area.AreaCode,
				Area:      util.NullString2Proto(area.Area),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send area stream request: %v - %v", err, stream.RecvMsg(nil))
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
