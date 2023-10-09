package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountInventory struct {
	Id           int64           `json:"id"`
	Uuid         uuid.UUID       `json:"uuid"`
	AccountId    int64           `json:"accountId"`
	RepositoryId sql.NullInt64   `json:"repositoryId"`
	BarCode      sql.NullString  `json:"barCode"`
	Code         string          `json:"code"`
	Quantity     decimal.Decimal `json:"quantity"`
	UnitPrice    decimal.Decimal `json:"unitPrice"`
	BookValue    decimal.Decimal `json:"bookValue"`
	Discount     decimal.Decimal `json:"discount"`
	TaxRate      decimal.Decimal `json:"taxRate"`
	Remarks      string          `json:"remarks"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
}
