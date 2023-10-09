package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ActorGroupAction struct {
	Uuid               uuid.UUID      `json:"uuid"`
	ActorGroupId       int64          `json:"actorGroupId"`
	ActorGroup         string         `json:"actorGroup"`
	TicketTypeActionId int64          `json:"ticketTypeActionId"`
	TicketType         string         `json:"ticketType"`
	ActionDesc         string         `json:"actionDesc"`
	OtherInfo          sql.NullString `json:"otherInfo"`
}
