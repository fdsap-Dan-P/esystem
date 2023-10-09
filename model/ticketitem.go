package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TicketItem struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	TicketId  int64          `json:"ticketId"`
	ItemCode  int64          `json:"itemCode"`
	ItemId    int64          `json:"itemId"`
	Item      string         `json:"item"`
	StatusId  int64          `json:"statusId"`
	Remarks   string         `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

type TicketItemSpecsDate struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"ticketItemId"`
	SpecsCode    string    `json:"specsCode"`
	SpecsId      int64     `json:"specsId"`
	Value        time.Time `json:"value"`
	Value2       time.Time `json:"value2"`
}

type TicketItemSpecsNumber struct {
	Uuid         uuid.UUID       `json:"uuid"`
	TicketItemId int64           `json:"ticketItemId"`
	SpecsCode    string          `json:"specsCode"`
	SpecsId      int64           `json:"specsId"`
	Value        decimal.Decimal `json:"value"`
	Value2       decimal.Decimal `json:"value2"`
	MeasureId    sql.NullInt64   `json:"measureId"`
}

type TicketItemSpecsRef struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"ticketItemId"`
	SpecsCode    string    `json:"specsCode"`
	SpecsId      int64     `json:"specsId"`
	RefId        int64     `json:"refId"`
}

type TicketItemSpecsString struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"ticketItemId"`
	SpecsCode    string    `json:"specsCode"`
	SpecsId      int64     `json:"specsId"`
	Value        string    `json:"value"`
}
