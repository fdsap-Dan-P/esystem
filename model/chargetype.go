package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ChargeType struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ChargeType   string         `json:"chargeType"`
	UnrealizedId int64          `json:"unrealizedId"`
	RealizedId   int64          `json:"realizedId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
