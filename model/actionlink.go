package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ActionLink struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	EventName    string         `json:"eventName"`
	TypeId       int64          `json:"typeId"`
	EndPointCall sql.NullString `json:"endPointCall"`
	ServerId     sql.NullInt64  `json:"serverId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
