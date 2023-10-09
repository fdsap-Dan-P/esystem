package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccessRole struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	AccessName  string         `json:"accessName"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}
