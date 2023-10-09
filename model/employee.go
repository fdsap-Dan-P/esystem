package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Employee struct {
	Id          int64           `json:"id"`
	Uuid        uuid.UUID       `json:"uuid"`
	Iiid        int64           `json:"iiid"`
	CentralId   int64           `json:"centralId"`
	EmployeeNo  string          `json:"employeeNo"`
	BasicPay    decimal.Decimal `json:"basicPay"`
	DateHired   sql.NullTime    `json:"dateHired"`
	DateRegular sql.NullTime    `json:"dateRegular"`
	JobGrade    int16           `json:"jobGrade"`
	JobStep     int16           `json:"jobStep"`
	LevelId     sql.NullInt64   `json:"levelId"`
	OfficeId    int64           `json:"officeId"`
	PositionId  int64           `json:"positionId"`
	StatusCode  int64           `json:"statusCode"`
	SuperiorId  sql.NullInt64   `json:"superiorId"`
	TypeId      sql.NullInt64   `json:"typeId"`
	OtherInfo   sql.NullString  `json:"otherInfo"`
}
