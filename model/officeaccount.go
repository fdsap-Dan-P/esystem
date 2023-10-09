package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OfficeAccount struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	OfficeId         int64           `json:"officeId"`
	TypeId           int64           `json:"typeId"`
	Currency         string          `json:"currency"`
	PartitionId      sql.NullInt64   `json:"partitionId"`
	Balance          decimal.Decimal `json:"balance"`
	PendingTrnAmt    decimal.Decimal `json:"pendingTrnAmt"`
	Budget           decimal.Decimal `json:"budget"`
	LastActivityDate sql.NullTime    `json:"lastActivityDate"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}
