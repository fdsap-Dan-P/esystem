package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TicketTypeStatus struct {
	Id                    int64          `json:"id"`
	Uuid                  uuid.UUID      `json:"uuid"`
	ProductTicketTypeId   int64          `json:"productTicketTypeId"`
	StatusId              int64          `json:"statusId"`
	TicketTypeActionArray []int64        `json:"ticketTypeActionArray"`
	OtherInfo             sql.NullString `json:"otherInfo"`
}
type TicketTypeAction struct {
	Id                  int64          `json:"id"`
	Uuid                uuid.UUID      `json:"uuid"`
	ProductTicketTypeId int64          `json:"productTicketTypeId"`
	ActionId            int64          `json:"actionId"`
	Actiondesc          string         `json:"actiondesc"`
	ActionLinkId        sql.NullInt64  `json:"actionLinkId"`
	OtherInfo           sql.NullString `json:"otherInfo"`
}

type TicketActionConditionDate struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"ticketTypeStatusId"`
	ItemCode           string    `json:"itemCode"`
	ItemId             int64     `json:"itemId"`
	ConditionId        int64     `json:"conditionId"`
	Value              time.Time `json:"value"`
	Value2             time.Time `json:"value2"`
}

type TicketActionConditionNumber struct {
	Uuid               uuid.UUID       `json:"uuid"`
	TicketTypeStatusId int64           `json:"ticketTypeStatusId"`
	ItemCode           string          `json:"itemCode"`
	ItemId             int64           `json:"itemId"`
	ConditionId        int64           `json:"conditionId"`
	Value              decimal.Decimal `json:"value"`
	Value2             decimal.Decimal `json:"value2"`
	MeasureId          sql.NullInt64   `json:"measureId"`
}

type TicketActionConditionRef struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"ticketTypeStatusId"`
	ItemCode           string    `json:"itemCode"`
	ItemId             int64     `json:"itemId"`
	ConditionId        int64     `json:"conditionId"`
	RefId              int64     `json:"refId"`
}

type TicketActionConditionString struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"ticketTypeStatusId"`
	ItemCode           string    `json:"itemCode"`
	ItemId             int64     `json:"itemId"`
	ConditionId        int64     `json:"conditionId"`
	Value              string    `json:"value"`
}
