package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ReferenceType struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	Code        sql.NullString `json:"code"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}
