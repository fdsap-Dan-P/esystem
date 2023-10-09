package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	// "log"

	"simplebank/model"

	"github.com/google/uuid"
)

const createReference = `-- name: CreateReference :one
INSERT INTO Reference(
  Code, Short_Name, Title, Parent_Id, Type_Id, Remark, Other_Info
  ) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
ON CONFLICT(lower(Ref_Type), Code) DO UPDATE SET
  Short_Name = EXCLUDED.Short_Name, 
  Title = EXCLUDED.Title, 
  Parent_Id = EXCLUDED.Parent_Id, 
  Type_Id = EXCLUDED.Type_Id, 
  Remark = EXCLUDED.Remark, 
  Other_Info = EXCLUDED.Other_Info
RETURNING id, uuid, code, short_name, title, parent_id, Type_Id, remark, other_info
`

type ReferenceRequest struct {
	Id        int64          `json:"id"`
	Code      int64          `json:"code"`
	ShortName sql.NullString `json:"shortName"`
	Title     string         `json:"title"`
	ParentId  sql.NullInt64  `json:"parentId"`
	TypeId    int64          `json:"TypeId"`
	Remark    sql.NullString `json:"remark"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesReference) CreateReference(ctx context.Context, arg ReferenceRequest) (model.Reference, error) {
	row_ := q.db.QueryRowContext(ctx, createReference,
		arg.Code,
		arg.ShortName,
		arg.Title,
		arg.ParentId,
		arg.TypeId,
		arg.Remark,
		arg.OtherInfo,
	)
	var i model.Reference
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Title,
		&i.ParentId,
		&i.TypeId,
		&i.Remark,
		&i.OtherInfo,
	)
	return i, err
}

const deleteReference = `-- name: DeleteReference :exec
DELETE FROM Reference
WHERE id = $1
`

func (q *QueriesReference) DeleteReference(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReference, id)
	return err
}

const getReference = `-- name: GetReference :one
SELECT 
  Id, UUId, Code, Short_Name, Title, Parent_Id, Type_Id, Remark, Other_Info
FROM Reference
WHERE id = $1 LIMIT 1
`

func (q *QueriesReference) GetReference(ctx context.Context, id int64) (model.Reference, error) {
	row_ := q.db.QueryRowContext(ctx, getReference, id)
	var i model.Reference
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Title,
		&i.ParentId,
		&i.TypeId,
		&i.Remark,
		&i.OtherInfo,
	)
	return i, err
}

func populateData(q *QueriesReference, ctx context.Context, sql string) (ReferenceInfo, error) {
	var i ReferenceInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.Code,
		&i.RefType,
		&i.ParentId,
		&i.ParentUuid,
		&i.ParentCode,
		&i.Parent,
		&i.ShortName,
		&i.Title,
		&i.ParentTypeId,
		&i.ParentRefTypeuuid,
		&i.ParentRefType,
		&i.Remark,
		&i.ModCtr,
		&i.OtherInfo,
		&i.Created,
		&i.Updated,
		&i.vecSimpleName,
	)
	return i, err
}

func populateDatas(q *QueriesReference, ctx context.Context, sql string) ([]ReferenceInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []ReferenceInfo{}
	for rows.Next() {
		var i ReferenceInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TypeId,
			&i.Code,
			&i.RefType,
			&i.ParentId,
			&i.ParentUuid,
			&i.ParentCode,
			&i.Parent,
			&i.ShortName,
			&i.Title,
			&i.ParentTypeId,
			&i.ParentRefTypeuuid,
			&i.ParentRefType,
			&i.Remark,
			&i.ModCtr,
			&i.OtherInfo,
			&i.Created,
			&i.Updated,
			&i.vecSimpleName,
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

const selectSQL = `-- name: selectSQL
SELECT 
  rf.Id, mr.UUId,
  typ.Id AS Type_Id, rf.Code, typ.Title AS Ref_Type,

  p.Id AS Parent_Id, p.UUId AS Parent_Uuid, p.Code AS Parent_Code,
  p.Title AS Parent, rf.Short_Name, rf.Title,

  ptyp.Id AS Parent_Type_Id, ptyp.UUId AS Parent_Ref_TypeUUId,
  ptyp.Title AS Parent_Ref_Type,
  rf.Remark,

  mr.Mod_Ctr,
  rf.Other_Info,
  mr.Created,
  mr.Updated, rf.Vec_Simple_Name
FROM Reference rf
INNER JOIN Main_Record mr on mr.UUId = rf.UUId
INNER JOIN Reference_Type typ ON rf.Type_Id = typ.Id
LEFT JOIN Reference p ON rf.Parent_Id = p.Id
LEFT JOIN Reference_Type ptyp ON p.Type_Id = ptyp.Id
`

type ReferenceInfo struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	TypeId            int64          `json:"TypeId"`
	Code              int64          `json:"code"`
	RefType           string         `json:"refType"`
	ParentId          sql.NullInt64  `json:"parentId"`
	ParentUuid        uuid.UUID      `json:"parentUuid"`
	ParentCode        sql.NullInt64  `json:"parent_code"`
	Parent            sql.NullString `json:"parent"`
	ShortName         sql.NullString `json:"shortName"`
	Title             string         `json:"title"`
	ParentTypeId      sql.NullInt64  `json:"parent_TypeId"`
	ParentRefTypeuuid uuid.UUID      `json:"parentTefTypeUuid"`
	ParentRefType     sql.NullString `json:"parentTefType"`
	Remark            sql.NullString `json:"remark"`
	ModCtr            int64          `json:"modCtr"`
	OtherInfo         sql.NullString `json:"otherInfo"`
	Created           sql.NullTime   `json:"created"`
	Updated           sql.NullTime   `json:"updated"`
	vecSimpleName     string         `json:"vecSimpleName"`
}

func (q *QueriesReference) GetReferenceInfo(ctx context.Context, id int64) (ReferenceInfo, error) {
	return populateData(q, ctx,
		fmt.Sprintf("%s WHERE rf.Id = %v", selectSQL, id))
}

func (q *QueriesReference) GetReferenceInfobyCode(ctx context.Context, code sql.NullString) (ReferenceInfo, error) {
	return populateData(q, ctx,
		fmt.Sprintf("%s WHERE rf.code = %v", selectSQL, code))
}

func (q *QueriesReference) GetReferenceInfobyUuId(ctx context.Context, uuid uuid.UUID) (ReferenceInfo, error) {
	return populateData(q, ctx,
		fmt.Sprintf("%s WHERE rf.uuid = '%v'", selectSQL, uuid))
}

func (q *QueriesReference) GetReferenceInfobyTitle(
	ctx context.Context, refType string, parentId int64, title string) (ReferenceInfo, error) {
	log.Println(fmt.Sprintf(`%s WHERE LOWER(typ.Title) = '%s'
		and COALESCE(rf.Parent_Id,0) = %v
		and LOWER(rf.Title) = '%v' `, selectSQL, strings.ToLower(refType), parentId, strings.ToLower(title)))
	return populateData(q, ctx,
		fmt.Sprintf(`%s WHERE LOWER(typ.Title) = '%v'
		and COALESCE(rf.Parent_Id,0) = %v
		and LOWER(rf.Title) = '%v' `, selectSQL, strings.ToLower(refType), parentId, strings.ToLower(title)))
}

type ListReferenceParams struct {
	RefType string `json:"ref_type"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *QueriesReference) ListReference(ctx context.Context, arg ListReferenceParams) ([]ReferenceInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE lower(typ.Title) = '%s' LIMIT %d OFFSET %d",
			selectSQL, strings.ToLower(arg.RefType), arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE lower(typ.Title) = '%s' ", selectSQL, strings.ToLower(arg.RefType))
	}
	return populateDatas(q, ctx, sql)
}

