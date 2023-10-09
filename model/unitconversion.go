package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UnitConversion struct {
	Id        int64           `json:"id"`
	Uuid      uuid.UUID       `json:"uuid"`
	TypeId    int64           `json:"typeId"`
	FromId    int64           `json:"fromId"`
	ToId      int64           `json:"toId"`
	Value     decimal.Decimal `json:"value"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}
