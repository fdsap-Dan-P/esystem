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

func (server *DumpServer) CreateDumpUsersList(stream pb.DumpService_CreateDumpUsersListServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpUsersList:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		usersList := req.GetUsersList()

		log.Printf("received UsersList request: id = %s", usersList)

		usersListReq := model.UsersList{
			UserId:              usersList.UserId,
			AccessCode:          util.NullProto2Int64(usersList.AccessCode),
			LName:               usersList.LName,
			FName:               usersList.FName,
			MName:               usersList.MName,
			DateHired:           util.NullProto2Time(usersList.DateHired),
			BirthDay:            util.NullProto2Time(usersList.BirthDay),
			DateGiven:           util.NullProto2Time(usersList.DateGiven),
			DateExpired:         util.NullProto2Time(usersList.DateExpired),
			Address:             util.NullProto2String(usersList.Address),
			Position:            util.NullProto2String(usersList.Position),
			AreaCode:            util.NullProto2Int64(usersList.AreaCode),
			ManCode:             util.NullProto2Int64(usersList.ManCode),
			AddInfo:             util.NullProto2String(usersList.AddInfo),
			Passwd:              usersList.Passwd,
			Attempt:             util.NullProto2Int64(usersList.Attempt),
			DateLocked:          util.NullProto2Time(usersList.DateLocked),
			Remarks:             util.NullProto2String(usersList.Remarks),
			Picture:             usersList.Picture,
			IsLoggedIn:          usersList.IsLoggedIn,
			AccountExpirationDt: util.NullProto2Time(usersList.AccountExpirationDt),
		}
		err = server.dump.CreateUsersList(context.Background(), usersListReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create UsersList: %v", err))
		}

		res := &pb.CreateDumpUsersListResponse{
			ModCtr: []int64{usersList.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateUsersList(context.Background(), stream.)
		// log.Printf("CreateDumpUsersList: %v", req)
	}
	return nil
}
