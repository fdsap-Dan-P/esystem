package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createAccountSubject = `-- name: CreateAccountSubject: one
INSERT INTO Account_Subject(
   uuid, account_id, section_subject_id, subject_id, 
   ratings_1st_qtr, ratings_2nd_qtr, ratings_3rd_qtr, ratings_4th_qtr, ratings_final, 
   attendance_ctr, absent_ctr, late_ctr, status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
ON CONFLICT(UUID)
DO UPDATE SET
	account_id =  EXCLUDED.account_id,
	section_subject_id =  EXCLUDED.section_subject_id,
	subject_id =  EXCLUDED.subject_id,
	ratings_1st_qtr =  EXCLUDED.ratings_1st_qtr,
	ratings_2nd_qtr =  EXCLUDED.ratings_2nd_qtr,
	ratings_3rd_qtr =  EXCLUDED.ratings_3rd_qtr,
	ratings_4th_qtr =  EXCLUDED.ratings_4th_qtr,
	ratings_final =  EXCLUDED.ratings_final,
	attendance_ctr =  EXCLUDED.attendance_ctr,
	absent_ctr =  EXCLUDED.absent_ctr,
	late_ctr =  EXCLUDED.late_ctr,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  id, uuid, account_id, section_subject_id, subject_id, 
  ratings_1st_qtr, ratings_2nd_qtr, ratings_3rd_qtr, ratings_4th_qtr, ratings_final, 
  attendance_ctr, absent_ctr, late_ctr, status_id, remarks, other_info`

type AccountSubjectRequest struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	AccountId        int64           `json:"accountId"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	SubjectId        int64           `json:"subjectId"`
	Ratings1stQtr    decimal.Decimal `json:"ratings1stQtr"`
	Ratings2ndQtr    decimal.Decimal `json:"ratings2ndQtr"`
	Ratings3rdQtr    decimal.Decimal `json:"ratings3rdQtr"`
	Ratings4thQtr    decimal.Decimal `json:"ratings4thQtr"`
	RatingsFinal     decimal.Decimal `json:"ratingsFinal"`
	AttendanceCtr    int64           `json:"attendanceCtr"`
	AbsentCtr        int64           `json:"absentCtr"`
	LateCtr          int64           `json:"lateCtr"`
	StatusId         int64           `json:"statusId"`
	Remarks          string          `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

