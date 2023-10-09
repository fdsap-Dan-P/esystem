package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeSource struct {
	Uuid      uuid.UUID       `json:"uuid"`
	Iiid      int64           `json:"iiid"`
	Series    int16           `json:"series"`
	Source    string          `json:"source"`
	TypeId    int64           `json:"typeId"`
	MinIncome decimal.Decimal `json:"minIncome"`
	MaxIncome decimal.Decimal `json:"maxIncome"`
	Remarks   sql.NullString  `json:"remarks"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}
