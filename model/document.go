package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Document struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	Code        string         `json:"code"`
	OfficeId    sql.NullInt64  `json:"officeId"`
	DocDate     sql.NullTime   `json:"docDate"`
	FilePath    string         `json:"filePath"`
	Thumbnail   []byte         `json:"thumbnail"`
	DoctypeId   int64          `json:"doctypeId"`
	ServerId    int64          `json:"serverId"`
	SecretKey   sql.NullString `json:"secretKey"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

type DocumentAccess struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	RoleId     int64          `json:"roleId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

type DocumentUser struct {
	Uuid       uuid.UUID      `json:"uuid"`
	DocumentId int64          `json:"documentId"`
	UserId     int64          `json:"userId"`
	AccessCode string         `json:"accessCode"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}
