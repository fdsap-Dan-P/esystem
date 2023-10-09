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

func (server *DumpServer) CreateDumpJnlDetails(stream pb.DumpService_CreateDumpJnlDetailsServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpJnlDetails:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		jnlDetails := req.GetJnlDetails()

		log.Printf("received JnlDetails request: id = %s", jnlDetails)

		jnlDetailsReq := model.JnlDetails{
			ModCtr:    jnlDetails.ModCtr,
			BrCode:    jnlDetails.BrCode,
			ModAction: jnlDetails.ModAction,
			Acc:       jnlDetails.Acc,
			Trn:       jnlDetails.Trn,
			Series:    util.NullProto2Int64(jnlDetails.Series),
			Debit:     util.NullProto2Decimal(jnlDetails.Debit),
			Credit:    util.NullProto2Decimal(jnlDetails.Credit),
		}
		err = server.dump.CreateJnlDetails(context.Background(), jnlDetailsReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create JnlDetails: %v", err))
		}

		res := &pb.CreateDumpJnlDetailsResponse{
			ModCtr: []int64{jnlDetails.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateJnlDetails(context.Background(), stream.)
		// log.Printf("CreateDumpJnlDetails: %v", req)
	}
	return nil
}
