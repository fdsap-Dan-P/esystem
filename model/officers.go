package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Officer struct {
	Uuid        uuid.UUID      `json:"uuid"`
	OfficeId    int64          `json:"officeId"`
	Position    sql.NullString `json:"position"`
	PeriodStart time.Time      `json:"periodStart"`
	PeriodEnd   sql.NullTime   `json:"periodEnd"`
	StatusId    int64          `json:"statusId"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}
