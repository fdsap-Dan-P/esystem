package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

const createCustomer = `-- name: CreateCustomer: one
INSERT INTO Customer (
	CID, Center_Code, Title, Cname, FName, MName, MaidenFName, MaidenLName, 
	MaidenMName, Sex, DoBirth, Birth_Place, CivilStatus, CustType, Remarks, 
	Status, Classification, DepoType, subclassification, Pledge_Amount, 
	Mutual_Amount, Pangarap_Amount, Katuparan_Amount, Insurance_Amount, 
	AccPledge, AccMutual, AccPang, AcctKatuparan, AcctInsurance, LoanLmt, 
	CreditLmt, DoRecognized, DoResigned, DoEntry, GoldenLifeDate, Restricted, 
	Borrower, CoMaker, Guarantor, DOSRI, IDCode1, IDNum1, IDCode2, IDNum2, 
	Contact1, Contact2, ContPhone1, Reffered1, Reffered2, Reffered3, Education, 
	Validity1, Validity2, BusinessType, accountNumber, IIID, Religion,
	REG_AREA_CODE, BNK_BRANCH_CODE, GROUP_CODE, SEQ_NUM, REP_CODE, NGO_OFF_CODE
) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
  $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
  $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
  $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, 
  $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, 
  $51, $52, $53, $54, $55, $56, $57, '01', 'E3','1','a',0,'E3') 
`

type CustomerRequest struct {
	CID               int64               `json:"CID"`
	CenterCode        sql.NullString      `json:"centerCode"`
	Title             sql.NullInt64       `json:"title"`
	LName             sql.NullString      `json:"lName"`
	FName             sql.NullString      `json:"fName"`
	MName             sql.NullString      `json:"mName"`
	MaidenFName       sql.NullString      `json:"maidenFName"`
	MaidenLName       sql.NullString      `json:"maidenLName"`
	MaidenMName       sql.NullString      `json:"maidenMName"`
	Sex               sql.NullString      `json:"sex"`
	BirthDate         sql.NullTime        `json:"birthDate"`
	BirthPlace        sql.NullString      `json:"birthPlace"`
	CivilStatus       sql.NullInt64       `json:"civilStatus"`
	CustType          sql.NullInt64       `json:"custType"`
	Remarks           sql.NullString      `json:"remarks"`
	Status            sql.NullInt64       `json:"status"`
	Classification    sql.NullInt64       `json:"classification"`
	DepoType          sql.NullString      `json:"depoType"`
	SubClassification sql.NullInt64       `json:"subClassification"`
	PledgeAmount      decimal.NullDecimal `json:"pledgeAmount"`
	MutualAmount      decimal.NullDecimal `json:"mutualAmount"`
	PangarapAmount    decimal.NullDecimal `json:"pangarapAmount"`
	KatuparanAmount   decimal.NullDecimal `json:"katuparanAmount"`
	InsuranceAmount   decimal.NullDecimal `json:"insuranceAmount"`
	AccPledge         decimal.NullDecimal `json:"accPledge"`
	AccMutual         decimal.NullDecimal `json:"accMutual"`
	AccPang           decimal.NullDecimal `json:"accPang"`
	AccKatuparan      decimal.NullDecimal `json:"accKatuparan"`
	AccInsurance      decimal.NullDecimal `json:"accInsurance"`
	LoanLimit         decimal.NullDecimal `json:"loanLimit"`
	CreditLimit       decimal.NullDecimal `json:"creditLimit"`
	DateRecognized    sql.NullTime        `json:"dateRecognized"`
	DateResigned      sql.NullTime        `json:"dateResigned"`
	DateEntry         sql.NullTime        `json:"dateEntry"`
	GoldenLifeDate    sql.NullTime        `json:"goldenLifeDate"`
	Restricted        sql.NullString      `json:"restricted"`
	Borrower          sql.NullString      `json:"borrower"`
	CoMaker           sql.NullString      `json:"coMaker"`
	Guarantor         sql.NullString      `json:"guarantor"`
	DOSRI             sql.NullInt64       `json:"dOSRI"`
	IDCode1           sql.NullInt64       `json:"iDCode1"`
	IDNum1            sql.NullString      `json:"iDNum1"`
	IDCode2           sql.NullInt64       `json:"iDCode2"`
	IDNum2            sql.NullString      `json:"iDNum2"`
	Contact1          sql.NullString      `json:"contact1"`
	Contact2          sql.NullString      `json:"contact2"`
	Phone1            sql.NullString      `json:"phone1"`
	Reffered1         sql.NullString      `json:"reffered1"`
	Reffered2         sql.NullString      `json:"reffered2"`
	Reffered3         sql.NullString      `json:"reffered3"`
	Education         sql.NullInt64       `json:"education"`
	Validity1         sql.NullTime        `json:"validity1"`
	Validity2         sql.NullTime        `json:"validity2"`
	BusinessType      sql.NullInt64       `json:"businessType"`
	AccountNumber     sql.NullString      `json:"accountNumber"`
	IIID              sql.NullInt64       `json:"iIID"`
	Religion          sql.NullInt64       `json:"religion"`
}

