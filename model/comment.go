package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Comment struct {
	Uuid       uuid.UUID      `json:"uuid"`
	RecordUuid uuid.UUID      `json:"recordUuid"`
	UserId     int64          `json:"userId"`
	Comment    string         `json:"comment"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}
