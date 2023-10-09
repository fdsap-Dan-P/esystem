package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserProduct struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ProductId int64          `json:"productId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
