package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Fxrate struct {
	Uuid         uuid.UUID       `json:"uuid"`
	BuyRate      decimal.Decimal `json:"buyRate"`
	CutofDate    time.Time       `json:"cutofDate"`
	SellRate     decimal.Decimal `json:"sellRate"`
	BaseCurrency string          `json:"baseCurrency"`
	Currency     string          `json:"currency"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
}
