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

func (server *DumpServer) CreateDumpCenterWorker(stream pb.DumpService_CreateDumpCenterWorkerServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCenterWorker:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		centerWorker := req.GetCenterWorker()

		log.Printf("received CenterWorker request: id = %s", centerWorker)

		centerWorkerReq := model.CenterWorker{
			ModCtr:      centerWorker.ModCtr,
			BrCode:      centerWorker.BrCode,
			ModAction:   centerWorker.ModAction,
			AOID:        centerWorker.AOID,
			Lname:       util.NullProto2String(centerWorker.Lname),
			FName:       util.NullProto2String(centerWorker.FName),
			Mname:       util.NullProto2String(centerWorker.Mname),
			PhoneNumber: util.NullProto2String(centerWorker.PhoneNumber),
		}
		err = server.dump.CreateCenterWorker(context.Background(), centerWorkerReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create CenterWorker: %v", err))
		}

		res := &pb.CreateDumpCenterWorkerResponse{
			ModCtr: []int64{centerWorker.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCenterWorker(context.Background(), stream.)
		// log.Printf("CreateDumpCenterWorker: %v", req)
	}
	return nil
}
