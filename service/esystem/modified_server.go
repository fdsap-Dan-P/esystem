package service

import (
	"context"

	pb "simplebank/pb"
	"simplebank/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb.DumpServiceServer = (*DumpServer)(nil)

func (server *DumpServer) GetModifiedTable(d *pb.GetModifiedTableRequest, stream pb.DumpService_GetModifiedTableServer) error {

	modList, err := server.dump.ListModifiedTable(context.Background(), d.BrCode)

	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "Cannot Get Modified Table: %v", err))
	}

	for _, mod := range modList {
		if err := stream.Send(&pb.DumpModifiedTable{
			BrCode:     mod.BrCode,
			TableName:  mod.DumpTable,
			LastModCtr: mod.LastModCtr,
		}); err != nil {
			return err
		}
	}
	return nil
}
