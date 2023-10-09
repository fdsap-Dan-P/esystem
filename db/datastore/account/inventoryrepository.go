package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"simplebank/model"
)

const createInventoryRepository = `-- name: CreateInventoryRepository: one
INSERT INTO Inventory_Repository (
    Central_Office_Id, Repository_Code, Repository, Office_Id, 
	Custodian_Id, Geography_Id, Location_Description, Remarks, Other_Info
 ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
 RETURNING 
   Id, UUID, Central_Office_Id, Repository_Code, Repository, Office_Id, 
   Custodian_Id, Geography_Id, Location_Description, Remarks, Other_Info
`

type InventoryRepositoryRequest struct {
	Id                  int64          `json:"Id"`
	Uuid                uuid.UUID      `json:"UUID"`
	CentralOfficeId     int64          `json:"CentralOfficeId"`
	RepositoryCode      string         `json:"RepositoryCode"`
	Repository          string         `json:"Repository"`
	OfficeId            int64          `json:"OfficeId"`
	CustodianId         sql.NullInt64  `json:"CustodianId"`
	GeographyId         sql.NullInt64  `json:"GeographyId"`
	LocationDescription sql.NullString `json:"LocationDescription"`
	Remarks             sql.NullString `json:"Remarks"`
	OtherInfo           sql.NullString `json:"OtherInfo"`
}

func (q *QueriesAccount) CreateInventoryRepository(ctx context.Context, arg InventoryRepositoryRequest) (model.InventoryRepository, error) {
	row := q.db.QueryRowContext(ctx, createInventoryRepository,
		arg.CentralOfficeId,
		arg.RepositoryCode,
		arg.Repository,
		arg.OfficeId,
		arg.CustodianId,
		arg.GeographyId,
		arg.LocationDescription,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.InventoryRepository
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.RepositoryCode,
		&i.Repository,
		&i.OfficeId,
		&i.CustodianId,
		&i.GeographyId,
		&i.LocationDescription,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteInventoryRepository = `-- name: DeleteInventoryRepository :exec
DELETE FROM Inventory_Repository
WHERE id = $1
`

func (q *QueriesAccount) DeleteInventoryRepository(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteInventoryRepository, id)
	return err
}

type InventoryRepositoryInfo struct {
	Id                  int64          `json:"Id"`
	Uuid                uuid.UUID      `json:"UUID"`
	CentralOfficeId     int64          `json:"CentralOfficeId"`
	RepositoryCode      string         `json:"RepositoryCode"`
	Repository          string         `json:"Repository"`
	OfficeId            int64          `json:"OfficeId"`
	CustodianId         sql.NullInt64  `json:"CustodianId"`
	GeographyId         sql.NullInt64  `json:"GeographyId"`
	LocationDescription sql.NullString `json:"LocationDescription"`
	Remarks             sql.NullString `json:"Remarks"`
	OtherInfo           sql.NullString `json:"OtherInfo"`
	ModCtr              int64          `json:"modCtr"`
	Created             sql.NullTime   `json:"created"`
	Updated             sql.NullTime   `json:"updated"`

	// Child                []model.InventoryRepository        `json:"child"`
	// InventorySpecsNumber []model.InventorySpecsNumber `json:"inventorySpecsNumber"`
	// InventorySpecsDate   []model.InventorySpecsDate   `json:"inventorySpecsDate"`
	// InventorySpecsString []model.InventorySpecsString `json:"inventorySpecsString"`
}

func populateInventoryRepository(q *QueriesAccount, ctx context.Context, sql string) (InventoryRepositoryInfo, error) {
	var i InventoryRepositoryInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.RepositoryCode,
		&i.Repository,
		&i.OfficeId,
		&i.CustodianId,
		&i.GeographyId,
		&i.LocationDescription,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateInventoryRepositorys(q *QueriesAccount, ctx context.Context, sql string) ([]InventoryRepositoryInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []InventoryRepositoryInfo{}
	for rows.Next() {
		var i InventoryRepositoryInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CentralOfficeId,
			&i.RepositoryCode,
			&i.Repository,
			&i.OfficeId,
			&i.CustodianId,
			&i.GeographyId,
			&i.LocationDescription,
			&i.Remarks,
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

const inventoryRepositorySQL = `-- name: inventoryRepositorySQL
SELECT 
  d.Id, mr.Uuid, d.Central_Office_Id, d.Repository_Code, d.Repository, 
  d.Office_Id, d.Custodian_Id, d.Geography_Id, d.Location_Description, 
  d.Remarks, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Repository d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
`

func (q *QueriesAccount) GetInventoryRepository(ctx context.Context, id int64) (InventoryRepositoryInfo, error) {
	return populateInventoryRepository(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", inventoryRepositorySQL, id))
}

func (q *QueriesAccount) GetInventoryRepositorybyUuid(ctx context.Context, uuid uuid.UUID) (InventoryRepositoryInfo, error) {
	return populateInventoryRepository(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", inventoryRepositorySQL, uuid))
}

func (q *QueriesAccount) GetInventoryRepositorybyName(ctx context.Context, name string) (InventoryRepositoryInfo, error) {
	return populateInventoryRepository(q, ctx, fmt.Sprintf("%s WHERE lower(d.Item_Name) = %v", inventoryRepositorySQL, strings.ToLower(name)))
}

type ListInventoryRepositoryParams struct {
	OfficeId int64 `json:"officeId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventoryRepository(ctx context.Context, arg ListInventoryRepositoryParams) ([]InventoryRepositoryInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Office_Id = %v LIMIT %d OFFSET %d",
			inventoryRepositorySQL, arg.OfficeId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Office_Id = %v ", inventoryRepositorySQL, arg.OfficeId)
	}
	return populateInventoryRepositorys(q, ctx, sql)
}

const updateInventoryRepository = `-- name: UpdateInventoryRepository :one
UPDATE Inventory_Repository SET 
  Central_Office_Id = $2,
  Repository_Code = $3,
  Repository = $4,
  Office_Id = $5,
  Custodian_Id = $6,
  Geography_Id = $7,
  Location_Description = $8,
  Remarks = $9,
  Other_Info = $10
WHERE id = $1
RETURNING 
  Id, Uuid, Central_Office_Id, Repository_Code, Repository, Office_Id, 
  Custodian_Id, Geography_Id, Location_Description, Remarks, Other_Info
`

func (q *QueriesAccount) UpdateInventoryRepository(ctx context.Context, arg InventoryRepositoryRequest) (model.InventoryRepository, error) {
	row := q.db.QueryRowContext(ctx, updateInventoryRepository,
		arg.Id,
		arg.CentralOfficeId,
		arg.RepositoryCode,
		arg.Repository,
		arg.OfficeId,
		arg.CustodianId,
		arg.GeographyId,
		arg.LocationDescription,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.InventoryRepository
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.RepositoryCode,
		&i.Repository,
		&i.OfficeId,
		&i.CustodianId,
		&i.GeographyId,
		&i.LocationDescription,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
