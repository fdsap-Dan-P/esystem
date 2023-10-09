package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Ids struct {
	Uuid             uuid.UUID      `json:"uuid"`
	Iiid             int64          `json:"iiid"`
	Series           int16          `json:"series"`
	IdNumber         string         `json:"idNumber"`
	RegistrationDate sql.NullTime   `json:"registrationDate"`
	ValidityDate     sql.NullTime   `json:"validityDate"`
	TypeId           int64          `json:"typeId"`
	OtherInfo        sql.NullString `json:"otherInfo"`
}
