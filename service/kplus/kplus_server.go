package service

import (
	"context"
	"database/sql"
	"log"

	pb "simplebank/pb"
	"simplebank/util"

	kplus "simplebank/db/datastore/kplus"
	dsUsr "simplebank/db/datastore/user"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type KPlusServiceServer struct {
	pb.UnimplementedKPlusServiceServer
	db *sql.DB
	// jwtManager *JWTManager
	kplus     *kplus.QueriesKPlus
	userStore dsUsr.QueriesUser //.UserStore
}

func NewKPlusServiceServer(db *sql.DB, userStore dsUsr.QueriesUser) *KPlusServiceServer {
	log.Println("NewKplusServer..........")
	return &KPlusServiceServer{
		db: db,
		// jwtManager: jwtManager,
		userStore: userStore,
		kplus:     kplus.New(db),
	}
}

func (server *KPlusServiceServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: util.Concatinate("Hello ", in.Name)}, nil
}

// func (server *KPlusServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
// 	log.Printf("Login %v", req.GetLoginName())
// 	// user, err := server.userStore.GetUserbyName(ctx, req.GetLoginName())
// 	// if err != nil {
// 	// 	return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
// 	// }

// 	// // log.Println("Login 1", user.LoginName)

// 	// if !user.IsCorrectPassword(req.GetPassword()) {
// 	// 	return nil, status.Errorf(codes.NotFound, "incorrect username/password")
// 	// }
// 	// token, err := server.jwtManager.Generate(user)
// 	// if err != nil {
// 	// 	return nil, status.Errorf(codes.Internal, "cannot generate access token")
// 	// }

// 	token := "test Token"
// 	res := &pb.LoginResponse{AccessToken: token}
// 	return res, nil
// }

func (server *KPlusServiceServer) SearchCustomerCID(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSCustomerResponse, error) {

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

func (server *KPlusServiceServer) CustSavingsList(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSCustSavingsListResponse, error) {

	cid := r.GetCid()
	log.Printf("received SearchCustomerCID request: cid = %v", cid)

	cus, err := server.kplus.SavingsList(context.Background(), kplus.SavingsListParams{INAIIID: cid})
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

func (server *KPlusServiceServer) GenerateColShtperCID(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSGenerateColShtperCIDResponse, error) {

	inaiiid := r.GetCid()

	log.Printf("received GetTransactionHistory request: acc:%v ", inaiiid)

	data, err := server.kplus.ColSht(context.Background(),
		kplus.ColShtParams{INAIIID: inaiiid})
	if err != nil {
		return &pb.KPLUSGenerateColShtperCIDResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to Get Collection Sheet: %v", err))
	}

	res := []*pb.KPLUSColShtperCID{}
	for _, d := range data {
		r := &pb.KPLUSColShtperCID{
			INAIIID:         d.INAIIID,
			BrCode:          d.BrCode,
			AppType:         d.AppType,
			Code:            d.Code,
			Status:          d.Status,
			StatusDesc:      d.StatusDesc,
			Acc:             d.Acc,
			Iiid:            d.Iiid,
			CustomerId:      d.CustomerId,
			CentralOfficeId: d.CentralOfficeId,
			CID:             d.CID,
			UM:              d.UM,
			ClientName:      d.ClientName,
			CenterCode:      d.CenterCode,
			CenterName:      d.CenterName,
			ManCode:         d.ManCode,
			Unit:            d.Unit,
			AreaCode:        d.AreaCode,
			Area:            d.Area,
			StaffName:       d.StaffName,
			AcctType:        d.AcctType,
			AcctDesc:        d.AcctDesc,
			DisbDate:        d.DisbDate,
			DateStart:       d.DateStart,
			Maturity:        d.Maturity,
			Principal:       d.Principal,
			Interest:        d.Interest,
			Gives:           d.Gives,
			IbalPrin:        d.IbalPrin,
			IbalInt:         d.IbalInt,
			BalPrin:         d.BalPrin,
			BalInt:          d.BalInt,
			Amort:           d.Amort,
			DuePrin:         d.DuePrin,
			DueInt:          d.DueInt,
			LoanBal:         d.LoanBal,
			SaveBal:         d.SaveBal,
			WaivedInt:       d.WaivedInt,
			UnPaidCtr:       d.UnPaidCtr,
			WritenOff:       d.WritenOff,
			Classification:  d.Classification,
			ClassDesc:       d.ClassDesc,
			WriteOff:        d.WriteOff,
			Pay:             d.Pay,
			Withdraw:        d.Withdraw,
			Type:            d.Type,
			OrgName:         d.OrgName,
			OrgAddress:      d.OrgAddress,
			MeetingDate:     d.MeetingDate,
			MeetingDay:      d.MeetingDay,
			SharesOfStock:   d.SharesOfStock,
			DateEstablished: d.DateEstablished,
			Uuid:            d.Uuid,
		}
		res = append(res, r)
	}
	return &pb.KPLUSGenerateColShtperCIDResponse{KPLUSColShtperCID: res}, nil
}

func (server *KPlusServiceServer) GetTransactionHistory(ctx context.Context, r *pb.KPLUSGetTransactionHistoryRequest) (*pb.KPLUSGetTransactionHistoryResponse, error) {

	acc := r.GetAcc()
	dateFrom := util.SetDate(r.GetDateFrom())
	dateTo := util.SetDate(r.GetDateTo())

	log.Printf("received GetTransactionHistory request: acc:%v, from:%v to:%v ", acc, dateFrom, dateTo)

	data, err := server.kplus.Transaction(context.Background(),
		kplus.TransactionParams{Acc: acc, DateFrom: dateFrom, DateTo: dateTo})
	if err != nil {
		return &pb.KPLUSGetTransactionHistoryResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to GetTransactionHistory: %v", err))
	}

	res := []*pb.KPLUSTransaction{}
	for _, d := range data {
		r := &pb.KPLUSTransaction{
			Acc:          d.Acc,
			Trndate:      d.TrnDate,
			TrnHeadId:    d.TrnHeadId,
			Trn:          d.Trn,
			AlternateKey: d.AlternateKey,
			Prin:         d.Prin,
			Intr:         d.Intr,
			TrnAmount:    d.TrnAmount,
			Balance:      d.Balance,
			Particulars:  d.Particulars,
			TrnType:      d.TrnType,
			Username:     d.Username,
			IsFinancial:  d.IsFinancial,
		}
		res = append(res, r)
	}
	return &pb.KPLUSGetTransactionHistoryResponse{Transaction: res}, nil
}
