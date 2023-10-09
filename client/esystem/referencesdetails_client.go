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

func (dumpClient *DumpClient) CreateDumpReferencesDetails(referencesDetailsList []local.ReferencesDetailsInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpReferencesDetails(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for referencesDetails: %v", err)
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
	for _, referencesDetails := range referencesDetailsList {
		req := &pb.CreateDumpReferencesDetailsRequest{
			ReferencesDetails: &pb.DumpReferencesDetails{
				ModCtr:             referencesDetails.ModCtr,
				BrCode:             referencesDetails.BrCode,
				ModAction:          referencesDetails.ModAction,
				ID:                 referencesDetails.ID,
				RefID:              referencesDetails.RefID,
				PurposeDescription: util.NullString2Proto(referencesDetails.PurposeDescription),
				ParentID:           util.NullInt642Proto(referencesDetails.ParentID),
				CodeID:             util.NullInt642Proto(referencesDetails.CodeID),
				Stat:               util.NullInt642Proto(referencesDetails.Stat),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send referencesDetails stream request: %v - %v", err, stream.RecvMsg(nil))
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
