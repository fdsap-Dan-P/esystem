package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createPersonalInfo = `-- name: CreatePersonalInfo: one
INSERT INTO Personal_Info (
  id, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26
) 
ON CONFLICT(Id)
DO UPDATE SET
  maiden_lname = excluded.maiden_lname,
  maiden_fname = excluded.maiden_fname,
  maiden_mname = excluded.maiden_mname,
  spouse_fname = excluded.spouse_fname,
  spouse_lname = excluded.spouse_lname,
  spouse_mname = excluded.spouse_mname,
  mother_maiden_lname = excluded.mother_maiden_lname,
  mother_maiden_fname = excluded.mother_maiden_fname,
  mother_maiden_mname = excluded.mother_maiden_mname,
  Marriage_Date = excluded.Marriage_Date, 
  isAdopted = excluded.isAdopted, 
  Known_Language = excluded.Known_Language, 
  Industry_Id = excluded.Industry_Id, 
  Nationality_Id = excluded.Nationality_Id, 
  Occupation_Id = excluded.Occupation_Id,   
  education_id = excluded.education_id,
  Religion_Id = excluded.Religion_Id, 
  Sector_Id = excluded.Sector_Id, 
  Source_Income_Id = excluded.Source_Income_Id, 
  Disability_Id = excluded.Disability_Id,
  Business_Name = excluded.Business_Name,
  Business_Address = excluded.Business_Address,
  Business_Address_Id = excluded.Business_Address_Id,
  Business_Position = excluded.Business_Position,
  Other_Info = excluded.Other_Info
RETURNING 
  id, UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info
`

type PersonalInfoRequest struct {
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
	KnownLanguage     sql.NullString `json:"knownLanguage"`
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

func (q *QueriesIdentity) CreatePersonalInfo(ctx context.Context, arg PersonalInfoRequest) (model.PersonalInfo, error) {
	row := q.db.QueryRowContext(ctx, createPersonalInfo,
		arg.Id,
		arg.MaidenLname,
		arg.MaidenFname,
		arg.MaidenMname,
		arg.SpouseLname,
		arg.SpouseFname,
		arg.SpouseMname,
		arg.MotherMaidenLname,
		arg.MotherMaidenFname,
		arg.MotherMaidenMname,
		arg.MarriageDate,
		arg.Isadopted,
		arg.KnownLanguage,
		arg.IndustryId,
		arg.NationalityId,
		arg.OccupationId,
		arg.EducationId,
		arg.ReligionId,
		arg.SectorId,
		arg.SourceIncomeId,
		arg.DisabilityId,
		arg.BusinessName,
		arg.BusinessAddress,
		arg.BusinessAddressId,
		arg.BusinessPosition,
		arg.OtherInfo,
	)
	var i model.PersonalInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.MaidenLname,
		&i.MaidenFname,
		&i.MaidenMname,
		&i.SpouseLname,
		&i.SpouseFname,
		&i.SpouseMname,
		&i.MotherMaidenLname,
		&i.MotherMaidenFname,
		&i.MotherMaidenMname,
		&i.MarriageDate,
		&i.Isadopted,
		&i.KnownLanguage,
		&i.IndustryId,
		&i.NationalityId,
		&i.OccupationId,
		&i.EducationId,
		&i.ReligionId,
		&i.SectorId,
		&i.SourceIncomeId,
		&i.DisabilityId,
		&i.BusinessName,
		&i.BusinessAddress,
		&i.BusinessAddressId,
		&i.BusinessPosition,
		&i.OtherInfo,
	)
	return i, err
}

const deletePersonalInfo = `-- name: DeletePersonalInfo :exec
DELETE FROM Personal_Info
WHERE id = $1
`

func (q *QueriesIdentity) DeletePersonalInfo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePersonalInfo, id)
	return err
}

type PersonalInfoInfo struct {
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
	KnownLanguage     sql.NullString `json:"knownLanguage"`
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
	ModCtr            int64          `json:"modCtr"`
	Created           sql.NullTime   `json:"created"`
	Updated           sql.NullTime   `json:"updated"`
}

const getPersonalInfo = `-- name: GetPersonalInfo :one
SELECT 
  id, mr.UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Personal_Info d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetPersonalInfo(ctx context.Context, id int64) (PersonalInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getPersonalInfo, id)
	var i PersonalInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.MaidenLname,
		&i.MaidenFname,
		&i.MaidenMname,
		&i.SpouseLname,
		&i.SpouseFname,
		&i.SpouseMname,
		&i.MotherMaidenLname,
		&i.MotherMaidenFname,
		&i.MotherMaidenMname,
		&i.MarriageDate,
		&i.Isadopted,
		&i.KnownLanguage,
		&i.IndustryId,
		&i.NationalityId,
		&i.OccupationId,
		&i.EducationId,
		&i.ReligionId,
		&i.SectorId,
		&i.SourceIncomeId,
		&i.DisabilityId,
		&i.BusinessName,
		&i.BusinessAddress,
		&i.BusinessAddressId,
		&i.BusinessPosition,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getPersonalInfobyUuId = `-- name: GetPersonalInfobyUuId :one
