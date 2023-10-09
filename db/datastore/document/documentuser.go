package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createDocumentUser = `-- name: CreateDocumentUser: one
INSERT INTO Document_User(
   document_id, user_id, access_code, other_info )
VALUES ($1, $2, $3, $4)
RETURNING UUID, document_id, user_id, access_code, other_info`

type DocumentUserRequest struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	UserId     int64          `json:"userId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

func (q *QueriesDocument) CreateDocumentUser(ctx context.Context, arg DocumentUserRequest) (model.DocumentUser, error) {
	row := q.db.QueryRowContext(ctx, createDocumentUser,
		arg.DocumentId,
		arg.UserId,
		arg.AccessCode,
		arg.OtherInfo,
	)
	var i model.DocumentUser
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.UserId,
		&i.AccessCode,
		&i.OtherInfo,
	)
	return i, err
}

const deleteDocumentUser = `-- name: DeleteDocumentUser :exec
DELETE FROM Document_User
WHERE uuid = $1
`

func (q *QueriesDocument) DeleteDocumentUser(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDocumentUser, uuid)
	return err
}

type DocumentUserInfo struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	UserId     int64          `json:"userId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
	ModCtr     int64          `json:"modCtr"`
	Created    sql.NullTime   `json:"created"`
	Updated    sql.NullTime   `json:"updated"`
}

const documentUserSQL = `-- name: DocumentUserSQL :one
SELECT
mr.UUID, document_id, user_id, access_code, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Document_User d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateDocumentUser(q *QueriesDocument, ctx context.Context, sql string) (DocumentUserInfo, error) {
	var i DocumentUserInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.UserId,
		&i.AccessCode,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateDocumentUsers(q *QueriesDocument, ctx context.Context, sql string) ([]DocumentUserInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DocumentUserInfo{}
	for rows.Next() {
		var i DocumentUserInfo
		err := rows.Scan(
			&i.Uuid,
			&i.DocumentId,
			&i.UserId,
			&i.AccessCode,
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

func (q *QueriesDocument) GetDocumentUser(ctx context.Context, id int64) (DocumentUserInfo, error) {
	return populateDocumentUser(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", documentUserSQL, id))
}

func (q *QueriesDocument) GetDocumentUserbyUuid(ctx context.Context, uuid uuid.UUID) (DocumentUserInfo, error) {
	return populateDocumentUser(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", documentUserSQL, uuid))
}

type ListDocumentUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDocument) ListDocumentUser(ctx context.Context, arg ListDocumentUserParams) ([]DocumentUserInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			documentUserSQL, arg.Limit, arg.Offset)
	} else {
		sql = documentUserSQL
	}
	return populateDocumentUsers(q, ctx, sql)
}

const updateDocumentUser = `-- name: UpdateDocumentUser :one
UPDATE Document_User SET 
	document_id = $2,
	user_id = $3,
	access_code = $4,
	other_info = $5
WHERE uuid = $1
RETURNING uuid, document_id, user_id, access_code, other_info
`

func (q *QueriesDocument) UpdateDocumentUser(ctx context.Context, arg DocumentUserRequest) (model.DocumentUser, error) {
	row := q.db.QueryRowContext(ctx, updateDocumentUser,
		arg.Uuid,
		arg.DocumentId,
		arg.UserId,
		arg.AccessCode,
		arg.OtherInfo,
	)
	var i model.DocumentUser
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.UserId,
		&i.AccessCode,
		&i.OtherInfo,
	)
	return i, err
}
