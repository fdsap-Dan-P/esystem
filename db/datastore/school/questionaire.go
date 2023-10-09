package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createQuestionaire = `-- name: CreateQuestionaire: one
INSERT INTO Questionaire(
   uuid, code, version, title, type_id, subject_id, date_revised, 
   office_id, author_id, status_id, point_equivalent, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
ON CONFLICT (UUID)
DO UPDATE SET
	code =  EXCLUDED.code,
	version =  EXCLUDED.version,
	title =  EXCLUDED.title,
	type_id =  EXCLUDED.type_id,
	subject_id =  EXCLUDED.subject_id,
	date_revised =  EXCLUDED.date_revised,
	office_id =  EXCLUDED.office_id,
	author_id =  EXCLUDED.author_id,
	status_id =  EXCLUDED.status_id,
	point_equivalent =  EXCLUDED.point_equivalent,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  id, uuid, code, version, title, type_id, subject_id, date_revised, 
  office_id, author_id, status_id, point_equivalent, remarks, other_info`

type QuestionaireRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	Code            string         `json:"code"`
	Version         int64          `json:"version"`
	Title           string         `json:"title"`
	TypeId          int64          `json:"typeId"`
	SubjectId       sql.NullInt64  `json:"subjectId"`
	DateRevised     time.Time      `json:"dateRevised"`
	OfficeId        sql.NullInt64  `json:"officeId"`
	AuthorId        sql.NullInt64  `json:"authorId"`
	StatusId        int64          `json:"statusId"`
	PointEquivalent sql.NullString `json:"pointEquivalent"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateQuestionaire(ctx context.Context, arg QuestionaireRequest) (model.Questionaire, error) {
	row := q.db.QueryRowContext(ctx, createQuestionaire,
		arg.Uuid,
		arg.Code,
		arg.Version,
		arg.Title,
		arg.TypeId,
		arg.SubjectId,
		arg.DateRevised,
		arg.OfficeId,
		arg.AuthorId,
		arg.StatusId,
		arg.PointEquivalent,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Questionaire
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Version,
		&i.Title,
		&i.TypeId,
		&i.SubjectId,
		&i.DateRevised,
		&i.OfficeId,
		&i.AuthorId,
		&i.StatusId,
		&i.PointEquivalent,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteQuestionaire = `-- name: DeleteQuestionaire :exec
DELETE FROM Questionaire
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteQuestionaire(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteQuestionaire, uuid)
	return err
}

type QuestionaireInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	Code            string         `json:"code"`
	Version         int64          `json:"version"`
	Title           string         `json:"title"`
	TypeId          int64          `json:"typeId"`
	SubjectId       sql.NullInt64  `json:"subjectId"`
	DateRevised     time.Time      `json:"dateRevised"`
	OfficeId        sql.NullInt64  `json:"officeId"`
	AuthorId        sql.NullInt64  `json:"authorId"`
	StatusId        int64          `json:"statusId"`
	PointEquivalent sql.NullString `json:"pointEquivalent"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const questionaireSQL = `-- name: QuestionaireSQL :one
SELECT
  id, mr.UUID, code, version, title, type_id, subject_id, 
  date_revised, office_id, author_id, status_id, point_equivalent, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Questionaire d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateQuestionaire(q *QueriesSchool, ctx context.Context, sql string) (QuestionaireInfo, error) {
	var i QuestionaireInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Version,
		&i.Title,
		&i.TypeId,
		&i.SubjectId,
		&i.DateRevised,
		&i.OfficeId,
		&i.AuthorId,
		&i.StatusId,
		&i.PointEquivalent,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateQuestionaires(q *QueriesSchool, ctx context.Context, sql string) ([]QuestionaireInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuestionaireInfo{}
	for rows.Next() {
		var i QuestionaireInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.Version,
			&i.Title,
			&i.TypeId,
			&i.SubjectId,
			&i.DateRevised,
			&i.OfficeId,
			&i.AuthorId,
			&i.StatusId,
			&i.PointEquivalent,
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

func (q *QueriesSchool) GetQuestionaire(ctx context.Context, id int64) (QuestionaireInfo, error) {
	return populateQuestionaire(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", questionaireSQL, id))
}

func (q *QueriesSchool) GetQuestionairebyUuid(ctx context.Context, uuid uuid.UUID) (QuestionaireInfo, error) {
	return populateQuestionaire(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", questionaireSQL, uuid))
}

type ListQuestionaireParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSchool) ListQuestionaire(ctx context.Context, arg ListQuestionaireParams) ([]QuestionaireInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			questionaireSQL, arg.Limit, arg.Offset)
	} else {
		sql = questionaireSQL
	}
	return populateQuestionaires(q, ctx, sql)
}

const updateQuestionaire = `-- name: UpdateQuestionaire :one
UPDATE Questionaire SET 
	uuid = $2,
	code = $3,
	version = $4,
	title = $5,
	type_id = $6,
	subject_id = $7,
	date_revised = $8,
	office_id = $9,
	author_id = $10,
	status_id = $11,
	point_equivalent = $12,
	remarks = $13,
	other_info = $14
WHERE id = $1
RETURNING 
  id, uuid, code, version, title, type_id, subject_id, date_revised, office_id, 
  author_id, status_id, point_equivalent, remarks, other_info
`

func (q *QueriesSchool) UpdateQuestionaire(ctx context.Context, arg QuestionaireRequest) (model.Questionaire, error) {
	row := q.db.QueryRowContext(ctx, updateQuestionaire,
		arg.Id,
		arg.Uuid,
		arg.Code,
		arg.Version,
		arg.Title,
		arg.TypeId,
		arg.SubjectId,
		arg.DateRevised,
		arg.OfficeId,
		arg.AuthorId,
		arg.StatusId,
		arg.PointEquivalent,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Questionaire
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Version,
		&i.Title,
		&i.TypeId,
		&i.SubjectId,
		&i.DateRevised,
		&i.OfficeId,
		&i.AuthorId,
		&i.StatusId,
		&i.PointEquivalent,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
