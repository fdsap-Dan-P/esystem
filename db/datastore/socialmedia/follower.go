package db

import (
	"context"
	"time"

	"simplebank/model"

	"github.com/google/uuid"
)

const createFollower = `-- name: CreateFollower: one
INSERT INTO Follower (
  UUID, User_Id, Follower_Id, Date_Followed, is_Follower
) VALUES (
  $1, $2, $3, $4, $5
) 
ON CONFLICT(User_Id, Follower_Id) DO UPDATE SET
	Date_Followed = excluded.Date_Followed,
	is_Follower = excluded.is_Follower
RETURNING Uuid, User_Id, Follower_Id, Date_Followed, is_Follower
`

type FollowerRequest struct {
	Uuid         uuid.UUID `json:"uuid"`
	UserId       int64     `json:"userId"`
	FollowerId   int64     `json:"followerId"`
	DateFollowed time.Time `json:"dateFollowed"`
	IsFollower   bool      `json:"isFollower"`
}

func (q *QueriesSocialMedia) CreateFollower(ctx context.Context, arg FollowerRequest) (model.Follower, error) {
	row := q.db.QueryRowContext(ctx, createFollower,
		arg.Uuid,
		arg.UserId,
		arg.FollowerId,
		arg.DateFollowed,
		arg.IsFollower,
	)
	var i model.Follower
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.FollowerId,
		&i.DateFollowed,
		&i.IsFollower,
	)
	return i, err
}

const deleteFollower = `-- name: DeleteFollower :exec
DELETE FROM Follower
WHERE user_id = $1 and follower_id = $2
`

func (q *QueriesSocialMedia) DeleteFollower(ctx context.Context, userId int64, followedId int64) error {
	_, err := q.db.ExecContext(ctx, deleteFollower, userId, followedId)
	return err
}

type FollowerInfo struct {
	Uuid         uuid.UUID `json:"uuid"`
	UserId       int64     `json:"userId"`
	FollowerId   int64     `json:"followerId"`
	DateFollowed time.Time `json:"dateFollowed"`
	IsFollower   bool      `json:"isFollower"`
}

const getFollower = `-- name: GetFollower :one
SELECT UUID, User_Id, Follower_Id, Date_Followed, is_Follower
FROM Follower d 
WHERE User_Id = $1 and Follower_Id = $2 LIMIT 1
`

func (q *QueriesSocialMedia) GetFollower(ctx context.Context, userId int64, followedId int64) (FollowerInfo, error) {
	row := q.db.QueryRowContext(ctx, getFollower, userId, followedId)
	var i FollowerInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.FollowerId,
		&i.DateFollowed,
		&i.IsFollower,
	)
	return i, err
}

const listFollower = `-- name: ListFollower:many
SELECT 
	UUID, User_Id, Follower_Id, Date_Followed, is_Follower
FROM Follower d 
WHERE User_Id = $1
ORDER BY Follower_Id
LIMIT $2
OFFSET $3
`

type ListFollowerParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesSocialMedia) ListFollower(ctx context.Context, arg ListFollowerParams) ([]FollowerInfo, error) {
	rows, err := q.db.QueryContext(ctx, listFollower, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FollowerInfo{}
	for rows.Next() {
		var i FollowerInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.FollowerId,
			&i.DateFollowed,
			&i.IsFollower,
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

const updateFollower = `-- name: UpdateFollower :one
UPDATE Follower SET 
    User_Id = $2,
	Follower_Id = $3,
	Date_Followed = $4,
	is_Follower = $5
WHERE Uuid = $1
RETURNING Uuid, User_Id, Follower_Id, Date_Followed, is_Follower
`

func (q *QueriesSocialMedia) UpdateFollower(ctx context.Context, arg FollowerRequest) (model.Follower, error) {
	row := q.db.QueryRowContext(ctx, updateFollower,
		arg.Uuid,
		arg.UserId,
		arg.FollowerId,
		arg.DateFollowed,
		arg.IsFollower,
	)
	var i model.Follower
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.FollowerId,
		&i.DateFollowed,
		&i.IsFollower,
	)
	return i, err
}
