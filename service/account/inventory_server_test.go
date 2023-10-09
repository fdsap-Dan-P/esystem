package service

import (
	"context"
	"log"
	pb "simplebank/pb"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInventoryItem(t *testing.T) {
	// log.Printf("Test Start %v", 1)

	// laptop := sample.NewLaptop()

	// create grpc server
	req := &pb.GetInventoryItemRequest{Id: 25}
	invItem, e := testInventory.GetInventoryItem(context.Background(), req)
	require.NoError(t, e)

	log.Printf("TestLaptop2InventoryItem TestInventoryItem:  %+v", invItem)

	require.True(t, false)
}
