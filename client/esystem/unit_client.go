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

func (dumpClient *DumpClient) CreateDumpUnit(unitList []local.UnitInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpUnit(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for unit: %v", err)
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
	for _, unit := range unitList {
		req := &pb.CreateDumpUnitRequest{
			Unit: &pb.DumpUnit{
				ModCtr:      unit.ModCtr,
				BrCode:      unit.BrCode,
				ModAction:   unit.ModAction,
				UnitCode:    unit.UnitCode,
				Unit:        util.NullString2Proto(unit.Unit),
				AreaCode:    util.NullInt642Proto(unit.AreaCode),
				LName:       util.NullString2Proto(unit.LName),
				FName:       util.NullString2Proto(unit.FName),
				MName:       util.NullString2Proto(unit.MName),
				VatReg:      util.NullString2Proto(unit.VatReg),
				UnitAddress: util.NullString2Proto(unit.UnitAddress),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send unit stream request: %v - %v", err, stream.RecvMsg(nil))
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
