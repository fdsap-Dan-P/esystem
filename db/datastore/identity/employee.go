package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createEmployee = `-- name: CreateEmployee: one
INSERT INTO Employee (
IIId, Central_Office_Id, Employee_No, Basic_Pay, Date_Hired, 
Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, Position_Id, 
Status_Code, Superior_Id, Type_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
$13, $14, $15) 
ON CONFLICT(UUID)
DO UPDATE SET
    iiid =  EXCLUDED.iiid,
	central_Office_id =  EXCLUDED.central_Office_id,
	employee_no =  EXCLUDED.employee_no,
	basic_pay =  EXCLUDED.basic_pay,
	date_hired =  EXCLUDED.date_hired,
	date_regular =  EXCLUDED.date_regular,
	job_grade =  EXCLUDED.job_grade,
	job_step =  EXCLUDED.job_step,
	level_id =  EXCLUDED.level_id,
	office_id =  EXCLUDED.office_id,
	position_id =  EXCLUDED.position_id,
	status_Code =  EXCLUDED.status_Code,
	superior_id =  EXCLUDED.superior_id,
	type_id =  EXCLUDED.type_id,
	other_info =  EXCLUDED.other_info

RETURNING Id, UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, Date_Hired, 
Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, Position_Id, 
Status_Code, Superior_Id, Type_Id, Other_Info
`

