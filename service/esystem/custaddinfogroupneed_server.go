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

func (server *DumpServer) CreateDumpCustAddInfoGroupNeed(stream pb.DumpService_CreateDumpCustAddInfoGroupNeedServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCustAddInfoGroupNeed:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		custAddInfoGroupNeed := req.GetCustAddInfoGroupNeed()

		log.Printf("received CustAddInfoGroupNeed request: id = %s", custAddInfoGroupNeed)

		custAddInfoGroupNeedReq := model.CustAddInfoGroupNeed{
			ModCtr:      custAddInfoGroupNeed.ModCtr,
			BrCode:      custAddInfoGroupNeed.BrCode,
			ModAction:   custAddInfoGroupNeed.ModAction,
			InfoGroup:   custAddInfoGroupNeed.InfoGroup,
			InfoCode:    custAddInfoGroupNeed.InfoCode,
			InfoProcess: util.NullProto2String(custAddInfoGroupNeed.InfoProcess),
		}
		err = server.dump.CreateCustAddInfoGroupNeed(context.Background(), custAddInfoGroupNeedReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfoGroupNeed: %v", err))
		}

		res := &pb.CreateDumpCustAddInfoGroupNeedResponse{
			ModCtr: []int64{custAddInfoGroupNeed.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCustAddInfoGroupNeed(context.Background(), stream.)
		// log.Printf("CreateDumpCustAddInfoGroupNeed: %v", req)
	}
	return nil
}
