package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccessProduct struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	ProductId int64          `json:"productId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
