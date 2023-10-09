package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Modified struct {
	ModCtr    int64          `json:"modCtr"`
	Uuid      uuid.UUID      `json:"uuid"`
	Updated   sql.NullTime   `json:"updated"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
