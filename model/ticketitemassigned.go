package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TicketItemAssigned struct {
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	UserId       int64          `json:"userId"`
	AssignedById int64          `json:"assignedById"`
	AssignedDate time.Time      `json:"assignedDate"`
	Remarks      string         `json:"remarks"`
	StatusId     int64          `json:"statusId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
