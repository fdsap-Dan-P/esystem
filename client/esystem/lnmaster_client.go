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

func (dumpClient *DumpClient) CreateDumpLnMaster(lnMasterList []local.LnMasterInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpLnMaster(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for lnMaster: %v", err)
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
	for _, lnMaster := range lnMasterList {
		req := &pb.CreateDumpLnMasterRequest{
			LnMaster: &pb.DumpLnMaster{
				ModCtr:      lnMaster.ModCtr,
				BrCode:      lnMaster.BrCode,
				ModAction:   lnMaster.ModAction,
				CID:         lnMaster.CID,
				Acc:         lnMaster.Acc,
				AcctType:    util.NullInt642Proto(lnMaster.AcctType),
				DisbDate:    util.NullTime2Proto(lnMaster.DisbDate),
				Principal:   util.NullDecimal2Proto(lnMaster.Principal),
				Interest:    util.NullDecimal2Proto(lnMaster.Interest),
				NetProceed:  util.NullDecimal2Proto(lnMaster.NetProceed),
				Gives:       util.NullInt642Proto(lnMaster.Gives),
				Frequency:   util.NullInt642Proto(lnMaster.Frequency),
				AnnumDiv:    util.NullInt642Proto(lnMaster.AnnumDiv),
				Prin:        util.NullDecimal2Proto(lnMaster.Prin),
				IntR:        util.NullDecimal2Proto(lnMaster.IntR),
				WaivedInt:   util.NullDecimal2Proto(lnMaster.WaivedInt),
				WeeksPaid:   util.NullInt642Proto(lnMaster.WeeksPaid),
				DoMaturity:  util.NullTime2Proto(lnMaster.DoMaturity),
				ConIntRate:  util.NullDecimal2Proto(lnMaster.ConIntRate),
				Status:      util.NullString2Proto(lnMaster.Status),
				Cycle:       util.NullInt642Proto(lnMaster.Cycle),
				LNGrpCode:   util.NullInt642Proto(lnMaster.LNGrpCode),
				Proff:       util.NullInt642Proto(lnMaster.Proff),
				FundSource:  util.NullString2Proto(lnMaster.FundSource),
				DOSRI:       util.NullBool2Proto(lnMaster.DOSRI),
				LnCategory:  util.NullInt642Proto(lnMaster.LnCategory),
				OpenDate:    util.NullTime2Proto(lnMaster.OpenDate),
				LastTrnDate: util.NullTime2Proto(lnMaster.LastTrnDate),
				DisbBy:      util.NullString2Proto(lnMaster.DisbBy),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send lnMaster stream request: %v - %v", err, stream.RecvMsg(nil))
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
