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

func (dumpClient *DumpClient) CreateDumpAddresses(addressesList []local.AddressesInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpAddresses(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for addresses: %v", err)
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

			log.Printf("received response: %v", res)
		}
	}()

	// send requests
	for _, addresses := range addressesList {
		req := &pb.CreateDumpAddressesRequest{
			Addresses: &pb.DumpAddresses{
				ModCtr:         addresses.ModCtr,
				BrCode:         addresses.BrCode,
				ModAction:      addresses.ModAction,
				CID:            addresses.CID,
				SeqNum:         addresses.SeqNum,
				AddressDetails: util.NullString2Proto(addresses.AddressDetails),
				Barangay:       util.NullString2Proto(addresses.Barangay),
				City:           util.NullString2Proto(addresses.City),
				Province:       util.NullString2Proto(addresses.Province),
				Phone1:         util.NullString2Proto(addresses.Phone1),
				Phone2:         util.NullString2Proto(addresses.Phone2),
				Phone3:         util.NullString2Proto(addresses.Phone3),
				Phone4:         util.NullString2Proto(addresses.Phone4),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send addresses stream request: %v - %v", err, stream.RecvMsg(nil))
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
