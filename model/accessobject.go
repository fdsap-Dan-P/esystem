package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccessObject struct {
	Uuid      uuid.UUID       `json:"uuid"`
	RoleId    int64           `json:"roleId"`
	ObjectId  int64           `json:"objectId"`
	Allow     sql.NullBool    `json:"allow"`
	MaxValue  decimal.Decimal `json:"maxValue"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}
