package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createDocumentAccess = `-- name: CreateDocumentAccess: one
INSERT INTO Document_Access(
   document_id, role_id, access_code, other_info )
VALUES ($1, $2, $3, $4)
RETURNING UUID, document_id, role_id, access_code, other_info`

type DocumentAccessRequest struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	RoleId     int64          `json:"roleId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

func (q *QueriesDocument) CreateDocumentAccess(ctx context.Context, arg DocumentAccessRequest) (model.DocumentAccess, error) {
	row := q.db.QueryRowContext(ctx, createDocumentAccess,
		arg.DocumentId,
		arg.RoleId,
		arg.AccessCode,
		arg.OtherInfo,
	)
	var i model.DocumentAccess
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.RoleId,
		&i.AccessCode,
		&i.OtherInfo,
	)
	return i, err
}

const deleteDocumentAccess = `-- name: DeleteDocumentAccess :exec
DELETE FROM Document_Access
WHERE uuid = $1
`

func (q *QueriesDocument) DeleteDocumentAccess(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDocumentAccess, uuid)
	return err
}

type DocumentAccessInfo struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	RoleId     int64          `json:"roleId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
	ModCtr     int64          `json:"modCtr"`
	Created    sql.NullTime   `json:"created"`
	Updated    sql.NullTime   `json:"updated"`
}

const documentAccessSQL = `-- name: DocumentAccessSQL :one
SELECT
mr.UUID, document_id, role_id, access_code, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Document_Access d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateDocumentAccess(q *QueriesDocument, ctx context.Context, sql string) (DocumentAccessInfo, error) {
	var i DocumentAccessInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.RoleId,
		&i.AccessCode,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateDocumentAccesss(q *QueriesDocument, ctx context.Context, sql string) ([]DocumentAccessInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DocumentAccessInfo{}
	for rows.Next() {
		var i DocumentAccessInfo
		err := rows.Scan(
			&i.Uuid,
			&i.DocumentId,
			&i.RoleId,
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

func (q *QueriesDocument) GetDocumentAccess(ctx context.Context, id int64) (DocumentAccessInfo, error) {
	return populateDocumentAccess(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", documentAccessSQL, id))
}

func (q *QueriesDocument) GetDocumentAccessbyUuid(ctx context.Context, uuid uuid.UUID) (DocumentAccessInfo, error) {
	return populateDocumentAccess(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", documentAccessSQL, uuid))
}

type ListDocumentAccessParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDocument) ListDocumentAccess(ctx context.Context, arg ListDocumentAccessParams) ([]DocumentAccessInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			documentAccessSQL, arg.Limit, arg.Offset)
	} else {
		sql = documentAccessSQL
	}
	return populateDocumentAccesss(q, ctx, sql)
}

const updateDocumentAccess = `-- name: UpdateDocumentAccess :one
UPDATE Document_Access SET 
document_id = $2,
role_id = $3,
access_code = $4,
other_info = $5
WHERE uuid = $1
RETURNING uuid, document_id, role_id, access_code, other_info
`

func (q *QueriesDocument) UpdateDocumentAccess(ctx context.Context, arg DocumentAccessRequest) (model.DocumentAccess, error) {
	row := q.db.QueryRowContext(ctx, updateDocumentAccess,
		arg.Uuid,
		arg.DocumentId,
		arg.RoleId,
		arg.AccessCode,
		arg.OtherInfo,
	)
	var i model.DocumentAccess
	err := row.Scan(
		&i.Uuid,
		&i.DocumentId,
		&i.RoleId,
		&i.AccessCode,
		&i.OtherInfo,
	)
	return i, err
}
