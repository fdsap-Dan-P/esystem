package model

import (
	"github.com/google/uuid"
)

type TicketAccountTypeGroup struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeId       int64     `json:"ticketTypeId"`
	AccountTypeGroupId int64     `json:"accountTypeGroupId"`
	StatusId           int64     `json:"statusId"`
}
