package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Employment struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	Company       string         `json:"company"`
	Title         string         `json:"title"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Remarks       sql.NullString `json:"remarks"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
