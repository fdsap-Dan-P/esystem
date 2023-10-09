package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Relation struct {
	Uuid         uuid.UUID      `json:"uuid"`
	Iiid         int64          `json:"iiid"`
	Series       int16          `json:"series"`
	RelationIiid int64          `json:"relationIiid"`
	TypeId       int64          `json:"typeId"`
	RelationDate sql.NullTime   `json:"relationDate"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
