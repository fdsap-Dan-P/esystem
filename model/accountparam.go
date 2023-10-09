package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountParam struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	DateImplemented time.Time      `json:"dateImplemented"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type AccountParamString struct {
	Uuid     uuid.UUID `json:"uuid"`
	ParamId  int64     `json:"paramId"`
	ItemCode string    `json:"itemCode"`
	ItemId   int64     `json:"itemId"`
	Value    string    `json:"value"`
}

type AccountParamNumber struct {
	Uuid      uuid.UUID       `json:"uuid"`
	ParamId   int64           `json:"paramId"`
	ItemCode  string          `json:"itemCode"`
	ItemId    int64           `json:"itemId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

type AccountParamDate struct {
	Uuid     uuid.UUID `json:"uuid"`
	ParamId  int64     `json:"paramId"`
	ItemCode string    `json:"itemCode"`
	ItemId   int64     `json:"itemId"`
	Value    time.Time `json:"value"`
	Value2   time.Time `json:"value2"`
}

type AccountParamRef struct {
	Uuid     uuid.UUID `json:"uuid"`
	ParamId  int64     `json:"paramId"`
	ItemCode string    `json:"itemCode"`
	ItemId   int64     `json:"itemId"`
	RefId    int64     `json:"refId"`
}
