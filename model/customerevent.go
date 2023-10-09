package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CustomerEvent struct {
	Uuid       uuid.UUID      `json:"uuid"`
	TrnHeadId  int64          `json:"trnHeadId"`
	CustomerId int64          `json:"customerId"`
	TypeId     int64          `json:"typeId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}
