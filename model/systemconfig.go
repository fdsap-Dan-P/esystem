package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type SystemConfig struct {
	Uuid         uuid.UUID      `json:"uuid"`
	OfficeId     int64          `json:"officeId"`
	GlDate       time.Time      `json:"glDate"`
	LastAccruals time.Time      `json:"lastAccruals"`
	LastMonthEnd time.Time      `json:"lastMonthEnd"`
	NextMonthEnd time.Time      `json:"nextMonthEnd"`
	SystemDate   time.Time      `json:"systemDate"`
	RunState     int16          `json:"runState"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
