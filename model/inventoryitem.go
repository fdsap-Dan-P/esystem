package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InventoryItem struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	ItemCode        string         `json:"itemCode"`
	BarCode         sql.NullString `json:"barCode"`
	ItemName        string         `json:"itemName"`
	UniqueVariation string         `json:"uniqueVariation"`
	ParentId        sql.NullInt64  `json:"parentId"`
	GenericNameId   sql.NullInt64  `json:"genericNameId"`
	BrandNameId     sql.NullInt64  `json:"brandNameId"`
	MeasureId       int64          `json:"measureId"`
	ImageId         sql.NullInt64  `json:"imageId"`
	Remarks         string         `json:"remarks"`
	VecSimpleName   sql.NullString `json:"vecSimpleName"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type InventorySpecsString struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventoryItemId"`
	SpecsCode       string    `json:"specsCode"`
	SpecsId         int64     `json:"specsId"`
	Value           string    `json:"value"`
}

type InventorySpecsNumber struct {
	Uuid            uuid.UUID       `json:"uuid"`
	InventoryItemId int64           `json:"inventoryItemId"`
	SpecsCode       string          `json:"specsCode"`
	SpecsId         int64           `json:"specsId"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
}

type InventorySpecsDate struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventoryItemId"`
	SpecsCode       string    `json:"specsCode"`
	SpecsId         int64     `json:"specsId"`
	Value           time.Time `json:"value"`
	Value2          time.Time `json:"value2"`
}

type InventorySpecsRef struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventoryItemId"`
	SpecsCode       string    `json:"specsCode"`
	SpecsId         int64     `json:"specsId"`
	RefId           int64     `json:"refId"`
}
