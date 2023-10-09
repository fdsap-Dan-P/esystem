package service

import (
	"context"
	"log"
	pb "simplebank/pb"
	"simplebank/util"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (server *KPlusServiceServer) K2CCallBackRef(ctx context.Context, r *pb.KPLUSCallBackRefRequest) (*pb.KPLUSResponse, error) {

	prNo := r.GetPrNumber()
	log.Printf("received K2CCallBackRef request: cid = %v", prNo)

	dt, err := server.kplus.CallBackRef(context.Background(), prNo)
	if err != nil {
		return &pb.KPLUSResponse{}, util.LogError(status.Errorf(codes.Unknown, "Unable to K2CCallBackRef: %v", err))
	}

	res := &pb.KPLUSResponse{
		RetCode:   dt.RetCode,
		Message:   dt.Message,
		Reference: dt.Reference,
	}

	return res, nil
}
