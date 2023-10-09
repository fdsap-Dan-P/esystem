package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createEmployment = `-- name: CreateEmployment: one
INSERT INTO Employment (
IIId, Series, Company, Title, Address_Detail, 
Address_URL, Geography_Id, Start_Date, End_Date, Period_Date, Remarks, 
Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING UUId, IIId, Series, Company, Title, Address_Detail, 
Address_URL, Geography_Id, Start_Date, End_Date, Period_Date, Remarks, 
Other_Info
`

type EmploymentRequest struct {
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	Company       string         `json:"company"`
	Title         string         `json:"title"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Remarks       sql.NullString `json:"remarks"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateEmployment(ctx context.Context, arg EmploymentRequest) (model.Employment, error) {
	row := q.db.QueryRowContext(ctx, createEmployment,
		arg.Iiid,
		arg.Series,
		arg.Company,
		arg.Title,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.StartDate,
		arg.EndDate,
		arg.PeriodDate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Employment
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Company,
		&i.Title,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteEmployment = `-- name: DeleteEmployment :exec
DELETE FROM Employment
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteEmployment(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEmployment, uuid)
	return err
}

type EmploymentInfo struct {
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	Company       string         `json:"company"`
	Title         string         `json:"title"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Remarks       sql.NullString `json:"remarks"`
	OtherInfo     sql.NullString `json:"otherInfo"`
	ModCtr        int64          `json:"modCtr"`
	Created       sql.NullTime   `json:"created"`
	Updated       sql.NullTime   `json:"updated"`
}

const getEmployment = `-- name: GetEmployment :one
SELECT 
mr.UUId, IIId, 
Series, Company, Title, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEmployment(ctx context.Context, uuid uuid.UUID) (EmploymentInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployment, uuid)
	var i EmploymentInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Company,
		&i.Title,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getEmploymentbyUuId = `-- name: GetEmploymentbyUuId :one
SELECT 
mr.UUId, IIId, 
Series, Company, Title, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEmploymentbyUuId(ctx context.Context, uuid uuid.UUID) (EmploymentInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmploymentbyUuId, uuid)
	var i EmploymentInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Company,
		&i.Title,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

// const getEmploymentbyName = `-- name: GetEmploymentbyName :one
// SELECT
// mr.UUId, IIId,
// Series, Company, Title, Address_Detail, Address_URL, Geography_Id,
// Start_Date, End_Date, Period_Date, Remarks, Other_Info
// ,mr.Mod_Ctr, mr.Created, mr.Updated
// FROM Employment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
// WHERE Title = $1 LIMIT 1
// `

// func (q *QueriesIdentity) GetEmploymentbyName(ctx context.Context, name string) (EmploymentInfo, error) {
// 	row := q.db.QueryRowContext(ctx, getEmploymentbyName, name)
// 	var i EmploymentInfo
// 	err := row.Scan(
// 		&i.Uuid,
// 		&i.Iiid,
// 		&i.Series,
// 		&i.Company,
// 		&i.Title,
// 		&i.AddressDetail,
// 		&i.AddressUrl,
// 		&i.GeographyId,
// 		&i.StartDate,
// 		&i.EndDate,
// 		&i.PeriodDate,
// 		&i.Remarks,
// 		&i.OtherInfo,

// 		&i.ModCtr,
// 		&i.Created,
// 		&i.Updated,
// 	)
// 	return i, err
// }

const listEmployment = `-- name: ListEmployment:many
SELECT 
mr.UUId, IIId, 
Series, Company, Title, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListEmploymentParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListEmployment(ctx context.Context, arg ListEmploymentParams) ([]EmploymentInfo, error) {
	rows, err := q.db.QueryContext(ctx, listEmployment, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EmploymentInfo{}
	for rows.Next() {
		var i EmploymentInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.Company,
			&i.Title,
			&i.AddressDetail,
			&i.AddressUrl,
			&i.GeographyId,
			&i.StartDate,
			&i.EndDate,
			&i.PeriodDate,
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

const updateEmployment = `-- name: UpdateEmployment :one
UPDATE Employment SET 
IIId = $2,
Series = $3,
Company = $4,
Title = $5,
Address_Detail = $6,
Address_URL = $7,
Geography_Id = $8,
Start_Date = $9,
End_Date = $10,
Period_Date = $11,
Remarks = $12,
Other_Info = $13
WHERE uuid = $1
RETURNING UUId, IIId, Series, Company, Title, Address_Detail, 
Address_URL, Geography_Id, Start_Date, End_Date, Period_Date, Remarks, 
Other_Info
`

func (q *QueriesIdentity) UpdateEmployment(ctx context.Context, arg EmploymentRequest) (model.Employment, error) {
	row := q.db.QueryRowContext(ctx, updateEmployment,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.Company,
		arg.Title,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.StartDate,
		arg.EndDate,
		arg.PeriodDate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Employment
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Company,
		&i.Title,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
