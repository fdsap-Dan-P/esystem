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

func (server *DumpServer) CreateDumpReferencesDetails(stream pb.DumpService_CreateDumpReferencesDetailsServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpReferencesDetails:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		referencesDetails := req.GetReferencesDetails()

		log.Printf("received ReferencesDetails request: id = %s", referencesDetails)

		referencesDetailsReq := model.ReferencesDetails{
			ModCtr:             referencesDetails.ModCtr,
			BrCode:             referencesDetails.BrCode,
			ModAction:          referencesDetails.ModAction,
			ID:                 referencesDetails.ID,
			RefID:              referencesDetails.RefID,
			PurposeDescription: util.NullProto2String(referencesDetails.PurposeDescription),
			ParentID:           util.NullProto2Int64(referencesDetails.ParentID),
			CodeID:             util.NullProto2Int64(referencesDetails.CodeID),
			Stat:               util.NullProto2Int64(referencesDetails.Stat),
		}
		err = server.dump.CreateReferencesDetails(context.Background(), referencesDetailsReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create ReferencesDetails: %v", err))
		}

		res := &pb.CreateDumpReferencesDetailsResponse{
			ModCtr: []int64{referencesDetails.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateReferencesDetails(context.Background(), stream.)
		// log.Printf("CreateDumpReferencesDetails: %v", req)
	}
	return nil
}
