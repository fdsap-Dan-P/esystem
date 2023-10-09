package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const getColSht = `-- name: GetColSht :one
SELECT 
  OrgParms.DefBranch_Code, AppType, Code, cast(Status as int) Status, 
  OrgParms.DefBranch_Code + CASE WHEN Code = 0 THEN 'MBA' + Cast(CID as VarChar(10)) ELSE Acc END Acc, CID, 
  UM, ClientName, Center_Code, Center_Name, ManCode, 
  Unit, AreaCode, Area, StaffName, AcctType, 
  AcctDesc, DisbDate, DateStart, Maturity, Principal, 
  Interest, Gives, BalPrin, BalInt, Amort, 
  DuePrin, DueInt, LoanBal, SaveBal, WaivedInt, 
  UnPaidCtr, WritenOff, c.OrgName, c.OrgAddress, MeetingDate, 
  MeetingDay, SharesOfStock, DateEstablished, Classification, WriteOff
FROM colsht c, Orgparms
`

func scanRowColSht(row *sql.Row) (ColSht, error) {
	var i ColSht
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

func scanRowsColSht(rows *sql.Rows) ([]ColSht, error) {
	items := []ColSht{}
	for rows.Next() {
		var i ColSht
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

func (q *QueriesLocal) GetColSht(ctx context.Context) ([]ColSht, error) {
	rows, err := q.db.QueryContext(ctx, getColSht)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsColSht(rows)
}

func (q *QueriesLocal) GetColShtPerCID(ctx context.Context, cid int64) ([]ColSht, error) {
	sql := fmt.Sprintf(
		`%v WHERE cid = $1`,
		getColSht)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsColSht(rows)
}

func (q *QueriesLocal) ColShtCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, getColSht, filenamePath)
}
