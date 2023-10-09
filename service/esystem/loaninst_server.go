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

func (server *DumpServer) CreateDumpLoanInst(stream pb.DumpService_CreateDumpLoanInstServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpLoanInst:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		loanInst := req.GetLoanInst()

		log.Printf("received LoanInst request: id = %s", loanInst)

		loanInstReq := model.LoanInst{
			ModCtr:    loanInst.ModCtr,
			BrCode:    loanInst.BrCode,
			ModAction: loanInst.ModAction,
			Acc:       loanInst.Acc,
			Dnum:      loanInst.Dnum,
			DueDate:   loanInst.DueDate.AsTime(),
			InstFlag:  loanInst.InstFlag,
			DuePrin:   util.Proto2Decimal(loanInst.DuePrin),
			DueInt:    util.Proto2Decimal(loanInst.DueInt),
			UpInt:     util.Proto2Decimal(loanInst.UpInt),
		}
		err = server.dump.CreateLoanInst(context.Background(), loanInstReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create LoanInst: %v", err))
		}

		res := &pb.CreateDumpLoanInstResponse{
			ModCtr: []int64{loanInst.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateLoanInst(context.Background(), stream.)
		// log.Printf("CreateDumpLoanInst: %v", req)
	}
	return nil
}
