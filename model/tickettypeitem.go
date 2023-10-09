package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type TicketTypeItem struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	TicketTypeId int64          `json:"ticketTypeId"`
	ItemId       int64          `json:"itemId"`
	Item         string         `json:"item"`
	StatusId     int64          `json:"statusId"`
	Remarks      string         `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
