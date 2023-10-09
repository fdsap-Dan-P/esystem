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

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func (dumpClient *DumpClient) CreateDumpColSht(colShtList []local.ColSht) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ctx, cancel, reset := WithTimeoutReset(context.Background(), time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpColSht(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for colSht: %v", err)
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
			log.Printf("received response: %v", res)
		}
	}()

	// send requests
	for _, colSht := range colShtList {
		req := &pb.CreateDumpColShtRequest{
			ColSht: &pb.DumpColSht{
				BrCode:          colSht.BrCode,
				AppType:         colSht.AppType,
				Code:            colSht.Code,
				Status:          colSht.Status,
				Acc:             colSht.Acc,
				CID:             colSht.CID,
				UM:              colSht.UM,
				ClientName:      colSht.ClientName,
				CenterCode:      colSht.CenterCode,
				CenterName:      colSht.CenterName,
				ManCode:         colSht.ManCode,
				Unit:            colSht.Unit,
				AreaCode:        colSht.AreaCode,
				Area:            colSht.Area,
				StaffName:       colSht.StaffName,
				AcctType:        colSht.AcctType,
				AcctDesc:        colSht.AcctDesc,
				DisbDate:        timestamp.New(colSht.DisbDate),
				DateStart:       timestamp.New(colSht.DateStart),
				Maturity:        timestamp.New(colSht.Maturity),
				Principal:       util.Decimal2Proto(colSht.Principal),
				Interest:        util.Decimal2Proto(colSht.Interest),
				Gives:           colSht.Gives,
				BalPrin:         util.Decimal2Proto(colSht.BalPrin),
				BalInt:          util.Decimal2Proto(colSht.BalInt),
				Amort:           util.Decimal2Proto(colSht.Amort),
				DuePrin:         util.Decimal2Proto(colSht.DuePrin),
				DueInt:          util.Decimal2Proto(colSht.DueInt),
				LoanBal:         util.Decimal2Proto(colSht.LoanBal),
				SaveBal:         util.Decimal2Proto(colSht.SaveBal),
				WaivedInt:       util.Decimal2Proto(colSht.WaivedInt),
				UnPaidCtr:       colSht.UnPaidCtr,
				WrittenOff:      colSht.WrittenOff,
				OrgName:         colSht.OrgName,
				OrgAddress:      colSht.OrgAddress,
				MeetingDate:     timestamp.New(colSht.MeetingDate),
				MeetingDay:      colSht.MeetingDay,
				SharesOfStock:   util.Decimal2Proto(colSht.SharesOfStock),
				DateEstablished: timestamp.New(colSht.DateEstablished),
				Classification:  colSht.Classification,
				WriteOff:        colSht.WriteOff,
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send colSht stream request: %v - %v", err, stream.RecvMsg(nil))
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
