package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccountClass struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	ProductId int64          `json:"productId"`
	GroupId   int64          `json:"groupId"`
	ClassId   int64          `json:"classId"`
	CurId     int64          `json:"curId"`
	NoncurId  sql.NullInt64  `json:"noncurId"`
	BsAccId   sql.NullInt64  `json:"bsAccId"`
	IsAccId   sql.NullInt64  `json:"isAccId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}
