package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createMultiplePaymentReceipt = `-- name: CreateMultiplePaymentReceipt: one
INSERT INTO MultiplePaymentReceipt(
	ClientCode, ClientCID, Remarks, AmtPaid, PaymentDate, 
	PayingrepresentativeCID, Acc, OrNo, UserName, TermID, 
	Cancelled, Trn, PrNo
	)
  VALUES(
	'A', $3, 'Multiple Payment', $7, $1, 
	$3, '', $2, 'sa', $6, 
	0, 0, $4)
;
`

type MultiplePaymentReceiptRequest struct {
	TrnDate  time.Time       `json:"trnDate"`
	OrNo     int64           `json:"orNo"`
	CID      int64           `json:"cID"`
	PrNo     int64           `json:"prNo"`
	UserName string          `json:"userName"`
	TermId   string          `json:"termId"`
	AmtPaid  decimal.Decimal `json:"amtPaid"`
}

func (q *QueriesLocal) CreateMultiplePaymentReceipt(ctx context.Context, arg MultiplePaymentReceiptRequest) error {
	_, err := q.db.ExecContext(ctx, createMultiplePaymentReceipt,
		arg.TrnDate,
		arg.OrNo,
		arg.CID,
		arg.PrNo,
		arg.UserName,
		arg.TermId,
		arg.AmtPaid,
	)
	return err
}

const deleteMultiplePaymentReceipt = `-- name: DeleteMultiplePaymentReceipt :exec
DELETE FROM MultiplePaymentReceipt WHERE OrNo = $1
`

func (q *QueriesLocal) DeleteMultiplePaymentReceipt(ctx context.Context, orNo int64) error {
	_, err := q.db.ExecContext(ctx, deleteMultiplePaymentReceipt, orNo)
	return err
}

const getMultiplePaymentReceipt = `-- name: GetMultiplePaymentReceipt :one
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, PaymentDate TrnDate, OrNo, CID, PrNo, UserName, TermId, AmtPaid
FROM OrgParms, 
 (SELECT
    PaymentDate, OrNo, max(PayingRepresentativeCID) CID, Max(PrNo) PrNo, trim(UserName) UserName, Max(trim(TermId)) TermId, Sum(AmtPaid) AmtPaid
  FROM MultiplePaymentReceipt
  GROUP BY PaymentDate, OrNo, UserName
) d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.OrNo 
`

func scanRowMultiplePaymentReceipt(row *sql.Row) (MultiplePaymentReceipt, error) {
	var i MultiplePaymentReceipt
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.TrnDate,
		&i.OrNo,
		&i.CID,
		&i.PrNo,
		&i.UserName,
		&i.TermId,
		&i.AmtPaid,
	)
	return i, err
}

func scanRowsMultiplePaymentReceipt(rows *sql.Rows) ([]MultiplePaymentReceipt, error) {
	items := []MultiplePaymentReceipt{}
	for rows.Next() {
		var i MultiplePaymentReceipt
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.TrnDate,
			&i.OrNo,
			&i.CID,
			&i.PrNo,
			&i.UserName,
			&i.TermId,
			&i.AmtPaid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *QueriesLocal) GetMultiplePaymentReceipt(ctx context.Context, orNo int64) (MultiplePaymentReceipt, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'MultiplePaymentReceipt' AND Uploaded = 0 and OrNo = $1", getMultiplePaymentReceipt)
	row := q.db.QueryRowContext(ctx, sql, orNo)
	return scanRowMultiplePaymentReceipt(row)
}

type ListMultiplePaymentReceiptParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ListMultiplePaymentReceipt(ctx context.Context) ([]MultiplePaymentReceipt, error) {
	sql := fmt.Sprintf(`%v WHERE m.TableName = 'MultiplePaymentReceipt' AND Uploaded = 0`, getMultiplePaymentReceipt)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsMultiplePaymentReceipt(rows)
}

func (q *QueriesLocal) MultiplePaymentReceiptCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
  SELECT 
	0 ModCtr, OrgParms.DefBranch_Code BrCode, PaymentDate TrnDate, OrNo, CID, PrNo, UserName, TermId, AmtPaid
  FROM OrgParms, 
   (SELECT
	PaymentDate, OrNo, max(PayingRepresentativeCID) CID, Max(PrNo) PrNo, trim(UserName) UserName, Max(trim(TermId)) TermId, Sum(AmtPaid) AmtPaid
	FROM MultiplePaymentReceipt
	GROUP BY PaymentDate, OrNo, UserName
  ) d
`, filenamePath)
}

const updateMultiplePaymentReceipt = `-- name: UpdateMultiplePaymentReceipt :one
UPDATE MultiplePaymentReceipt SET 
PaymentDate = $1,
  OrNo = $2,
  ClientCID = $3,
  PayingrepresentativeCID = $3,
  PrNo = $4,
  UserName = $5,
  TermId = $6,
  AmtPaid = $7
WHERE OrNo = $2
`

func (q *QueriesLocal) UpdateMultiplePaymentReceipt(ctx context.Context, arg MultiplePaymentReceiptRequest) error {
	_, err := q.db.ExecContext(ctx, updateMultiplePaymentReceipt,
		arg.TrnDate,
		arg.OrNo,
		arg.CID,
		arg.PrNo,
		arg.UserName,
		arg.TermId,
		arg.AmtPaid,
	)
	return err
}
