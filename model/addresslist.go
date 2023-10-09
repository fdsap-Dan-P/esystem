package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AddressList struct {
	Uuid        uuid.UUID      `json:"uuid"`
	Iiid        int64          `json:"iiid"`
	Series      int16          `json:"series"`
	Detail      sql.NullString `json:"detail"`
	Url         sql.NullString `json:"url"`
	TypeId      int64          `json:"typeId"`
	GeographyId sql.NullInt64  `json:"geographyId"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}
