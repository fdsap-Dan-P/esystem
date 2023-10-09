package db

import (
	"context"
	"log"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestColSht(t *testing.T) {
	// Get Data
	var cid = int64(400200)
	getData1, err := testQueriesLocal.GetColShtPerCID(context.Background(), cid)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)
	require.Equal(t, cid, getData1[0].CID)
	log.Printf("ColSht: %v", getData1)
	// require.Equal(t, 0, 1)
}
