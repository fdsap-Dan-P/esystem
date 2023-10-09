package db

import (
	"context"
	"log"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrgParms(t *testing.T) {

	// Get Data
	getData1, err1 := testQueriesLocal.GetOrgParms(context.Background())
	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	log.Printf("TestOrgParms %v", getData1)

	// require.True(t, false)

}