SELECT 
  id, mr.UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, 
  business_name, business_address, business_address_Id, business_position, other_info,   
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Personal_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetPersonalInfobyUuId(ctx context.Context, uuid uuid.UUID) (PersonalInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getPersonalInfobyUuId, uuid)
	var i PersonalInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.MaidenLname,
		&i.MaidenFname,
		&i.MaidenMname,
		&i.SpouseLname,
		&i.SpouseFname,
		&i.SpouseMname,
		&i.MotherMaidenLname,
		&i.MotherMaidenFname,
		&i.MotherMaidenMname,
		&i.MarriageDate,
		&i.Isadopted,
		&i.KnownLanguage,
		&i.IndustryId,
		&i.NationalityId,
		&i.OccupationId,
		&i.EducationId,
		&i.ReligionId,
		&i.SectorId,
		&i.SourceIncomeId,
		&i.DisabilityId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getPersonalInfobyName = `-- name: GetPersonalInfobyName :one
SELECT 
  id, mr.UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Personal_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesIdentity) GetPersonalInfobyName(ctx context.Context, name string) (PersonalInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getPersonalInfobyName, name)
	var i PersonalInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.MaidenLname,
		&i.MaidenFname,
		&i.MaidenMname,
		&i.SpouseLname,
		&i.SpouseFname,
		&i.SpouseMname,
		&i.MotherMaidenLname,
		&i.MotherMaidenFname,
		&i.MotherMaidenMname,
		&i.MarriageDate,
		&i.Isadopted,
		&i.KnownLanguage,
		&i.IndustryId,
		&i.NationalityId,
		&i.OccupationId,
		&i.EducationId,
		&i.ReligionId,
		&i.SectorId,
		&i.SourceIncomeId,
		&i.DisabilityId,
		&i.BusinessName,
		&i.BusinessAddress,
		&i.BusinessAddressId,
		&i.BusinessPosition,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listPersonalInfo = `-- name: ListPersonalInfo:many
SELECT 
  id, mr.UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Personal_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPersonalInfoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListPersonalInfo(ctx context.Context, arg ListPersonalInfoParams) ([]PersonalInfoInfo, error) {
	rows, err := q.db.QueryContext(ctx, listPersonalInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PersonalInfoInfo{}
	for rows.Next() {
		var i PersonalInfoInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.MaidenLname,
			&i.MaidenFname,
			&i.MaidenMname,
			&i.SpouseLname,
			&i.SpouseFname,
			&i.SpouseMname,
			&i.MotherMaidenLname,
			&i.MotherMaidenFname,
			&i.MotherMaidenMname,
			&i.MarriageDate,
			&i.Isadopted,
			&i.KnownLanguage,
			&i.IndustryId,
			&i.NationalityId,
			&i.OccupationId,
			&i.EducationId,
			&i.ReligionId,
			&i.SectorId,
			&i.SourceIncomeId,
			&i.DisabilityId,
			&i.BusinessName,
			&i.BusinessAddress,
			&i.BusinessAddressId,
			&i.BusinessPosition,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePersonalInfo = `-- name: UpdatePersonalInfo :one
UPDATE Personal_Info SET 
  maiden_lname = $2,
  maiden_fname = $3,
  maiden_mname = $4,
  spouse_fname = $5,
  spouse_lname = $6,
  spouse_mname = $7,
  mother_maiden_lname = $8,
  mother_maiden_fname = $9,
  mother_maiden_mname = $10,
  marriage_date = $11,
  isadopted = $12,
  known_language = $13,
  industry_id = $14,
  nationality_id = $15,
  occupation_id = $16,
  education_id = $17,
  religion_id = $18,
  sector_id = $19,
  source_income_id = $20,
  disability_id = $21,
  Business_Name = $22,
  Business_Address = $23,
  Business_Address_Id = $24,
  Business_Position = $25,
  other_info = $26
WHERE id = $1
RETURNING 
  id, UUID, maiden_lname, maiden_fname, maiden_mname, 
  spouse_fname, spouse_lname, spouse_mname, 
  mother_maiden_lname, mother_maiden_fname, mother_maiden_mname, 
  marriage_date, isadopted, known_language, industry_id, nationality_id, 
  occupation_id, education_id, religion_id, sector_id, source_income_id, disability_id, 
  business_name, business_address, business_address_Id, business_position, other_info
`

func (q *QueriesIdentity) UpdatePersonalInfo(ctx context.Context, arg PersonalInfoRequest) (model.PersonalInfo, error) {
	row := q.db.QueryRowContext(ctx, updatePersonalInfo,
		arg.Id,
		arg.MaidenLname,
		arg.MaidenFname,
		arg.MaidenMname,
		arg.SpouseLname,
		arg.SpouseFname,
		arg.SpouseMname,
		arg.MotherMaidenLname,
		arg.MotherMaidenFname,
		arg.MotherMaidenMname,
		arg.MarriageDate,
		arg.Isadopted,
		arg.KnownLanguage,
		arg.IndustryId,
		arg.NationalityId,
		arg.OccupationId,
		arg.EducationId,
		arg.ReligionId,
		arg.SectorId,
		arg.SourceIncomeId,
		arg.DisabilityId,
		arg.BusinessName,
		arg.BusinessAddress,
		arg.BusinessAddressId,
		arg.BusinessPosition,
		arg.OtherInfo,
	)
	var i model.PersonalInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.MaidenLname,
		&i.MaidenFname,
		&i.MaidenMname,
		&i.SpouseLname,
		&i.SpouseFname,
		&i.SpouseMname,
		&i.MotherMaidenLname,
		&i.MotherMaidenFname,
		&i.MotherMaidenMname,
		&i.MarriageDate,
		&i.Isadopted,
		&i.KnownLanguage,
		&i.IndustryId,
		&i.NationalityId,
		&i.OccupationId,
		&i.EducationId,
		&i.ReligionId,
		&i.SectorId,
		&i.SourceIncomeId,
		&i.DisabilityId,
		&i.BusinessName,
		&i.BusinessAddress,
		&i.BusinessAddressId,
		&i.BusinessPosition,
		&i.OtherInfo,
	)
	return i, err
}
