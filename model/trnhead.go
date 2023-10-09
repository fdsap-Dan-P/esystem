package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TrnHead struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	TrnSerial       string         `json:"trnSerial"`
	TicketId        int64          `json:"ticketId"`
	TrnDate         time.Time      `json:"trnDate"`
	TypeId          int64          `json:"typeId"`
	Particular      sql.NullString `json:"particular"`
	OfficeId        int64          `json:"officeId"`
	UserId          int64          `json:"userId"`
	TransactingIiid sql.NullInt64  `json:"transactingIiid"`
	Orno            sql.NullString `json:"orno"`
	Isfinal         sql.NullBool   `json:"isfinal"`
	Ismanual        sql.NullBool   `json:"ismanual"`
	AlternateTrn    sql.NullString `json:"alternateTrn"`
	Reference       string         `json:"reference"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type TrnHeadSpecsString struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

type TrnHeadSpecsNumber struct {
	Uuid      uuid.UUID       `json:"uuid"`
	TrnHeadId int64           `json:"trnHeadId"`
	SpecsCode string          `json:"specsCode"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

type TrnHeadSpecsDate struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

type TrnHeadSpecsRef struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}
