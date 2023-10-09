package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	TicketDate      time.Time      `json:"ticketDate"`
	PostedbyId      int64          `json:"postedbyId"`
	StatusId        int64          `json:"statusId"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}
