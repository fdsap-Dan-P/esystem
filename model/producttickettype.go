package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ProductTicketType struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	ProductId       int64          `json:"productId"`
	TicketTypeId    int64          `json:"ticketTypeId"`
	StatusId        int64          `json:"statusId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}
