package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Questionaire struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	Code            string         `json:"code"`
	Version         int64          `json:"version"`
	Title           string         `json:"title"`
	TypeId          int64          `json:"typeId"`
	SubjectId       sql.NullInt64  `json:"subjectId"`
	DateRevised     time.Time      `json:"dateRevised"`
	OfficeId        sql.NullInt64  `json:"officeId"`
	AuthorId        sql.NullInt64  `json:"authorId"`
	StatusId        int64          `json:"statusId"`
	PointEquivalent sql.NullString `json:"pointEquivalent"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}
