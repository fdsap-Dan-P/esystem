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

func (server *DumpServer) CreateDumpAddresses(stream pb.DumpService_CreateDumpAddressesServer) error {
	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("CreateDumpAddresses:no more data")
			break
		}
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		addresses := req.GetAddresses()

		log.Printf("received Addresses request: id = %s", addresses)

		addressesReq := model.Addresses{
			ModCtr:         addresses.ModCtr,
			BrCode:         addresses.BrCode,
			ModAction:      addresses.ModAction,
			CID:            addresses.CID,
			SeqNum:         addresses.SeqNum,
			AddressDetails: util.NullProto2String(addresses.AddressDetails),
			Barangay:       util.NullProto2String(addresses.Barangay),
			City:           util.NullProto2String(addresses.City),
			Province:       util.NullProto2String(addresses.Province),
			Phone1:         util.NullProto2String(addresses.Phone1),
			Phone2:         util.NullProto2String(addresses.Phone2),
			Phone3:         util.NullProto2String(addresses.Phone3),
			Phone4:         util.NullProto2String(addresses.Phone4),
		}
		err = server.dump.CreateAddresses(context.Background(), addressesReq)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "Unable to create Addresses: %v", err))
		}

		res := &pb.CreateDumpAddressesResponse{
			ModCtr: []int64{addresses.ModCtr},
		}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
		//err := dump.CreateAddresses(context.Background(), stream.)
		// log.Printf("CreateDumpAddresses: %v", req)
	}
	return nil
}
