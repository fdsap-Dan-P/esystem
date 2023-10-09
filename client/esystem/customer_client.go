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

func (dumpClient *DumpClient) CreateDumpCustomer(customerList []local.CustomerInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpCustomer(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for customer: %v", err)
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
	for _, customer := range customerList {
		req := &pb.CreateDumpCustomerRequest{
			Customer: &pb.DumpCustomer{
				ModCtr:            customer.ModCtr,
				BrCode:            customer.BrCode,
				ModAction:         customer.ModAction,
				CID:               customer.CID,
				CenterCode:        util.NullString2Proto(customer.CenterCode),
				Title:             util.NullInt642Proto(customer.Title),
				LName:             util.NullString2Proto(customer.LName),
				FName:             util.NullString2Proto(customer.FName),
				MName:             util.NullString2Proto(customer.MName),
				MaidenFName:       util.NullString2Proto(customer.MaidenFName),
				MaidenLName:       util.NullString2Proto(customer.MaidenLName),
				MaidenMName:       util.NullString2Proto(customer.MaidenMName),
				Sex:               util.NullString2Proto(customer.Sex),
				BirthDate:         util.NullTime2Proto(customer.BirthDate),
				BirthPlace:        util.NullString2Proto(customer.BirthPlace),
				CivilStatus:       util.NullInt642Proto(customer.CivilStatus),
				CustType:          util.NullInt642Proto(customer.CustType),
				Remarks:           util.NullString2Proto(customer.Remarks),
				Status:            util.NullInt642Proto(customer.Status),
				Classification:    util.NullInt642Proto(customer.Classification),
				DepoType:          util.NullString2Proto(customer.DepoType),
				SubClassification: util.NullInt642Proto(customer.SubClassification),
				PledgeAmount:      util.NullDecimal2Proto(customer.PledgeAmount),
				MutualAmount:      util.NullDecimal2Proto(customer.MutualAmount),
				PangarapAmount:    util.NullDecimal2Proto(customer.PangarapAmount),
				KatuparanAmount:   util.NullDecimal2Proto(customer.KatuparanAmount),
				InsuranceAmount:   util.NullDecimal2Proto(customer.InsuranceAmount),
				AccPledge:         util.NullDecimal2Proto(customer.AccPledge),
				AccMutual:         util.NullDecimal2Proto(customer.AccMutual),
				AccPang:           util.NullDecimal2Proto(customer.AccPang),
				AccKatuparan:      util.NullDecimal2Proto(customer.AccKatuparan),
				AccInsurance:      util.NullDecimal2Proto(customer.AccInsurance),
				LoanLimit:         util.NullDecimal2Proto(customer.LoanLimit),
				CreditLimit:       util.NullDecimal2Proto(customer.CreditLimit),
				DateRecognized:    util.NullTime2Proto(customer.DateRecognized),
				DateResigned:      util.NullTime2Proto(customer.DateResigned),
				DateEntry:         util.NullTime2Proto(customer.DateEntry),
				GoldenLifeDate:    util.NullTime2Proto(customer.GoldenLifeDate),
				Restricted:        util.NullString2Proto(customer.Restricted),
				Borrower:          util.NullString2Proto(customer.Borrower),
				CoMaker:           util.NullString2Proto(customer.CoMaker),
				Guarantor:         util.NullString2Proto(customer.Guarantor),
				DOSRI:             util.NullInt642Proto(customer.DOSRI),
				IDCode1:           util.NullInt642Proto(customer.IDCode1),
				IDNum1:            util.NullString2Proto(customer.IDNum1),
				IDCode2:           util.NullInt642Proto(customer.IDCode2),
				IDNum2:            util.NullString2Proto(customer.IDNum2),
				Contact1:          util.NullString2Proto(customer.Contact1),
				Contact2:          util.NullString2Proto(customer.Contact2),
				Phone1:            util.NullString2Proto(customer.Phone1),
				Reffered1:         util.NullString2Proto(customer.Reffered1),
				Reffered2:         util.NullString2Proto(customer.Reffered2),
				Reffered3:         util.NullString2Proto(customer.Reffered3),
				Education:         util.NullInt642Proto(customer.Education),
				Validity1:         util.NullTime2Proto(customer.Validity1),
				Validity2:         util.NullTime2Proto(customer.Validity2),
				BusinessType:      util.NullInt642Proto(customer.BusinessType),
				AccountNumber:     util.NullString2Proto(customer.AccountNumber),
				IIID:              util.NullInt642Proto(customer.IIID),
				Religion:          util.NullInt642Proto(customer.Religion),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send customer stream request: %v - %v", err, stream.RecvMsg(nil))
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
