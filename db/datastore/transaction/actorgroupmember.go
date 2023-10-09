package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createActorGroupMember = `-- name: CreateActorGroupMember: one
INSERT INTO Actor_Group_Member(
   uuid, actor_group_id, actor_group, user_id, other_info )
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(UUID) 
DO UPDATE SET
	actor_group_id =  EXCLUDED.actor_group_id,
	actor_group =  EXCLUDED.actor_group,
	user_id =  EXCLUDED.user_id,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, actor_group_id, actor_group, user_id, other_info`

type ActorGroupMemberRequest struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ActorGroupId int64          `json:"actorGroupId"`
	ActorGroup   string         `json:"actorGroup"`
	UserId       int64          `json:"userId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateActorGroupMember(ctx context.Context, arg ActorGroupMemberRequest) (model.ActorGroupMember, error) {
	row := q.db.QueryRowContext(ctx, createActorGroupMember,
		arg.Uuid,
		arg.ActorGroupId,
		arg.ActorGroup,
		arg.UserId,
		arg.OtherInfo,
	)
	var i model.ActorGroupMember
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.UserId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteActorGroupMember = `-- name: DeleteActorGroupMember :exec
DELETE FROM Actor_Group_Member
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteActorGroupMember(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteActorGroupMember, uuid)
	return err
}

type ActorGroupMemberInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ActorGroupId int64          `json:"actorGroupId"`
	ActorGroup   string         `json:"actorGroup"`
	UserId       int64          `json:"userId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

const actorGroupMemberSQL = `-- name: ActorGroupMemberSQL :one
SELECT
  Id, mr.UUID, actor_group_id, actor_group, user_id, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Actor_Group_Member d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateActorGroupMember(q *QueriesTransaction, ctx context.Context, sql string) (ActorGroupMemberInfo, error) {
	var i ActorGroupMemberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.UserId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateActorGroupMembers(q *QueriesTransaction, ctx context.Context, sql string) ([]ActorGroupMemberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActorGroupMemberInfo{}
	for rows.Next() {
		var i ActorGroupMemberInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ActorGroupId,
			&i.ActorGroup,
			&i.UserId,
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

func (q *QueriesTransaction) GetActorGroupMember(ctx context.Context, id int64) (ActorGroupMemberInfo, error) {
	return populateActorGroupMember(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", actorGroupMemberSQL, id))
}

func (q *QueriesTransaction) GetActorGroupMemberbyUuid(ctx context.Context, uuid uuid.UUID) (ActorGroupMemberInfo, error) {
	return populateActorGroupMember(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", actorGroupMemberSQL, uuid))
}

type ListActorGroupMemberParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesTransaction) ListActorGroupMember(ctx context.Context, arg ListActorGroupMemberParams) ([]ActorGroupMemberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			actorGroupMemberSQL, arg.Limit, arg.Offset)
	} else {
		sql = actorGroupMemberSQL
	}
	return populateActorGroupMembers(q, ctx, sql)
}

const updateActorGroupMember = `-- name: UpdateActorGroupMember :one
UPDATE Actor_Group_Member SET 
uuid = $2,
actor_group_id = $3,
actor_group = $4,
user_id = $5,
other_info = $6
WHERE id = $1
RETURNING id, uuid, actor_group_id, actor_group, user_id, other_info
`

func (q *QueriesTransaction) UpdateActorGroupMember(ctx context.Context, arg ActorGroupMemberRequest) (model.ActorGroupMember, error) {
	row := q.db.QueryRowContext(ctx, updateActorGroupMember,
		arg.Id,
		arg.Uuid,
		arg.ActorGroupId,
		arg.ActorGroup,
		arg.UserId,
		arg.OtherInfo,
	)
	var i model.ActorGroupMember
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.UserId,
		&i.OtherInfo,
	)
	return i, err
}
