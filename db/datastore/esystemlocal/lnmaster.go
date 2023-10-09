package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

const createLnMaster = `-- name: CreateLnMaster: one
INSERT INTO LnMaster (
	CID, Acc, AcctType, DisbDate, 
	Principal, Interest, NetProceed, Gives, Frequency, AnnumDiv, Prin, IntR, 
	WaivedInt, WeeksPaid, DoMaturity, ConIntRate, Status, Cycle, LNGrpCode, 
	Proff, FundSource, DOSRI, LnCategory, DOPEN, DOLASTTRN, DisbBy
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, 
	$15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
`

type LnMasterRequest struct {
	CID         int64               `json:"CID"`
	Acc         string              `json:"acc"`
	AcctType    sql.NullInt64       `json:"acctType"`
	DisbDate    sql.NullTime        `json:"disbDate"`
	Principal   decimal.NullDecimal `json:"principal"`
	Interest    decimal.NullDecimal `json:"interest"`
	NetProceed  decimal.NullDecimal `json:"netProceed"`
	Gives       sql.NullInt64       `json:"gives"`
	Frequency   sql.NullInt64       `json:"frequency"`
	AnnumDiv    sql.NullInt64       `json:"annumDiv"`
	Prin        decimal.NullDecimal `json:"prin"`
	IntR        decimal.NullDecimal `json:"intR"`
	WaivedInt   decimal.NullDecimal `json:"waivedInt"`
	WeeksPaid   sql.NullInt64       `json:"weeksPaid"`
	DoMaturity  sql.NullTime        `json:"doMaturity"`
	ConIntRate  decimal.NullDecimal `json:"conIntRate"`
	Status      sql.NullString      `json:"status"`
	Cycle       sql.NullInt64       `json:"cycle"`
	LNGrpCode   sql.NullInt64       `json:"lNGrpCode"`
	Proff       sql.NullInt64       `json:"proff"`
	FundSource  sql.NullString      `json:"fundSource"`
	DOSRI       sql.NullBool        `json:"DOSRI"`
	LnCategory  sql.NullInt64       `json:"lnCategory"`
	OpenDate    sql.NullTime        `json:"openDate"`
	LastTrnDate sql.NullTime        `json:"lastTrnDate"`
	DisbBy      sql.NullString      `json:"disbBy"`
}

