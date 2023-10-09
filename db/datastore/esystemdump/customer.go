package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCustomer = `-- name: CreateCustomer: one
INSERT INTO esystemdump.Customer(
   ModCtr, BrCode, ModAction, CID, CenterCode, Title, 
   LName, FName, MName, MaidenFName, MaidenLName, MaidenMName, 
   Sex, BirthDate, BirthPlace, CivilStatus, CustType, Remarks, Status, 
   Classification, DepoType, SubClassification, PledgeAmount, MutualAmount, 
   PangarapAmount, KatuparanAmount, InsuranceAmount, AccPledge, AccMutual, 
   AccPang, AccKatuparan, AccInsurance, LoanLimit, CreditLimit, DateRecognized, 
   DateResigned, DateEntry, GoldenLifeDate, Restricted, Borrower, CoMaker, Guarantor, 
   DOSRI, IDCode1, IDNum1, IDCode2, IDNum2, Contact1, Contact2, Phone1, 
   Reffered1, Reffered2, Reffered3, Education, Validity1, Validity2, BusinessType, 
   AccountNumber, IIID, Religion )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
	$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, 
	$41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60)
ON CONFLICT (brCode, cID, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	CenterCode =  EXCLUDED.CenterCode,
	Title =  EXCLUDED.Title,
	LName =  EXCLUDED.LName,
	FName =  EXCLUDED.FName,
	MName =  EXCLUDED.MName,
	MaidenFName =  EXCLUDED.MaidenFName,
	MaidenLName =  EXCLUDED.MaidenLName,
	MaidenMName =  EXCLUDED.MaidenMName,
	Sex =  EXCLUDED.Sex,
	BirthDate =  EXCLUDED.BirthDate,
	BirthPlace =  EXCLUDED.BirthPlace,
	CivilStatus =  EXCLUDED.CivilStatus,
	CustType =  EXCLUDED.CustType,
	Remarks =  EXCLUDED.Remarks,
	Status =  EXCLUDED.Status,
	Classification =  EXCLUDED.Classification,
	DepoType =  EXCLUDED.DepoType,
	SubClassification =  EXCLUDED.SubClassification,
	PledgeAmount =  EXCLUDED.PledgeAmount,
	MutualAmount =  EXCLUDED.MutualAmount,
	PangarapAmount =  EXCLUDED.PangarapAmount,
	KatuparanAmount =  EXCLUDED.KatuparanAmount,
	InsuranceAmount =  EXCLUDED.InsuranceAmount,
	AccPledge =  EXCLUDED.AccPledge,
	AccMutual =  EXCLUDED.AccMutual,
	AccPang =  EXCLUDED.AccPang,
	AccKatuparan =  EXCLUDED.AccKatuparan,
	AccInsurance =  EXCLUDED.AccInsurance,
	LoanLimit =  EXCLUDED.LoanLimit,
	CreditLimit =  EXCLUDED.CreditLimit,
	DateRecognized =  EXCLUDED.DateRecognized,
	DateResigned =  EXCLUDED.DateResigned,
	DateEntry =  EXCLUDED.DateEntry,
	GoldenLifeDate =  EXCLUDED.GoldenLifeDate,
	Restricted =  EXCLUDED.Restricted,
	Borrower =  EXCLUDED.Borrower,
	CoMaker =  EXCLUDED.CoMaker,
	Guarantor =  EXCLUDED.Guarantor,
	DOSRI =  EXCLUDED.DOSRI,
	IDCode1 =  EXCLUDED.IDCode1,
	IDNum1 =  EXCLUDED.IDNum1,
	IDCode2 =  EXCLUDED.IDCode2,
	IDNum2 =  EXCLUDED.IDNum2,
	Contact1 =  EXCLUDED.Contact1,
	Contact2 =  EXCLUDED.Contact2,
	Phone1 =  EXCLUDED.Phone1,
	Reffered1 =  EXCLUDED.Reffered1,
	Reffered2 =  EXCLUDED.Reffered2,
	Reffered3 =  EXCLUDED.Reffered3,
	Education =  EXCLUDED.Education,
	Validity1 =  EXCLUDED.Validity1,
	Validity2 =  EXCLUDED.Validity2,
	BusinessType =  EXCLUDED.BusinessType,
	AccountNumber =  EXCLUDED.AccountNumber,
	IIID =  EXCLUDED.IIID,
	Religion =  EXCLUDED.Religion
`

