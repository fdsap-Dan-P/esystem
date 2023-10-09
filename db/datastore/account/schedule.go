package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
	"simplebank/util"

	"github.com/shopspring/decimal"
)

const createSchedule = `-- name: CreateSchedule: one
INSERT INTO Schedule (
  Account_Id, Series, Due_Date, Due_Prin, Due_Int, 
  End_Prin, End_Int, Carrying_Value, Realizable, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) 
ON CONFLICT(Account_Id, Series) DO UPDATE SET
Account_Id = excluded.Account_Id,
Series = excluded.Series,
Due_Date  = excluded.Due_Date ,
Due_Prin  = excluded.Due_Prin ,
Due_Int  = excluded.Due_Int ,
End_Prin  = excluded.End_Prin ,
End_Int  = excluded.End_Int ,
Carrying_Value  = excluded.Carrying_Value ,
Realizable = excluded.Realizable,
Other_Info = excluded.Other_Info

RETURNING Id, UUId, Account_Id, Series, Due_Date, Due_Prin, Due_Int, 
End_Prin, End_Int, Carrying_Value, Realizable, Other_Info
`

type ScheduleRequest struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	AccountId     int64           `json:"accountId"`
	Series        int16           `json:"series"`
	DueDate       time.Time       `json:"dueDate"`
	DuePrin       decimal.Decimal `json:"duePrin"`
	DueInt        decimal.Decimal `json:"dueInt"`
	EndPrin       decimal.Decimal `json:"endPrin"`
	EndInt        decimal.Decimal `json:"endInt"`
	CarryingValue decimal.Decimal `json:"carryingValue"`
	Realizable    decimal.Decimal `json:"realizable"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccount) CreateSchedule(ctx context.Context, arg ScheduleRequest) (model.Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
		arg.AccountId,
		arg.Series,
		arg.DueDate,
		arg.DuePrin,
		arg.DueInt,
		arg.EndPrin,
		arg.EndInt,
		arg.CarryingValue,
		arg.Realizable,
		arg.OtherInfo,
	)
	var i model.Schedule
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.DueDate,
		&i.DuePrin,
		&i.DueInt,
		&i.EndPrin,
		&i.EndInt,
		&i.CarryingValue,
		&i.Realizable,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSchedule = `-- name: DeleteSchedule :exec
DELETE FROM Schedule
WHERE id = $1
`

func (q *QueriesAccount) DeleteSchedule(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSchedule, id)
	return err
}

type ScheduleInfo struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	Acc           string          `json:"acc"`
	AccountId     int64           `json:"accountId"`
	Series        int16           `json:"series"`
	DueDate       time.Time       `json:"dueDate"`
	DuePrin       decimal.Decimal `json:"duePrin"`
	DueInt        decimal.Decimal `json:"dueInt"`
	EndPrin       decimal.Decimal `json:"endPrin"`
	EndInt        decimal.Decimal `json:"endInt"`
	CarryingValue decimal.Decimal `json:"carryingValue"`
	Realizable    decimal.Decimal `json:"realizable"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
	ModCtr        int64           `json:"modCtr"`
	Created       sql.NullTime    `json:"created"`
	Updated       sql.NullTime    `json:"updated"`
}

const getSchedule = `-- name: GetSchedule :one
SELECT 
  d.Id, mr.UUID, acc.Acc, d.Account_Id, d.Series, d.Due_Date, d.Due_Prin, d.Due_Int, 
  d.End_Prin, d.End_Int, d.Carrying_Value, d.Realizable, d.Other_Info, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Schedule d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Account acc on acc.ID = d.Account_Id
`

// WHERE id = $1 LIMIT 1
func populateSchedule(q *QueriesAccount, ctx context.Context,
	sql string, param ...interface{}) (map[string]map[int16]ScheduleInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := make(map[string]map[int16]ScheduleInfo)
	for rows.Next() {
		var i ScheduleInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Acc,
			&i.AccountId,
			&i.Series,
			&i.DueDate,
			&i.DuePrin,
			&i.DueInt,
			&i.EndPrin,
			&i.EndInt,
			&i.CarryingValue,
			&i.Realizable,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
			return items, err
		}
		if items[i.Acc] == nil {
			items[i.Acc] = make(map[int16]ScheduleInfo)
		}
		items[i.Acc][i.Series] = i
	}
	if err := rows.Close(); err != nil {
		return items, err
	}
	if err := rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}

