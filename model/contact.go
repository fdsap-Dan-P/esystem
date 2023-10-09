package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Contact struct {
	Uuid      uuid.UUID      `json:"uuid"`
	Iiid      int64          `json:"iiid"`
	Series    int16          `json:"series"`
	Contact   string         `json:"contact"`
	TypeId    int64          `json:"typeId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
