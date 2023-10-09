package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Educational struct {
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	LevelId       sql.NullInt64  `json:"levelId"`
	CourseTypeId  sql.NullInt64  `json:"courseTypeId"`
	CourseId      sql.NullInt64  `json:"courseId"`
	Course        string         `json:"course"`
	School        string         `json:"school"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Completed     sql.NullBool   `json:"completed"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