// func (q *QueriesKPlus) SavingsList(ctx context.Context, arg SavingsListParams) ([]Savings, error) {
// 	var sql string
// 	if arg.Limit != 0 {
// 		sql = fmt.Sprintf("%s WHERE lower(trim(INAIIID)) = '%v' LIMIT %d OFFSET %d",
// 			kPLUSCustSavingsListSQL, arg.INAIIID, arg.Limit, arg.Offset)
// 	} else {
// 		sql = fmt.Sprintf("%s WHERE lower(trim(INAIIID)) = '%v' ", kPLUSCustSavingsListSQL, arg.INAIIID)
// 	}
// 	return populateKPLUSCustSavingsList(q, ctx, sql)
// }

func (q *QueriesAccount) GetSchedule(ctx context.Context, id int64) (ScheduleInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.id = $1 Order By acc.Acc, Series`, getSchedule)
	log.Printf("sql: %v", script)
	items, err := populateSchedule(q, ctx, script, id)
	log.Printf("----------- items: %v err:%v", items, err)
	for _, val := range items {
		for _, v := range val {
			return v, err
		}
	}
	return ScheduleInfo{}, fmt.Errorf("schedule ID:%v not found", id)
}

// sql.ErrNoRows.Error()
// const getSchedule = `-- name: GetSchedulebyUuid :one
// SELECT
//   Id, mr.UUId, Account_Id, Series, Due_Date, Due_Prin, Due_Int,
//   End_Prin, End_Int, Carrying_Value, Realizable, Other_Info,
//   mr.Mod_Ctr, mr.Created, mr.Updated
// FROM Schedule d INNER JOIN Main_Record mr on mr.UUId = d.UUId
// WHERE mr.UUID = $1 LIMIT 1
// `

func (q *QueriesAccount) GetSchedulebyUuid(ctx context.Context, uuid uuid.UUID) (ScheduleInfo, error) {
	sql := fmt.Sprintf(`%v WHERE mr.uuid = $1`, getSchedule)
	items, err := populateSchedule(q, ctx, sql, uuid)

	for _, val := range items {
		for _, v := range val {
			return v, err
		}
	}
	return ScheduleInfo{}, fmt.Errorf("schedule UUID:%v not found", uuid)
}

// type ListScheduleParams struct {
// 	AccountId int64 `json:"accountId"`
// 	Limit     int32 `json:"limit"`
// 	Offset    int32 `json:"offset"`
// }

func (q *QueriesAccount) GetSchedulebyAccId(ctx context.Context, accountId int64) (map[int16]ScheduleInfo, error) {
	sql := fmt.Sprintf(`%v WHERE Account_Id = $1 ORDER BY Series
	`, getSchedule)
	dat, err := populateSchedule(q, ctx, sql, accountId)
	for _, val := range dat {
		return val, err
	}
	return make(map[int16]ScheduleInfo), err
}

func (q *QueriesAccount) GetSchedulebyAcc(ctx context.Context, accList []string) (map[string]map[int16]ScheduleInfo, error) {
	accs := util.String2SqlList(accList)
	sql := fmt.Sprintf(
		`%v INNER JOIN Account a on a.ID = d.Account_Id 
		WHERE a.Acc in %s 
		ORDER BY acc.Acc, Series
	`, getSchedule, accs)
	return populateSchedule(q, ctx, sql)
}

const updateSchedule = `-- name: UpdateSchedule :one
UPDATE Schedule SET 
  Account_Id = $2,
  Series = $3,
  Due_Date = $4,
  Due_Prin = $5,
  Due_Int = $6,
  End_Prin = $7,
  End_Int = $8,
  Carrying_Value = $9,
  Realizable = $10,
  Other_Info = $11
WHERE id = $1
RETURNING 
  Id, UUId, Account_Id, Series, Due_Date, Due_Prin, Due_Int, 
  End_Prin, End_Int, Carrying_Value, Realizable, Other_Info
`

func (q *QueriesAccount) UpdateSchedule(ctx context.Context, arg ScheduleRequest) (model.Schedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedule,
		arg.Id,
		arg.AccountId,
		arg.Series,
		arg.DueDate,
		arg.DuePrin,
		arg.DueInt,
		arg.EndPrin,
		arg.EndInt,
		arg.CarryingValue,
		arg.Realizable,
		arg.OtherInfo,
	)
	var i model.Schedule
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.DueDate,
		&i.DuePrin,
		&i.DueInt,
		&i.EndPrin,
		&i.EndInt,
		&i.CarryingValue,
		&i.Realizable,
		&i.OtherInfo,
	)
	return i, err
}
