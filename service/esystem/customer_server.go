package service

import (
	"context"
	"io"

	pb "simplebank/pb"
	"simplebank/util"

	"log"

	model "simplebank/db/datastore/esystemlocal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *DumpServer) CreateDumpCustomer(stream pb.DumpService_CreateDumpCustomerServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCustomer:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		customer := req.GetCustomer()

		log.Printf("received Customer request: id = %s", customer)

		customerReq := model.Customer{
			ModCtr:            customer.ModCtr,
			BrCode:            customer.BrCode,
			ModAction:         customer.ModAction,
			CID:               customer.CID,
			CenterCode:        util.NullProto2String(customer.CenterCode),
			Title:             util.NullProto2Int64(customer.Title),
			LName:             util.NullProto2String(customer.LName),
			FName:             util.NullProto2String(customer.FName),
			MName:             util.NullProto2String(customer.MName),
			MaidenFName:       util.NullProto2String(customer.MaidenFName),
			MaidenLName:       util.NullProto2String(customer.MaidenLName),
			MaidenMName:       util.NullProto2String(customer.MaidenMName),
			Sex:               util.NullProto2String(customer.Sex),
			BirthDate:         util.NullProto2Time(customer.BirthDate),
			BirthPlace:        util.NullProto2String(customer.BirthPlace),
			CivilStatus:       util.NullProto2Int64(customer.CivilStatus),
			CustType:          util.NullProto2Int64(customer.CustType),
			Remarks:           util.NullProto2String(customer.Remarks),
			Status:            util.NullProto2Int64(customer.Status),
			Classification:    util.NullProto2Int64(customer.Classification),
			DepoType:          util.NullProto2String(customer.DepoType),
			SubClassification: util.NullProto2Int64(customer.SubClassification),
			PledgeAmount:      util.NullProto2Decimal(customer.PledgeAmount),
			MutualAmount:      util.NullProto2Decimal(customer.MutualAmount),
			PangarapAmount:    util.NullProto2Decimal(customer.PangarapAmount),
			KatuparanAmount:   util.NullProto2Decimal(customer.KatuparanAmount),
			InsuranceAmount:   util.NullProto2Decimal(customer.InsuranceAmount),
			AccPledge:         util.NullProto2Decimal(customer.AccPledge),
			AccMutual:         util.NullProto2Decimal(customer.AccMutual),
			AccPang:           util.NullProto2Decimal(customer.AccPang),
			AccKatuparan:      util.NullProto2Decimal(customer.AccKatuparan),
			AccInsurance:      util.NullProto2Decimal(customer.AccInsurance),
			LoanLimit:         util.NullProto2Decimal(customer.LoanLimit),
			CreditLimit:       util.NullProto2Decimal(customer.CreditLimit),
			DateRecognized:    util.NullProto2Time(customer.DateRecognized),
			DateResigned:      util.NullProto2Time(customer.DateResigned),
			DateEntry:         util.NullProto2Time(customer.DateEntry),
			GoldenLifeDate:    util.NullProto2Time(customer.GoldenLifeDate),
			Restricted:        util.NullProto2String(customer.Restricted),
			Borrower:          util.NullProto2String(customer.Borrower),
			CoMaker:           util.NullProto2String(customer.CoMaker),
			Guarantor:         util.NullProto2String(customer.Guarantor),
			DOSRI:             util.NullProto2Int64(customer.DOSRI),
			IDCode1:           util.NullProto2Int64(customer.IDCode1),
			IDNum1:            util.NullProto2String(customer.IDNum1),
			IDCode2:           util.NullProto2Int64(customer.IDCode2),
			IDNum2:            util.NullProto2String(customer.IDNum2),
			Contact1:          util.NullProto2String(customer.Contact1),
			Contact2:          util.NullProto2String(customer.Contact2),
			Phone1:            util.NullProto2String(customer.Phone1),
			Reffered1:         util.NullProto2String(customer.Reffered1),
			Reffered2:         util.NullProto2String(customer.Reffered2),
			Reffered3:         util.NullProto2String(customer.Reffered3),
			Education:         util.NullProto2Int64(customer.Education),
			Validity1:         util.NullProto2Time(customer.Validity1),
			Validity2:         util.NullProto2Time(customer.Validity2),
			BusinessType:      util.NullProto2Int64(customer.BusinessType),
			AccountNumber:     util.NullProto2String(customer.AccountNumber),
			IIID:              util.NullProto2Int64(customer.IIID),
			Religion:          util.NullProto2Int64(customer.Religion),
		}
		err = server.dump.CreateCustomer(context.Background(), customerReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Customer: %v", err))
		}

		res := &pb.CreateDumpCustomerResponse{
			ModCtr: []int64{customer.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCustomer(context.Background(), stream.)
		// log.Printf("CreateDumpCustomer: %v", req)
	}
	return nil
}
