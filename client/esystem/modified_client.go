package client

import (
	"context"
	"io"
	"log"

	pb "simplebank/pb"
)

type modTableOrder struct {
	CentralTableName string
	LocalTableName   string
	ModCtr           int64
}

type modTableLocal struct {
	CentralTableName string
	ModCtr           int64
}

type modTableCentral struct {
	LocalTableName string
	ModCtr         int64
}

//  modTableOrderList := make(map[int16]modTableOrder)
//  modTableLocalList := make(map[string]modTableLocal)
//  modTableCentralList := make(map[string]modTableCentral)

// func (client *DumpClient) SetModCtrTableLocal(tabName string, modCtr int64) {
// 	// val, ok := client.ModTableLocalList[tabName]
// 	// if ok {
// 	// 	val.ModCtr = modCtr
// 	// 	client.ModTableLocalList[tabName] = val
// 	// } else {
// 	client.ModTableLocalList[tabName] = modTableLocal{
// 		ModCtr: modCtr,
// 	}
// 	// }
// }

// func (client *DumpClient) SetModCtrTableCentral(tabName string, modCtr int64) {
// 	client.ModTableCentralList[tabName] = modTableCentral{
// 		ModCtr: modCtr,
// 	}
// }

func (client *DumpClient) GetModifiedTable(ctx context.Context, in *pb.GetModifiedTableRequest) {

	stream, err := client.service.GetModifiedTable(context.Background(), in)
	if err != nil {
		log.Printf("With Error %v", err)
	}

	for {
		mod, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client: %v | Error: %v", client, err)
		}
		client.ModTableCentralList[mod.TableName] = mod.LastModCtr
	}
	// return
}
