package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createIdentityInfo = `-- name: CreateIdentityInfo: one
INSERT INTO Identity_Info (
isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, 
Contact_Id, Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
$13, $14, $15, $16, $17, $18, $19
) RETURNING Id, UUId, isPerson, Title, Last_Name, First_Name, Middle_Name, 
  Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, 
  Contact_Id, Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
`

type IdentityInfoRequest struct {
	Id                   int64          `json:"id"`
	Uuid                 uuid.UUID      `json:"uuid"`
	Isperson             bool           `json:"isperson"`
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
	IdentityMapId        sql.NullInt64  `json:"identityMapId"`
	AlternateId          sql.NullString `json:"alternateId"`
	Phone                sql.NullString `json:"phone"`
	Email                sql.NullString `json:"email"`
	OtherInfo            sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateIdentityInfo(ctx context.Context, arg IdentityInfoRequest) (model.IdentityInfo, error) {
	row := q.db.QueryRowContext(ctx, createIdentityInfo,
		arg.Isperson,
		arg.Title,
		arg.LastName,
		arg.FirstName,
		arg.MiddleName,
		arg.SuffixName,
		arg.ProfessionalSuffixes,
		arg.MotherMaidenName,
		arg.Birthday,
		arg.Sex,
		arg.GenderId,
		arg.CivilStatusId,
		arg.BirthPlaceId,
		arg.ContactId,
		arg.IdentityMapId,
		arg.AlternateId,
		arg.Phone,
		arg.Email,
		arg.OtherInfo,
	)
	var i model.IdentityInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,
	)
	return i, err
}

const deleteIdentityInfo = `-- name: DeleteIdentityInfo :exec
DELETE FROM Identity_Info
WHERE id = $1
`

func (q *QueriesIdentity) DeleteIdentityInfo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteIdentityInfo, id)
	return err
}

type IdentityInfoInfo struct {
	Id                   int64          `json:"id"`
	Uuid                 uuid.UUID      `json:"uuid"`
	Isperson             bool           `json:"isperson"`
	Title                sql.NullString `json:"title"`
	LastName             string         `json:"lastName"`
	FirstName            sql.NullString `json:"firstName"`
	MiddleName           sql.NullString `json:"MiddleName"`
	SuffixName           sql.NullString `json:"suffixName"`
	ProfessionalSuffixes sql.NullString `json:"professionalSuffixes"`
	MotherMaidenName     sql.NullString `json:"MotherMaidenName"`
	Birthday             sql.NullTime   `json:"birthday"`
	Sex                  sql.NullBool   `json:"sex"`
	GenderId             sql.NullInt64  `json:"genderId"`
	CivilStatusId        sql.NullInt64  `json:"civilStatusId"`
	BirthPlaceId         sql.NullInt64  `json:"birthPlaceId"`
	ContactId            sql.NullInt64  `json:"contactId"`
	IdentityMapId        sql.NullInt64  `json:"identityMapId"`
	AlternateId          sql.NullString `json:"alternateId"`
	Phone                sql.NullString `json:"phone"`
	Email                sql.NullString `json:"email"`
	OtherInfo            sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getIdentityInfo = `-- name: GetIdentityInfo :one
SELECT 
Id, mr.UUId, 
isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, 
Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, Contact_Id, 
Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIdentityInfo(ctx context.Context, id int64) (IdentityInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getIdentityInfo, id)
	var i IdentityInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIdentityInfobyAcc = `-- name: GetIdentityInfobyAcc :one
SELECT 
  Id, mr.UUId, 
  isPerson, Title, Last_Name, First_Name, Middle_Name, 
  Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, 
  Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, Contact_Id, 
  Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE acc = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIdentityInfobyAcc(ctx context.Context, acc string) (IdentityInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getIdentityInfobyAcc, acc)
	var i IdentityInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIdentityInfobyAltId = `-- name: GetIdentityInfobyAltId :one
SELECT 
Id, mr.UUId, 
isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, 
Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, Contact_Id, 
Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE d.Alternate_Id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIdentityInfobyAltId(ctx context.Context, altAcc string) (IdentityInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getIdentityInfobyAltId, altAcc)
	var i IdentityInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIdentityInfobyUuId = `-- name: GetIdentityInfobyUuId :one
SELECT 
Id, mr.UUId, 
isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, 
Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, Contact_Id, 
Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIdentityInfobyUuId(ctx context.Context, uuid uuid.UUID) (IdentityInfoInfo, error) {
	row := q.db.QueryRowContext(ctx, getIdentityInfobyUuId, uuid)
	var i IdentityInfoInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listIdentityInfo = `-- name: ListIdentityInfo:many
SELECT 
Id, mr.UUId, 
isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, 
Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, Contact_Id, 
Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Info d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListIdentityInfoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIdentityInfo(ctx context.Context, arg ListIdentityInfoParams) ([]IdentityInfoInfo, error) {
	rows, err := q.db.QueryContext(ctx, listIdentityInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IdentityInfoInfo{}
	for rows.Next() {
		var i IdentityInfoInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Isperson,
			&i.Title,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.SuffixName,
			&i.ProfessionalSuffixes,
			&i.MotherMaidenName,
			&i.Birthday,
			&i.Sex,
			&i.GenderId,
			&i.CivilStatusId,
			&i.BirthPlaceId,
			&i.ContactId,
			&i.IdentityMapId,
			&i.AlternateId,
			&i.Phone,
			&i.Email,
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

const updateIdentityInfo = `-- name: UpdateIdentityInfo :one
UPDATE Identity_Info SET 
	isPerson = $2,
	Title = $3,
	Last_Name = $4,
	First_Name = $5,
	Middle_Name = $6,
	Suffix_Name = $7, 
	Professional_Suffixes = $8, 
	Mother_Maiden_Name = $9,
	Birthday = $10,
	Sex = $11,
	Gender_id = $12,
	Civil_Status_Id = $13,
	Birth_Place_Id = $14,
	Contact_Id = $15,
	Identity_Map_Id = $16,
	Alternate_Id = $17,
	Phone = $18,
	Email = $19,
	Other_Info = $20
WHERE id = $1
RETURNING Id, UUId, isPerson, Title, Last_Name, First_Name, Middle_Name, 
Suffix_Name, Professional_Suffixes, Mother_Maiden_Name, Birthday, Sex, Gender_id, Civil_Status_Id, Birth_Place_Id, 
Contact_Id, Identity_Map_Id, Alternate_Id, Phone, Email, Other_Info
`

func (q *QueriesIdentity) UpdateIdentityInfo(ctx context.Context, arg IdentityInfoRequest) (model.IdentityInfo, error) {
	row := q.db.QueryRowContext(ctx, updateIdentityInfo,
		arg.Id,
		arg.Isperson,
		arg.Title,
		arg.LastName,
		arg.FirstName,
		arg.MiddleName,
		arg.SuffixName,
		arg.ProfessionalSuffixes,
		arg.MotherMaidenName,
		arg.Birthday,
		arg.Sex,
		arg.GenderId,
		arg.CivilStatusId,
		arg.BirthPlaceId,
		arg.ContactId,
		arg.IdentityMapId,
		arg.AlternateId,
		arg.Phone,
		arg.Email,
		arg.OtherInfo,
	)
	var i model.IdentityInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Isperson,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.SuffixName,
		&i.ProfessionalSuffixes,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.CivilStatusId,
		&i.BirthPlaceId,
		&i.ContactId,
		&i.IdentityMapId,
		&i.AlternateId,
		&i.Phone,
		&i.Email,
		&i.OtherInfo,
	)
	return i, err
}
