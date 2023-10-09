package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	CustomerId       int64           `json:"customerId"`
	Acc              string          `json:"acc"`
	AlternateAcc     sql.NullString  `json:"alternateAcc"`
	AccountName      string          `json:"accountName"`
	Balance          decimal.Decimal `json:"balance"`
	NonCurrent       decimal.Decimal `json:"nonCurrent"`
	ContractDate     sql.NullTime    `json:"contractDate"`
	Credit           decimal.Decimal `json:"credit"`
	Debit            decimal.Decimal `json:"debit"`
	Isbudget         sql.NullBool    `json:"isbudget"`
	LastActivityDate sql.NullTime    `json:"lastActivityDate"`
	OpenDate         time.Time       `json:"openDate"`
	PassbookLine     int16           `json:"passbookLine"`
	PendingTrnAmt    decimal.Decimal `json:"pendingTrnAmt"`
	Principal        decimal.Decimal `json:"principal"`
	ClassId          int64           `json:"classId"`
	AccountTypeId    int64           `json:"accountTypeId"`
	BudgetAccountId  sql.NullInt64   `json:"budgetAccountId"`
	CategoryId       sql.NullInt64   `json:"categoryId"`
	Currency         string          `json:"currency"`
	OfficeId         int64           `json:"officeId"`
	ReferredbyId     sql.NullInt64   `json:"referredbyId"`
	StatusCode       int64           `json:"statusCode"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

type AccountSpecsString struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

type AccountSpecsNumber struct {
	Uuid      uuid.UUID       `json:"uuid"`
	AccountId int64           `json:"accountId"`
	SpecsCode string          `json:"specsCode"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

type AccountSpecsDate struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

type AccountSpecsRef struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsCode string    `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}
