package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CoaParent struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	Acc       string         `json:"acc"`
	CoaSeq    sql.NullInt64  `json:"coaSeq"`
	Title     string         `json:"title"`
	ParentId  sql.NullInt64  `json:"parentId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
