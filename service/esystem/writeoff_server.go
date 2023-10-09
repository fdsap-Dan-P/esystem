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

func (server *DumpServer) CreateDumpWriteoff(stream pb.DumpService_CreateDumpWriteoffServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpWriteoff:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		writeoff := req.GetWriteoff()

		log.Printf("received Writeoff request: id = %s", writeoff)

		writeoffReq := model.Writeoff{
			ModCtr:     writeoff.ModCtr,
			BrCode:     writeoff.BrCode,
			ModAction:  writeoff.ModAction,
			Acc:        writeoff.Acc,
			DisbDate:   writeoff.DisbDate.AsTime(),
			Principal:  util.Proto2Decimal(writeoff.Principal),
			Interest:   util.Proto2Decimal(writeoff.Interest),
			BalPrin:    util.Proto2Decimal(writeoff.BalPrin),
			BalInt:     util.Proto2Decimal(writeoff.BalInt),
			TrnDate:    writeoff.TrnDate.AsTime(),
			AcctType:   writeoff.AcctType,
			Print:      util.NullProto2String(writeoff.Print),
			PostedBy:   util.NullProto2String(writeoff.PostedBy),
			VerifiedBy: util.NullProto2String(writeoff.VerifiedBy),
		}
		err = server.dump.CreateWriteoff(context.Background(), writeoffReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Writeoff: %v", err))
		}

		res := &pb.CreateDumpWriteoffResponse{
			ModCtr: []int64{writeoff.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateWriteoff(context.Background(), stream.)
		// log.Printf("CreateDumpWriteoff: %v", req)
	}
	return nil
}