func (q *QueriesLocal) CreateLnMaster(ctx context.Context, arg LnMasterRequest) error {
	_, err := q.db.ExecContext(ctx, createLnMaster,
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
DELETE FROM LnMaster WHERE Acc = $1
`

func (q *QueriesLocal) DeleteLnMaster(ctx context.Context, Acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnMaster, Acc)
	return err
}

type LnMasterInfo struct {
	ModCtr      int64               `json:"modCtr"`
	BrCode      string              `json:"brCode"`
	ModAction   string              `json:"modAction"`
	CID         int64               `json:"CID"`
	Acc         string              `json:"acc"`
	AcctType    sql.NullInt64       `json:"acctType"`
	DisbDate    sql.NullTime        `json:"disbDate"`
	Principal   decimal.NullDecimal `json:"principal"`
	Interest    decimal.NullDecimal `json:"interest"`
	NetProceed  decimal.NullDecimal `json:"netProceed"`
	Gives       sql.NullInt64       `json:"gives"`
	Frequency   sql.NullInt64       `json:"frequency"`
	AnnumDiv    sql.NullInt64       `json:"annumDiv"`
	Prin        decimal.NullDecimal `json:"prin"`
	IntR        decimal.NullDecimal `json:"intR"`
	WaivedInt   decimal.NullDecimal `json:"waivedInt"`
	WeeksPaid   sql.NullInt64       `json:"weeksPaid"`
	DoMaturity  sql.NullTime        `json:"doMaturity"`
	ConIntRate  decimal.NullDecimal `json:"conIntRate"`
	Status      sql.NullString      `json:"status"`
	Cycle       sql.NullInt64       `json:"cycle"`
	LNGrpCode   sql.NullInt64       `json:"lNGrpCode"`
	Proff       sql.NullInt64       `json:"proff"`
	FundSource  sql.NullString      `json:"fundSource"`
	DOSRI       sql.NullBool        `json:"DOSRI"`
	LnCategory  sql.NullInt64       `json:"lnCategory"`
	OpenDate    sql.NullTime        `json:"openDate"`
	LastTrnDate sql.NullTime        `json:"lastTrnDate"`
	DisbBy      sql.NullString      `json:"disbBy"`
}

// -- name: GetLnMaster :one
const getLnMaster = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, CID, Acc, AcctType, dbo.plaindate(DisbDate) DisbDate, 
  Principal, Interest, NetProceed, Gives, Frequency, AnnumDiv, Prin, IntR, 
  WaivedInt, WeeksPaid, dbo.plaindate(DoMaturity) DoMaturity, ConIntRate, Status, Cycle, LNGrpCode, 
  Proff, FundSource, DOSRI, LnCategory, dbo.plaindate(DOPEN) DOPEN, dbo.plaindate(DOLASTTRN) DOLASTTRN, 
  DisbBy
FROM OrgParms, LnMaster d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Acc
`

func scanRowLnMaster(row *sql.Row) (LnMasterInfo, error) {
	var i LnMasterInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
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

func scanRowsLnMaster(rows *sql.Rows) ([]LnMasterInfo, error) {
	items := []LnMasterInfo{}
	for rows.Next() {
		var i LnMasterInfo
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

func (q *QueriesLocal) GetLnMaster(ctx context.Context, Acc string) (LnMasterInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'LnMaster' AND Uploaded = 0 and Acc = $1", getLnMaster)
	row := q.db.QueryRowContext(ctx, sql, Acc)
	return scanRowLnMaster(row)
}

type ListLnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) LnMasterCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, CID, Acc, AcctType, dbo.plaindate(DisbDate) DisbDate, 
  Principal, Interest, NetProceed, Gives, Frequency, AnnumDiv, Prin, IntR, 
  WaivedInt, WeeksPaid, dbo.plaindate(DoMaturity) DoMaturity, ConIntRate, Status, Cycle, LNGrpCode, 
  Proff, FundSource, DOSRI, LnCategory, dbo.plaindate(DOPEN) DOPEN, dbo.plaindate(DOLASTTRN) DOLASTTRN, 
  DisbBy
FROM OrgParms, LnMaster d
`, filenamePath)
}

func (q *QueriesLocal) ListLnMaster(ctx context.Context) ([]LnMasterInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'LnMaster' AND Uploaded = 0`,
		getLnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnMaster(rows)
}

// -- name: UpdateLnMaster :one
const updateLnMaster = `
UPDATE LnMaster SET 
	CID = $2,
	AcctType = $3,
	DISBDATE = $4,
	Principal = $5,
	Interest = $6,
	NetProceed = $7,
	Gives = $8,
	Frequency = $9,
	AnnumDiv = $10,
	Prin = $11,
	IntR = $12,
	WaivedInt = $13,
	WeeksPaid = $14,
	DoMaturity = $15,
	ConIntRate = $16,
	Status = $17,
	Cycle = $18,
	LNGrpCode = $19,
	Proff = $20,
	FundSource = $21,
	DOSRI = $22,
	LnCategory = $23,
	DOpen = $24,
	DOLastTrn = $25,
	DisbBy = $26
WHERE Acc = $1`

func (q *QueriesLocal) UpdateLnMaster(ctx context.Context, arg LnMasterRequest) error {
	_, err := q.db.ExecContext(ctx, updateLnMaster,
		arg.Acc,
		arg.CID,
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

const Table = `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`

// func (q *QueriesLocal) Csv2LnMaster() {

// 	filename := "foo.csv"
//     dbconn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
//     if err != nil {
//         panic(err)
//     }
//     defer dbconn.Release()
//     f, err := os.Open(filename)
//     if err != nil {
//         panic(err)
//     }
//     defer func() { _ = f.Close() }()
//     res, err := dbconn.Conn().PgConn().CopyFrom(context.Background(), f, "COPY csv_test FROM STDIN (FORMAT csv)")
//     if err != nil {
//         panic(err)
//     }
//     fmt.Print(res.RowsAffected())

// 	// r := csv.NewReader(strings.NewReader(Table))
// 	// r.Comma = ','
// 	// r.Comment = '#'

// 	// header, err := r.Read()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// scan, err := NewScanner(header, &Person{})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// for {
// 	// 	record, err := r.Read()
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	var person Person
// 	// 	if err := scan(record, &person); err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Println(person.FirstName, person.LastName)
// 	// }
// }
