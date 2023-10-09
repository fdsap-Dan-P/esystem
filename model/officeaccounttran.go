package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OfficeAccountTran struct {
	Uuid            uuid.UUID       `json:"uuid"`
	TrnHeadId       int64           `json:"trnHeadId"`
	Series          int16           `json:"series"`
	OfficeAccountId int64           `json:"officeAccountId"`
	TrnAmt          decimal.Decimal `json:"trnAmt"`
	OtherInfo       sql.NullString  `json:"otherInfo"`
}
