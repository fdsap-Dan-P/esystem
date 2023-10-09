package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTrnHeadRelation = `-- name: CreateTrnHeadRelation: one
INSERT INTO Trn_Head_Relation (
Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
) VALUES (
$1, $2, $3, $4, $5
) RETURNING UUId, Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
`

type TrnHeadRelationRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	TrnHeadId int64          `json:"trnHeadId"`
	RelatedId int64          `json:"relatedId"`
	TypeId    int64          `json:"typeId"`
	Remarks   sql.NullString `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTrnHeadRelation(ctx context.Context, arg TrnHeadRelationRequest) (model.TrnHeadRelation, error) {
	row := q.db.QueryRowContext(ctx, createTrnHeadRelation,
		arg.TrnHeadId,
		arg.RelatedId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TrnHeadRelation
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.RelatedId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTrnHeadRelation = `-- name: DeleteTrnHeadRelation :exec
DELETE FROM Trn_Head_Relation
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTrnHeadRelation(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHeadRelation, uuid)
	return err
}

type TrnHeadRelationInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	TrnHeadId int64          `json:"trnHeadId"`
	RelatedId int64          `json:"relatedId"`
	TypeId    int64          `json:"typeId"`
	Remarks   sql.NullString `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getTrnHeadRelation = `-- name: GetTrnHeadRelation :one
SELECT 
mr.UUId, 
Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesTransaction) GetTrnHeadRelation(ctx context.Context, uuid uuid.UUID) (TrnHeadRelationInfo, error) {
	row := q.db.QueryRowContext(ctx, getTrnHeadRelation, uuid)
	var i TrnHeadRelationInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.RelatedId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getTrnHeadRelationbyUuid = `-- name: GetTrnHeadRelationbyUuid :one
SELECT 
mr.UUId, 
Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesTransaction) GetTrnHeadRelationbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadRelationInfo, error) {
	row := q.db.QueryRowContext(ctx, getTrnHeadRelationbyUuid, uuid)
	var i TrnHeadRelationInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.RelatedId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listTrnHeadRelation = `-- name: ListTrnHeadRelation:many
SELECT 
mr.UUId, 
Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Trn_Head_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListTrnHeadRelationParams struct {
	TrnHeadId int64 `json:"trnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHeadRelation(ctx context.Context, arg ListTrnHeadRelationParams) ([]TrnHeadRelationInfo, error) {
	rows, err := q.db.QueryContext(ctx, listTrnHeadRelation, arg.TrnHeadId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TrnHeadRelationInfo{}
	for rows.Next() {
		var i TrnHeadRelationInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.RelatedId,
			&i.TypeId,
			&i.Remarks,
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

const updateTrnHeadRelation = `-- name: UpdateTrnHeadRelation :one
UPDATE Trn_Head_Relation SET 
Trn_Head_Id = $2,
Related_Id = $3,
Type_Id = $4,
Remarks = $5,
Other_Info = $6
WHERE uuid = $1
RETURNING UUId, Trn_Head_Id, Related_Id, Type_Id, Remarks, Other_Info
`

func (q *QueriesTransaction) UpdateTrnHeadRelation(ctx context.Context, arg TrnHeadRelationRequest) (model.TrnHeadRelation, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHeadRelation,

		arg.Uuid,
		arg.TrnHeadId,
		arg.RelatedId,
		arg.TypeId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TrnHeadRelation
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.RelatedId,
		&i.TypeId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