func (q *QueriesLocal) CreateCustomer(ctx context.Context, arg CustomerRequest) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
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
DELETE FROM Customer WHERE CID = $1
`

func (q *QueriesLocal) DeleteCustomer(ctx context.Context, CID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, CID)
	return err
}

type CustomerInfo struct {
	ModCtr            int64               `json:"modCtr"`
	BrCode            string              `json:"brCode"`
	ModAction         string              `json:"modAction"`
	CID               int64               `json:"CID"`
	CenterCode        sql.NullString      `json:"centerCode"`
	Title             sql.NullInt64       `json:"title"`
	LName             sql.NullString      `json:"lName"`
	FName             sql.NullString      `json:"fName"`
	MName             sql.NullString      `json:"mName"`
	MaidenFName       sql.NullString      `json:"maidenFName"`
	MaidenLName       sql.NullString      `json:"maidenLName"`
	MaidenMName       sql.NullString      `json:"maidenMName"`
	Sex               sql.NullString      `json:"sex"`
	BirthDate         sql.NullTime        `json:"birthDate"`
	BirthPlace        sql.NullString      `json:"birthPlace"`
	CivilStatus       sql.NullInt64       `json:"civilStatus"`
	CustType          sql.NullInt64       `json:"custType"`
	Remarks           sql.NullString      `json:"remarks"`
	Status            sql.NullInt64       `json:"status"`
	Classification    sql.NullInt64       `json:"classification"`
	DepoType          sql.NullString      `json:"depoType"`
	SubClassification sql.NullInt64       `json:"subClassification"`
	PledgeAmount      decimal.NullDecimal `json:"pledgeAmount"`
	MutualAmount      decimal.NullDecimal `json:"mutualAmount"`
	PangarapAmount    decimal.NullDecimal `json:"pangarapAmount"`
	KatuparanAmount   decimal.NullDecimal `json:"katuparanAmount"`
	InsuranceAmount   decimal.NullDecimal `json:"insuranceAmount"`
	AccPledge         decimal.NullDecimal `json:"accPledge"`
	AccMutual         decimal.NullDecimal `json:"accMutual"`
	AccPang           decimal.NullDecimal `json:"accPang"`
	AccKatuparan      decimal.NullDecimal `json:"accKatuparan"`
	AccInsurance      decimal.NullDecimal `json:"accInsurance"`
	LoanLimit         decimal.NullDecimal `json:"loanLimit"`
	CreditLimit       decimal.NullDecimal `json:"creditLimit"`
	DateRecognized    sql.NullTime        `json:"dateRecognized"`
	DateResigned      sql.NullTime        `json:"dateResigned"`
	DateEntry         sql.NullTime        `json:"dateEntry"`
	GoldenLifeDate    sql.NullTime        `json:"goldenLifeDate"`
	Restricted        sql.NullString      `json:"restricted"`
	Borrower          sql.NullString      `json:"borrower"`
	CoMaker           sql.NullString      `json:"coMaker"`
	Guarantor         sql.NullString      `json:"guarantor"`
	DOSRI             sql.NullInt64       `json:"dOSRI"`
	IDCode1           sql.NullInt64       `json:"iDCode1"`
	IDNum1            sql.NullString      `json:"iDNum1"`
	IDCode2           sql.NullInt64       `json:"iDCode2"`
	IDNum2            sql.NullString      `json:"iDNum2"`
	Contact1          sql.NullString      `json:"contact1"`
	Contact2          sql.NullString      `json:"contact2"`
	Phone1            sql.NullString      `json:"phone1"`
	Reffered1         sql.NullString      `json:"reffered1"`
	Reffered2         sql.NullString      `json:"reffered2"`
	Reffered3         sql.NullString      `json:"reffered3"`
	Education         sql.NullInt64       `json:"education"`
	Validity1         sql.NullTime        `json:"validity1"`
	Validity2         sql.NullTime        `json:"validity2"`
	BusinessType      sql.NullInt64       `json:"businessType"`
	AccountNumber     sql.NullString      `json:"accountNumber"`
	IIID              sql.NullInt64       `json:"iIID"`
	Religion          sql.NullInt64       `json:"religion"`
}

// -- name: GetCustomer :one
const getCustomer = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, 
  CID, Center_Code, Title, Cname, FName, MName, MaidenFName, MaidenLName, 
  MaidenMName, Sex, DoBirth, Birth_Place, CivilStatus, CustType, Remarks, 
  Status, Classification, DepoType, subclassification, Pledge_Amount, 
  Mutual_Amount, Pangarap_Amount, Katuparan_Amount, Insurance_Amount, 
  AccPledge, AccMutual, AccPang, AcctKatuparan, AcctInsurance, LoanLmt, 
  CreditLmt, DoRecognized, DoResigned, DoEntry, GoldenLifeDate, Restricted, 
  Borrower, CoMaker, Guarantor, DOSRI, IDCode1, IDNum1, IDCode2, IDNum2, 
  Contact1, Contact2, ContPhone1, Reffered1, Reffered2, Reffered3, Education, 
  Validity1, Validity2, BusinessType, accountNumber, IIID, Religion
FROM OrgParms, Customer d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.CID 
`

