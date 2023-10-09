package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createAnswer = `-- name: CreateAnswer: one
INSERT INTO Answer(
   uuid, subject_event_id, question_id, answers, points, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (UUID)
DO UPDATE SET
	subject_event_id =  EXCLUDED.subject_event_id,
	question_id =  EXCLUDED.question_id,
	answers =  EXCLUDED.answers,
	points =  EXCLUDED.points,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, subject_event_id, question_id, answers, points, remarks, other_info`

type AnswerRequest struct {
	Id             int64           `json:"id"`
	Uuid           uuid.UUID       `json:"uuid"`
	SubjectEventId int64           `json:"subjectEventId"`
	QuestionId     int64           `json:"questionId"`
	Answers        sql.NullString  `json:"answers"`
	Points         decimal.Decimal `json:"points"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}

func (q *QueriesSchool) CreateAnswer(ctx context.Context, arg AnswerRequest) (model.Answer, error) {
	row := q.db.QueryRowContext(ctx, createAnswer,
		arg.Uuid,
		arg.SubjectEventId,
		arg.QuestionId,
		arg.Answers,
		arg.Points,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Answer
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectEventId,
		&i.QuestionId,
		&i.Answers,
		&i.Points,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAnswer = `-- name: DeleteAnswer :exec
DELETE FROM Answer
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteAnswer(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAnswer, uuid)
	return err
}

type AnswerInfo struct {
	Id             int64           `json:"id"`
	Uuid           uuid.UUID       `json:"uuid"`
	SubjectEventId int64           `json:"subjectEventId"`
	QuestionId     int64           `json:"questionId"`
	Answers        sql.NullString  `json:"answers"`
	Points         decimal.Decimal `json:"points"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
	ModCtr         int64           `json:"modCtr"`
	Created        sql.NullTime    `json:"created"`
	Updated        sql.NullTime    `json:"updated"`
}

const answerSQL = `-- name: AnswerSQL :one
SELECT
  id, mr.UUID, subject_event_id, question_id, answers, points, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Answer d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAnswer(q *QueriesSchool, ctx context.Context, sql string) (AnswerInfo, error) {
	var i AnswerInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectEventId,
		&i.QuestionId,
		&i.Answers,
		&i.Points,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAnswers(q *QueriesSchool, ctx context.Context, sql string) ([]AnswerInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnswerInfo{}
	for rows.Next() {
		var i AnswerInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SubjectEventId,
			&i.QuestionId,
			&i.Answers,
			&i.Points,
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

func (q *QueriesSchool) GetAnswer(ctx context.Context, id int64) (AnswerInfo, error) {
	return populateAnswer(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", answerSQL, id))
}

func (q *QueriesSchool) GetAnswerbyUuid(ctx context.Context, uuid uuid.UUID) (AnswerInfo, error) {
	return populateAnswer(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", answerSQL, uuid))
}

type ListAnswerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSchool) ListAnswer(ctx context.Context, arg ListAnswerParams) ([]AnswerInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			answerSQL, arg.Limit, arg.Offset)
	} else {
		sql = answerSQL
	}
	return populateAnswers(q, ctx, sql)
}

const updateAnswer = `-- name: UpdateAnswer :one
UPDATE Answer SET 
uuid = $2,
subject_event_id = $3,
question_id = $4,
answers = $5,
points = $6,
remarks = $7,
other_info = $8
WHERE id = $1
RETURNING id, uuid, subject_event_id, question_id, answers, points, remarks, other_info
`

func (q *QueriesSchool) UpdateAnswer(ctx context.Context, arg AnswerRequest) (model.Answer, error) {
	row := q.db.QueryRowContext(ctx, updateAnswer,
		arg.Id,
		arg.Uuid,
		arg.SubjectEventId,
		arg.QuestionId,
		arg.Answers,
		arg.Points,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Answer
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SubjectEventId,
		&i.QuestionId,
		&i.Answers,
		&i.Points,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
