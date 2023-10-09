package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserAccountType struct {
	Uuid          uuid.UUID      `json:"uuid"`
	UserId        int64          `json:"userId"`
	AccountTypeId int64          `json:"accountTypeId"`
	Allow         bool           `json:"allow"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
