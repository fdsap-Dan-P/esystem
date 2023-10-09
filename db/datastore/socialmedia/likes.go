package db

import (
	"context"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createLikes = `-- name: CreateLikes: one
INSERT INTO Likes (UUId, User_Id, Mood, Date_Liked) VALUES (
  $1, $2, $3, $4
) 
ON CONFLICT(UUId, User_Id) DO UPDATE SET
  Mood = excluded.Mood,
	Date_Liked = excluded.Date_Liked
RETURNING UUId, User_Id, Mood, Date_Liked
`

type LikesRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserId    int64     `json:"userId"`
	Mood      int32     `json:"mood"`
	DateLiked time.Time `json:"dateLiked"`
}

func (q *QueriesSocialMedia) CreateLikes(ctx context.Context, arg LikesRequest) (model.Likes, error) {
	row := q.db.QueryRowContext(ctx, createLikes,
		arg.Uuid,
		arg.UserId,
		arg.Mood,
		arg.DateLiked,
	)
	var i model.Likes
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Mood,
		&i.DateLiked,
	)
	return i, err
}

const deleteLikes = `-- name: DeleteLikes :exec
DELETE FROM Likes
WHERE uuid = $1 and User_Id = $2
`

func (q *QueriesSocialMedia) DeleteLikes(ctx context.Context, uuid uuid.UUID, userId int64) error {
	_, err := q.db.ExecContext(ctx, deleteLikes, uuid, userId)
	return err
}

type LikesInfo struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserId    int64     `json:"userId"`
	Mood      int32     `json:"mood"`
	DateLiked time.Time `json:"dateLiked"`
}

const getLikes = `-- name: GetLikes :one
SELECT UUId, User_Id, Mood, Date_Liked
FROM Likes d 
WHERE uuid = $1 and User_Id = $2 LIMIT 1
`

func (q *QueriesSocialMedia) GetLikes(ctx context.Context, uuid uuid.UUID, userId int64) (LikesInfo, error) {
	row := q.db.QueryRowContext(ctx, getLikes, uuid, userId)
	var i LikesInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Mood,
		&i.DateLiked,
	)
	return i, err
}

const listLikes = `-- name: ListLikes:many
SELECT UUId, User_Id, Mood, Date_Liked FROM Likes d 
WHERE UUId = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListLikesParams struct {
	Uuid   uuid.UUID `json:"uuid"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *QueriesSocialMedia) ListLikes(ctx context.Context, arg ListLikesParams) ([]LikesInfo, error) {
	rows, err := q.db.QueryContext(ctx, listLikes, arg.Uuid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LikesInfo{}
	for rows.Next() {
		var i LikesInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.Mood,
			&i.DateLiked,
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

const updateLikes = `-- name: UpdateLikes :one
UPDATE Likes SET 
	Mood = $3,
	Date_Liked = $4
WHERE uuid = $1 and User_Id = $2
RETURNING UUId, User_Id, Mood, Date_Liked
`

func (q *QueriesSocialMedia) UpdateLikes(ctx context.Context, arg LikesRequest) (model.Likes, error) {
	row := q.db.QueryRowContext(ctx, updateLikes,
		arg.Uuid,
		arg.UserId,
		arg.Mood,
		arg.DateLiked,
	)
	var i model.Likes
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.Mood,
		&i.DateLiked,
	)
	return i, err
}
