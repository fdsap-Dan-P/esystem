package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createUsersList = `-- name: CreateUsersList: one
IF EXISTS (SELECT UserId from UsersList WHERE UserId = $1)
BEGIN
  UPDATE UsersList SET 
	AccessCode = $2,
	LName = $3,
	FName = $4,
	MName = $5,
	DateHired = $6,
	BirthDay = $7,
	DateGiven = $8,
	DateExpired = $9,
	Address = $10,
	Position = $11,
	AreaCode = $12,
	ManCode = $13,
	AddInfo = $14,
	Passwd = $15,
	Attempt = $16,
	DateLocked = $17,
	Remarks = $18,
	Picture = $19,
	isLoggedIn = $20,
	AccountExpirationDt = $21
  WHERE UserId = $1
END ELSE BEGIN
  INSERT INTO UsersList (
    UserID, AccessCode, LName, FName, MName, DateHired, BirthDay, 
    DateGiven, DateExpired, Address, Position, AreaCode, ManCode, AddInfo, 
    Passwd, Attempt, DateLocked, Remarks, Picture, IsLoggedIn, AccountExpirationDt
) 
VALUES 
 ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) 
END
`

type UsersListRequest struct {
	UserId              string         `json:"userId"`
	AccessCode          sql.NullInt64  `json:"accessCode"`
	LName               string         `json:"lName"`
	FName               string         `json:"fName"`
	MName               string         `json:"mName"`
	DateHired           sql.NullTime   `json:"dateHired"`
	BirthDay            sql.NullTime   `json:"birthDay"`
	DateGiven           sql.NullTime   `json:"dateGiven"`
	DateExpired         sql.NullTime   `json:"dateExpired"`
	Address             sql.NullString `json:"address"`
	Position            sql.NullString `json:"position"`
	AreaCode            sql.NullInt64  `json:"areaCode"`
	ManCode             sql.NullInt64  `json:"manCode"`
	AddInfo             sql.NullString `json:"addInfo"`
	Passwd              []byte         `json:"passwd"`
	Attempt             sql.NullInt64  `json:"attempt"`
	DateLocked          sql.NullTime   `json:"dateLocked"`
	Remarks             sql.NullString `json:"remarks"`
	Picture             []byte         `json:"picture"`
	IsLoggedIn          bool           `json:"isLoggedIn"`
	AccountExpirationDt sql.NullTime   `json:"accountExpirationDt"`
}

func (q *QueriesLocal) CreateUsersList(ctx context.Context, arg UsersListRequest) error {
	log.Println(arg)
	_, err := q.db.ExecContext(ctx, createUsersList,
		arg.UserId,
		arg.AccessCode,
		arg.LName,
		arg.FName,
		arg.MName,
		arg.DateHired,
		arg.BirthDay,
		arg.DateGiven,
		arg.DateExpired,
		arg.Address,
		arg.Position,
		arg.AreaCode,
		arg.ManCode,
		arg.AddInfo,
		arg.Passwd,
		arg.Attempt,
		arg.DateLocked,
		arg.Remarks,
		arg.Picture,
		arg.IsLoggedIn,
		arg.AccountExpirationDt,
	)
	return err
}

const deleteUsersList = `-- name: DeleteUsersList :exec
DELETE FROM UsersList WHERE UserId = $1
`

func (q *QueriesLocal) DeleteUsersList(ctx context.Context, UserId string) error {
	_, err := q.db.ExecContext(ctx, deleteUsersList, UserId)
	return err
}

type UsersListInfo struct {
	ModCtr              int64          `json:"modCtr"`
	BrCode              string         `json:"brCode"`
	ModAction           string         `json:"modAction"`
	UserId              string         `json:"userId"`
	AccessCode          sql.NullInt64  `json:"accessCode"`
	LName               string         `json:"lName"`
	FName               string         `json:"fName"`
	MName               string         `json:"mName"`
	DateHired           sql.NullTime   `json:"dateHired"`
	BirthDay            sql.NullTime   `json:"birthDay"`
	DateGiven           sql.NullTime   `json:"dateGiven"`
	DateExpired         sql.NullTime   `json:"dateExpired"`
	Address             sql.NullString `json:"address"`
	Position            sql.NullString `json:"position"`
	AreaCode            sql.NullInt64  `json:"areaCode"`
	ManCode             sql.NullInt64  `json:"manCode"`
	AddInfo             sql.NullString `json:"addInfo"`
	Passwd              []byte         `json:"passwd"`
	Attempt             sql.NullInt64  `json:"attempt"`
	DateLocked          sql.NullTime   `json:"dateLocked"`
	Remarks             sql.NullString `json:"remarks"`
	Picture             []byte         `json:"picture"`
	IsLoggedIn          bool           `json:"isLoggedIn"`
	AccountExpirationDt sql.NullTime   `json:"accountExpirationDt"`
}

