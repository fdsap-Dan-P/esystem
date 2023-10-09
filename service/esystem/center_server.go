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

func (server *DumpServer) CreateDumpCenter(stream pb.DumpService_CreateDumpCenterServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpCenter:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		center := req.GetCenter()

		log.Printf("received Center request: id = %s", center)

		centerReq := model.Center{
			ModCtr:          center.ModCtr,
			BrCode:          center.BrCode,
			ModAction:       center.ModAction,
			CenterCode:      center.CenterCode,
			CenterName:      util.NullProto2String(center.CenterName),
			CenterAddress:   util.NullProto2String(center.CenterAddress),
			MeetingDay:      util.NullProto2Int64(center.MeetingDay),
			Unit:            util.NullProto2Int64(center.Unit),
			DateEstablished: util.NullProto2Time(center.DateEstablished),
		}
		err = server.dump.CreateCenter(context.Background(), centerReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Center: %v", err))
		}

		res := &pb.CreateDumpCenterResponse{
			ModCtr: []int64{center.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateCenter(context.Background(), stream.)
		// log.Printf("CreateDumpCenter: %v", req)
	}
	return nil
}