func (q *QueriesDump) CreateCustomer(ctx context.Context, arg model.Customer) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.CenterCode,
		arg.Title,
		arg.LName,
		arg.FName,
		arg.MName,
		arg.MaidenFName,
		arg.MaidenLName,
		arg.MaidenMName,
		arg.Sex,
		arg.BirthDate,
		arg.BirthPlace,
		arg.CivilStatus,
		arg.CustType,
		arg.Remarks,
		arg.Status,
		arg.Classification,
		arg.DepoType,
		arg.SubClassification,
		arg.PledgeAmount,
		arg.MutualAmount,
		arg.PangarapAmount,
		arg.KatuparanAmount,
		arg.InsuranceAmount,
		arg.AccPledge,
		arg.AccMutual,
		arg.AccPang,
		arg.AccKatuparan,
		arg.AccInsurance,
		arg.LoanLimit,
		arg.CreditLimit,
		arg.DateRecognized,
		arg.DateResigned,
		arg.DateEntry,
		arg.GoldenLifeDate,
		arg.Restricted,
		arg.Borrower,
		arg.CoMaker,
		arg.Guarantor,
		arg.DOSRI,
		arg.IDCode1,
		arg.IDNum1,
		arg.IDCode2,
		arg.IDNum2,
		arg.Contact1,
		arg.Contact2,
		arg.Phone1,
		arg.Reffered1,
		arg.Reffered2,
		arg.Reffered3,
		arg.Education,
		arg.Validity1,
		arg.Validity2,
		arg.BusinessType,
		arg.AccountNumber,
		arg.IIID,
		arg.Religion,
	)
	return err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM esystemdump.Customer WHERE BrCode = $1 and CID = $2
`

func (q *QueriesDump) DeleteCustomer(ctx context.Context, brCode string, cID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, brCode, cID)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT
	ModCtr, BrCode, ModAction, CID, CenterCode, Title, LName, FName, MName, MaidenFName, MaidenLName, MaidenMName, Sex, BirthDate, BirthPlace, 
	CivilStatus, CustType, Remarks, Status, Classification, DepoType, SubClassification, PledgeAmount, MutualAmount, PangarapAmount, 
	KatuparanAmount, InsuranceAmount, AccPledge, AccMutual, AccPang, AccKatuparan, AccInsurance, LoanLimit, CreditLimit, 
	DateRecognized, DateResigned, DateEntry, GoldenLifeDate, Restricted, Borrower, CoMaker, Guarantor, DOSRI, IDCode1, IDNum1, IDCode2, IDNum2, 
	Contact1, Contact2, Phone1, Reffered1, Reffered2, Reffered3, Education, Validity1, Validity2, BusinessType, AccountNumber, IIID, Religion
FROM esystemdump.Customer
`

func scanRowCustomer(row *sql.Row) (model.Customer, error) {
	var i model.Customer
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.CenterCode,
		&i.Title,
		&i.LName,
		&i.FName,
		&i.MName,
		&i.MaidenFName,
		&i.MaidenLName,
		&i.MaidenMName,
		&i.Sex,
		&i.BirthDate,
		&i.BirthPlace,
		&i.CivilStatus,
		&i.CustType,
		&i.Remarks,
		&i.Status,
		&i.Classification,
		&i.DepoType,
		&i.SubClassification,
		&i.PledgeAmount,
		&i.MutualAmount,
		&i.PangarapAmount,
		&i.KatuparanAmount,
		&i.InsuranceAmount,
		&i.AccPledge,
		&i.AccMutual,
		&i.AccPang,
		&i.AccKatuparan,
		&i.AccInsurance,
		&i.LoanLimit,
		&i.CreditLimit,
		&i.DateRecognized,
		&i.DateResigned,
		&i.DateEntry,
		&i.GoldenLifeDate,
		&i.Restricted,
		&i.Borrower,
		&i.CoMaker,
		&i.Guarantor,
		&i.DOSRI,
		&i.IDCode1,
		&i.IDNum1,
		&i.IDCode2,
		&i.IDNum2,
		&i.Contact1,
		&i.Contact2,
		&i.Phone1,
		&i.Reffered1,
		&i.Reffered2,
		&i.Reffered3,
		&i.Education,
		&i.Validity1,
		&i.Validity2,
		&i.BusinessType,
		&i.AccountNumber,
		&i.IIID,
		&i.Religion,
	)
	return i, err
}

