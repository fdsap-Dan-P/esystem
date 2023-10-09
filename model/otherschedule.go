package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OtherSchedule struct {
	Uuid       uuid.UUID       `json:"uuid"`
	AccountId  int64           `json:"accountId"`
	ChargeId   int64           `json:"chargeId"`
	Series     int16           `json:"series"`
	DueDate    time.Time       `json:"dueDate"`
	DueAmt     decimal.Decimal `json:"dueAmt"`
	Realizable decimal.Decimal `json:"realizable"`
	EndBal     decimal.Decimal `json:"endBal"`
	OtherInfo  sql.NullString  `json:"otherInfo"`
}
