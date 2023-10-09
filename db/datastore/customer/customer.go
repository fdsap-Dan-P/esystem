package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createCustomer = `-- name: CreateCustomer: one
INSERT INTO Customer (
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, 
Credit_Limit, Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, 
Office_Id, Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
$13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
) RETURNING Id, UUId, IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, 
Credit_Limit, Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, 
Office_Id, Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
`

type CustomerRequest struct {
	Id                  int64           `json:"id"`
	Uuid                uuid.UUID       `json:"uuid"`
	Iiid                int64           `json:"iiid"`
	CentralOfficeId     int64           `json:"centralOfficeId"`
	Cid                 int64           `json:"Cid"`
	CustomerAltId       sql.NullString  `json:"customerAltId"`
	DebitLimit          decimal.Decimal `json:"debitLimit"`
	CreditLimit         decimal.Decimal `json:"creditLimit"`
	DateEntry           sql.NullTime    `json:"dateEntry"`
	DateRecognized      sql.NullTime    `json:"dateRecognized"`
	DateResigned        sql.NullTime    `json:"dateResigned"`
	Resigned            sql.NullBool    `json:"resigned"`
	ReasonResigned      sql.NullString  `json:"reasonResigned"`
	LastActivityDate    sql.NullTime    `json:"lastActivityDate"`
	Dosri               bool            `json:"dosri"`
	ClassificationId    sql.NullInt64   `json:"classificationId"`
	CustomerGroupId     sql.NullInt64   `json:"customerGroupId"`
	OfficeId            int64           `json:"officeId"`
	RestrictionId       sql.NullInt64   `json:"restrictionId"`
	RiskClassId         sql.NullInt64   `json:"riskClassId"`
	StatusCode          int64           `json:"statusCode"`
	IndustryId          sql.NullInt64   `json:"industryId"`
	SubClassificationId sql.NullInt64   `json:"subClassificationId"`
	Remarks             sql.NullString  `json:"remarks"`
	OtherInfo           sql.NullString  `json:"otherInfo"`
}

