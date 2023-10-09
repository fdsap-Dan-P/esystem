package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	model "simplebank/db/datastore/esystemlocal"
)

const createUsersList = `-- name: CreateUsersList: one
  INSERT INTO esystemdump.UsersList (
    ModCtr, BrCode, ModAction, UserId, AccessCode, LName, FName, MName, DateHired, BirthDay, 
    DateGiven, DateExpired, Address, Position, AreaCode, ManCode, AddInfo, 
    Passwd, Attempt, DateLocked, Remarks, Picture, IsLoggedIn, AccountExpirationDt
) 
VALUES 
 ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
  $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
  $21, $22, $23, $24) 
ON CONFLICT (brCode, UserId, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	AccessCode =  EXCLUDED.AccessCode,
	LName =  EXCLUDED.LName,
	FName =  EXCLUDED.FName,
	MName =  EXCLUDED.MName,
	DateHired =  EXCLUDED.DateHired,
	BirthDay =  EXCLUDED.BirthDay,
	DateGiven =  EXCLUDED.DateGiven,
	DateExpired =  EXCLUDED.DateExpired,
	Address =  EXCLUDED.Address,
	Position =  EXCLUDED.Position,
	AreaCode =  EXCLUDED.AreaCode,
	ManCode =  EXCLUDED.ManCode,
	AddInfo =  EXCLUDED.AddInfo,
	Passwd =  EXCLUDED.Passwd,
	Attempt =  EXCLUDED.Attempt,
	DateLocked =  EXCLUDED.DateLocked,
	Remarks =  EXCLUDED.Remarks,
	Picture =  EXCLUDED.Picture,
	IsLoggedIn =  EXCLUDED.IsLoggedIn,
	AccountExpirationDt =  EXCLUDED.AccountExpirationDt
  `

func (q *QueriesDump) CreateUsersList(ctx context.Context, arg model.UsersList) error {
	log.Println(arg)
	_, err := q.db.ExecContext(ctx, createUsersList,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
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
DELETE FROM esystemdump.UsersList WHERE UserId = $1
`

func (q *QueriesDump) DeleteUsersList(ctx context.Context, brCode string, UserId string) error {
	_, err := q.db.ExecContext(ctx, deleteUsersList, UserId)
	return err
}

// -- name: GetUsersList :one
const getUsersList = `
SELECT 
  ModCtr, BrCode, ModAction,
  UserId, AccessCode, LName, FName, MName, DateHired, BirthDay, 
  DateGiven, DateExpired, Address, Position, AreaCode, ManCode, AddInfo, 
  Passwd, Attempt, DateLocked, Remarks, Picture, IsLoggedIn, AccountExpirationDt
FROM esystemdump.UsersList d
`

func scanRowUsersList(row *sql.Row) (model.UsersList, error) {
	var i model.UsersList
	err := row.Scan(
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
	)
	return i, err
}

func scanRowsUsersList(rows *sql.Rows) ([]model.UsersList, error) {
	items := []model.UsersList{}
	for rows.Next() {
		var i model.UsersList
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

func (q *QueriesDump) GetUsersList(ctx context.Context, brCode string, userId string) (model.UsersList, error) {
	sql := fmt.Sprintf("%s WHERE UserId = $1", getUsersList)
	log.Println(sql)
	row := q.db.QueryRowContext(ctx, sql, userId)
	return scanRowUsersList(row)
}

type ListUsersListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListUsersList(ctx context.Context, lastModCtr int64) ([]model.UsersList, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getUsersList)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsUsersList(rows)
}

// -- name: UpdateUsersList :one
const updateUsersList = `
UPDATE esystemdump.UsersList SET 
	ModCtr = $1,
	AccessCode = $5,
	LName = $6,
	FName = $7,
	MName = $8,
	DateHired = $9,
	BirthDay = $10,
	DateGiven = $11,
	DateExpired = $12,
	Address = $13,
	Position = $14,
	AreaCode = $15,
	ManCode = $16,
	AddInfo = $17,
	Passwd = $18,
	Attempt = $19,
	DateLocked = $20,
	Remarks = $21,
	Picture = $22,
	isLoggedIn = $23,
	AccountExpirationDt = $24
WHERE BrCode = $2 and UserId = $4 and ModAction = $3 `

func (q *QueriesDump) UpdateUsersList(ctx context.Context, arg model.UsersList) error {
	_, err := q.db.ExecContext(ctx, updateUsersList,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
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
