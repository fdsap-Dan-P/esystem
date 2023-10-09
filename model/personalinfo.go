package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type PersonalInfo struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	MaidenLname       sql.NullString `json:"maidenLname"`
	MaidenFname       sql.NullString `json:"maidenFname"`
	MaidenMname       sql.NullString `json:"maidenMname"`
	SpouseLname       sql.NullString `json:"spouseLname"`
	SpouseFname       sql.NullString `json:"spouseFname"`
	SpouseMname       sql.NullString `json:"spouseMname"`
	MotherMaidenLname sql.NullString `json:"motherMaidenLname"`
	MotherMaidenFname sql.NullString `json:"motherMaidenFname"`
	MotherMaidenMname sql.NullString `json:"motherMaidenMname"`
	MarriageDate      sql.NullTime   `json:"marriageDate"`
	Isadopted         sql.NullBool   `json:"isadopted"`
	SpouseName        sql.NullString `json:"spouseName"`
	KnownLanguage     sql.NullString `json:"knownLanguage"`
	Disability        sql.NullInt64  `json:"disability"`
	IndustryId        sql.NullInt64  `json:"industryId"`
	NationalityId     sql.NullInt64  `json:"nationalityId"`
	OccupationId      sql.NullInt64  `json:"occupationId"`
	EducationId       sql.NullInt64  `json:"educationId"`
	ReligionId        sql.NullInt64  `json:"religionId"`
	SectorId          sql.NullInt64  `json:"sectorId"`
	SourceIncomeId    sql.NullInt64  `json:"sourceIncomeId"`
	DisabilityId      sql.NullInt64  `json:"disabilityId"`
	BusinessName      sql.NullString `json:"businessName"`
	BusinessAddress   sql.NullString `json:"businessAddress"`
	BusinessAddressId sql.NullInt64  `json:"businessAddressId"`
	BusinessPosition  sql.NullString `json:"businessPosition"`
	OtherInfo         sql.NullString `json:"otherInfo"`
}
