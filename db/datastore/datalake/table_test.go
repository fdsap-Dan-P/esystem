package db

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {

	testListTable(t, ListTableParams{
		Limit:  5,
		Offset: 0,
	})

}

func testListTable(t *testing.T, arg ListTableParams) {

	table, err := testQueriesLocal.ListTable(context.Background(), "db2inst1")
	require.NoError(t, err)
	// fmt.Printf("%+v\n", table)
	require.NotEmpty(t, table)

}
