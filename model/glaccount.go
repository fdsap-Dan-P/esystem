package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type GlAccount struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	Balance       decimal.Decimal `json:"balance"`
	PendingTrnAmt decimal.Decimal `json:"pendingTrnAmt"`
	AccountTypeId sql.NullInt64   `json:"accountTypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	Remark        sql.NullString  `json:"remark"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}