// const searchReference = `-- name: ListReference :many
// SELECT id, uuid, code, short_name, title, parent_id, Type_Id, remark, other_info
// FROM Reference rf
// WHERE $1
// ORDER BY id
// `

type FilterParams struct {
	Filter string `json:"filter"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type ReferenceSearchParams struct {
	Search string `json:"filter"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *QueriesReference) GetReferenceFilter(ctx context.Context, arg FilterParams) ([]ReferenceInfo, error) {
	sql := fmt.Sprintf("SELECT * FROM ( %s ) d WHERE %s ORDER BY id", selectSQL, arg.Filter)
	log.Println(sql)
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sql, arg.Limit, arg.Offset)
	}
	return populateDatas(q, ctx, sql)
}

func (q *QueriesReference) ReferenceSearch(ctx context.Context, arg ReferenceSearchParams) ([]ReferenceInfo, error) {
	sql := fmt.Sprintf("%s FROM ( %s ) d, %s) s WHERE Vec_Simple_Name @@ loc_or ORDER BY %s",
		"SELECT d.* ",
		selectSQL,
		fmt.Sprintf("(SELECT plainto_tsquery_or('%s') loc_or", arg.Search),
		"ts_rank_cd(Vec_Simple_Name, loc_or, 32) desc")
	log.Println(sql)
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sql, arg.Limit, arg.Offset)
	}
	return populateDatas(q, ctx, sql)
}

const updateReference = `-- name: UpdateReference :one
UPDATE Reference SET 
  Code        = $2,
  Short_Name  = $3,
  Title       = $4,
  Parent_Id   = $5,
  Type_Id     = $6,
  Remark      = $7,
  Other_Info  = $8
WHERE id = $1
RETURNING id, uuid, code, short_name, title, parent_id, Type_Id, remark, other_info
`

func (q *QueriesReference) UpdateReference(ctx context.Context, arg ReferenceRequest) (model.Reference, error) {
	row_ := q.db.QueryRowContext(ctx, updateReference,
		arg.Id,
		arg.Code,
		arg.ShortName,
		arg.Title,
		arg.ParentId,
		arg.TypeId,
		arg.Remark,
		arg.OtherInfo,
	)
	var i model.Reference
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Title,
		&i.ParentId,
		&i.TypeId,
		&i.Remark,
		&i.OtherInfo,
	)
	return i, err
}
