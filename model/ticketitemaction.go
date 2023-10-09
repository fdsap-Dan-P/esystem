package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TicketItemAction struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	TrnHeadId    int64          `json:"trnHeadId"`
	UserId       int64          `json:"userId"`
	ActionId     int64          `json:"actionId"`
	ActionDate   time.Time      `json:"actionDate"`
	Remarks      sql.NullString `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
