package client

import (
	"context"
	"fmt"
	"log"

	// kplus "simplebank/db/datastore/kplus"
	pb "simplebank/pb"
)

func (kPlusClient *KPlusClient) SearchCustomerCID(ctx context.Context, cid int64) (*pb.KPLUSCustomerResponse, error) {

	log.Printf("client.SearchCustomerCID %v", cid)
	in := &pb.KPLUSCustomerRequest{Cid: cid}
	cus, err := kPlusClient.service.SearchCustomerCID(ctx, in)
	log.Printf("client.SearchCustomerCID err %v", err)
	if err != nil {
		return &pb.KPLUSCustomerResponse{}, fmt.Errorf("cannot SearchCustomerCID: %v", err)
	}
	log.Printf("client.SearchCustomerCID cus %v", cus)

	return cus, err
}

func (kPlusClient *KPlusClient) CustSavingsList(ctx context.Context, cid int64) (*pb.KPLUSCustSavingsListResponse, error) {

	log.Printf("client.CustSavingsList %v", cid)
	in := &pb.KPLUSCustomerRequest{Cid: cid}
	data, err := kPlusClient.service.CustSavingsList(ctx, in)
	log.Printf("client.CustSavingsList err %v", err)
	if err != nil {
		return &pb.KPLUSCustSavingsListResponse{}, fmt.Errorf("cannot CustSavingsList: %v", err)
	}
	log.Printf("client.CustSavingsList cus %v", data)

	return data, err
}