// -- name: GetUsersList :one
const getUsersList = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, 
  UserID, AccessCode, LName, FName, MName, DateHired, BirthDay, 
  DateGiven, DateExpired, Address, Position, AreaCode, ManCode, AddInfo, 
  Passwd, Attempt, DateLocked, Remarks, Picture, 
  CASE WHEN isnull(IsLoggedIn,0) = 1 THEN 1 ELSE 0 END IsLoggedIn, AccountExpirationDt
FROM OrgParms, UsersList d
INNER JOIN Modified m on m.UniqueKeyString1 = d.UserId
`

func scanRowUsersList(row *sql.Row) (UsersListInfo, error) {
	var i UsersListInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.UserId,
		&i.AccessCode,
		&i.LName,
		&i.FName,
		&i.MName,
		&i.DateHired,
		&i.BirthDay,
		&i.DateGiven,
		&i.DateExpired,
		&i.Address,
		&i.Position,
		&i.AreaCode,
		&i.ManCode,
		&i.AddInfo,
		&i.Passwd,
		&i.Attempt,
		&i.DateLocked,
		&i.Remarks,
		&i.Picture,
		&i.IsLoggedIn,
		&i.AccountExpirationDt,
	)
	return i, err
}

func scanRowsUsersList(rows *sql.Rows) ([]UsersListInfo, error) {
	items := []UsersListInfo{}
	for rows.Next() {
		var i UsersListInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.UserId,
			&i.AccessCode,
			&i.LName,
			&i.FName,
			&i.MName,
			&i.DateHired,
			&i.BirthDay,
			&i.DateGiven,
			&i.DateExpired,
			&i.Address,
			&i.Position,
			&i.AreaCode,
			&i.ManCode,
			&i.AddInfo,
			&i.Passwd,
			&i.Attempt,
			&i.DateLocked,
			&i.Remarks,
			&i.Picture,
			&i.IsLoggedIn,
			&i.AccountExpirationDt,
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

func (q *QueriesLocal) GetUsersList(ctx context.Context, userId string) (UsersListInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'UsersList' AND Uploaded = 0 and UserId = $1", getUsersList)
	log.Println(sql)
	row := q.db.QueryRowContext(ctx, sql, userId)
	return scanRowUsersList(row)
}

type ListUsersListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) UsersListCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, 
  UserID, AccessCode, LName, FName, MName, DateHired, BirthDay, 
  DateGiven, DateExpired, Address, Position, AreaCode, ManCode, AddInfo, 
  Attempt, DateLocked, Remarks, CASE WHEN isnull(IsLoggedIn,0) = 1 THEN 1 ELSE 0 END IsLoggedIn, 
  AccountExpirationDt
FROM OrgParms, UsersList d
`, filenamePath)
}

func (q *QueriesLocal) ListUsersList(ctx context.Context) ([]UsersListInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'UsersList' AND Uploaded = 0`,
		getUsersList)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsUsersList(rows)
}

// -- name: UpdateUsersList :one
const updateUsersList = `
UPDATE UsersList SET 
	AccessCode = $2,
	LName = $3,
	FName = $4,
	MName = $5,
	DateHired = $6,
	BirthDay = $7,
	DateGiven = $8,
	DateExpired = $9,
	Address = $10,
	Position = $11,
	AreaCode = $12,
	ManCode = $13,
	AddInfo = $14,
	Passwd = $15,
	Attempt = $16,
	DateLocked = $17,
	Remarks = $18,
	Picture = $19,
	isLoggedIn = $20,
	AccountExpirationDt = $21
WHERE UserId = $1`

func (q *QueriesLocal) UpdateUsersList(ctx context.Context, arg UsersListRequest) error {
	_, err := q.db.ExecContext(ctx, updateUsersList,
		arg.UserId,
		arg.AccessCode,
		arg.LName,
		arg.FName,
		arg.MName,
		arg.DateHired,
		arg.BirthDay,
		arg.DateGiven,
		arg.DateExpired,
		arg.Address,
		arg.Position,
		arg.AreaCode,
		arg.ManCode,
		arg.AddInfo,
		arg.Passwd,
		arg.Attempt,
		arg.DateLocked,
		arg.Remarks,
		arg.Picture,
		arg.IsLoggedIn,
		arg.AccountExpirationDt,
	)
	return err
}