func scanRowCustomer(row *sql.Row) (CustomerInfo, error) {
	var i CustomerInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
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

func scanRowsCustomer(rows *sql.Rows) ([]CustomerInfo, error) {
	items := []CustomerInfo{}
	for rows.Next() {
		var i CustomerInfo
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

func (q *QueriesLocal) GetCustomer(ctx context.Context, CID int64) (CustomerInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Customer' AND Uploaded = 0 and CID = $1", getCustomer)
	row := q.db.QueryRowContext(ctx, sql, CID)
	return scanRowCustomer(row)
}

type ListCustomerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CustomerCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
	0 ModCtr, OrgParms.DefBranch_Code BrCode, 
	CID, Center_Code, Title, Cname, FName, MName, MaidenFName, MaidenLName, 
	MaidenMName, Sex, DoBirth, Birth_Place, CivilStatus, CustType, Remarks, 
	Status, Classification, DepoType, subclassification, Pledge_Amount, 
	Mutual_Amount, Pangarap_Amount, Katuparan_Amount, Insurance_Amount, 
	AccPledge, AccMutual, AccPang, AcctKatuparan, AcctInsurance, LoanLmt, 
	CreditLmt, DoRecognized, DoResigned, DoEntry, GoldenLifeDate, Restricted, 
	Borrower, CoMaker, Guarantor, DOSRI, IDCode1, IDNum1, IDCode2, IDNum2, 
	Contact1, Contact2, ContPhone1, Reffered1, Reffered2, Reffered3, Education, 
	Validity1, Validity2, BusinessType, accountNumber, IIID, Religion
FROM OrgParms, Customer d
`, filenamePath)
}

func (q *QueriesLocal) ListCustomer(ctx context.Context) ([]CustomerInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Customer' AND Uploaded = 0`,
		getCustomer)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustomer(rows)
}

// -- name: UpdateCustomer :one
const updateCustomer = `
UPDATE Customer SET 
	Center_Code = $2,
	Title = $3,
	Cname = $4,
	FName = $5,
	MName = $6,
	MaidenFName = $7,
	MaidenLName = $8,
	MaidenMName = $9,
	Sex = $10,
	DoBirth = $11,
	Birth_Place = $12,
	CivilStatus = $13,
	CustType = $14,
	Remarks = $15,
	Status = $16,
	Classification = $17,
	DepoType = $18,
	subclassification = $19,
	Pledge_Amount = $20,
	Mutual_Amount = $21,
	Pangarap_Amount = $22,
	Katuparan_Amount = $23,
	Insurance_Amount = $24,
	AccPledge = $25,
	AccMutual = $26,
	AccPang = $27,
	AcctKatuparan = $28,
	AcctInsurance = $29,
	LoanLmt = $30,
	CreditLmt = $31,
	DoRecognized = $32,
	DoResigned = $33,
	DoEntry = $34,
	GoldenLifeDate = $35,
	Restricted = $36,
	Borrower = $37,
	CoMaker = $38,
	Guarantor = $39,
	DOSRI = $40,
	IDCode1 = $41,
	IDNum1 = $42,
	IDCode2 = $43,
	IDNum2 = $44,
	Contact1 = $45,
	Contact2 = $46,
	ContPhone1 = $47,
	Reffered1 = $48,
	Reffered2 = $49,
	Reffered3 = $50,
	Education = $51,
	Validity1 = $52,
	Validity2 = $53,
	BusinessType = $54,
	accountNumber = $55,
	IIID = $56,
	Religion = $57
WHERE CID = $1`

func (q *QueriesLocal) UpdateCustomer(ctx context.Context, arg CustomerRequest) error {
	_, err := q.db.ExecContext(ctx, updateCustomer,
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
