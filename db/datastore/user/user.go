package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
	"simplebank/util"
)

const createUser = `-- name: CreateUser: one
INSERT INTO Users (
  IIID, Login_Name, Display_Name, Access_Role_Id, Status_Code, 
  Date_Given, Date_Expired, Date_Locked, Password_Changed_At, Hashed_Password, 
  Attempt, Isloggedin, Thumbnail, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING Id, UUID, IIID, Login_Name, Display_Name, Access_Role_Id, Status_Code, Date_Given, 
Date_Expired, Date_Locked, Password_Changed_At, Hashed_Password, Attempt, Isloggedin, 
Thumbnail, Other_Info
`

type UserRequest struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	Iiid              int64          `json:"iiid"`
	LoginName         string         `json:"loginName" binding:"required,alphanum"`
	DisplayName       sql.NullString `json:"DisplayName"`
	AccessRoleId      int64          `json:"accessRoleId"`
	StatusCode        int64          `json:"statusCode"`
	DateGiven         sql.NullTime   `json:"dateGiven"`
	DateExpired       sql.NullTime   `json:"dateExpired"`
	DateLocked        sql.NullTime   `json:"dateLocked"`
	PasswordChangedAt sql.NullTime   `json:"passwordChangedAt"`
	Password          string         `json:"password" binding:"required,min=6"`
	Attempt           int16          `json:"attempt"`
	Isloggedin        sql.NullBool   `json:"isloggedin"`
	Thumbnail         []byte         `json:"thumbnail"`
	OtherInfo         sql.NullString `json:"otherInfo"`
}

// IsCorrectPassword checks if the provided password is correct or not
func (user *UserInfo) IsCorrectPassword(password string) bool {
	// err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return util.CheckPassword(password, user.HashedPassword) == nil
}

func (q *QueriesUser) CreateUser(ctx context.Context, arg UserRequest) (model.Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Iiid,
		arg.LoginName,
		arg.DisplayName,
		arg.AccessRoleId,
		arg.StatusCode,
		arg.DateGiven,
		arg.DateExpired,
		arg.DateLocked,
		arg.PasswordChangedAt,
		model.HashedPassword(arg.LoginName, arg.Password),
		arg.Attempt,
		arg.Isloggedin,
		arg.Thumbnail,
		arg.OtherInfo,
	)
	var i model.Users
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.LoginName,
		&i.DisplayName,
		&i.AccessRoleId,
		&i.StatusCode,
		&i.DateGiven,
		&i.DateExpired,
		&i.DateLocked,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.Attempt,
		&i.Isloggedin,
		&i.Thumbnail,
		&i.OtherInfo,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1
`

func (q *QueriesUser) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

type UserInfo struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	Iiid              int64          `json:"iiid"`
	LoginName         string         `json:"loginName"`
	DisplayName       sql.NullString `json:"displayName"`
	Title             sql.NullInt64  `json:"title"`
	LastName          string         `json:"last_name"`
	FirstName         sql.NullString `json:"first_name"`
	MiddleName        sql.NullString `json:"middle_name"`
	MotherMaidenName  sql.NullString `json:"mother_maiden_name"`
	Birthday          sql.NullTime   `json:"birthday"`
	Sex               sql.NullBool   `json:"sex"`
	GenderId          sql.NullInt64  `json:"genderId"`
	AccessRoleId      int64          `json:"accessRoleId"`
	AccessRole        string         `json:"accessRole"`
	StatusCode        int64          `json:"statusCode"`
	Status            string         `json:"status"`
	DateGiven         sql.NullTime   `json:"dateGiven"`
	DateExpired       sql.NullTime   `json:"dateExpired"`
	DateLocked        sql.NullTime   `json:"dateLocked"`
	PasswordChangedAt sql.NullTime   `json:"passwordChangedAt"`
	HashedPassword    []byte         `json:"hashedPassword"`
	Attempt           int16          `json:"attempt"`
	Isloggedin        sql.NullBool   `json:"isloggedin"`
	Thumbnail         []byte         `json:"thumbnail"`
	OtherInfo         sql.NullString `json:"otherInfo"`
	ModCtr            int64          `json:"modCtr"`
	Created           sql.NullTime   `json:"created"`
	Updated           sql.NullTime   `json:"updated"`
}

const getUser = `-- name: GetUser :one
SELECT 
  d.Id, mr.UUId, d.IIId, d.Login_Name,  d.Display_Name, 
  ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
  ii.Birthday, ii.Sex, ii.Gender_id,
  d.Access_Role_Id, r.Access_Name Access_Role, d.Status_Code, s.Title Status, d.Date_Given, d.Date_Expired, d.Date_Locked, 
  d.Password_Changed_At, d.Hashed_Password, d.Attempt, d.Isloggedin, Thumbnail, d.Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users d 
