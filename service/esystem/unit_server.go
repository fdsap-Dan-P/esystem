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

func (server *DumpServer) CreateDumpUnit(stream pb.DumpService_CreateDumpUnitServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpUnit:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		unit := req.GetUnit()

		log.Printf("received Unit request: id = %s", unit)

		unitReq := model.Unit{
			ModCtr:      unit.ModCtr,
			BrCode:      unit.BrCode,
			ModAction:   unit.ModAction,
			UnitCode:    unit.UnitCode,
			Unit:        util.NullProto2String(unit.Unit),
			AreaCode:    util.NullProto2Int64(unit.AreaCode),
			FName:       util.NullProto2String(unit.FName),
			LName:       util.NullProto2String(unit.LName),
			MName:       util.NullProto2String(unit.MName),
			VatReg:      util.NullProto2String(unit.VatReg),
			UnitAddress: util.NullProto2String(unit.UnitAddress),
		}
		err = server.dump.CreateUnit(context.Background(), unitReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Unit: %v", err))
		}

		res := &pb.CreateDumpUnitResponse{
			ModCtr: []int64{unit.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateUnit(context.Background(), stream.)
		// log.Printf("CreateDumpUnit: %v", req)
	}
	return nil
}
