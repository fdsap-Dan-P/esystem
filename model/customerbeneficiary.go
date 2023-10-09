package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CustomerBeneficiary struct {
	Uuid           uuid.UUID      `json:"uuid"`
	CustomerId     int64          `json:"customerId"`
	Series         int16          `json:"series"`
	Iiid           int64          `json:"iiid"`
	TypeId         int64          `json:"typeId"`
	RelationTypeId int64          `json:"relationTypeId"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}
