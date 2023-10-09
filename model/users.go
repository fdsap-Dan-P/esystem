package model

import (
	"database/sql"
	"simplebank/util"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const Seed4Password = "17OV3pR0@raMminG"

type Users struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	Iiid              int64          `json:"iiid"`
	LoginName         string         `json:"loginName"`
	DisplayName       sql.NullString `json:"displayName"`
	AccessRoleId      int64          `json:"accessRoleId"`
	StatusCode        int64          `json:"statusCode"`
	DateGiven         sql.NullTime   `json:"dateGiven"`
	DateExpired       sql.NullTime   `json:"dateExpired"`
	DateLocked        sql.NullTime   `json:"dateLocked"`
	PasswordChangedAt sql.NullTime   `json:"passwordChangedAt"`
	HashedPassword    []byte         `json:"hashedPassword"`
	Attempt           int16          `json:"attempt"`
	Isloggedin        sql.NullBool   `json:"isloggedin"`
	Thumbnail         []byte         `json:"thumbnail"`
	OtherInfo         sql.NullString `json:"otherInfo"`
}

// NewUser returns a new user

func HashedPassword(userName string, pass string) []byte {
	// encode := util.Encode(userName+Seed4Password, pass)
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(encode), bcrypt.DefaultCost)
	// if err != nil {
	// 	return []byte{}
	// }
	hash, _ := util.HashPassword(userName+Seed4Password, pass)
	return []byte(hash)
}

func NewUser(
	iiid int64,
	loginName string,
	displayName sql.NullString,
	accessRoleId int64,
	statusCode int64,
	dateGiven sql.NullTime,
	dateExpired sql.NullTime,
	dateLocked sql.NullTime,
	passwordChangedAt sql.NullTime,
	pass string,
	attempt int16,
	isloggedin sql.NullBool,
	otherInfo sql.NullString,
	thumbnail []byte) (*Users, error) {

	user := &Users{
		Iiid:              iiid,
		LoginName:         loginName,
		DisplayName:       displayName,
		AccessRoleId:      accessRoleId,
		StatusCode:        statusCode,
		DateGiven:         dateGiven,
		DateExpired:       dateExpired,
		DateLocked:        dateLocked,
		PasswordChangedAt: passwordChangedAt,
		HashedPassword:    HashedPassword(loginName, pass),
		Attempt:           attempt,
		Isloggedin:        isloggedin,
		OtherInfo:         otherInfo,
		Thumbnail:         thumbnail,
	}
	return user, nil
}

// IsCorrectPassword checks if the provided password is correct or not
func (user *Users) IsCorrectPassword(password string) bool {
	// err := bcrypt.CompareHashAndPassword([]byte(users.HashedPassword), []byte(password))
	return util.CheckPassword(password, user.HashedPassword) == nil
}

// Clone returns a clone of this user
func (users *Users) Clone() *Users {
	return &Users{
		Iiid:              users.Iiid,
		LoginName:         users.LoginName,
		DisplayName:       users.DisplayName,
		AccessRoleId:      users.AccessRoleId,
		StatusCode:        users.StatusCode,
		DateGiven:         users.DateGiven,
		DateExpired:       users.DateExpired,
		DateLocked:        users.DateLocked,
		PasswordChangedAt: users.PasswordChangedAt,
		HashedPassword:    users.HashedPassword,
		Attempt:           users.Attempt,
		Isloggedin:        users.Isloggedin,
		Thumbnail:         users.Thumbnail,
		OtherInfo:         users.OtherInfo,
	}
}

type UserSpecsString struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserCode  string    `json:"userCode"`
	UserId    int64     `json:"userId"`
	SpecsCode int64     `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

type UserSpecsNumber struct {
	Uuid      uuid.UUID       `json:"uuid"`
	UserId    int64           `json:"userId"`
	UserCode  string          `json:"userCode"`
	SpecsCode int64           `json:"specsCode"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"MeasureId"`
}

type UserSpecsDate struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserId    int64     `json:"userId"`
	UserCode  string    `json:"userCode"`
	SpecsCode int64     `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

type UserSpecsRef struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserId    int64     `json:"userId"`
	UserCode  string    `json:"userCode"`
	SpecsCode int64     `json:"specsCode"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}
