package db

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestModifiedTable(t *testing.T) {
	ModifiedTable, err := testQueriesDump.ListModifiedTable(context.Background(), "01")
	require.NoError(t, err)
	require.NotEmpty(t, ModifiedTable)
}
