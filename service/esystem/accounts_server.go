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

func (server *DumpServer) CreateDumpAccounts(stream pb.DumpService_CreateDumpAccountsServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpAccounts:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		accounts := req.GetAccounts()

		log.Printf("received Accounts request: id = %s", accounts)

		accountsReq := model.Accounts{
			ModCtr:    accounts.ModCtr,
			BrCode:    accounts.BrCode,
			ModAction: accounts.ModAction,
			Acc:       accounts.Acc,
			Title:     accounts.Title,
			Category:  accounts.Category,
			Type:      accounts.Type,
			MainCD:    util.NullProto2String(accounts.MainCD),
			Parent:    util.NullProto2String(accounts.Parent),
		}
		err = server.dump.CreateAccounts(context.Background(), accountsReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Accounts: %v", err))
		}

		res := &pb.CreateDumpAccountsResponse{
			ModCtr: []int64{accounts.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateAccounts(context.Background(), stream.)
		// log.Printf("CreateDumpAccounts: %v", req)
	}
	return nil
}
