package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createRelation = `-- name: CreateRelation: one
INSERT INTO Relation (
  IIID, Series, Relation_IIID, Type_ID, Relation_Date, Other_Info) 
VALUES ($1, $2, $3, $4, $5, $6) 
ON CONFLICT(IIID, Relation_IIID, Type_ID) 
DO UPDATE SET
	IIID = excluded.IIID,  
	Series = excluded.Series,   
	Relation_IIID = excluded.Relation_IIID,   
	Type_ID = excluded.Type_ID,   
	Relation_Date = excluded.Relation_Date
RETURNING UUID, IIID, Series, Relation_IIID, Type_ID, Relation_Date, Other_Info
`

type RelationRequest struct {
	Uuid         uuid.UUID      `json:"uuid"`
	Iiid         int64          `json:"iiid"`
	Series       int16          `json:"series"`
	RelationIiid int64          `json:"relationIiid"`
	TypeId       int64          `json:"typeId"`
	RelationDate sql.NullTime   `json:"relationDate"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateRelation(ctx context.Context, arg RelationRequest) (model.Relation, error) {
	row := q.db.QueryRowContext(ctx, createRelation,
		arg.Iiid,
		arg.Series,
		arg.RelationIiid,
		arg.TypeId,
		arg.RelationDate,
		arg.OtherInfo,
	)
	var i model.Relation
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.RelationIiid,
		&i.TypeId,
		&i.RelationDate,
		&i.OtherInfo,
	)
	return i, err
}

const deleteRelation = `-- name: DeleteRelation :exec
DELETE FROM Relation
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteRelation(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRelation, uuid)
	return err
}

type RelationInfo struct {
	Uuid         uuid.UUID      `json:"uuid"`
	Iiid         int64          `json:"iiid"`
	Series       int16          `json:"series"`
	RelationIiid int64          `json:"relationIiid"`
	TypeId       int64          `json:"typeId"`
	RelationDate sql.NullTime   `json:"relationDate"`
	OtherInfo    sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getRelation = `-- name: GetRelation :one
SELECT 
  mr.UUID, IIID, Series, Relation_IIID, Type_ID, Relation_Date, Other_Info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUID
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetRelation(ctx context.Context, uuid uuid.UUID) (RelationInfo, error) {
	row := q.db.QueryRowContext(ctx, getRelation, uuid)
	var i RelationInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.RelationIiid,
		&i.TypeId,
		&i.RelationDate,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getRelationbyUuId = `-- name: GetRelationbyUuId :one
SELECT 
mr.UUID, IIID, Series, Relation_IIID, Type_ID, Relation_Date, Other_Info, 
mr.Mod_Ctr, mr.Created, mr.Updated
FROM Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUID
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetRelationbyUuId(ctx context.Context, uuid uuid.UUID) (RelationInfo, error) {
	row := q.db.QueryRowContext(ctx, getRelationbyUuId, uuid)
	var i RelationInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.RelationIiid,
		&i.TypeId,
		&i.RelationDate,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

// const getRelationbyName = `-- name: GetRelationbyName :one
// SELECT
// mr.UUId, IIId,
// Series, Relation_IIId, Type_Id, Relation_Date, Other_Info
// ,mr.Mod_Ctr, mr.Created, mr.Updated
// FROM Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUId
// WHERE Title = $1 LIMIT 1
// `

// func (q *QueriesIdentity) GetRelationbyName(ctx context.Context, name string) (RelationInfo, error) {
// 	row := q.db.QueryRowContext(ctx, getRelationbyName, name)
// 	var i RelationInfo
// 	err := row.Scan(
// 		&i.Uuid,
// 		&i.Iiid,
// 		&i.Series,
// 		&i.RelationIiid,
// 		&i.TypeId,
// 		&i.RelationDate,
// 		&i.OtherInfo,

// 		&i.ModCtr,
// 		&i.Created,
// 		&i.Updated,
// 	)
// 	return i, err
// }

const listRelation = `-- name: ListRelation:many
SELECT 
  mr.UUID, IIID, Series, Relation_IIID, Type_ID, Relation_Date, Other_Info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Relation d INNER JOIN Main_Record mr on mr.UUId = d.UUID
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListRelationParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListRelation(ctx context.Context, arg ListRelationParams) ([]RelationInfo, error) {
	rows, err := q.db.QueryContext(ctx, listRelation, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RelationInfo{}
	for rows.Next() {
		var i RelationInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.RelationIiid,
			&i.TypeId,
			&i.RelationDate,
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

const updateRelation = `-- name: UpdateRelation :one
UPDATE Relation SET 
IIID = $2,
Series = $3,
Relation_IIID = $4,
Type_ID = $5,
Relation_Date = $6,
Other_Info = $7
WHERE uuid = $1
RETURNING UUID, IIID, Series, Relation_IIID, Type_ID, Relation_Date, 
Other_Info
`

func (q *QueriesIdentity) UpdateRelation(ctx context.Context, arg RelationRequest) (model.Relation, error) {
	row := q.db.QueryRowContext(ctx, updateRelation,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.RelationIiid,
		arg.TypeId,
		arg.RelationDate,
		arg.OtherInfo,
	)
	var i model.Relation
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.RelationIiid,
		&i.TypeId,
		&i.RelationDate,
		&i.OtherInfo,
	)
	return i, err
}
