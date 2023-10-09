package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	model "simplebank/db/datastore/esystemlocal"
)

const createColSht = `-- name: CreateColSht: one
INSERT INTO esystemdump.ColSht(
  BrCode, AppType, Code, Status, Acc, CID, 
  UM, ClientName, CenterCode, CenterName, ManCode, 
  Unit, AreaCode, Area, StaffName, AcctType, 
  AcctDesc, DisbDate, DateStart, Maturity, Principal, 
  Interest, Gives, BalPrin, BalInt, Amort, 
  DuePrin, DueInt, LoanBal, SaveBal, WaivedInt, 
  UnPaidCtr, WrittenOff, OrgName, OrgAddress, MeetingDate, 
  MeetingDay, SharesOfStock, DateEstablished, Classification, WriteOff
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, 
  $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41
)
ON CONFLICT (Acc)
DO UPDATE SET
  BrCode =  EXCLUDED.BrCode,
  AppType =  EXCLUDED.AppType,
  Code =  EXCLUDED.Code,
  Status =  EXCLUDED.Status,
  CID =  EXCLUDED.CID,
  UM =  EXCLUDED.UM,
  ClientName =  EXCLUDED.ClientName,
  CenterCode =  EXCLUDED.CenterCode,
  CenterName =  EXCLUDED.CenterName,
  ManCode =  EXCLUDED.ManCode,
  Unit =  EXCLUDED.Unit,
  AreaCode =  EXCLUDED.AreaCode,
  Area =  EXCLUDED.Area,
  StaffName =  EXCLUDED.StaffName,
  AcctType =  EXCLUDED.AcctType,
  AcctDesc =  EXCLUDED.AcctDesc,
  DisbDate =  EXCLUDED.DisbDate,
  DateStart =  EXCLUDED.DateStart,
  Maturity =  EXCLUDED.Maturity,
  Principal =  EXCLUDED.Principal,
  Interest =  EXCLUDED.Interest,
  Gives =  EXCLUDED.Gives,
  BalPrin =  EXCLUDED.BalPrin,
  BalInt =  EXCLUDED.BalInt,
  Amort =  EXCLUDED.Amort,
  DuePrin =  EXCLUDED.DuePrin,
  DueInt =  EXCLUDED.DueInt,
  LoanBal =  EXCLUDED.LoanBal,
  SaveBal =  EXCLUDED.SaveBal,
  WaivedInt =  EXCLUDED.WaivedInt,
  UnPaidCtr =  EXCLUDED.UnPaidCtr,
  WrittenOff =  EXCLUDED.WrittenOff,
  OrgName =  EXCLUDED.OrgName,
  OrgAddress =  EXCLUDED.OrgAddress,
  MeetingDate =  EXCLUDED.MeetingDate,
  MeetingDay =  EXCLUDED.MeetingDay,
  SharesOfStock =  EXCLUDED.SharesOfStock,
  DateEstablished =  EXCLUDED.DateEstablished,
  Classification =  EXCLUDED.Classification,
  WriteOff =  EXCLUDED.WriteOff
`

func (q *QueriesDump) CreateColSht(ctx context.Context, arg model.ColSht) error {
	_, err := q.db.ExecContext(ctx, createColSht,
		arg.BrCode,
		arg.AppType,
		arg.Code,
		arg.Status,
		arg.Acc,
		arg.CID,
		arg.UM,
		arg.ClientName,
		arg.CenterCode,
		arg.CenterName,
		arg.ManCode,
		arg.Unit,
		arg.AreaCode,
		arg.Area,
		arg.StaffName,
		arg.AcctType,
		arg.AcctDesc,
		arg.DisbDate,
		arg.DateStart,
		arg.Maturity,
		arg.Principal,
		arg.Interest,
		arg.Gives,
		arg.BalPrin,
		arg.BalInt,
		arg.Amort,
		arg.DuePrin,
		arg.DueInt,
		arg.LoanBal,
		arg.SaveBal,
		arg.WaivedInt,
		arg.UnPaidCtr,
		arg.WrittenOff,
		arg.OrgName,
		arg.OrgAddress,
		arg.MeetingDate,
		arg.MeetingDay,
		arg.SharesOfStock,
		arg.DateEstablished,
		arg.Classification,
		arg.WriteOff,
	)
	return err
}

const deleteColSht = `-- name: DeleteColSht :exec
  DELETE FROM esystemdump.ColSht WHERE BrCode = $1
`

