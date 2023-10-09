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

func (dumpClient *DumpClient) CreateDumpAccounts(accountsList []local.AccountsInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	stream, err := dumpClient.service.CreateDumpAccounts(ctx)
	if err != nil {
		return fmt.Errorf("cannot Create Dump for accounts: %v", err)
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
	for _, accounts := range accountsList {
		req := &pb.CreateDumpAccountsRequest{
			Accounts: &pb.DumpAccounts{
				ModCtr:    accounts.ModCtr,
				ModAction: accounts.ModAction,
				BrCode:    accounts.BrCode,
				Acc:       accounts.Acc,
				Title:     accounts.Title,
				Category:  accounts.Category,
				Type:      accounts.Type,
				MainCD:    util.NullString2Proto(accounts.MainCD),
				Parent:    util.NullString2Proto(accounts.Parent),
			},
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send accounts stream request: %v - %v", err, stream.RecvMsg(nil))
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
