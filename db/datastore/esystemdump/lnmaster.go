package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createLnMaster = `-- name: CreateLnMaster: one
INSERT INTO esystemdump.LnMaster(
   ModCtr, BrCode, ModAction, CID, Acc, AcctType, DisbDate, Principal, Interest, NetProceed, 
   Gives, Frequency, AnnumDiv, Prin, IntR, WaivedInt, WeeksPaid, DoMaturity, ConIntRate, Status, 
   Cycle, LNGrpCode, Proff, FundSource, DOSRI, LnCategory, OpenDate, LastTrnDate, DisbBy )
VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
	$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
	$21, $22, $23, $24, $25, $26, $27, $28, $29)
ON CONFLICT (brCode, acc, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	CID =  EXCLUDED.CID,
	AcctType =  EXCLUDED.AcctType,
	DisbDate =  EXCLUDED.DisbDate,
	Principal =  EXCLUDED.Principal,
	Interest =  EXCLUDED.Interest,
	NetProceed =  EXCLUDED.NetProceed,
	Gives =  EXCLUDED.Gives,
	Frequency =  EXCLUDED.Frequency,
	AnnumDiv =  EXCLUDED.AnnumDiv,
	Prin =  EXCLUDED.Prin,
	IntR =  EXCLUDED.IntR,
	WaivedInt =  EXCLUDED.WaivedInt,
	WeeksPaid =  EXCLUDED.WeeksPaid,
	DoMaturity =  EXCLUDED.DoMaturity,
	ConIntRate =  EXCLUDED.ConIntRate,
	Status =  EXCLUDED.Status,
	Cycle =  EXCLUDED.Cycle,
	LNGrpCode =  EXCLUDED.LNGrpCode,
	Proff =  EXCLUDED.Proff,
	FundSource =  EXCLUDED.FundSource,
	DOSRI =  EXCLUDED.DOSRI,
	LnCategory =  EXCLUDED.LnCategory,
	OpenDate =  EXCLUDED.OpenDate,
	LastTrnDate =  EXCLUDED.LastTrnDate,
	DisbBy =  EXCLUDED.DisbBy
`

func (q *QueriesDump) CreateLnMaster(ctx context.Context, arg model.LnMaster) error {
	_, err := q.db.ExecContext(ctx, createLnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.Acc,
		arg.AcctType,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.NetProceed,
		arg.Gives,
		arg.Frequency,
		arg.AnnumDiv,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.WeeksPaid,
		arg.DoMaturity,
		arg.ConIntRate,
		arg.Status,
		arg.Cycle,
		arg.LNGrpCode,
		arg.Proff,
		arg.FundSource,
		arg.DOSRI,
		arg.LnCategory,
		arg.OpenDate,
		arg.LastTrnDate,
		arg.DisbBy,
	)
	return err
}

const deleteLnMaster = `-- name: DeleteLnMaster :exec
DELETE FROM esystemdump.LnMaster WHERE BrCode = $1 and Acc = $2
`

func (q *QueriesDump) DeleteLnMaster(ctx context.Context, brCode string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnMaster, brCode, acc)
	return err
}

const getLnMaster = `-- name: GetLnMaster :one
SELECT
ModCtr, BrCode, ModAction, CID, Acc, AcctType, DisbDate, Principal, Interest, NetProceed, Gives, Frequency, AnnumDiv, Prin, IntR, WaivedInt, WeeksPaid, DoMaturity, ConIntRate, Status, Cycle, LNGrpCode, Proff, FundSource, DOSRI, LnCategory, OpenDate, LastTrnDate, DisbBy
FROM esystemdump.LnMaster
`

func scanRowLnMaster(row *sql.Row) (model.LnMaster, error) {
	var i model.LnMaster
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.Acc,
		&i.AcctType,
		&i.DisbDate,
		&i.Principal,
		&i.Interest,
		&i.NetProceed,
		&i.Gives,
		&i.Frequency,
		&i.AnnumDiv,
		&i.Prin,
		&i.IntR,
		&i.WaivedInt,
		&i.WeeksPaid,
		&i.DoMaturity,
		&i.ConIntRate,
		&i.Status,
		&i.Cycle,
		&i.LNGrpCode,
		&i.Proff,
		&i.FundSource,
		&i.DOSRI,
		&i.LnCategory,
		&i.OpenDate,
		&i.LastTrnDate,
		&i.DisbBy,
	)
	return i, err
}

func scanRowsLnMaster(rows *sql.Rows) ([]model.LnMaster, error) {
	items := []model.LnMaster{}
	for rows.Next() {
		var i model.LnMaster
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.Acc,
			&i.AcctType,
			&i.DisbDate,
			&i.Principal,
			&i.Interest,
			&i.NetProceed,
			&i.Gives,
			&i.Frequency,
			&i.AnnumDiv,
			&i.Prin,
			&i.IntR,
			&i.WaivedInt,
			&i.WeeksPaid,
			&i.DoMaturity,
			&i.ConIntRate,
			&i.Status,
			&i.Cycle,
			&i.LNGrpCode,
			&i.Proff,
			&i.FundSource,
			&i.DOSRI,
			&i.LnCategory,
			&i.OpenDate,
			&i.LastTrnDate,
			&i.DisbBy,
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

func (q *QueriesDump) GetLnMaster(ctx context.Context, brCode string, acc string) (model.LnMaster, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2", getLnMaster)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc)
	return scanRowLnMaster(row)
}

type ListLnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListLnMaster(ctx context.Context, lastModCtr int64) ([]model.LnMaster, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getLnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnMaster(rows)
}

const updateLnMaster = `-- name: UpdateLnMaster :one
UPDATE esystemdump.LnMaster SET 
	ModCtr = $1,
	CID = $4,
	AcctType = $6,
	DisbDate = $7,
	Principal = $8,
	Interest = $9,
	NetProceed = $10,
	Gives = $11,
	Frequency = $12,
	AnnumDiv = $13,
	Prin = $14,
	IntR = $15,
	WaivedInt = $16,
	WeeksPaid = $17,
	DoMaturity = $18,
	ConIntRate = $19,
	Status = $20,
	Cycle = $21,
	LNGrpCode = $22,
	Proff = $23,
	FundSource = $24,
	DOSRI = $25,
	LnCategory = $26,
	OpenDate = $27,
	LastTrnDate = $28,
	DisbBy = $29
WHERE BrCode = $2 and Acc = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateLnMaster(ctx context.Context, arg model.LnMaster) error {
	_, err := q.db.ExecContext(ctx, updateLnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.Acc,
		arg.AcctType,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.NetProceed,
		arg.Gives,
		arg.Frequency,
		arg.AnnumDiv,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.WeeksPaid,
		arg.DoMaturity,
		arg.ConIntRate,
		arg.Status,
		arg.Cycle,
		arg.LNGrpCode,
		arg.Proff,
		arg.FundSource,
		arg.DOSRI,
		arg.LnCategory,
		arg.OpenDate,
		arg.LastTrnDate,
		arg.DisbBy,
	)
	return err
}
