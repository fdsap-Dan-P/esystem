package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserObject struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ObjectId  int64          `json:"objectId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
