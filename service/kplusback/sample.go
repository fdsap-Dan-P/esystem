package service

// import (
// 	"context"
// 	"log"
// 	dskPlus "simplebank/db/datastore/kplus"
// 	pb "simplebank/pb"
// 	"simplebank/util"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func (server *KPlusServer) GenerateColShtperCID(ctx context.Context, r *pb.KPLUSCustomerRequest) (*pb.KPLUSGenerateColShtperCIDResponse, error) {

// 	cid := r.GetCid()

// 	log.Printf("received GenerateColShtperCID request: cid:%v, from:%v to:%v ", cid)

// 	data, err := server.kplus.Transaction(context.Background(),
// 		dskPlus.TransactionParams{Cid: cid, DateFrom: dateFrom, DateTo: dateTo})
// 	if err != nil {
// 		return &pb.KPLUSGenerateColShtperCIDResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to GenerateColShtperCID: %v", err))
// 	}

// 	res := []*pb.KPLUSGenerateColShtperCIDResponse{}
// 	for _, d := range data {
// 		r := &pb.KPLUSGenerateColShtperCIDResponse{
// 			BrCode:          d.BrCode,
// 			AppType:         d.AppType,
// 			Code:            d.Code,
// 			Status:          d.Status,
// 			StatusDesc:      d.StatusDesc,
// 			Acc:             d.Acc,
// 			Iiid:            d.Iiid,
// 			CustomerId:      d.CustomerId,
// 			CentralOfficeId: d.CentralOfficeId,
// 			CID:             d.CID,
// 			UM:              d.UM,
// 			ClientName:      d.ClientName,
// 			CenterCode:      d.CenterCode,
// 			CenterName:      d.CenterName,
// 			ManCode:         d.ManCode,
// 			Unit:            d.Unit,
// 			AreaCode:        d.AreaCode,
// 			Area:            d.Area,
// 			StaffName:       d.StaffName,
// 			AcctType:        d.AcctType,
// 			AcctDesc:        d.AcctDesc,
// 			DisbDate:        d.DisbDate,
// 			DateStart:       d.DateStart,
// 			Maturity:        d.Maturity,
// 			Principal:       util.NullProto2Decimal(d.Principal),
// 			Interest:        util.NullProto2Decimal(d.Interest),
// 			Gives:           d.Gives,
// 			IbalPrin:        util.NullProto2Decimal(d.IbalPrin),
// 			IbalInt:         util.NullProto2Decimal(d.IbalInt),
// 			BalPrin:         util.NullProto2Decimal(d.BalPrin),
// 			BalInt:          util.NullProto2Decimal(d.BalInt),
// 			Amort:           util.NullProto2Decimal(d.Amort),
// 			DuePrin:         util.NullProto2Decimal(d.DuePrin),
// 			DueInt:          util.NullProto2Decimal(d.DueInt),
// 			LoanBal:         util.NullProto2Decimal(d.LoanBal),
// 			SaveBal:         util.NullProto2Decimal(d.SaveBal),
// 			WaivedInt:       util.NullProto2Decimal(d.WaivedInt),
// 			UnPaidCtr:       d.UnPaidCtr,
// 			WritenOff:       d.WritenOff,
// 			Classification:  d.Classification,
// 			ClassDesc:       d.ClassDesc,
// 			WriteOff:        d.WriteOff,
// 			Pay:             util.NullProto2Decimal(d.Pay),
// 			Withdraw:        util.NullProto2Decimal(d.Withdraw),
// 			Type:            d.Type,
// 			OrgName:         d.OrgName,
// 			OrgAddress:      d.OrgAddress,
// 			MeetingDate:     d.MeetingDate,
// 			MeetingDay:      d.MeetingDay,
// 			SharesOfStock:   util.NullProto2Decimal(d.SharesOfStock),
// 			DateEstablished: d.DateEstablished,
// 			Uuid:            d.Uuid,
// 		}
// 		res = append(res, r)
// 	}
// 	return &pb.KPLUSGenerateColShtperCIDResponse{Transaction: res}, nil
// }
