package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccessConfigDate struct {
	Uuid       uuid.UUID `json:"uuid"`
	RoleId     int64     `json:"roleId"`
	ConfigCode string    `json:"configCode"`
	ConfigId   int64     `json:"configId"`
	Value      time.Time `json:"value"`
	Value2     time.Time `json:"value2"`
}

type AccessConfigNumber struct {
	Uuid       uuid.UUID       `json:"uuid"`
	RoleId     int64           `json:"roleId"`
	ConfigCode string          `json:"configCode"`
	ConfigId   int64           `json:"configId"`
	Value      decimal.Decimal `json:"value"`
	Value2     decimal.Decimal `json:"value2"`
	MeasureId  sql.NullInt64   `json:"measureId"`
}

type AccessConfigRef struct {
	Uuid       uuid.UUID `json:"uuid"`
	RoleId     int64     `json:"roleId"`
	ConfigCode string    `json:"configCode"`
	ConfigId   int64     `json:"configId"`
	RefId      int64     `json:"refId"`
}

type AccessConfigString struct {
	Uuid       uuid.UUID `json:"uuid"`
	RoleId     int64     `json:"roleId"`
	ConfigCode string    `json:"configCode"`
	ConfigId   int64     `json:"configId"`
	Value      string    `json:"value"`
}