func (q *QueriesDump) DeleteColSht(ctx context.Context, brCode string) error {
	_, err := q.db.ExecContext(ctx, deleteColSht, brCode)
	return err
}

const getColSht = `-- name: GetColSht :one
SELECT
  BrCode, AppType, Code, Status, Acc, CID, 
  UM, ClientName, CenterCode, CenterName, ManCode, 
  Unit, AreaCode, Area, StaffName, AcctType, 
  AcctDesc, DisbDate, DateStart, Maturity, Principal, 
  Interest, Gives, BalPrin, BalInt, Amort, 
  DuePrin, DueInt, LoanBal, SaveBal, WaivedInt, 
  UnPaidCtr, WrittenOff, OrgName, OrgAddress, MeetingDate, 
  MeetingDay, SharesOfStock, DateEstablished, Classification, WriteOff
FROM esystemdump.ColSht
`

func scanRowColSht(row *sql.Row) (model.ColSht, error) {
	var i model.ColSht
	err := row.Scan(
		&i.BrCode,
		&i.AppType,
		&i.Code,
		&i.Status,
		&i.Acc,
		&i.CID,
		&i.UM,
		&i.ClientName,
		&i.CenterCode,
		&i.CenterName,
		&i.ManCode,
		&i.Unit,
		&i.AreaCode,
		&i.Area,
		&i.StaffName,
		&i.AcctType,
		&i.AcctDesc,
		&i.DisbDate,
		&i.DateStart,
		&i.Maturity,
		&i.Principal,
		&i.Interest,
		&i.Gives,
		&i.BalPrin,
		&i.BalInt,
		&i.Amort,
		&i.DuePrin,
		&i.DueInt,
		&i.LoanBal,
		&i.SaveBal,
		&i.WaivedInt,
		&i.UnPaidCtr,
		&i.WrittenOff,
		&i.OrgName,
		&i.OrgAddress,
		&i.MeetingDate,
		&i.MeetingDay,
		&i.SharesOfStock,
		&i.DateEstablished,
		&i.Classification,
		&i.WriteOff,
	)
	return i, err
}

