package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type OrgParmsInfo struct {
	BrCode      string    `json:"brCode"`
	EbSysDate   time.Time `json:"ebSysDate"`
	RunState    int64     `json:"runState"`
	OrgAddress  string    `json:"orgAddress"`
	TaxInfo     string    `json:"taxInfo"`
	DefCity     string    `json:"defCity"`
	DefProvince string    `json:"defProvince"`
	DefCountry  string    `json:"defCountry"`
	DefZip      string    `json:"defZip"`
	WaivableInt bool      `json:"waivableInt"`
	DBVersion   string    `json:"dBVersion"`
	ESystemVer  []byte    `json:"eSystemVer"`
	NewBrCode   int64     `json:"newBrCode"`
}

func scanRowOrgParms(row *sql.Row) (OrgParmsInfo, error) {
	var i OrgParmsInfo
	err := row.Scan(
		&i.BrCode,
		&i.EbSysDate,
		&i.RunState,
		&i.OrgAddress,
		&i.TaxInfo,
		&i.DefCity,
		&i.DefProvince,
		&i.DefCountry,
		&i.DefZip,
		&i.WaivableInt,
		&i.DBVersion,
		&i.ESystemVer,
		&i.NewBrCode,
	)
	return i, err
}

// -- name: GetOrgParms :one
const getOrgParms = `
SELECT 
  DefBranch_Code BrCode, EbSysDate, RunState, OrgAddress, TaxInfo, DefCity, DefProvince, 
  DefCountry, DefZip, WaivableInt, DBVersion, ESystemVer, NewBrCode
FROM OrgParms 
`

func (q *QueriesLocal) GetOrgParms(ctx context.Context) (OrgParmsInfo, error) {
	log.Printf("getOrgParms ctx:%v", ctx)
	row := q.db.QueryRowContext(ctx, getOrgParms)
	log.Printf("getOrgParms %v", row)
	return scanRowOrgParms(row)
}
