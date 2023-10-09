package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createComment = `-- name: CreateComment: one
INSERT INTO Comment (
   Record_UUID, User_Id, Comment, Other_Info
) VALUES (
$1, $2, $3, $4
) RETURNING UUID, Record_UUId, User_Id, Comment, Other_Info
`

type CommentRequest struct {
	Uuid       uuid.UUID      `json:"uuid"`
	RecordUuid uuid.UUID      `json:"recordUuid"`
	UserId     int64          `json:"userId"`
	Comment    string         `json:"comment"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

func (q *QueriesSocialMedia) CreateComment(ctx context.Context, arg CommentRequest) (model.Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment,
		arg.RecordUuid,
		arg.UserId,
		arg.Comment,
		arg.OtherInfo,
	)
	var i model.Comment
	err := row.Scan(
		&i.Uuid,
		&i.RecordUuid,
		&i.UserId,
		&i.Comment,
		&i.OtherInfo,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM Comment
WHERE uuid = $1
`

func (q *QueriesSocialMedia) DeleteComment(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteComment, uuid)
	return err
}

type CommentInfo struct {
	Uuid       uuid.UUID      `json:"uuid"`
	RecordUuid uuid.UUID      `json:"recordUuid"`
	UserId     int64          `json:"userId"`
	Comment    string         `json:"comment"`
	OtherInfo  sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getComment = `-- name: GetComment :one
SELECT 
mr.UUId, Record_UUId, User_Id, Comment, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Comment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesSocialMedia) GetComment(ctx context.Context, uuid uuid.UUID) (CommentInfo, error) {
	row := q.db.QueryRowContext(ctx, getComment, uuid)
	var i CommentInfo
	err := row.Scan(
		&i.Uuid,
		&i.RecordUuid,
		&i.UserId,
		&i.Comment,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCommentbyUuId = `-- name: GetCommentbyUuId :one
SELECT 
mr.UUId, Record_UUId, User_Id, Comment, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Comment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesSocialMedia) GetCommentbyUuId(ctx context.Context, uuid uuid.UUID) (CommentInfo, error) {
	row := q.db.QueryRowContext(ctx, getCommentbyUuId, uuid)
	var i CommentInfo
	err := row.Scan(
		&i.Uuid,
		&i.RecordUuid,
		&i.UserId,
		&i.Comment,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listComment = `-- name: ListComment:many
SELECT 
mr.UUId, Record_UUId, User_Id, Comment, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Comment d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Record_UUId = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListCommentParams struct {
	RecordUuid uuid.UUID `json:"recordUuid"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

func (q *QueriesSocialMedia) ListComment(ctx context.Context, arg ListCommentParams) ([]CommentInfo, error) {
	rows, err := q.db.QueryContext(ctx, listComment, arg.RecordUuid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CommentInfo{}
	for rows.Next() {
		var i CommentInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.RecordUuid,
			&i.UserId,
			&i.Comment,
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

const updateComment = `-- name: UpdateComment :one
UPDATE Comment SET 
Record_UUId = $2,
User_Id = $3,
Comment = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, Record_UUId, User_Id, Comment, Other_Info
`

func (q *QueriesSocialMedia) UpdateComment(ctx context.Context, arg CommentRequest) (model.Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment,

		arg.Uuid,
		arg.RecordUuid,
		arg.UserId,
		arg.Comment,
		arg.OtherInfo,
	)
	var i model.Comment
	err := row.Scan(
		&i.Uuid,
		&i.RecordUuid,
		&i.UserId,
		&i.Comment,
		&i.OtherInfo,
	)
	return i, err
}
