package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"simplebank/model"
)

const createPost = `-- name: CreatePost: one
INSERT INTO Post (
User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING UUId, User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
`

type PostRequest struct {
	Uuid         uuid.UUID       `json:"uuid"`
	UserId       int64           `json:"userId"`
	Caption      string          `json:"caption"`
	MessageBody  string          `json:"messageBody"`
	Url          string          `json:"url"`
	ImageUri     string          `json:"imageUri"`
	ThumbnailUri string          `json:"thumbnailUri"`
	Keywords     []string        `json:"keywords"`
	Mood         model.MoodState `json:"mood"`
	MoodEmoji    string          `json:"moodEmoji"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
}

func (q *QueriesSocialMedia) CreatePost(ctx context.Context, arg PostRequest) (model.Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.UserId,
		arg.Caption,
		arg.MessageBody,
		arg.Url,
		arg.ImageUri,
		arg.ThumbnailUri,
		pq.Array(arg.Keywords),
		arg.Mood,
		arg.MoodEmoji,
		arg.OtherInfo,
	)
	var i model.Post
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Caption,
		&i.MessageBody,
		&i.Url,
		&i.ImageUri,
		&i.ThumbnailUri,
		pq.Array(&i.Keywords),
		&i.Mood,
		&i.MoodEmoji,
		&i.OtherInfo,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM Post
WHERE uuid = $1
`

func (q *QueriesSocialMedia) DeletePost(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePost, uuid)
	return err
}

type PostInfo struct {
	Uuid         uuid.UUID       `json:"uuid"`
	UserId       int64           `json:"userId"`
	Caption      string          `json:"caption"`
	MessageBody  string          `json:"messageBody"`
	Url          string          `json:"url"`
	ImageUri     string          `json:"imageUri"`
	ThumbnailUri string          `json:"thumbnailUri"`
	Keywords     []string        `json:"keywords"`
	Mood         model.MoodState `json:"mood"`
	MoodEmoji    string          `json:"moodEmoji"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
	ModCtr       int64           `json:"modCtr"`
	Created      sql.NullTime    `json:"created"`
	Updated      sql.NullTime    `json:"updated"`
}

const getPost = `-- name: GetPost :one
SELECT 
mr.UUId, User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Post d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesSocialMedia) GetPost(ctx context.Context, uuid uuid.UUID) (PostInfo, error) {
	row := q.db.QueryRowContext(ctx, getPost, uuid)
	var i PostInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Caption,
		&i.MessageBody,
		&i.Url,
		&i.ImageUri,
		&i.ThumbnailUri,
		pq.Array(&i.Keywords),
		&i.Mood,
		&i.MoodEmoji,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getPostbyUuId = `-- name: GetPostbyUuId :one
SELECT 
mr.UUId, User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Post d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesSocialMedia) GetPostbyUuId(ctx context.Context, uuid uuid.UUID) (PostInfo, error) {
	row := q.db.QueryRowContext(ctx, getPostbyUuId, uuid)
	var i PostInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Caption,
		&i.MessageBody,
		&i.Url,
		&i.ImageUri,
		&i.ThumbnailUri,
		pq.Array(&i.Keywords),
		&i.Mood,
		&i.MoodEmoji,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listPost = `-- name: ListPost:many
SELECT 
mr.UUId, User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Post d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListPostParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSocialMedia) ListPost(ctx context.Context, arg ListPostParams) ([]PostInfo, error) {
	rows, err := q.db.QueryContext(ctx, listPost, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PostInfo{}
	for rows.Next() {
		var i PostInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.Caption,
			&i.MessageBody,
			&i.Url,
			&i.ImageUri,
			&i.ThumbnailUri,
			pq.Array(&i.Keywords),
			&i.Mood,
			&i.MoodEmoji,
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

const updatePost = `-- name: UpdatePost :one
UPDATE Post SET 
User_Id = $2,
Caption = $3,
Message_Body = $4,
URL = $5,
Image_URI = $6,
Thumbnail_URI = $7,
Keywords = $8,
Mood = $9,
Mood_Emoji = $10,
Other_Info = $11
WHERE uuid = $1
RETURNING UUId, User_Id, Caption, Message_Body, URL, Image_URI, 
Thumbnail_URI, Keywords, Mood, Mood_Emoji, Other_Info
`

func (q *QueriesSocialMedia) UpdatePost(ctx context.Context, arg PostRequest) (model.Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,

		arg.Uuid,
		arg.UserId,
		arg.Caption,
		arg.MessageBody,
		arg.Url,
		arg.ImageUri,
		arg.ThumbnailUri,
		pq.Array(arg.Keywords),
		arg.Mood,
		arg.MoodEmoji,
		arg.OtherInfo,
	)
	var i model.Post
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Caption,
		&i.MessageBody,
		&i.Url,
		&i.ImageUri,
		&i.ThumbnailUri,
		pq.Array(&i.Keywords),
		&i.Mood,
		&i.MoodEmoji,
		&i.OtherInfo,
	)
	return i, err
}
