package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserOffice struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	OfficeId  int64          `json:"officeId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
