package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InventoryDetail struct {
	Id                 int64           `json:"id"`
	Uuid               uuid.UUID       `json:"uuid"`
	AccountInventoryId int64           `json:"accountInventoryId"`
	InventoryItemId    int64           `json:"inventoryItemId"`
	RepositoryId       sql.NullInt64   `json:"repositoryId"`
	SupplierId         sql.NullInt64   `json:"supplierId"`
	UnitPrice          decimal.Decimal `json:"unitPrice"`
	BookValue          decimal.Decimal `json:"bookValue"`
	Unit               decimal.Decimal `json:"unit"`
	MeasureId          int64           `json:"measureId"`
	BatchNumber        sql.NullString  `json:"batchNumber"`
	DateManufactured   sql.NullTime    `json:"dateManufactured"`
	DateExpired        sql.NullTime    `json:"dateExpired"`
	Remarks            string          `json:"remarks"`
	OtherInfo          sql.NullString  `json:"otherInfo"`
}
