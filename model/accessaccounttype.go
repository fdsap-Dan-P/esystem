package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccessAccountType struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	TypeId    int64          `json:"typeId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
