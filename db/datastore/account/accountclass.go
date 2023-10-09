package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountClass = `-- name: CreateAccountClass: one
INSERT INTO Account_Class 
 (Product_Id, Group_Id, Class_Id, Cur_Id, NonCur_Id, 
  BS_Acc_Id, IS_Acc_Id, Other_Info) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
ON CONFLICT(Product_ID, Group_ID, Class_ID)
DO UPDATE SET
	Cur_Id = EXCLUDED.Cur_Id,
	NonCur_Id = EXCLUDED.NonCur_Id, 
	BS_Acc_Id = EXCLUDED.BS_Acc_Id, 
	IS_Acc_Id = EXCLUDED.IS_Acc_Id, 
	Other_Info  = EXCLUDED.Other_Info
RETURNING 
  Id, UUId, Product_Id, Group_Id, Class_Id, Cur_Id, NonCur_Id, 
	BS_Acc_Id, IS_Acc_Id, Other_Info
`

type AccountClassRequest struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	ProductId int64          `json:"productId"`
	GroupId   int64          `json:"groupId"`
	ClassId   int64          `json:"classId"`
	CurId     int64          `json:"curId"`
	NoncurId  sql.NullInt64  `json:"noncurId"`
	BsAccId   sql.NullInt64  `json:"bsAccId"`
	IsAccId   sql.NullInt64  `json:"isAccId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountClass(ctx context.Context, arg AccountClassRequest) (model.AccountClass, error) {
	row := q.db.QueryRowContext(ctx, createAccountClass,
		arg.ProductId,
		arg.GroupId,
		arg.ClassId,
		arg.CurId,
		arg.NoncurId,
		arg.BsAccId,
		arg.IsAccId,
		arg.OtherInfo,
	)
	var i model.AccountClass
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.ClassId,
		&i.CurId,
		&i.NoncurId,
		&i.BsAccId,
		&i.IsAccId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountClass = `-- name: DeleteAccountClass :exec
DELETE FROM Account_Class
WHERE id = $1
`

func (q *QueriesAccount) DeleteAccountClass(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountClass, id)
	return err
}

type AccountClassInfo struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	ProductId int64          `json:"productId"`
	GroupId   int64          `json:"groupId"`
	ClassId   int64          `json:"classId"`
	CurId     int64          `json:"curId"`
	NoncurId  sql.NullInt64  `json:"noncurId"`
	BsAccId   sql.NullInt64  `json:"bsAccId"`
	IsAccId   sql.NullInt64  `json:"isAccId"`
	OtherInfo sql.NullString `json:"otherInfo"`
	ModCtr    int64          `json:"modCtr"`
	Created   sql.NullTime   `json:"created"`
	Updated   sql.NullTime   `json:"updated"`
}

const getAccountClass = `-- name: GetAccountClass :one
SELECT 
Id, mr.UUId, Product_Id, Group_Id, Class_Id, 
Cur_Id, NonCur_Id, BS_Acc_Id, IS_Acc_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Class d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountClass(ctx context.Context, id int64) (AccountClassInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountClass, id)
	var i AccountClassInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.ClassId,
		&i.CurId,
		&i.NoncurId,
		&i.BsAccId,
		&i.IsAccId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountClassbyUuid = `-- name: GetAccountClassbyUuid :one
SELECT 
Id, mr.UUId, Product_Id, Group_Id, Class_Id, 
Cur_Id, NonCur_Id, BS_Acc_Id, IS_Acc_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Class d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountClassbyUuid(ctx context.Context, uuid uuid.UUID) (AccountClassInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountClassbyUuid, uuid)
	var i AccountClassInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.ClassId,
		&i.CurId,
		&i.NoncurId,
		&i.BsAccId,
		&i.IsAccId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountClassbyKeys = `-- name: GetAccountClassbyKeys :one
SELECT 
  Id, mr.UUId, Product_Id, Group_Id, Class_Id, 
  Cur_Id, NonCur_Id, BS_Acc_Id, IS_Acc_Id, Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Class d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE 
   Product_ID  = $1 and Group_ID  = $2 and Class_ID = $3
LIMIT 1
`

func (q *QueriesAccount) GetAccountClassbyKeys(
	ctx context.Context, productID int64, groupID int64, classID int64) (AccountClassInfo, error) {

	row := q.db.QueryRowContext(
		// ctx, getAccountClassbyKeys, 2, 2, 2)
		ctx, getAccountClassbyKeys, productID, groupID, classID)
	var i AccountClassInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.ClassId,
		&i.CurId,
		&i.NoncurId,
		&i.BsAccId,
		&i.IsAccId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccountClass = `-- name: ListAccountClass:many
SELECT 
Id, mr.UUId, Product_Id, Group_Id, Class_Id, 
Cur_Id, NonCur_Id, BS_Acc_Id, IS_Acc_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Class d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountClassParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountClass(ctx context.Context, arg ListAccountClassParams) ([]AccountClassInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccountClass, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountClassInfo{}
	for rows.Next() {
		var i AccountClassInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ProductId,
			&i.GroupId,
			&i.ClassId,
			&i.CurId,
			&i.NoncurId,
			&i.BsAccId,
			&i.IsAccId,
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

const updateAccountClass = `-- name: UpdateAccountClass :one
UPDATE Account_Class SET 
Product_Id = $2,
Group_Id = $3,
Class_Id = $4,
Cur_Id = $5,
NonCur_Id = $6,
BS_Acc_Id = $7,
IS_Acc_Id = $8,
Other_Info = $9
WHERE id = $1
RETURNING Id, UUId, Product_Id, Group_Id, Class_Id, Cur_Id, NonCur_Id, 
BS_Acc_Id, IS_Acc_Id, Other_Info
`

func (q *QueriesAccount) UpdateAccountClass(ctx context.Context, arg AccountClassRequest) (model.AccountClass, error) {
	row := q.db.QueryRowContext(ctx, updateAccountClass,
		arg.Id,
		arg.ProductId,
		arg.GroupId,
		arg.ClassId,
		arg.CurId,
		arg.NoncurId,
		arg.BsAccId,
		arg.IsAccId,
		arg.OtherInfo,
	)
	var i model.AccountClass
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.ClassId,
		&i.CurId,
		&i.NoncurId,
		&i.BsAccId,
		&i.IsAccId,
		&i.OtherInfo,
	)
	return i, err
}
