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

func (server *DumpServer) CreateDumpCustAddInfoList(stream pb.DumpService_CreateDumpCustAddInfoListServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCustAddInfoList:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		custAddInfoList := req.GetCustAddInfoList()

		log.Printf("received CustAddInfoList request: id = %s", custAddInfoList)

		custAddInfoListReq := model.CustAddInfoList{
			ModCtr:     custAddInfoList.ModCtr,
			BrCode:     custAddInfoList.BrCode,
			ModAction:  custAddInfoList.ModAction,
			InfoCode:   custAddInfoList.InfoCode,
			InfoOrder:  custAddInfoList.InfoOrder,
			Title:      custAddInfoList.Title,
			InfoType:   custAddInfoList.InfoType,
			InfoLen:    custAddInfoList.InfoLen,
			InfoFormat: custAddInfoList.InfoFormat,
			InputType:  custAddInfoList.InputType,
			InfoSource: custAddInfoList.InfoSource,
		}
		err = server.dump.CreateCustAddInfoList(context.Background(), custAddInfoListReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfoList: %v", err))
		}

		res := &pb.CreateDumpCustAddInfoListResponse{
			ModCtr: []int64{custAddInfoList.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCustAddInfoList(context.Background(), stream.)
		// log.Printf("CreateDumpCustAddInfoList: %v", req)
	}
	return nil
}
