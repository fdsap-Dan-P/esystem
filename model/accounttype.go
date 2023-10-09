package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccountType struct {
	Id              int64     `json:"id"`
	Uuid            uuid.UUID `json:"uuid"`
	CentralOfficeId int64     `json:"centralOfficeId"`
	Code            int64     `json:"code"`
	AccountType     string    `json:"accountType"`
	ProductId       int64     `json:"productId"`
	GroupId         int64     `json:"groupId"`
	// Iiid            sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	FilterType    int64          `json:"filterType"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

type AccountTypeGroup struct {
	Id               int64     `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	ProductId        int64     `json:"productId"`
	GroupId          int64     `json:"groupId"`
	AccountTypeGroup string    `json:"accountTypeGroup"`
	// Iiid             sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

type AccountTypeFilter struct {
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Allow           bool           `json:"allow"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}