func scanRowsCustomer(rows *sql.Rows) ([]model.Customer, error) {
	items := []model.Customer{}
	for rows.Next() {
		var i model.Customer
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.CenterCode,
			&i.Title,
			&i.LName,
			&i.FName,
			&i.MName,
			&i.MaidenFName,
			&i.MaidenLName,
			&i.MaidenMName,
			&i.Sex,
			&i.BirthDate,
			&i.BirthPlace,
			&i.CivilStatus,
			&i.CustType,
			&i.Remarks,
			&i.Status,
			&i.Classification,
			&i.DepoType,
			&i.SubClassification,
			&i.PledgeAmount,
			&i.MutualAmount,
			&i.PangarapAmount,
			&i.KatuparanAmount,
			&i.InsuranceAmount,
			&i.AccPledge,
			&i.AccMutual,
			&i.AccPang,
			&i.AccKatuparan,
			&i.AccInsurance,
			&i.LoanLimit,
			&i.CreditLimit,
			&i.DateRecognized,
			&i.DateResigned,
			&i.DateEntry,
			&i.GoldenLifeDate,
			&i.Restricted,
			&i.Borrower,
			&i.CoMaker,
			&i.Guarantor,
			&i.DOSRI,
			&i.IDCode1,
			&i.IDNum1,
			&i.IDCode2,
			&i.IDNum2,
			&i.Contact1,
			&i.Contact2,
			&i.Phone1,
			&i.Reffered1,
			&i.Reffered2,
			&i.Reffered3,
			&i.Education,
			&i.Validity1,
			&i.Validity2,
			&i.BusinessType,
			&i.AccountNumber,
			&i.IIID,
			&i.Religion,
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

func (q *QueriesDump) GetCustomer(ctx context.Context, brCode string, cID int64) (model.Customer, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and CID = $2", getCustomer)
	row := q.db.QueryRowContext(ctx, sql, brCode, cID)
	return scanRowCustomer(row)
}

type ListCustomerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCustomer(ctx context.Context, lastModCtr int64) ([]model.Customer, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCustomer)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustomer(rows)
}

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE esystemdump.Customer SET 
	ModCtr = $1,
	CenterCode = $5,
	Title = $6,
	LName = $7,
	FName = $8,
	MName = $9,
	MaidenFName = $10,
	MaidenLName = $11,
	MaidenMName = $12,
	Sex = $13,
	BirthDate = $14,
	BirthPlace = $15,
	CivilStatus = $16,
	CustType = $17,
	Remarks = $18,
	Status = $19,
	Classification = $20,
	DepoType = $21,
	SubClassification = $22,
	PledgeAmount = $23,
	MutualAmount = $24,
	PangarapAmount = $25,
	KatuparanAmount = $26,
	InsuranceAmount = $27,
	AccPledge = $28,
	AccMutual = $29,
	AccPang = $30,
	AccKatuparan = $31,
	AccInsurance = $32,
	LoanLimit = $33,
	CreditLimit = $34,
	DateRecognized = $35,
	DateResigned = $36,
	DateEntry = $37,
	GoldenLifeDate = $38,
	Restricted = $39,
	Borrower = $40,
	CoMaker = $41,
	Guarantor = $42,
	DOSRI = $43,
	IDCode1 = $44,
	IDNum1 = $45,
	IDCode2 = $46,
	IDNum2 = $47,
	Contact1 = $48,
	Contact2 = $49,
	Phone1 = $50,
	Reffered1 = $51,
	Reffered2 = $52,
	Reffered3 = $53,
	Education = $54,
	Validity1 = $55,
	Validity2 = $56,
	BusinessType = $57,
	AccountNumber = $58,
	IIID = $59,
	Religion = $60
WHERE BrCode = $2 and CID = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateCustomer(ctx context.Context, arg model.Customer) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.CenterCode,
		arg.Title,
		arg.LName,
		arg.FName,
		arg.MName,
		arg.MaidenFName,
		arg.MaidenLName,
		arg.MaidenMName,
		arg.Sex,
		arg.BirthDate,
		arg.BirthPlace,
		arg.CivilStatus,
		arg.CustType,
		arg.Remarks,
		arg.Status,
		arg.Classification,
		arg.DepoType,
		arg.SubClassification,
		arg.PledgeAmount,
		arg.MutualAmount,
		arg.PangarapAmount,
		arg.KatuparanAmount,
		arg.InsuranceAmount,
		arg.AccPledge,
		arg.AccMutual,
		arg.AccPang,
		arg.AccKatuparan,
		arg.AccInsurance,
		arg.LoanLimit,
		arg.CreditLimit,
		arg.DateRecognized,
		arg.DateResigned,
		arg.DateEntry,
		arg.GoldenLifeDate,
		arg.Restricted,
		arg.Borrower,
		arg.CoMaker,
		arg.Guarantor,
		arg.DOSRI,
		arg.IDCode1,
		arg.IDNum1,
		arg.IDCode2,
		arg.IDNum2,
		arg.Contact1,
		arg.Contact2,
		arg.Phone1,
		arg.Reffered1,
		arg.Reffered2,
		arg.Reffered3,
		arg.Education,
		arg.Validity1,
		arg.Validity2,
		arg.BusinessType,
		arg.AccountNumber,
		arg.IIID,
		arg.Religion,
	)
	return err
}
