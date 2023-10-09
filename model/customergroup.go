package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type CustomerGroup struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Code            string         `json:"code"`
	TypeId          int64          `json:"typeId"`
	GroupName       string         `json:"groupName"`
	ShortName       string         `json:"shortName"`
	DateStablished  sql.NullTime   `json:"dateStablished"`
	MeetingDay      sql.NullInt16  `json:"meetingDay"`
	OfficeId        int64          `json:"officeId"`
	OfficerId       sql.NullInt64  `json:"officerId"`
	ParentId        sql.NullInt64  `json:"parentId"`
	AlternateId     sql.NullString `json:"alternateId"`
	AddressDetail   sql.NullString `json:"addressDetail"`
	AddressUrl      sql.NullString `json:"addressUrl"`
	GeographyId     sql.NullInt64  `json:"geographyId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}
