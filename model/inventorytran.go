package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InventoryTran struct {
	Uuid              uuid.UUID       `json:"uuid"`
	TrnHeadId         int64           `json:"trnHeadId"`
	Series            int16           `json:"series"`
	InventoryDetailId int64           `json:"inventoryDetailId"`
	RepositoryId      int64           `json:"repositoryId"`
	Quantity          decimal.Decimal `json:"quantity"`
	UnitPrice         decimal.Decimal `json:"unitPrice"`
	Discount          decimal.Decimal `json:"discount"`
	TaxAmt            decimal.Decimal `json:"taxAmt"`
	NetTrnAmt         decimal.Decimal `json:"netTrnAmt"`
	OtherInfo         sql.NullString  `json:"otherInfo"`
}
