package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createAddresses = `-- name: CreateAddresses: one
INSERT INTO esystemdump.Addresses(
   ModCtr, BrCode, ModAction, CID, SeqNum, AddressDetails, Barangay, City, Province, Phone1, Phone2, Phone3, Phone4 )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
ON CONFLICT (brCode, seqNum, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  CID =  EXCLUDED.CID,
  AddressDetails =  EXCLUDED.AddressDetails,
  Barangay =  EXCLUDED.Barangay,
  City =  EXCLUDED.City,
  Province =  EXCLUDED.Province,
  Phone1 =  EXCLUDED.Phone1,
  Phone2 =  EXCLUDED.Phone2,
  Phone3 =  EXCLUDED.Phone3,
  Phone4 =  EXCLUDED.Phone4
`

func (q *QueriesDump) CreateAddresses(ctx context.Context, arg model.Addresses) error {
	_, err := q.db.ExecContext(ctx, createAddresses,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.SeqNum,
		arg.AddressDetails,
		arg.Barangay,
		arg.City,
		arg.Province,
		arg.Phone1,
		arg.Phone2,
		arg.Phone3,
		arg.Phone4,
	)
	return err
}

const deleteAddresses = `-- name: DeleteAddresses :exec
DELETE FROM esystemdump.Addresses WHERE BrCode = $1 and SeqNum = $2
`

func (q *QueriesDump) DeleteAddresses(ctx context.Context, brCode string, seqNum int64) error {
	_, err := q.db.ExecContext(ctx, deleteAddresses, brCode, seqNum)
	return err
}

const getAddresses = `-- name: GetAddresses :one
SELECT
  ModCtr, BrCode, ModAction, CID, SeqNum, AddressDetails, Barangay, City, Province, Phone1, Phone2, Phone3, Phone4
FROM esystemdump.Addresses
`

func scanRowAddresses(row *sql.Row) (model.Addresses, error) {
	var i model.Addresses
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.SeqNum,
		&i.AddressDetails,
		&i.Barangay,
		&i.City,
		&i.Province,
		&i.Phone1,
		&i.Phone2,
		&i.Phone3,
		&i.Phone4,
	)
	return i, err
}

func scanRowsAddresses(rows *sql.Rows) ([]model.Addresses, error) {
	items := []model.Addresses{}
	for rows.Next() {
		var i model.Addresses
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.SeqNum,
			&i.AddressDetails,
			&i.Barangay,
			&i.City,
			&i.Province,
			&i.Phone1,
			&i.Phone2,
			&i.Phone3,
			&i.Phone4,
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

func (q *QueriesDump) GetAddresses(ctx context.Context, brCode string, seqNum int64) (model.Addresses, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and SeqNum = $2", getAddresses)
	row := q.db.QueryRowContext(ctx, sql, brCode, seqNum)
	return scanRowAddresses(row)
}

type ListAddressesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListAddresses(ctx context.Context, lastModCtr int64) ([]model.Addresses, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getAddresses)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsAddresses(rows)
}

const updateAddresses = `-- name: UpdateAddresses :one
UPDATE esystemdump.Addresses SET 
	ModCtr = $1,
	CID = $4,
	AddressDetails = $6,
	Barangay = $7,
	City = $8,
	Province = $9,
	Phone1 = $10,
	Phone2 = $11,
	Phone3 = $12,
	Phone4 = $13
WHERE BrCode = $2 and ModAction = $3 and SeqNum = $5
`

func (q *QueriesDump) UpdateAddresses(ctx context.Context, arg model.Addresses) error {
	_, err := q.db.ExecContext(ctx, updateAddresses,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.SeqNum,
		arg.AddressDetails,
		arg.Barangay,
		arg.City,
		arg.Province,
		arg.Phone1,
		arg.Phone2,
		arg.Phone3,
		arg.Phone4,
	)
	return err
}