func (q *QueriesCustomer) CreateCustomer(ctx context.Context, arg CustomerRequest) (model.Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.Iiid,
		arg.CentralOfficeId,
		arg.Cid,
		arg.CustomerAltId,
		arg.DebitLimit,
		arg.CreditLimit,
		arg.DateEntry,
		arg.DateRecognized,
		arg.DateResigned,
		arg.Resigned,
		arg.ReasonResigned,
		arg.LastActivityDate,
		arg.Dosri,
		arg.ClassificationId,
		arg.CustomerGroupId,
		arg.OfficeId,
		arg.RestrictionId,
		arg.RiskClassId,
		arg.StatusCode,
		arg.IndustryId,
		arg.SubClassificationId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Customer
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM Customer
WHERE id = $1
`

func (q *QueriesCustomer) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

type CustomerInfo struct {
	Id                  int64           `json:"id"`
	Uuid                uuid.UUID       `json:"uuid"`
	Iiid                int64           `json:"iiid"`
	CentralOfficeId     int64           `json:"centralOfficeId"`
	Cid                 int64           `json:"Cid"`
	CustomerAltId       sql.NullString  `json:"customerAltId"`
	DebitLimit          decimal.Decimal `json:"debitLimit"`
	CreditLimit         decimal.Decimal `json:"creditLimit"`
	DateEntry           sql.NullTime    `json:"dateEntry"`
	DateRecognized      sql.NullTime    `json:"dateRecognized"`
	DateResigned        sql.NullTime    `json:"dateResigned"`
	Resigned            sql.NullBool    `json:"resigned"`
	ReasonResigned      sql.NullString  `json:"reasonResigned"`
	LastActivityDate    sql.NullTime    `json:"lastActivityDate"`
	Dosri               bool            `json:"dosri"`
	ClassificationId    sql.NullInt64   `json:"classificationId"`
	CustomerGroupId     sql.NullInt64   `json:"customerGroupId"`
	OfficeId            int64           `json:"officeId"`
	RestrictionId       sql.NullInt64   `json:"restrictionId"`
	RiskClassId         sql.NullInt64   `json:"riskClassId"`
	StatusCode          int64           `json:"statusCode"`
	IndustryId          sql.NullInt64   `json:"industryId"`
	SubClassificationId sql.NullInt64   `json:"subClassificationId"`
	Remarks             sql.NullString  `json:"remarks"`
	OtherInfo           sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getCustomer = `-- name: GetCustomer :one
SELECT 
Id, mr.UUId, 
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, Credit_Limit, 
Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, Office_Id, 
Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomer(ctx context.Context, id int64) (CustomerInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i CustomerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerbyUuId = `-- name: GetCustomerbyUuId :one
SELECT 
Id, mr.UUId, 
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, Credit_Limit, 
Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, Office_Id, 
Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerbyUuId(ctx context.Context, uuid uuid.UUID) (CustomerInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerbyUuId, uuid)
	var i CustomerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerbyAltId = `-- name: GetCustomerbyAltId :one
SELECT 
Id, mr.UUId, 
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, Credit_Limit, 
Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned, 
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, Office_Id, 
Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Customer_Alt_Id = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerbyAltId(ctx context.Context, altId string) (CustomerInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerbyAltId, altId)
	var i CustomerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerbyCid = `-- name: GetCustomerbyCid :one
SELECT 
Id, mr.UUId, 
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, Credit_Limit, 
Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, Office_Id, 
Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Central_Office_Id = $1 and Cid = $2 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerbyCid(ctx context.Context, centralOfficeId int64, Cid int64) (CustomerInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerbyCid, centralOfficeId, Cid)
	var i CustomerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listCustomer = `-- name: ListCustomer:many
SELECT 
Id, mr.UUId, 
IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, Credit_Limit, 
Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, Office_Id, 
Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListCustomerParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomer(ctx context.Context, arg ListCustomerParams) ([]CustomerInfo, error) {
	rows, err := q.db.QueryContext(ctx, listCustomer, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CustomerInfo{}
	for rows.Next() {
		var i CustomerInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Iiid,
			&i.CentralOfficeId,
			&i.Cid,
			&i.CustomerAltId,
			&i.DebitLimit,
			&i.CreditLimit,
			&i.DateEntry,
			&i.DateRecognized,
			&i.DateResigned,
			&i.Resigned,
			&i.ReasonResigned,
			&i.LastActivityDate,
			&i.Dosri,
			&i.ClassificationId,
			&i.CustomerGroupId,
			&i.OfficeId,
			&i.RestrictionId,
			&i.RiskClassId,
			&i.StatusCode,
			&i.IndustryId,
			&i.SubClassificationId,
			&i.Remarks,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
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

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE Customer SET 
IIId = $2,
Central_Office_Id = $3,
Cid = $4,
Customer_Alt_Id = $5,
Debit_Limit = $6,
Credit_Limit = $7,
Date_Entry = $8,
Date_Recognized = $9,
Date_Resigned = $10,
Resigned = $11,
Reason_Resigned = $12,
Last_Activity_Date = $13,
dosri = $14,
Classification_Id = $15,
Customer_Group_Id = $16,
Office_Id = $17,
Restriction_Id = $18,
Risk_Class_Id = $19,
Status_Code = $20,
Industry_Id = $21,
Sub_Classification_Id = $22,
Remarks = $23,
Other_Info = $24

WHERE id = $1
RETURNING Id, UUId, IIId, Central_Office_Id, Cid, Customer_Alt_Id, Debit_Limit, 
Credit_Limit, Date_Entry, Date_Recognized, Date_Resigned, Resigned, Reason_Resigned,
Last_Activity_Date, dosri, Classification_Id, Customer_Group_Id, 
Office_Id, Restriction_Id, Risk_Class_Id, Status_Code, Industry_Id, Sub_Classification_Id, Remarks, Other_Info
`

func (q *QueriesCustomer) UpdateCustomer(ctx context.Context, arg CustomerRequest) (model.Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.Id,
		arg.Iiid,
		arg.CentralOfficeId,
		arg.Cid,
		arg.CustomerAltId,
		arg.DebitLimit,
		arg.CreditLimit,
		arg.DateEntry,
		arg.DateRecognized,
		arg.DateResigned,
		arg.Resigned,
		arg.ReasonResigned,
		arg.LastActivityDate,
		arg.Dosri,
		arg.ClassificationId,
		arg.CustomerGroupId,
		arg.OfficeId,
		arg.RestrictionId,
		arg.RiskClassId,
		arg.StatusCode,
		arg.IndustryId,
		arg.SubClassificationId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Customer
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralOfficeId,
		&i.Cid,
		&i.CustomerAltId,
		&i.DebitLimit,
		&i.CreditLimit,
		&i.DateEntry,
		&i.DateRecognized,
		&i.DateResigned,
		&i.Resigned,
		&i.ReasonResigned,
		&i.LastActivityDate,
		&i.Dosri,
		&i.ClassificationId,
		&i.CustomerGroupId,
		&i.OfficeId,
		&i.RestrictionId,
		&i.RiskClassId,
		&i.StatusCode,
		&i.IndustryId,
		&i.SubClassificationId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
