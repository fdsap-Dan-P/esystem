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

func (server *DumpServer) CreateDumpJnlHeaders(stream pb.DumpService_CreateDumpJnlHeadersServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpJnlHeaders:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		jnlHeaders := req.GetJnlHeaders()

		log.Printf("received JnlHeaders request: id = %s", jnlHeaders)

		jnlHeadersReq := model.JnlHeaders{
			ModCtr:      jnlHeaders.ModCtr,
			BrCode:      jnlHeaders.BrCode,
			ModAction:   jnlHeaders.ModAction,
			Trn:         jnlHeaders.Trn,
			TrnDate:     jnlHeaders.TrnDate.AsTime(),
			Particulars: jnlHeaders.Particulars,
			UserName:    util.NullProto2String(jnlHeaders.UserName),
			Code:        jnlHeaders.Code,
		}
		err = server.dump.CreateJnlHeaders(context.Background(), jnlHeadersReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create JnlHeaders: %v", err))
		}

		res := &pb.CreateDumpJnlHeadersResponse{
			ModCtr: []int64{jnlHeaders.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateJnlHeaders(context.Background(), stream.)
		// log.Printf("CreateDumpJnlHeaders: %v", req)
	}
	return nil
}
