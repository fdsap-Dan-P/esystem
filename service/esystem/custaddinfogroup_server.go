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

func (server *DumpServer) CreateDumpCustAddInfoGroup(stream pb.DumpService_CreateDumpCustAddInfoGroupServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCustAddInfoGroup:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		custAddInfoGroup := req.GetCustAddInfoGroup()

		log.Printf("received CustAddInfoGroup request: id = %s", custAddInfoGroup)

		custAddInfoGroupReq := model.CustAddInfoGroup{
			ModCtr:     custAddInfoGroup.ModCtr,
			BrCode:     custAddInfoGroup.BrCode,
			ModAction:  custAddInfoGroup.ModAction,
			InfoGroup:  custAddInfoGroup.InfoGroup,
			GroupTitle: util.NullProto2String(custAddInfoGroup.GroupTitle),
			Remarks:    util.NullProto2String(custAddInfoGroup.Remarks),
			ReqOnEntry: util.NullProto2Bool(custAddInfoGroup.ReqOnEntry),
			ReqOnExit:  util.NullProto2Bool(custAddInfoGroup.ReqOnExit),
			Link2Loan:  util.NullProto2Int64(custAddInfoGroup.Link2Loan),
			Link2Save:  util.NullProto2Int64(custAddInfoGroup.Link2Save),
		}
		err = server.dump.CreateCustAddInfoGroup(context.Background(), custAddInfoGroupReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create CustAddInfoGroup: %v", err))
		}

		res := &pb.CreateDumpCustAddInfoGroupResponse{
			ModCtr: []int64{custAddInfoGroup.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCustAddInfoGroup(context.Background(), stream.)
		// log.Printf("CreateDumpCustAddInfoGroup: %v", req)
	}
	return nil
}
