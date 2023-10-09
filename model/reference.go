package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Reference struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Code          int64          `json:"code"`
	ShortName     sql.NullString `json:"shortName"`
	Title         string         `json:"title"`
	ParentId      sql.NullInt64  `json:"parentId"`
	TypeId        int64          `json:"typeId"`
	Remark        sql.NullString `json:"remark"`
	VecSimpleName sql.NullString `json:"vecSimpleName"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