INNER JOIN Main_Record mr on mr.UUId = d.UUID
INNER JOIN Identity_Info ii on ii.Id = d.IIID
INNER JOIN Access_Role r on r.Id = d.Access_Role_Id
INNER JOIN Reference s on s.Id = d.Status_Code
WHERE d.id = $1 LIMIT 1
`

func (q *QueriesUser) GetUser(ctx context.Context, id int64) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i UserInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.LoginName,
		&i.DisplayName,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.AccessRoleId,
		&i.AccessRole,
		&i.StatusCode,
		&i.Status,
		&i.DateGiven,
		&i.DateExpired,
		&i.DateLocked,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.Attempt,
		&i.Isloggedin,
		&i.Thumbnail,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserbyUuId = `-- name: GetUserbyUuId :one
SELECT 
  d.Id, mr.UUId, d.IIId, d.Login_Name,  d.Display_Name, 
  ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
  ii.Birthday, ii.Sex, ii.Gender_id,
  d.Access_Role_Id, r.Access_Name Access_Role, d.Status_Code, s.Title Status, d.Date_Given, d.Date_Expired, d.Date_Locked, 
  d.Password_Changed_At, d.Hashed_Password, d.Attempt, d.Isloggedin, Thumbnail, d.Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users d 
INNER JOIN Main_Record mr on mr.UUId = d.UUID
INNER JOIN Identity_Info ii on ii.Id = d.IIID
INNER JOIN Access_Role r on r.Id = d.Access_Role_Id
INNER JOIN Reference s on s.Id = d.Status_Code
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetUserbyUuId(ctx context.Context, uuid uuid.UUID) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserbyUuId, uuid)
	var i UserInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.LoginName,
		&i.DisplayName,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.AccessRoleId,
		&i.AccessRole,
		&i.StatusCode,
		&i.Status,
		&i.DateGiven,
		&i.DateExpired,
		&i.DateLocked,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.Attempt,
		&i.Isloggedin,
		&i.Thumbnail,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserbyName = `-- name: GetUserbyName :one
SELECT 
  d.Id, mr.UUId, d.IIId, d.Login_Name,  d.Display_Name, 
  ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
  ii.Birthday, ii.Sex, ii.Gender_id,
  d.Access_Role_Id, r.Access_Name Access_Role, d.Status_Code, s.Title Status, d.Date_Given, d.Date_Expired, d.Date_Locked, 
  d.Password_Changed_At, d.Hashed_Password, d.Attempt, d.Isloggedin, Thumbnail, d.Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users d 
INNER JOIN Main_Record mr on mr.UUId = d.UUID
INNER JOIN Identity_Info ii on ii.Id = d.IIID
INNER JOIN Access_Role r on r.Id = d.Access_Role_Id
INNER JOIN Reference s on s.Id = d.Status_Code
WHERE lower(Login_Name) = lower($1) LIMIT 1
`

