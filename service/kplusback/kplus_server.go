package service

import (
	"context"

	pb "simplebank/pb"
	"simplebank/util"

	"log"

	dskPlus "simplebank/db/datastore/kplus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *KPlusServer) SearchCustomerCID(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSCustomerResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   util.ToInt64(cus.INAIIID), // map INAIIID to Cid
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) CustSavingsList(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSCustSavingsListResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SavingsList(context.Background(), dskPlus.SavingsListParams{INAIIID: cid})
	if err != nil {
		return &pb.KPLUSCustSavingsListResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to CustSavingsList: %v", err))
	}

	res := []*pb.KPLUSCustSavingsList{}
	for _, d := range cus {
		r := &pb.KPLUSCustSavingsList{
			Cid:        util.ToInt64(d.INAIIID), // map INAIIID to Cid
			Acc:        d.Acc,
			AcctType:   d.AcctType,
			AccDesc:    d.AccDesc,
			Dopen:      d.Dopen,
			StatusDesc: d.StatusDesc,
			Balance:    d.Balance,
			Status:     d.Status,
		}
		res = append(res, r)
	}
	return &pb.KPLUSCustSavingsListResponse{CustSavingsList: res}, nil
}

func (server *KPlusServer) GetTransactionHistory(ctx context.Context, r *pb.KPLUSGetTransactionHistoryRequest) (*pb.KPLUSGetTransactionHistoryResponse, error) {

	acc := r.GetAcc()
	dateFrom := util.SetDate(r.GetDateFrom())
	dateTo := util.SetDate(r.GetDateTo())

	log.Printf("received GetTransactionHistory request: acc:%v, from:%v to:%v ", acc, dateFrom, dateTo)

	data, err := server.kplus.Transaction(context.Background(),
		dskPlus.TransactionParams{Acc: acc, DateFrom: dateFrom, DateTo: dateTo})
	if err != nil {
		return &pb.KPLUSGetTransactionHistoryResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to GetTransactionHistory: %v", err))
	}

	res := []*pb.KPLUSTransaction{}
	for _, d := range data {
		r := &pb.KPLUSTransaction{
			Acc:         d.Acc,
			Trn:         d.Trn,
			Prin:        d.Prin,
			Intr:        d.Intr,
			TrnAmount:   d.TrnAmount,
			Balance:     d.Balance,
			Particulars: d.Particulars,
			Username:    d.Username,
		}
		res = append(res, r)
	}
	return &pb.KPLUSGetTransactionHistoryResponse{Transaction: res}, nil
}

/*

func (server *KPlusServer) K2CCallBackRef(ctx context.Context, r *pb.KPLUSCallBackRefRequest) (*pb.KPLUSResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) GetReferences(ctx context.Context, r *pb.KPLUSGetReferencesRequest) (*pb.KPLUSGetReferencesResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) SearchLoanList(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSSearchLoanListResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) LoanInfo(ctx context.Context, r *pb.KPLUSAccRequest) (*pb.KPLUSLoanInfoResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) GetSavingForSuperApp(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSGetSavingResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}

func (server *KPlusServer) FundTransferRequest(ctx context.Context, r *pb.KPLUSFundTransferRequest) (*pb.KPLUSResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SearchCustomerCID(context.Background(), cid)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to SearchCustomerCID: %v", err))
	}

	res := &pb.KPLUSCustomerResponse{
		Cid:                   cus.Cid,
		LastName:              cus.LastName,
		FirstName:             cus.FirstName,
		MiddleName:            cus.MiddleName,
		MaidenFName:           cus.MaidenFName,
		MaidenLName:           cus.MaidenLName,
		MaidenMName:           cus.MaidenMName,
		DoBirth:               cus.DoBirth,
		BirthPlace:            cus.BirthPlace,
		Sex:                   cus.Sex,
		CivilStatus:           cus.CivilStatus,
		Title:                 cus.Title,
		Status:                cus.Status,
		StatusDesc:            cus.StatusDesc,
		Classification:        cus.Classification,
		ClassificationDesc:    cus.ClassificationDesc,
		SubClassification:     cus.SubClassification,
		SubClassificationDesc: cus.SubClassificationDesc,
		Business:              cus.Business,
		DoEntry:               cus.DoEntry,
		DoRecognized:          cus.DoRecognized,
		DoResigned:            cus.DoResigned,
		BrCode:                cus.BrCode,
		BranchName:            cus.BranchName,
		UnitCode:              cus.UnitCode,
		UnitName:              cus.UnitName,
		CenterCode:            cus.CenterCode,
		CenterName:            cus.CenterName,
		Dosri:                 cus.Dosri,
		Reffered:              cus.Reffered,
		Remarks:               cus.Remarks,
		AccountNumber:         cus.AccountNumber,
		SearchName:            cus.SearchName,
		MemberMaidenFName:     cus.MemberMaidenFName,
		MemberMaidenLName:     cus.MemberMaidenLName,
		MemberMaidenMName:     cus.MemberMaidenMName,
	}

	return res, nil
}
*/
