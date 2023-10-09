package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type OfficeAccountType struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	OfficeAccountType string         `json:"officeAccountType"`
	CoaId             int64          `json:"coaId"`
	OtherInfo         sql.NullString `json:"otherInfo"`
}
