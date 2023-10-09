package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Customer struct {
	Id                  int64           `json:"id"`
	Uuid                uuid.UUID       `json:"uuid"`
	Iiid                int64           `json:"iiid"`
	CentralOfficeId     int64           `json:"centralOfficeId"`
	Cid                 int64           `json:"Cid"`
	CustomerAltId       sql.NullString  `json:"customerAltId"`
	DebitLimit          decimal.Decimal `json:"debitLimit"`
	CreditLimit         decimal.Decimal `json:"creditLimit"`
	DateEntry           sql.NullTime    `json:"dateEntry"`
	DateRecognized      sql.NullTime    `json:"dateRecognized"`
	DateResigned        sql.NullTime    `json:"dateResigned"`
	Resigned            sql.NullBool    `json:"resigned"`
	ReasonResigned      sql.NullString  `json:"reasonResigned"`
	LastActivityDate    sql.NullTime    `json:"lastActivityDate"`
	Dosri               bool            `json:"dosri"`
	ClassificationId    sql.NullInt64   `json:"classificationId"`
	CustomerGroupId     sql.NullInt64   `json:"customerGroupId"`
	OfficeId            int64           `json:"officeId"`
	RestrictionId       sql.NullInt64   `json:"restrictionId"`
	RiskClassId         sql.NullInt64   `json:"riskClassId"`
	StatusCode          int64           `json:"statusCode"`
	IndustryId          sql.NullInt64   `json:"industryId"`
	SubClassificationId sql.NullInt64   `json:"subClassificationId"`
	Remarks             sql.NullString  `json:"remarks"`
	OtherInfo           sql.NullString  `json:"otherInfo"`
}

type CustomerSpecsString struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsCode  int64     `json:"specsCode"`
	SpecsId    int64     `json:"specsId"`
	Value      string    `json:"value"`
}

type CustomerSpecsNumber struct {
	Uuid       uuid.UUID       `json:"uuid"`
	CustomerId int64           `json:"customerId"`
	SpecsCode  int64           `json:"specsCode"`
	SpecsId    int64           `json:"specsId"`
	Value      decimal.Decimal `json:"value"`
	Value2     decimal.Decimal `json:"value2"`
	MeasureId  sql.NullInt64   `json:"measureId"`
}

type CustomerSpecsDate struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsCode  int64     `json:"specsCode"`
	SpecsId    int64     `json:"specsId"`
	Value      time.Time `json:"value"`
	Value2     time.Time `json:"value2"`
}

type CustomerSpecsRef struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsCode  int64     `json:"specsCode"`
	SpecsId    int64     `json:"specsId"`
	RefId      int64     `json:"refId"`
}