func (q *QueriesSchool) CreateAccountSubject(ctx context.Context, arg AccountSubjectRequest) (model.AccountSubject, error) {
	row := q.db.QueryRowContext(ctx, createAccountSubject,
		arg.Uuid,
		arg.AccountId,
		arg.SectionSubjectId,
		arg.SubjectId,
		arg.Ratings1stQtr,
		arg.Ratings2ndQtr,
		arg.Ratings3rdQtr,
		arg.Ratings4thQtr,
		arg.RatingsFinal,
		arg.AttendanceCtr,
		arg.AbsentCtr,
		arg.LateCtr,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.AccountSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.SectionSubjectId,
		&i.SubjectId,
		&i.Ratings1stQtr,
		&i.Ratings2ndQtr,
		&i.Ratings3rdQtr,
		&i.Ratings4thQtr,
		&i.RatingsFinal,
		&i.AttendanceCtr,
		&i.AbsentCtr,
		&i.LateCtr,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountSubject = `-- name: DeleteAccountSubject :exec
DELETE FROM Account_Subject
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteAccountSubject(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountSubject, uuid)
	return err
}

type AccountSubjectInfo struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	AccountId        int64           `json:"accountId"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	SubjectId        int64           `json:"subjectId"`
	Ratings1stQtr    decimal.Decimal `json:"ratings1stQtr"`
	Ratings2ndQtr    decimal.Decimal `json:"ratings2ndQtr"`
	Ratings3rdQtr    decimal.Decimal `json:"ratings3rdQtr"`
	Ratings4thQtr    decimal.Decimal `json:"ratings4thQtr"`
	RatingsFinal     decimal.Decimal `json:"ratingsFinal"`
	AttendanceCtr    int64           `json:"attendanceCtr"`
	AbsentCtr        int64           `json:"absentCtr"`
	LateCtr          int64           `json:"lateCtr"`
	StatusId         int64           `json:"statusId"`
	Remarks          string          `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
	ModCtr           int64           `json:"modCtr"`
	Created          sql.NullTime    `json:"created"`
	Updated          sql.NullTime    `json:"updated"`
}

const accountSubjectSQL = `-- name: AccountSubjectSQL :one
SELECT
  id, mr.UUID, account_id, section_subject_id, subject_id, ratings_1st_qtr, ratings_2nd_qtr, ratings_3rd_qtr, ratings_4th_qtr, ratings_final, attendance_ctr, absent_ctr, late_ctr, status_id, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Subject d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAccountSubject(q *QueriesSchool, ctx context.Context, sql string) (AccountSubjectInfo, error) {
	var i AccountSubjectInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.SectionSubjectId,
		&i.SubjectId,
		&i.Ratings1stQtr,
		&i.Ratings2ndQtr,
		&i.Ratings3rdQtr,
		&i.Ratings4thQtr,
		&i.RatingsFinal,
		&i.AttendanceCtr,
		&i.AbsentCtr,
		&i.LateCtr,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountSubjects(q *QueriesSchool, ctx context.Context, sql string) ([]AccountSubjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountSubjectInfo{}
	for rows.Next() {
		var i AccountSubjectInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.AccountId,
			&i.SectionSubjectId,
			&i.SubjectId,
			&i.Ratings1stQtr,
			&i.Ratings2ndQtr,
			&i.Ratings3rdQtr,
			&i.Ratings4thQtr,
			&i.RatingsFinal,
			&i.AttendanceCtr,
			&i.AbsentCtr,
			&i.LateCtr,
			&i.StatusId,
			&i.Remarks,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

func (q *QueriesSchool) GetAccountSubject(ctx context.Context, id int64) (AccountSubjectInfo, error) {
	return populateAccountSubject(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", accountSubjectSQL, id))
}

func (q *QueriesSchool) GetAccountSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSubjectInfo, error) {
	return populateAccountSubject(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accountSubjectSQL, uuid))
}

type ListAccountSubjectParams struct {
	AccountId int64 `json:"accountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesSchool) ListAccountSubject(ctx context.Context, arg ListAccountSubjectParams) ([]AccountSubjectInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			accountSubjectSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(accountSubjectSQL)
	}
	return populateAccountSubjects(q, ctx, sql)
}

const updateAccountSubject = `-- name: UpdateAccountSubject :one
UPDATE Account_Subject SET 
	uuid = $2,
	account_id = $3,
	section_subject_id = $4,
	subject_id = $5,
	ratings_1st_qtr = $6,
	ratings_2nd_qtr = $7,
	ratings_3rd_qtr = $8,
	ratings_4th_qtr = $9,
	ratings_final = $10,
	attendance_ctr = $11,
	absent_ctr = $12,
	late_ctr = $13,
	status_id = $14,
	remarks = $15,
	other_info = $16
WHERE id = $1
RETURNING id, uuid, account_id, section_subject_id, subject_id, ratings_1st_qtr, ratings_2nd_qtr, ratings_3rd_qtr, ratings_4th_qtr, ratings_final, attendance_ctr, absent_ctr, late_ctr, status_id, remarks, other_info
`

func (q *QueriesSchool) UpdateAccountSubject(ctx context.Context, arg AccountSubjectRequest) (model.AccountSubject, error) {
	row := q.db.QueryRowContext(ctx, updateAccountSubject,
		arg.Id,
		arg.Uuid,
		arg.AccountId,
		arg.SectionSubjectId,
		arg.SubjectId,
		arg.Ratings1stQtr,
		arg.Ratings2ndQtr,
		arg.Ratings3rdQtr,
		arg.Ratings4thQtr,
		arg.RatingsFinal,
		arg.AttendanceCtr,
		arg.AbsentCtr,
		arg.LateCtr,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.AccountSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.SectionSubjectId,
		&i.SubjectId,
		&i.Ratings1stQtr,
		&i.Ratings2ndQtr,
		&i.Ratings3rdQtr,
		&i.Ratings4thQtr,
		&i.RatingsFinal,
		&i.AttendanceCtr,
		&i.AbsentCtr,
		&i.LateCtr,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
