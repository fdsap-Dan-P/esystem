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

func (server *DumpServer) CreateDumpLnBeneficiary(stream pb.DumpService_CreateDumpLnBeneficiaryServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpLnBeneficiary:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		lnBeneficiary := req.GetLnBeneficiary()

		log.Printf("received LnBeneficiary request: id = %s", lnBeneficiary)

		lnBeneficiaryReq := model.LnBeneficiary{
			ModCtr:     lnBeneficiary.ModCtr,
			BrCode:     lnBeneficiary.BrCode,
			Acc:        lnBeneficiary.Acc,
			BDay:       lnBeneficiary.BDay.AsTime(),
			EducLvl:    lnBeneficiary.EducLvl,
			Gender:     lnBeneficiary.Gender,
			LastName:   util.NullProto2String(lnBeneficiary.LastName),
			FirstName:  util.NullProto2String(lnBeneficiary.FirstName),
			MiddleName: util.NullProto2String(lnBeneficiary.MiddleName),
			Remarks:    util.NullProto2String(lnBeneficiary.Remarks),
		}
		err = server.dump.CreateLnBeneficiary(context.Background(), lnBeneficiaryReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create LnBeneficiary: %v", err))
		}

		res := &pb.CreateDumpLnBeneficiaryResponse{
			ModCtr: []int64{lnBeneficiary.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateLnBeneficiary(context.Background(), stream.)
		// log.Printf("CreateDumpLnBeneficiary: %v", req)
	}
	return nil
}
