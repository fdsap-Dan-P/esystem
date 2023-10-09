package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IdentityInfo struct {
	Id                   int64          `json:"id"`
	Uuid                 uuid.UUID      `json:"uuid"`
	IdentityMapId        sql.NullInt64  `json:"identityMapId"`
	Isperson             bool           `json:"isperson"`
	AlternateId          sql.NullString `json:"alternateId"`
	Title                sql.NullString `json:"title"`
	LastName             string         `json:"lastName"`
	FirstName            sql.NullString `json:"firstName"`
	MiddleName           sql.NullString `json:"middleName"`
	SuffixName           sql.NullString `json:"suffixName"`
	ProfessionalSuffixes sql.NullString `json:"professionalSuffixes"`
	MotherMaidenName     sql.NullString `json:"motherMaidenName"`
	Birthday             sql.NullTime   `json:"birthday"`
	Sex                  sql.NullBool   `json:"sex"`
	GenderId             sql.NullInt64  `json:"genderId"`
	CivilStatusId        sql.NullInt64  `json:"civilStatusId"`
	BirthPlaceId         sql.NullInt64  `json:"birthPlaceId"`
	ContactId            sql.NullInt64  `json:"contactId"`
	Phone                sql.NullString `json:"phone"`
	Email                sql.NullString `json:"email"`
	SimpleName           sql.NullString `json:"simpleName"`
	VecSimpleName        sql.NullString `json:"vecSimpleName"`
	VecFullSimpleName    sql.NullString `json:"vecFullSimpleName"`
	OtherInfo            sql.NullString `json:"otherInfo"`
}

type IdentitySpecsString struct {
	Uuid      uuid.UUID `json:"uuid"`
	Iiid      int64     `json:"iiid"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

type IdentitySpecsNumber struct {
	Uuid      uuid.UUID       `json:"uuid"`
	Iiid      int64           `json:"iiid"`
	SpecsCode string          `json:"specsCode"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

type IdentitySpecsDate struct {
	Uuid      uuid.UUID `json:"uuid"`
	Iiid      int64     `json:"iiid"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

type IdentitySpecsRef struct {
	Uuid      uuid.UUID `json:"uuid"`
	Iiid      int64     `json:"iiid"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}
