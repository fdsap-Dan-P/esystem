package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	local "simplebank/db/datastore/esystemlocal"
	pb "simplebank/pb"
	"simplebank/util"
)

func (dumpClient *DumpClient) CreateDumpUsersList(usersListList []local.UsersListInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpUsersList(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for usersList: %v", err)
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			QueriesLocal.UpdateModifiedTableUploaded(context.Background(), res.ModCtr, true)
			log.Printf("received response: %v", res)
		}
	}()

	// send requests
	for _, usersList := range usersListList {
		req := &pb.CreateDumpUsersListRequest{
			UsersList: &pb.DumpUsersList{
				UserId:              usersList.UserId,
				AccessCode:          util.NullInt642Proto(usersList.AccessCode),
				LName:               usersList.LName,
				FName:               usersList.FName,
				MName:               usersList.MName,
				DateHired:           util.NullTime2Proto(usersList.DateHired),
				BirthDay:            util.NullTime2Proto(usersList.BirthDay),
				DateGiven:           util.NullTime2Proto(usersList.DateGiven),
				DateExpired:         util.NullTime2Proto(usersList.DateExpired),
				Address:             util.NullString2Proto(usersList.Address),
				Position:            util.NullString2Proto(usersList.Position),
				AreaCode:            util.NullInt642Proto(usersList.AreaCode),
				ManCode:             util.NullInt642Proto(usersList.ManCode),
				AddInfo:             util.NullString2Proto(usersList.AddInfo),
				Passwd:              usersList.Passwd,
				Attempt:             util.NullInt642Proto(usersList.Attempt),
				DateLocked:          util.NullTime2Proto(usersList.DateLocked),
				Remarks:             util.NullString2Proto(usersList.Remarks),
				Picture:             usersList.Picture,
				IsLoggedIn:          usersList.IsLoggedIn,
				AccountExpirationDt: util.NullTime2Proto(usersList.AccountExpirationDt),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send usersList stream request: %v - %v", err, stream.RecvMsg(nil))
		}
		log.Print("sent request: ", req)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}