func (q *QueriesUser) GetUserbyName(ctx context.Context, name string) (UserInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserbyName, name)
	var i UserInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.LoginName,
		&i.DisplayName,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.AccessRoleId,
		&i.AccessRole,
		&i.StatusCode,
		&i.Status,
		&i.DateGiven,
		&i.DateExpired,
		&i.DateLocked,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.Attempt,
		&i.Isloggedin,
		&i.Thumbnail,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUser = `-- name: ListUser:many
SELECT 
  d.Id, mr.UUId, d.IIId, d.Login_Name,  d.Display_Name, 
  ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
  ii.Birthday, ii.Sex, ii.Gender_id,
  d.Access_Role_Id, r.Access_Name Access_Role, d.Status_Code, s.Title Status, d.Date_Given, d.Date_Expired, d.Date_Locked, 
  d.Password_Changed_At, d.Hashed_Password, d.Attempt, d.Isloggedin, Thumbnail, d.Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Users d 
INNER JOIN Main_Record mr on mr.UUId = d.UUID
INNER JOIN Identity_Info ii on ii.Id = d.IIID
INNER JOIN Access_Role r on r.Id = d.Access_Role_Id
INNER JOIN Reference s on s.Id = d.Status_Code
WHERE Iiid = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUserParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUser(ctx context.Context, arg ListUserParams) ([]UserInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUser, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserInfo{}
	for rows.Next() {
		var i UserInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Iiid,
			&i.LoginName,
			&i.DisplayName,
			&i.Title,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.MotherMaidenName,
			&i.Birthday,
			&i.Sex,
			&i.GenderId,
			&i.AccessRoleId,
			&i.AccessRole,
			&i.StatusCode,
			&i.Status,
			&i.DateGiven,
			&i.DateExpired,
			&i.DateLocked,
			&i.PasswordChangedAt,
			&i.HashedPassword,
			&i.Attempt,
			&i.Isloggedin,
			&i.Thumbnail,
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

const updateUser = `-- name: UpdateUser :one
UPDATE Users SET 
	IIId = $2,
	Login_Name = $3,
	Display_Name = $4,
	Access_Role_Id = $5,
	Status_Code = $6,
	Date_Given = $7,
	Date_Expired = $8,
	Date_Locked = $9,
	Password_Changed_At = $10,
	Hashed_Password = $11,
	Attempt = $12,
	Isloggedin = $13,
  Thumbnail = $14,
	Other_Info = $15
WHERE id = $1
RETURNING Id, UUId, IIId, Login_Name, Display_Name, Access_Role_Id, Status_Code, Date_Given, 
Date_Expired, Date_Locked, Password_Changed_At, Hashed_Password, Attempt, Isloggedin, 
Thumbnail, Other_Info
`

func (q *QueriesUser) UpdateUser(ctx context.Context, arg UserRequest) (model.Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Id,
		arg.Iiid,
		arg.LoginName,
		arg.DisplayName,
		arg.AccessRoleId,
		arg.StatusCode,
		arg.DateGiven,
		arg.DateExpired,
		arg.DateLocked,
		arg.PasswordChangedAt,
		model.HashedPassword(arg.LoginName, arg.Password),
		arg.Attempt,
		arg.Isloggedin,
		arg.Thumbnail,
		arg.OtherInfo,
	)
	var i model.Users
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Iiid,
		&i.LoginName,
		&i.DisplayName,
		&i.AccessRoleId,
		&i.StatusCode,
		&i.DateGiven,
		&i.DateExpired,
		&i.DateLocked,
		&i.PasswordChangedAt,
		&i.HashedPassword,
		&i.Attempt,
		&i.Isloggedin,
		&i.Thumbnail,
		&i.OtherInfo,
	)
	return i, err
}

const changePass = `-- name: ChangePass :one
UPDATE Users SET 
	Hashed_Password = $2
WHERE lower(Login_Name) = lower($1)
RETURNING id
`

func (q *QueriesUser) ChangePass(ctx context.Context, loginName string, password string) (bool, error) {
	row := q.db.QueryRowContext(ctx, changePass,
		loginName, model.HashedPassword(loginName, password))
	var i int64
	err := row.Scan(&i)
	return (err == nil), err
}
