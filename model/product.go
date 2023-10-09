package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Code          int64          `json:"code"`
	ProductName   string         `json:"productName"`
	Description   sql.NullString `json:"description"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
