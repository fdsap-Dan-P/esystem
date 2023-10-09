package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createAddresses = `-- name: CreateAddresses: one
INSERT INTO Addresses (
	CID, Address1, Address3, CityTown, StateProv, Phone1, Phone2, Phone3, Phone4
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
SELECT SCOPE_IDENTITY()
`

type AddressesRequest struct {
	CID            int64          `json:"CID"`
	SeqNum         int64          `json:"seqNum"`
	AddressDetails sql.NullString `json:"addressDetails"`
	Barangay       sql.NullString `json:"barangay"`
	City           sql.NullString `json:"city"`
	Province       sql.NullString `json:"province"`
	Phone1         sql.NullString `json:"phone1"`
	Phone2         sql.NullString `json:"phone2"`
	Phone3         sql.NullString `json:"phone3"`
	Phone4         sql.NullString `json:"phone4"`
}

func (q *QueriesLocal) CreateAddresses(ctx context.Context, arg AddressesRequest) (int64, error) {
	var id int64
	result := q.db.QueryRowContext(ctx, createAddresses,
		arg.CID,
		arg.AddressDetails,
		arg.Barangay,
		arg.City,
		arg.Province,
		arg.Phone1,
		arg.Phone2,
		arg.Phone3,
		arg.Phone4,
	)
	err := result.Scan(&id)
	return id, err
}

const deleteAddresses = `-- name: DeleteAddresses :exec
DELETE FROM Addresses WHERE SeqNum = $1
`

func (q *QueriesLocal) DeleteAddresses(ctx context.Context, seqNum int64) error {
	_, err := q.db.ExecContext(ctx, deleteAddresses, seqNum)
	return err
}

type AddressesInfo struct {
	ModCtr         int64          `json:"modCtr"`
	BrCode         string         `json:"brCode"`
	ModAction      string         `json:"modAction"`
	CID            int64          `json:"CID"`
	SeqNum         int64          `json:"seqNum"`
	AddressDetails sql.NullString `json:"addressDetails"`
	Barangay       sql.NullString `json:"barangay"`
	City           sql.NullString `json:"city"`
	Province       sql.NullString `json:"province"`
	Phone1         sql.NullString `json:"phone1"`
	Phone2         sql.NullString `json:"phone2"`
	Phone3         sql.NullString `json:"phone3"`
	Phone4         sql.NullString `json:"phone4"`
}

// -- name: GetAddresses :one
const getAddresses = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, CID, SeqNum, Address1, Address3, CityTown, StateProv, Phone1, Phone2, Phone3, Phone4
FROM OrgParms, Addresses d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.SeqNum 
`

func scanRowAddresses(row *sql.Row) (AddressesInfo, error) {
	var i AddressesInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
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

func scanRowsAddresses(rows *sql.Rows) ([]AddressesInfo, error) {
	items := []AddressesInfo{}
	for rows.Next() {
		var i AddressesInfo
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

func (q *QueriesLocal) GetAddresses(ctx context.Context, seqNum int64) (AddressesInfo, error) {
	sql := fmt.Sprintf("%s WHERE  m.TableName = 'Addresses' AND Uploaded = 0 and SeqNum = $1", getAddresses)
	row := q.db.QueryRowContext(ctx, sql, seqNum)
	return scanRowAddresses(row)
}

type ListAddressesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) AddressesCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, CID, SeqNum, Address1, Address3, CityTown, StateProv, Phone1, Phone2, Phone3, Phone4
FROM OrgParms, Addresses d
`, filenamePath)
}

func (q *QueriesLocal) ListAddresses(ctx context.Context) ([]AddressesInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE  m.TableName = 'Addresses' AND Uploaded = 0`,
		getAddresses)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsAddresses(rows)
}

// -- name: UpdateAddresses :one
const updateAddresses = `
UPDATE Addresses SET 
  CID = $2,
  Address1 = $3,
  Address3 = $4,
  CityTown = $5,
  StateProv = $6,
  Phone1 = $7,
  Phone2 = $8,
  Phone3 = $9,
  Phone4 = $10
WHERE SeqNum = $1`

func (q *QueriesLocal) UpdateAddresses(ctx context.Context, arg AddressesRequest) error {
	_, err := q.db.ExecContext(ctx, updateAddresses,
		arg.SeqNum,
		arg.CID,
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
