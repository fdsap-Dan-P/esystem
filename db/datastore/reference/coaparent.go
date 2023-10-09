package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCoaParent = `-- name: CreateCoaParent: one
INSERT INTO COA_Parent (
Acc, COA_Seq, Title, Parent_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5
) RETURNING Id, UUId, Acc, COA_Seq, Title, Parent_Id, Other_Info
`

type CoaParentRequest struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	Acc       string         `json:"acc"`
	CoaSeq    sql.NullInt64  `json:"coaSeq"`
	Title     string         `json:"title"`
	ParentId  sql.NullInt64  `json:"parentId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesReference) CreateCoaParent(ctx context.Context, arg CoaParentRequest) (model.CoaParent, error) {
	row := q.db.QueryRowContext(ctx, createCoaParent,
		arg.Acc,
		arg.CoaSeq,
		arg.Title,
		arg.ParentId,
		arg.OtherInfo,
	)
	var i model.CoaParent
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.CoaSeq,
		&i.Title,
		&i.ParentId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCoaParent = `-- name: DeleteCoaParent :exec
DELETE FROM COA_Parent
WHERE id = $1
`

func (q *QueriesReference) DeleteCoaParent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCoaParent, id)
	return err
}

type CoaParentInfo struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	Acc       string         `json:"acc"`
	CoaSeq    sql.NullInt64  `json:"coaSeq"`
	Title     string         `json:"title"`
	ParentId  sql.NullInt64  `json:"parentId"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getCoaParent = `-- name: GetCoaParent :one
SELECT 
Id, mr.UUId, 
Acc, COA_Seq, Title, Parent_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM COA_Parent d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesReference) GetCoaParent(ctx context.Context, id int64) (CoaParentInfo, error) {
	row := q.db.QueryRowContext(ctx, getCoaParent, id)
	var i CoaParentInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.CoaSeq,
		&i.Title,
		&i.ParentId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCoaParentbyUuId = `-- name: GetCoaParentbyUuId :one
SELECT 
Id, mr.UUId, 
Acc, COA_Seq, Title, Parent_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM COA_Parent d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetCoaParentbyUuId(ctx context.Context, uuid uuid.UUID) (CoaParentInfo, error) {
	row := q.db.QueryRowContext(ctx, getCoaParentbyUuId, uuid)
	var i CoaParentInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.CoaSeq,
		&i.Title,
		&i.ParentId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listCoaParent = `-- name: ListCoaParent:many
SELECT 
Id, mr.UUId, 
Acc, COA_Seq, Title, Parent_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM COA_Parent d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCoaParentParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesReference) ListCoaParent(ctx context.Context, arg ListCoaParentParams) ([]CoaParentInfo, error) {
	rows, err := q.db.QueryContext(ctx, listCoaParent, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CoaParentInfo{}
	for rows.Next() {
		var i CoaParentInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Acc,
			&i.CoaSeq,
			&i.Title,
			&i.ParentId,
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

const updateCoaParent = `-- name: UpdateCoaParent :one
UPDATE COA_Parent SET 
Acc = $2,
COA_Seq = $3,
Title = $4,
Parent_Id = $5,
Other_Info = $6
WHERE id = $1
RETURNING Id, UUId, Acc, COA_Seq, Title, Parent_Id, Other_Info
`

func (q *QueriesReference) UpdateCoaParent(ctx context.Context, arg CoaParentRequest) (model.CoaParent, error) {
	row := q.db.QueryRowContext(ctx, updateCoaParent,
		arg.Id,
		arg.Acc,
		arg.CoaSeq,
		arg.Title,
		arg.ParentId,
		arg.OtherInfo,
	)
	var i model.CoaParent
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.CoaSeq,
		&i.Title,
		&i.ParentId,
		&i.OtherInfo,
	)
	return i, err
}
