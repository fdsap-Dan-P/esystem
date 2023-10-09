package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodeIncrement(t *testing.T) {

	// Test Data
	trn, err := testQueriesLocal.CodeIncrement(context.Background(), "saTrnMaster", `1999-01-02`)
	require.NoError(t, err)
	require.Greater(t, trn, int64(0))
}
