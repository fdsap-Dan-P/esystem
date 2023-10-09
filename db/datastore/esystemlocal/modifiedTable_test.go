package db

import (
	"context"
	"log"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestModifiedTable(t *testing.T) {

	ModifiedTable, err := testQueriesLocal.ListModifiedTable(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, ModifiedTable)

	err = testQueriesLocal.UpdateModifiedTableUploaded(context.Background(), []int64{ModifiedTable[0].ModCtr}, true)
	require.NoError(t, err)

	d1, er := testQueriesLocal.GetModifiedTable(context.Background(), ModifiedTable[0].ModCtr)
	log.Printf("GetModifiedTable %v", d1)
	require.NoError(t, er)
	require.True(t, d1.Uploaded)

	err = testQueriesLocal.UpdateModifiedTableUploaded(context.Background(), []int64{ModifiedTable[0].ModCtr}, false)
	require.NoError(t, err)
	d1, er = testQueriesLocal.GetModifiedTable(context.Background(), ModifiedTable[0].ModCtr)
	require.NoError(t, er)
	require.False(t, d1.Uploaded)
}
