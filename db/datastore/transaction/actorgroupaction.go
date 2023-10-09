package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createActorGroupAction = `-- name: CreateActorGroupAction: one
INSERT INTO Actor_Group_Action(
   uuid, actor_group_id, actor_group, ticket_type_action_id, ticket_type, Action_Desc, other_info)
SELECT 
  $1, grp.Id actor_group_id, grp.Title actor_group, act.Id ticket_type_action_id, 
  ticType.Title ticket_type, actref.title Action_Desc, $4
FROM 
  Reference grp, Ticket_Type_Action act, 
  Product_Ticket_Type prod, Reference ticType, 
  Reference actref
WHERE 
  grp.ID = $2 and act.Id = $3 
  and prod.id = act.product_ticket_type_id 
  and ticType.ID = prod.ticket_type_id 
  and actref.ID = act.action_id 
ON CONFLICT (UUID)
DO UPDATE SET
	actor_group_id =  EXCLUDED.actor_group_id,
	actor_group =  EXCLUDED.actor_group,
	ticket_type_action_id =  EXCLUDED.ticket_type_action_id,
	ticket_type =  EXCLUDED.ticket_type,
	Action_Desc =  EXCLUDED.Action_Desc,
	other_info =  EXCLUDED.other_info
RETURNING UUID, actor_group_id, actor_group, ticket_type_action_id, ticket_type, Action_Desc, other_info`

type ActorGroupActionRequest struct {
	Uuid               uuid.UUID      `json:"uuid"`
	ActorGroupId       int64          `json:"actorGroupId"`
	TicketTypeActionId int64          `json:"ticketTypeActionId"`
	OtherInfo          sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateActorGroupAction(ctx context.Context, arg ActorGroupActionRequest) (model.ActorGroupAction, error) {
	row := q.db.QueryRowContext(ctx, createActorGroupAction,
		arg.Uuid,
		arg.ActorGroupId,
		arg.TicketTypeActionId,
		arg.OtherInfo,
	)
	var i model.ActorGroupAction
	err := row.Scan(
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.TicketTypeActionId,
		&i.TicketType,
		&i.ActionDesc,
		&i.OtherInfo,
	)
	return i, err
}

const deleteActorGroupAction = `-- name: DeleteActorGroupAction :exec
DELETE FROM Actor_Group_Action
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteActorGroupAction(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteActorGroupAction, uuid)
	return err
}

type ActorGroupActionInfo struct {
	Uuid               uuid.UUID      `json:"uuid"`
	ActorGroupId       int64          `json:"actorGroupId"`
	ActorGroup         string         `json:"actorGroup"`
	TicketTypeActionId int64          `json:"ticketTypeActionId"`
	TicketType         string         `json:"ticketType"`
	ActionDesc         string         `json:"actionDesc"`
	OtherInfo          sql.NullString `json:"otherInfo"`
	ModCtr             int64          `json:"modCtr"`
	Created            sql.NullTime   `json:"created"`
	Updated            sql.NullTime   `json:"updated"`
}

const actorGroupActionSQL = `-- name: ActorGroupActionSQL :one
SELECT
mr.UUID, actor_group_id, actor_group, ticket_type_action_id, ticket_type, Action_Desc, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Actor_Group_Action d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateActorGroupAction(q *QueriesTransaction, ctx context.Context, sql string) (ActorGroupActionInfo, error) {
	var i ActorGroupActionInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.TicketTypeActionId,
		&i.TicketType,
		&i.ActionDesc,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateActorGroupActions(q *QueriesTransaction, ctx context.Context, sql string) ([]ActorGroupActionInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActorGroupActionInfo{}
	for rows.Next() {
		var i ActorGroupActionInfo
		err := rows.Scan(
			&i.Uuid,
			&i.ActorGroupId,
			&i.ActorGroup,
			&i.TicketTypeActionId,
			&i.TicketType,
			&i.ActionDesc,
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

func (q *QueriesTransaction) GetActorGroupActionbyUuid(ctx context.Context, uuid uuid.UUID) (ActorGroupActionInfo, error) {
	return populateActorGroupAction(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", actorGroupActionSQL, uuid))
}

type ListActorGroupActionParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesTransaction) ListActorGroupAction(ctx context.Context, arg ListActorGroupActionParams) ([]ActorGroupActionInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			actorGroupActionSQL, arg.Limit, arg.Offset)
	} else {
		sql = actorGroupActionSQL
	}
	return populateActorGroupActions(q, ctx, sql)
}

const updateActorGroupAction = `-- name: UpdateActorGroupAction :one
UPDATE Actor_Group_Action SET 
	actor_group_id = a.actor_group_id2,
	actor_group = a.actor_group,
	ticket_type_action_id = a.ticket_type_action_id,
	ticket_type = a.ticket_type,
	Action_Desc = a.Action_Desc,
	other_info = $4
FROM	
 (SELECT grp.Id actor_group_id2, grp.Title actor_group, act.Id ticket_type_action_id, 
    ticType.Title ticket_type, actref.title Action_Desc
  FROM 
    Reference grp, Ticket_Type_Action act, 
    Product_Ticket_Type prod, Reference ticType, 
    Reference actref
  WHERE 
    grp.ID = $2 and act.Id = $3 
    and prod.id = act.product_ticket_type_id 
    and ticType.ID = prod.ticket_type_id 
    and actref.ID = act.action_id) a
  WHERE uuid = $1
  RETURNING uuid, Actor_Group_Action.actor_group_id, Actor_Group_Action.actor_group, Actor_Group_Action.ticket_type_action_id, 
    Actor_Group_Action.ticket_type, Actor_Group_Action.Action_Desc, Actor_Group_Action.other_info
`

func (q *QueriesTransaction) UpdateActorGroupAction(ctx context.Context, arg ActorGroupActionRequest) (model.ActorGroupAction, error) {
	row := q.db.QueryRowContext(ctx, updateActorGroupAction,
		arg.Uuid,
		arg.ActorGroupId,
		arg.TicketTypeActionId,
		arg.OtherInfo,
	)
	var i model.ActorGroupAction
	err := row.Scan(
		&i.Uuid,
		&i.ActorGroupId,
		&i.ActorGroup,
		&i.TicketTypeActionId,
		&i.TicketType,
		&i.ActionDesc,
		&i.OtherInfo,
	)
	return i, err
}
