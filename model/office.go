package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Office struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	Code           string         `json:"code"`
	ShortName      string         `json:"shortName"`
	OfficeName     string         `json:"officeName"`
	DateStablished sql.NullTime   `json:"dateStablished"`
	TypeId         int64          `json:"typeId"`
	ParentId       sql.NullInt64  `json:"parentId"`
	AlternateId    sql.NullString `json:"alternateId"`
	AddressDetail  sql.NullString `json:"addressDetail"`
	AddressUrl     sql.NullString `json:"addressUrl"`
	GeographyId    sql.NullInt64  `json:"geographyId"`
	CidSequence    sql.NullInt64  `json:"cidSequence"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}
