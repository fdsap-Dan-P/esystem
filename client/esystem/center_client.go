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

func (dumpClient *DumpClient) CreateDumpCenter(centerList []local.CenterInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCenter(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for center: %v", err)
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
	for _, center := range centerList {
		req := &pb.CreateDumpCenterRequest{
			Center: &pb.DumpCenter{
				ModCtr:          center.ModCtr,
				BrCode:          center.BrCode,
				ModAction:       center.ModAction,
				CenterCode:      center.CenterCode,
				CenterName:      util.NullString2Proto(center.CenterName),
				CenterAddress:   util.NullString2Proto(center.CenterAddress),
				MeetingDay:      util.NullInt642Proto(center.MeetingDay),
				Unit:            util.NullInt642Proto(center.Unit),
				DateEstablished: util.NullTime2Proto(center.DateEstablished),
				AOID:            util.NullInt642Proto(center.AOID),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send center stream request: %v - %v", err, stream.RecvMsg(nil))
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
