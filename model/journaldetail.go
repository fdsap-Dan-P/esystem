package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type JournalDetail struct {
	Uuid          uuid.UUID       `json:"uuid"`
	TrnHeadId     int64           `json:"trnHeadId"`
	Series        int16           `json:"series"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	AccountTypeId sql.NullInt64   `json:"accountTypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	TrnAmt        decimal.Decimal `json:"trnAmt"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}
