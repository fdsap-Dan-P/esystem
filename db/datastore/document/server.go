package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"simplebank/model"
)

const createServer = `-- name: CreateServer: one
INSERT INTO Server (
  Code, Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
ON CONFLICT(Code) DO UPDATE SET
  Connectivity = Connectivity, 
  Net_Address = Net_Address, 
  Certificate = Certificate, 
  HomePath = HomePath, 
  Description = Description, 
  Other_Info = Other_Info
RETURNING 
  Id, UUId, Code, Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info
`

type ServerRequest struct {
	Id           int64              `json:"id"`
	Uuid         uuid.UUID          `json:"uuid"`
	Code         string             `json:"code"`
	Connectivity model.Connectivity `json:"connectivity"`
	NetAddress   string             `json:"netAddress"`
	Certificate  sql.NullString     `json:"certificate"`
	HomePath     string             `json:"homePath"`
	Description  sql.NullString     `json:"description"`
	OtherInfo    sql.NullString     `json:"otherInfo"`
}

func (q *QueriesDocument) AddServer(
	ctx context.Context, arg ServerRequest, sql string) (model.Server, error) {
	row := q.db.QueryRowContext(ctx, sql,
		arg.Code,
		arg.Connectivity,
		arg.NetAddress,
		arg.Certificate,
		arg.HomePath,
		arg.Description,
		arg.OtherInfo,
	)
	var i model.Server
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Connectivity,
		&i.NetAddress,
		&i.Certificate,
		&i.HomePath,
		&i.Description,
		&i.OtherInfo,
	)
	return i, err
}

func (q *QueriesDocument) CreateServer(ctx context.Context, arg ServerRequest) (model.Server, error) {
	return q.AddServer(ctx, arg, createServer)
}

const deleteServer = `-- name: DeleteServer :exec
DELETE FROM Server
WHERE id = $1
`

func (q *QueriesDocument) DeleteServer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteServer, id)
	return err
}

type ServerInfo struct {
	Id           int64              `json:"id"`
	Uuid         uuid.UUID          `json:"uuid"`
	Code         string             `json:"code"`
	Connectivity model.Connectivity `json:"connectivity"`
	NetAddress   string             `json:"netAddress"`
	Certificate  sql.NullString     `json:"certificate"`
	HomePath     string             `json:"homePath"`
	Description  sql.NullString     `json:"description"`
	OtherInfo    sql.NullString     `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getServer = `-- name: GetServer :one
SELECT 
Id, mr.UUId, Code, 
Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Server d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesDocument) GetServer(ctx context.Context, id int64) (ServerInfo, error) {
	row := q.db.QueryRowContext(ctx, getServer, id)
	var i ServerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Connectivity,
		&i.NetAddress,
		&i.Certificate,
		&i.HomePath,
		&i.Description,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getServerbyUuId = `-- name: GetServerbyUuId :one
SELECT 
Id, mr.UUId, Code, 
Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Server d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesDocument) GetServerbyUuId(ctx context.Context, uuid uuid.UUID) (ServerInfo, error) {
	row := q.db.QueryRowContext(ctx, getServerbyUuId, uuid)
	var i ServerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Connectivity,
		&i.NetAddress,
		&i.Certificate,
		&i.HomePath,
		&i.Description,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getServerbyCode = `-- name: GetServerbyCode :one
SELECT 
Id, mr.UUId, Code, 
Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Server d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Code = $1 LIMIT 1
`

func (q *QueriesDocument) GetServerbyCode(ctx context.Context, code string) (ServerInfo, error) {
	row := q.db.QueryRowContext(ctx, getServerbyCode, code)
	var i ServerInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Connectivity,
		&i.NetAddress,
		&i.Certificate,
		&i.HomePath,
		&i.Description,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)

	// if err == sql.ErrNoRows {
	// 	log.Printf("GetServerbyCode 2 %+v", err)
	// 	return i, err
	// }
	log.Printf("GetServerbyCode 1 %+v", i)

	return i, err
}

const listServer = `-- name: ListServer:many
SELECT 
Id, mr.UUId, Code, 
Connectivity, Net_Address, Certificate, HomePath, Description, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Server d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListServerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDocument) ListServer(ctx context.Context, arg ListServerParams) ([]ServerInfo, error) {
	rows, err := q.db.QueryContext(ctx, listServer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServerInfo{}
	for rows.Next() {
		var i ServerInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.Connectivity,
			&i.NetAddress,
			&i.Certificate,
			&i.HomePath,
			&i.Description,
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

const updateServer = `-- name: UpdateServer :one
UPDATE Server SET 
Code = $2,
Connectivity = $3,
Net_Address = $4,
Certificate = $5,
HomePath = $6,
Description = $7,
Other_Info = $8
WHERE id = $1
RETURNING Id, UUId, Code, Connectivity, Net_Address, Certificate, HomePath, Description, 
Other_Info
`

func (q *QueriesDocument) UpdateServer(ctx context.Context, arg ServerRequest) (model.Server, error) {
	row := q.db.QueryRowContext(ctx, updateServer,
		arg.Id,
		arg.Code,
		arg.Connectivity,
		arg.NetAddress,
		arg.Certificate,
		arg.HomePath,
		arg.Description,
		arg.OtherInfo,
	)
	var i model.Server
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.Connectivity,
		&i.NetAddress,
		&i.Certificate,
		&i.HomePath,
		&i.Description,
		&i.OtherInfo,
	)
	return i, err
}
