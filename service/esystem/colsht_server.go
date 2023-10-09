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

func (server *DumpServer) CreateDumpColSht(stream pb.DumpService_CreateDumpColShtServer) error {
	data := []model.ColSht{}
	accList := []string{}
	var err error
	var buf int64 = 0
	for {
		err = util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			if buf > 0 {
				err = server.dump.BulkInsertColSht(context.Background(), data)
				if err != nil {
					return util.LogError(status.Errorf(codes.Unknown, "Unable to create ColSht: %v", err))
				}
				res := &pb.CreateDumpColShtResponse{Acc: accList}
				err = stream.Send(res)
				if err != nil {
					return util.LogError(status.Errorf(codes.Unknown, "Unable to create ColSht: %v", err))
				}
			}
			log.Print("CreateDumpColSht:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		colSht := req.GetColSht()

		log.Printf("received ColSht request: id = %s", colSht)

		colShtReq := model.ColSht{
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
			DisbDate:        colSht.DisbDate.AsTime(),
			DateStart:       colSht.DateStart.AsTime(),
			Maturity:        colSht.Maturity.AsTime(),
			Principal:       util.Proto2Decimal(colSht.Principal),
			Interest:        util.Proto2Decimal(colSht.Interest),
			Gives:           colSht.Gives,
			BalPrin:         util.Proto2Decimal(colSht.BalPrin),
			BalInt:          util.Proto2Decimal(colSht.BalInt),
			Amort:           util.Proto2Decimal(colSht.Amort),
			DuePrin:         util.Proto2Decimal(colSht.DuePrin),
			DueInt:          util.Proto2Decimal(colSht.DueInt),
			LoanBal:         util.Proto2Decimal(colSht.LoanBal),
			SaveBal:         util.Proto2Decimal(colSht.SaveBal),
			WaivedInt:       util.Proto2Decimal(colSht.WaivedInt),
			UnPaidCtr:       colSht.UnPaidCtr,
			WrittenOff:      colSht.WrittenOff,
			OrgName:         colSht.OrgName,
			OrgAddress:      colSht.OrgAddress,
			MeetingDate:     colSht.MeetingDate.AsTime(),
			MeetingDay:      colSht.MeetingDay,
			SharesOfStock:   util.Proto2Decimal(colSht.SharesOfStock),
			DateEstablished: colSht.DateEstablished.AsTime(),
			Classification:  colSht.Classification,
			WriteOff:        colSht.WriteOff,
		}
		data = append(data, colShtReq)
		accList = append(accList, colSht.Acc)
		buf++
		if buf > 1000 {
			err = server.dump.BulkInsertColSht(context.Background(), data)
			if err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "Unable to create ColSht: %v", err))
			}
			res := &pb.CreateDumpColShtResponse{Acc: accList}
			err = stream.Send(res)

			if err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "cannot send stream res:%v response: %v", res, err))
			}
			// Reset
			buf = 0
			data = []model.ColSht{}
			accList = []string{}
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateColSht(context.Background(), stream.)
		// log.Printf("CreateDumpColSht: %v", req)
	}

	return err
}
