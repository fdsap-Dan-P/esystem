package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type TrnHeadRelation struct {
	Uuid      uuid.UUID      `json:"uuid"`
	TrnHeadId int64          `json:"trnHeadId"`
	RelatedId int64          `json:"relatedId"`
	TypeId    int64          `json:"typeId"`
	Remarks   sql.NullString `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