func scanRowsColSht(rows *sql.Rows) ([]model.ColSht, error) {
	items := []model.ColSht{}
	for rows.Next() {
		var i model.ColSht
		if err := rows.Scan(
			&i.BrCode,
			&i.AppType,
			&i.Code,
			&i.Status,
			&i.Acc,
			&i.CID,
			&i.UM,
			&i.ClientName,
			&i.CenterCode,
			&i.CenterName,
			&i.ManCode,
			&i.Unit,
			&i.AreaCode,
			&i.Area,
			&i.StaffName,
			&i.AcctType,
			&i.AcctDesc,
			&i.DisbDate,
			&i.DateStart,
			&i.Maturity,
			&i.Principal,
			&i.Interest,
			&i.Gives,
			&i.BalPrin,
			&i.BalInt,
			&i.Amort,
			&i.DuePrin,
			&i.DueInt,
			&i.LoanBal,
			&i.SaveBal,
			&i.WaivedInt,
			&i.UnPaidCtr,
			&i.WrittenOff,
			&i.OrgName,
			&i.OrgAddress,
			&i.MeetingDate,
			&i.MeetingDay,
			&i.SharesOfStock,
			&i.DateEstablished,
			&i.Classification,
			&i.WriteOff,
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

func (q *QueriesDump) GetColSht(ctx context.Context, acc string) (model.ColSht, error) {
	sql := fmt.Sprintf("%s WHERE Acc = $1", getColSht)
	row := q.db.QueryRowContext(ctx, sql, acc)
	return scanRowColSht(row)
}

type ListColShtParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ColShtPerBranch(ctx context.Context, brCode string) ([]model.ColSht, error) {
	sql := fmt.Sprintf(`%v WHERE BrCode = $1`, getColSht)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, brCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsColSht(rows)
}

func (q *QueriesDump) ColShtPerCID(ctx context.Context, cid int64) ([]model.ColSht, error) {
	sql := fmt.Sprintf(`%v WHERE cid = $1`, getColSht)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsColSht(rows)
}

func (q *QueriesDump) ColShtPerCenter(ctx context.Context, cenCode string) ([]model.ColSht, error) {
	sql := fmt.Sprintf(`%v WHERE CenterCode = $1`, getColSht)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, cenCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsColSht(rows)
}

const updateColSht = `-- name: UpdateColSht :one
UPDATE esystemdump.ColSht SET 
  BrCode = $1,
  AppType = $2,
  Code = $3,
  Status = $4,
  Acc = $5,
  CID = $6,
  UM = $7,
  ClientName = $8,
  CenterCode = $9,
  CenterName = $10,
  ManCode = $11,
  Unit = $12,
  AreaCode = $13,
  Area = $14,
  StaffName = $15,
  AcctType = $16,
  AcctDesc = $17,
  DisbDate = $18,
  DateStart = $19,
  Maturity = $20,
  Principal = $21,
  Interest = $22,
  Gives = $23,
  BalPrin = $24,
  BalInt = $25,
  Amort = $26,
  DuePrin = $27,
  DueInt = $28,
  LoanBal = $29,
  SaveBal = $30,
  WaivedInt = $31,
  UnPaidCtr = $32,
  WrittenOff = $33,
  OrgName = $34,
  OrgAddress = $35,
  MeetingDate = $36,
  MeetingDay = $37,
  SharesOfStock = $38,
  DateEstablished = $39,
  Classification = $40,
  WriteOff = $41
WHERE Acc = $5
`

func (q *QueriesDump) UpdateColSht(ctx context.Context, arg model.ColSht) error {
	_, err := q.db.ExecContext(ctx, updateColSht,
		arg.BrCode,
		arg.AppType,
		arg.Code,
		arg.Status,
		arg.Acc,
		arg.CID,
		arg.UM,
		arg.ClientName,
		arg.CenterCode,
		arg.CenterName,
		arg.ManCode,
		arg.Unit,
		arg.AreaCode,
		arg.Area,
		arg.StaffName,
		arg.AcctType,
		arg.AcctDesc,
		arg.DisbDate,
		arg.DateStart,
		arg.Maturity,
		arg.Principal,
		arg.Interest,
		arg.Gives,
		arg.BalPrin,
		arg.BalInt,
		arg.Amort,
		arg.DuePrin,
		arg.DueInt,
		arg.LoanBal,
		arg.SaveBal,
		arg.WaivedInt,
		arg.UnPaidCtr,
		arg.WrittenOff,
		arg.OrgName,
		arg.OrgAddress,
		arg.MeetingDate,
		arg.MeetingDay,
		arg.SharesOfStock,
		arg.DateEstablished,
		arg.Classification,
		arg.WriteOff,
	)
	return err
}

func (q *QueriesDump) BulkInsertColSht(ctx context.Context, rows []model.ColSht) error {
	colCtr := 41
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*colCtr)
	i := 0
	for _, post := range rows {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*colCtr+1, i*colCtr+2, i*colCtr+3, i*colCtr+4, i*colCtr+5, i*colCtr+6,
			i*colCtr+7, i*colCtr+8, i*colCtr+9, i*colCtr+10, i*colCtr+11,
			i*colCtr+12, i*colCtr+13, i*colCtr+14, i*colCtr+15, i*colCtr+16,
			i*colCtr+17, i*colCtr+18, i*colCtr+19, i*colCtr+20, i*colCtr+21,
			i*colCtr+22, i*colCtr+23, i*colCtr+24, i*colCtr+25, i*colCtr+26,
			i*colCtr+27, i*colCtr+28, i*colCtr+29, i*colCtr+30, i*colCtr+31,
			i*colCtr+32, i*colCtr+33, i*colCtr+34, i*colCtr+35, i*colCtr+36,
			i*colCtr+37, i*colCtr+38, i*colCtr+39, i*colCtr+40, i*colCtr+41))
		valueArgs = append(valueArgs, post.BrCode)
		valueArgs = append(valueArgs, post.AppType)
		valueArgs = append(valueArgs, post.Code)
		valueArgs = append(valueArgs, post.Status)
		valueArgs = append(valueArgs, post.Acc)
		valueArgs = append(valueArgs, post.CID)
		valueArgs = append(valueArgs, post.UM)
		valueArgs = append(valueArgs, post.ClientName)
		valueArgs = append(valueArgs, post.CenterCode)
		valueArgs = append(valueArgs, post.CenterName)
		valueArgs = append(valueArgs, post.ManCode)
		valueArgs = append(valueArgs, post.Unit)
		valueArgs = append(valueArgs, post.AreaCode)
		valueArgs = append(valueArgs, post.Area)
		valueArgs = append(valueArgs, post.StaffName)
		valueArgs = append(valueArgs, post.AcctType)
		valueArgs = append(valueArgs, post.AcctDesc)
		valueArgs = append(valueArgs, post.DisbDate)
		valueArgs = append(valueArgs, post.DateStart)
		valueArgs = append(valueArgs, post.Maturity)
		valueArgs = append(valueArgs, post.Principal)
		valueArgs = append(valueArgs, post.Interest)
		valueArgs = append(valueArgs, post.Gives)
		valueArgs = append(valueArgs, post.BalPrin)
		valueArgs = append(valueArgs, post.BalInt)
		valueArgs = append(valueArgs, post.Amort)
		valueArgs = append(valueArgs, post.DuePrin)
		valueArgs = append(valueArgs, post.DueInt)
		valueArgs = append(valueArgs, post.LoanBal)
		valueArgs = append(valueArgs, post.SaveBal)
		valueArgs = append(valueArgs, post.WaivedInt)
		valueArgs = append(valueArgs, post.UnPaidCtr)
		valueArgs = append(valueArgs, post.WrittenOff)
		valueArgs = append(valueArgs, post.OrgName)
		valueArgs = append(valueArgs, post.OrgAddress)
		valueArgs = append(valueArgs, post.MeetingDate)
		valueArgs = append(valueArgs, post.MeetingDay)
		valueArgs = append(valueArgs, post.SharesOfStock)
		valueArgs = append(valueArgs, post.DateEstablished)
		valueArgs = append(valueArgs, post.Classification)
		valueArgs = append(valueArgs, post.WriteOff)
		i++
	}
	stmt := fmt.Sprintf(`
	INSERT INTO esystemdump.ColSht(
		BrCode, AppType, Code, Status, Acc, CID, 
		UM, ClientName, CenterCode, CenterName, ManCode, 
		Unit, AreaCode, Area, StaffName, AcctType, 
		AcctDesc, DisbDate, DateStart, Maturity, Principal, 
		Interest, Gives, BalPrin, BalInt, Amort, 
		DuePrin, DueInt, LoanBal, SaveBal, WaivedInt, 
		UnPaidCtr, WrittenOff, OrgName, OrgAddress, MeetingDate, 
		MeetingDay, SharesOfStock, DateEstablished, Classification, WriteOff
	  )
	  VALUES %s
	  ON CONFLICT (Acc)
	  DO UPDATE SET
		BrCode =  EXCLUDED.BrCode,
		AppType =  EXCLUDED.AppType,
		Code =  EXCLUDED.Code,
		Status =  EXCLUDED.Status,
		CID =  EXCLUDED.CID,
		UM =  EXCLUDED.UM,
		ClientName =  EXCLUDED.ClientName,
		CenterCode =  EXCLUDED.CenterCode,
		CenterName =  EXCLUDED.CenterName,
		ManCode =  EXCLUDED.ManCode,
		Unit =  EXCLUDED.Unit,
		AreaCode =  EXCLUDED.AreaCode,
		Area =  EXCLUDED.Area,
		StaffName =  EXCLUDED.StaffName,
		AcctType =  EXCLUDED.AcctType,
		AcctDesc =  EXCLUDED.AcctDesc,
		DisbDate =  EXCLUDED.DisbDate,
		DateStart =  EXCLUDED.DateStart,
		Maturity =  EXCLUDED.Maturity,
		Principal =  EXCLUDED.Principal,
		Interest =  EXCLUDED.Interest,
		Gives =  EXCLUDED.Gives,
		BalPrin =  EXCLUDED.BalPrin,
		BalInt =  EXCLUDED.BalInt,
		Amort =  EXCLUDED.Amort,
		DuePrin =  EXCLUDED.DuePrin,
		DueInt =  EXCLUDED.DueInt,
		LoanBal =  EXCLUDED.LoanBal,
		SaveBal =  EXCLUDED.SaveBal,
		WaivedInt =  EXCLUDED.WaivedInt,
		UnPaidCtr =  EXCLUDED.UnPaidCtr,
		WrittenOff =  EXCLUDED.WrittenOff,
		OrgName =  EXCLUDED.OrgName,
		OrgAddress =  EXCLUDED.OrgAddress,
		MeetingDate =  EXCLUDED.MeetingDate,
		MeetingDay =  EXCLUDED.MeetingDay,
		SharesOfStock =  EXCLUDED.SharesOfStock,
		DateEstablished =  EXCLUDED.DateEstablished,
		Classification =  EXCLUDED.Classification,
		WriteOff =  EXCLUDED.WriteOff;`, strings.Join(valueStrings, ","))

	// log.Println(stmt)
	_, err := q.db.ExecContext(ctx, stmt, valueArgs...)

	return err
}
