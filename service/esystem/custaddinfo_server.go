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

func (server *DumpServer) CreateDumpCustAddInfo(stream pb.DumpService_CreateDumpCustAddInfoServer) error {
	data := []model.CustAddInfo{}
	modCtrList := []int64{}
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
				err = server.dump.BulkInsertCustAddInfo(context.Background(), data)
				if err != nil {
					return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfo: %v", err))
				}
				res := &pb.CreateDumpCustAddInfoResponse{ModCtr: modCtrList}
				err = stream.Send(res)
				if err != nil {
					return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfo: %v", err))
				}
			}
			log.Print("CreateDumpCustAddInfo:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		custAddInfo := req.GetCustAddInfo()

		log.Printf("received CustAddInfo request: id = %s", custAddInfo)

		custAddInfoReq := model.CustAddInfo{
			ModCtr:    custAddInfo.ModCtr,
			BrCode:    custAddInfo.BrCode,
			ModAction: custAddInfo.ModAction,
			CID:       custAddInfo.CID,
			InfoDate:  custAddInfo.InfoDate.AsTime(),
			InfoCode:  custAddInfo.InfoCode,
			Info:      custAddInfo.Info,
			InfoValue: custAddInfo.InfoValue,
		}
		data = append(data, custAddInfoReq)
		modCtrList = append(modCtrList, custAddInfo.ModCtr)
		buf++
		if buf > 1000 {
			err = server.dump.BulkInsertCustAddInfo(context.Background(), data)
			if err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfo: %v", err))
			}
			res := &pb.CreateDumpCustAddInfoResponse{ModCtr: modCtrList}
			err = stream.Send(res)
			// Reset
			buf = 0
			data = []model.CustAddInfo{}
			modCtrList = []int64{}
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCustAddInfo(context.Background(), stream.)
		// log.Printf("CreateDumpCustAddInfo: %v", req)
	}

	return err
}