type EmployeeRequest struct {
	Id          int64           `json:"id"`
	Uuid        uuid.UUID       `json:"uuid"`
	Iiid        int64           `json:"iiid"`
	CentralId   int64           `json:"centralId"`
	EmployeeNo  string          `json:"employeeNo"`
	BasicPay    decimal.Decimal `json:"basicPay"`
	DateHired   sql.NullTime    `json:"dateHired"`
	DateRegular sql.NullTime    `json:"dateRegular"`
	JobGrade    int16           `json:"jobGrade"`
	JobStep     int16           `json:"jobStep"`
	LevelId     sql.NullInt64   `json:"levelId"`
	OfficeId    int64           `json:"officeId"`
	PositionId  int64           `json:"positionId"`
	StatusCode  int64           `json:"statusCode"`
	SuperiorId  sql.NullInt64   `json:"superiorId"`
	TypeId      sql.NullInt64   `json:"typeId"`
	OtherInfo   sql.NullString  `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateEmployee(ctx context.Context, arg EmployeeRequest) (model.Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.Iiid,
		arg.CentralId,
		arg.EmployeeNo,
		arg.BasicPay,
		arg.DateHired,
		arg.DateRegular,
		arg.JobGrade,
		arg.JobStep,
		arg.LevelId,
		arg.OfficeId,
		arg.PositionId,
		arg.StatusCode,
		arg.SuperiorId,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Employee
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralId,
		&i.EmployeeNo,
		&i.BasicPay,
		&i.DateHired,
		&i.DateRegular,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.OfficeId,
		&i.PositionId,
		&i.StatusCode,
		&i.SuperiorId,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM Employee
WHERE id = $1
`

func (q *QueriesIdentity) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

type EmployeeInfo struct {
	Id          int64           `json:"id"`
	Uuid        uuid.UUID       `json:"uuid"`
	Iiid        int64           `json:"iiid"`
	CentralId   int64           `json:"centralId"`
	EmployeeNo  string          `json:"employeeNo"`
	BasicPay    decimal.Decimal `json:"basicPay"`
	DateHired   sql.NullTime    `json:"dateHired"`
	DateRegular sql.NullTime    `json:"dateRegular"`
	JobGrade    int16           `json:"jobGrade"`
	JobStep     int16           `json:"jobStep"`
	LevelId     sql.NullInt64   `json:"levelId"`
	OfficeId    int64           `json:"officeId"`
	PositionId  int64           `json:"positionId"`
	StatusCode  int64           `json:"statusCode"`
	SuperiorId  sql.NullInt64   `json:"superiorId"`
	TypeId      sql.NullInt64   `json:"typeId"`
	OtherInfo   sql.NullString  `json:"otherInfo"`
	ModCtr      int64           `json:"modCtr"`
	Created     sql.NullTime    `json:"created"`
	Updated     sql.NullTime    `json:"updated"`
}

const getEmployee = `-- name: GetEmployee :one
SELECT 
Id, mr.UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, 
Date_Hired, Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, 
Position_Id, Status_Code, Superior_Id, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEmployee(ctx context.Context, id int64) (EmployeeInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i EmployeeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralId,
		&i.EmployeeNo,
		&i.BasicPay,
		&i.DateHired,
		&i.DateRegular,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.OfficeId,
		&i.PositionId,
		&i.StatusCode,
		&i.SuperiorId,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getEmployeebyUuId = `-- name: GetEmployeebyUuId :one
SELECT 
Id, mr.UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, 
Date_Hired, Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, 
Position_Id, Status_Code, Superior_Id, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEmployeebyUuId(ctx context.Context, uuid uuid.UUID) (EmployeeInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployeebyUuId, uuid)
	var i EmployeeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralId,
		&i.EmployeeNo,
		&i.BasicPay,
		&i.DateHired,
		&i.DateRegular,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.OfficeId,
		&i.PositionId,
		&i.StatusCode,
		&i.SuperiorId,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getEmployeebyName = `-- name: GetEmployeebyName :one
SELECT 
Id, mr.UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, 
Date_Hired, Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, 
Position_Id, Status_Code, Superior_Id, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Central_Office_Id = $1 and Employee_No = $2 LIMIT 1
`

func (q *QueriesIdentity) GetEmployeebyEmpNo(ctx context.Context, centralId int64, EmpNo string) (EmployeeInfo, error) {
	row := q.db.QueryRowContext(ctx, getEmployeebyName, centralId, EmpNo)
	var i EmployeeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralId,
		&i.EmployeeNo,
		&i.BasicPay,
		&i.DateHired,
		&i.DateRegular,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.OfficeId,
		&i.PositionId,
		&i.StatusCode,
		&i.SuperiorId,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listEmployee = `-- name: ListEmployee:many
SELECT 
Id, mr.UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, 
Date_Hired, Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, 
Position_Id, Status_Code, Superior_Id, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Employee d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Office_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListEmployeeParams struct {
	OfficeId int64 `json:"officeId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesIdentity) ListEmployee(ctx context.Context, arg ListEmployeeParams) ([]EmployeeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listEmployee, arg.OfficeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EmployeeInfo{}
	for rows.Next() {
		var i EmployeeInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Iiid,
			&i.CentralId,
			&i.EmployeeNo,
			&i.BasicPay,
			&i.DateHired,
			&i.DateRegular,
			&i.JobGrade,
			&i.JobStep,
			&i.LevelId,
			&i.OfficeId,
			&i.PositionId,
			&i.StatusCode,
			&i.SuperiorId,
			&i.TypeId,
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

const updateEmployee = `-- name: UpdateEmployee :one
UPDATE Employee SET 
IIId = $2,
Central_Office_Id = $3,
Employee_No = $4,
Basic_Pay = $5,
Date_Hired = $6,
Date_Regular = $7,
Job_Grade = $8,
Job_Step = $9,
Level_Id = $10,
Office_Id = $11,
Position_Id = $12,
Status_Code = $13,
Superior_Id = $14,
Type_Id = $15,
Other_Info = $16
WHERE id = $1
RETURNING Id, UUId, IIId, Central_Office_Id, Employee_No, Basic_Pay, Date_Hired, 
Date_Regular, Job_Grade, Job_Step, Level_Id, Office_Id, Position_Id, 
Status_Code, Superior_Id, Type_Id, Other_Info
`

func (q *QueriesIdentity) UpdateEmployee(ctx context.Context, arg EmployeeRequest) (model.Employee, error) {
	row := q.db.QueryRowContext(ctx, updateEmployee,
		arg.Id,
		arg.Iiid,
		arg.CentralId,
		arg.EmployeeNo,
		arg.BasicPay,
		arg.DateHired,
		arg.DateRegular,
		arg.JobGrade,
		arg.JobStep,
		arg.LevelId,
		arg.OfficeId,
		arg.PositionId,
		arg.StatusCode,
		arg.SuperiorId,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Employee
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.CentralId,
		&i.EmployeeNo,
		&i.BasicPay,
		&i.DateHired,
		&i.DateRegular,
		&i.JobGrade,
		&i.JobStep,
		&i.LevelId,
		&i.OfficeId,
		&i.PositionId,
		&i.StatusCode,
		&i.SuperiorId,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}
