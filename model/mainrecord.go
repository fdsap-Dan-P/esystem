package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type MainRecord struct {
	Uuid    uuid.UUID    `json:"uuid"`
	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}
