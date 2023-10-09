package db

import (
	"context"
)

const kPLUSCallBackRefSQL = `-- name: kPLUSCallBackRefSQL :one
SELECT trn_serial 
FROM trn_head th
INNER JOIN Users u on u.id = th.user_id 
WHERE reference  = $1 and u.login_name = 'konek2CARD'
`

func (q *QueriesKPlus) CallBackRef(ctx context.Context, prNo string) (KPLUSResponse, error) {
	var i KPLUSResponse
	var serial string
	row := q.db.QueryRowContext(ctx, kPLUSCallBackRefSQL, prNo)
	err := row.Scan(
		&serial,
	)
	if err == nil {
		i.RetCode = 1001
		i.Message = "Transaction Exist!"
		i.Reference = serial
	} else {
		i.RetCode = 0
		i.Message = "Transaction not Exist!"
		i.Reference = ""
		err = nil
	}
	return i, err

}
