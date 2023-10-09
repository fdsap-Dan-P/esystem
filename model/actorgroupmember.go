package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ActorGroupMember struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ActorGroupId int64          `json:"actorGroupId"`
	ActorGroup   string         `json:"actorGroup"`
	UserId       int64          `json:"userId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
