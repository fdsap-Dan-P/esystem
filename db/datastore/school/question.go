package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createQuestion = `-- name: CreateQuestion: one
INSERT INTO Question(
   uuid, questionaire_id, series, type_id, question_item, 
   choices, answer_type, parent_id, status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(UUID)
DO UPDATE SET
	questionaire_id =  EXCLUDED.questionaire_id,
	series =  EXCLUDED.series,
	type_id =  EXCLUDED.type_id,
	question_item =  EXCLUDED.question_item,
	choices =  EXCLUDED.choices,
	answer_type =  EXCLUDED.answer_type,
	parent_id =  EXCLUDED.parent_id,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  id, uuid, questionaire_id, series, type_id, question_item, 
  choices, answer_type, parent_id, status_id, remarks, other_info`

type QuestionRequest struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	QuestionaireId int64          `json:"questionaireId"`
	Series         int16          `json:"series"`
	TypeId         int64          `json:"typeId"`
	QuestionItem   string         `json:"questionItem"`
	Choices        sql.NullString `json:"choices"`
	AnswerType     sql.NullString `json:"answerType"`
	ParentId       sql.NullInt64  `json:"parentId"`
	StatusId       int64          `json:"statusId"`
	Remarks        sql.NullString `json:"remarks"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateQuestion(ctx context.Context, arg QuestionRequest) (model.Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion,
		arg.Uuid,
		arg.QuestionaireId,
		arg.Series,
		arg.TypeId,
		arg.QuestionItem,
		arg.Choices,
		arg.AnswerType,
		arg.ParentId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Question
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.QuestionaireId,
		&i.Series,
		&i.TypeId,
		&i.QuestionItem,
		&i.Choices,
		&i.AnswerType,
		&i.ParentId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
DELETE FROM Question
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteQuestion(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteQuestion, uuid)
	return err
}

type QuestionInfo struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	QuestionaireId int64          `json:"questionaireId"`
	Series         int16          `json:"series"`
	TypeId         int64          `json:"typeId"`
	QuestionItem   string         `json:"questionItem"`
	Choices        sql.NullString `json:"choices"`
	AnswerType     sql.NullString `json:"answerType"`
	ParentId       sql.NullInt64  `json:"parentId"`
	StatusId       int64          `json:"statusId"`
	Remarks        sql.NullString `json:"remarks"`
	OtherInfo      sql.NullString `json:"otherInfo"`
	ModCtr         int64          `json:"modCtr"`
	Created        sql.NullTime   `json:"created"`
	Updated        sql.NullTime   `json:"updated"`
}

const questionSQL = `-- name: QuestionSQL :one
SELECT
  id, mr.UUID, questionaire_id, series, type_id, question_item, 
  choices, answer_type, parent_id, status_id, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Question d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateQuestion(q *QueriesSchool, ctx context.Context, sql string) (QuestionInfo, error) {
	var i QuestionInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.QuestionaireId,
		&i.Series,
		&i.TypeId,
		&i.QuestionItem,
		&i.Choices,
		&i.AnswerType,
		&i.ParentId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateQuestions(q *QueriesSchool, ctx context.Context, sql string) ([]QuestionInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuestionInfo{}
	for rows.Next() {
		var i QuestionInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.QuestionaireId,
			&i.Series,
			&i.TypeId,
			&i.QuestionItem,
			&i.Choices,
			&i.AnswerType,
			&i.ParentId,
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

func (q *QueriesSchool) GetQuestion(ctx context.Context, id int64) (QuestionInfo, error) {
	return populateQuestion(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", questionSQL, id))
}

func (q *QueriesSchool) GetQuestionbyUuid(ctx context.Context, uuid uuid.UUID) (QuestionInfo, error) {
	return populateQuestion(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", questionSQL, uuid))
}

type ListQuestionParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSchool) ListQuestion(ctx context.Context, arg ListQuestionParams) ([]QuestionInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			questionSQL, arg.Limit, arg.Offset)
	} else {
		sql = questionSQL
	}
	return populateQuestions(q, ctx, sql)
}

const updateQuestion = `-- name: UpdateQuestion :one
UPDATE Question SET 
uuid = $2,
questionaire_id = $3,
series = $4,
type_id = $5,
question_item = $6,
choices = $7,
answer_type = $8,
parent_id = $9,
status_id = $10,
remarks = $11,
other_info = $12
WHERE id = $1
RETURNING id, uuid, questionaire_id, series, type_id, question_item, choices, answer_type, parent_id, status_id, remarks, other_info
`

func (q *QueriesSchool) UpdateQuestion(ctx context.Context, arg QuestionRequest) (model.Question, error) {
	row := q.db.QueryRowContext(ctx, updateQuestion,
		arg.Id,
		arg.Uuid,
		arg.QuestionaireId,
		arg.Series,
		arg.TypeId,
		arg.QuestionItem,
		arg.Choices,
		arg.AnswerType,
		arg.ParentId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Question
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.QuestionaireId,
		&i.Series,
		&i.TypeId,
		&i.QuestionItem,
		&i.Choices,
		&i.AnswerType,
		&i.ParentId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
