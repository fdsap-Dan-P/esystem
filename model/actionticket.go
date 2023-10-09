package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ActionTicket struct {
	Uuid             uuid.UUID      `json:"uuid"`
	TicketId         int64          `json:"ticketId"`
	TypeId           int64          `json:"typeId"`
	UserId           int64          `json:"userId"`
	ActionTicketDate sql.NullTime   `json:"actionTicketDate"`
	Reference        sql.NullString `json:"reference"`
	Remarks          sql.NullString `json:"remarks"`
	OtherInfo        sql.NullString `json:"otherInfo"`
}
