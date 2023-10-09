package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type EmployeeEvent struct {
	Uuid           uuid.UUID       `json:"uuid"`
	EmployeeId     int64           `json:"employeeId"`
	TicketId       int64           `json:"ticketId"`
	EventTypeId    int64           `json:"eventTypeId"`
	OfficeId       int64           `json:"officeId"`
	PositionId     int64           `json:"positionId"`
	BasicPay       decimal.Decimal `json:"basicPay"`
	StatusId       int64           `json:"statusId"`
	JobGrade       int16           `json:"jobGrade"`
	JobStep        int16           `json:"jobStep"`
	LevelId        sql.NullInt64   `json:"levelId"`
	EmployeeTypeId sql.NullInt64   `json:"employeeTypeId"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}
