package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type ChartofAccount struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Acc           string         `json:"acc"`
	Active        bool           `json:"active"`
	ContraAccount bool           `json:"contraAccount"`
	NormalBalance bool           `json:"normalBalance"`
	Title         string         `json:"title"`
	ParentId      int64          `json:"parentId"`
	ShortName     string         `json:"shortName"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}
