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

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE account
SET balance = balance + $1
WHERE id = $2
RETURNING 
  id, uuid, customer_id, acc, alternate_acc, account_name, 
  balance, non_current, contract_date, credit, debit, isbudget, 
  last_activity_date, open_date, passbook_line, pending_trn_amt, 
  principal, class_id, Account_Type_Id, budget_account_id, Category_ID,
  currency, office_id, referredby_id, status_Code, remarks, other_info
`

type AddAccountBalanceParams struct {
	Amount decimal.Decimal `json:"amount"`
	Id     int64           `json:"id"`
}

func (q *QueriesAccount) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (model.Account, error) {
	row_ := q.db.QueryRowContext(ctx, addAccountBalance, arg.Amount, arg.Id)
	var i model.Account
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.CustomerId,
		&i.Acc,
		&i.AlternateAcc,
		&i.AccountName,
		&i.Balance,
		&i.NonCurrent,
		&i.ContractDate,
		&i.Credit,
		&i.Debit,
		&i.Isbudget,
		&i.LastActivityDate,
		&i.OpenDate,
		&i.PassbookLine,
		&i.PendingTrnAmt,
		&i.Principal,
		&i.ClassId,
		&i.AccountTypeId,
		&i.BudgetAccountId,
		&i.CategoryId,
		&i.Currency,
		&i.OfficeId,
		&i.ReferredbyId,
		&i.StatusCode,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
    Customer_Id, Acc, Alternate_Acc, Account_Name, 
    Balance, Non_Current, Contract_Date, Credit, Debit, isBudget, 
    Last_Activity_Date, Open_Date, Passbook_Line, Pending_Trn_Amt, 
    Principal, Class_Id, Account_Type_Id, Budget_Account_Id, Category_ID,
    Currency, Office_Id , Referredby_Id, Status_Code, Remarks, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, 
$16, $17, $18, $19, $20, $21, $22, $23, $24, $25
) 
ON CONFLICT(Acc) DO UPDATE SET
  Customer_Id = EXCLUDED.Customer_Id,
  Alternate_Acc = EXCLUDED.Alternate_Acc,
  Account_Name = EXCLUDED.Account_Name,
  Balance = EXCLUDED.Balance,
  Non_Current = EXCLUDED.Non_Current,
  Contract_Date = EXCLUDED.Contract_Date,
  Credit = EXCLUDED.Credit,
  Debit = EXCLUDED.Debit,
  isBudget = EXCLUDED.isBudget,
  Last_Activity_Date = EXCLUDED.Last_Activity_Date,
  Open_Date = EXCLUDED.Open_Date,
  Passbook_Line = EXCLUDED.Passbook_Line,
  Pending_Trn_Amt = EXCLUDED.Pending_Trn_Amt,
  Principal = EXCLUDED.Principal,
  Class_Id = EXCLUDED.Class_Id,
  Account_Type_Id = EXCLUDED.Account_Type_Id,
  Budget_Account_Id = EXCLUDED.Budget_Account_Id,
  Category_ID = EXCLUDED.Category_ID,
  Currency = EXCLUDED.Currency,
  Office_Id = EXCLUDED.Office_Id,
  Referredby_Id = EXCLUDED.Referredby_Id,
  Status_Code = EXCLUDED.Status_Code,
  Remarks = EXCLUDED.Remarks,
  Other_Info = EXCLUDED.Other_Info
RETURNING 
  id, uuid, customer_id, acc, alternate_acc, account_name, 
  balance, non_current, contract_date, credit, debit, isbudget, 
  last_activity_date, open_date, passbook_line, pending_trn_amt, 
  principal, class_id, Account_Type_Id, budget_account_id, Category_Id,
  currency, office_id, referredby_id, status_Code, remarks, other_info
`

type AccountRequest struct {
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

func (q *QueriesAccount) CreateAccount(ctx context.Context, arg AccountRequest) (model.Account, error) {
	row_ := q.db.QueryRowContext(ctx, createAccount,
		arg.CustomerId,
		arg.Acc,
		arg.AlternateAcc,
		arg.AccountName,
		arg.Balance,
		arg.NonCurrent,
		arg.ContractDate,
		arg.Credit,
		arg.Debit,
		arg.Isbudget,
		arg.LastActivityDate,
		arg.OpenDate,
		arg.PassbookLine,
		arg.PendingTrnAmt,
		arg.Principal,
		arg.ClassId,
		arg.AccountTypeId,
		arg.BudgetAccountId,
		arg.CategoryId,
		arg.Currency,
		arg.OfficeId,
		arg.ReferredbyId,
		arg.StatusCode,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Account
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.CustomerId,
		&i.Acc,
		&i.AlternateAcc,
		&i.AccountName,
		&i.Balance,
		&i.NonCurrent,
		&i.ContractDate,
		&i.Credit,
		&i.Debit,
		&i.Isbudget,
		&i.LastActivityDate,
		&i.OpenDate,
		&i.PassbookLine,
		&i.PendingTrnAmt,
		&i.Principal,
		&i.ClassId,
		&i.AccountTypeId,
		&i.BudgetAccountId,
		&i.CategoryId,
		&i.Currency,
		&i.OfficeId,
		&i.ReferredbyId,
		&i.StatusCode,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1
`

func (q *QueriesAccount) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

type AccountInfo struct {
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
	ModCtr           int64           `json:"modCtr"`
	Created          sql.NullTime    `json:"created"`
	Updated          sql.NullTime    `json:"updated"`
}

type AccountStat struct {
	AccountId     int64           `json:"accountId"`
	Uuid          uuid.UUID       `json:"uuid"`
	CustomerId    int64           `json:"customerId"`
	Acc           string          `json:"acc"`
	AlternateAcc  sql.NullString  `json:"alternateAcc"`
	AccountTypeId int64           `json:"accountTypeId"`
	AccountType   string          `json:"accountType"`
	Principal     decimal.Decimal `json:"principal"`
	Interest      decimal.Decimal `json:"interest"`
	BalPrin       decimal.Decimal `json:"balPrin"`
	BalInt        decimal.Decimal `json:"balInt"`
	Waivable      decimal.Decimal `json:"waivable"`
	StatusCode    int64           `json:"statusCode"`
	Status        string          `json:"status"`
	Frequency     sql.NullInt64   `json:"frequency"`
	N             sql.NullInt64   `json:"n"`
	PaidN         sql.NullInt64   `json:"paidN"`
	DateStart     sql.NullTime    `json:"dateStart"`
	Maturity      sql.NullTime    `json:"maturity"`
	MeetingDate   time.Time       `json:"meetingDate"`
}

const getAccount = `-- name: GetAccount :one
SELECT 
  d.id, mr.uuid, customer_id, acc, alternate_acc, account_name, balance, non_current, 
  contract_date, credit, debit, isbudget, last_activity_date, open_date, 
  passbook_line, pending_trn_amt, principal, class_id, 
  Account_Type_Id, budget_account_id, Category_Id, currency, office_id, 
  referredby_id, status_Code, remarks,
  d.Other_Info, mr.Mod_Ctr, mr.Created, mr.Updated
FROM account d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

const getAccountStat = `-- name: GetAccount :one
SELECT 
  acc.Id Account_Id, acc.UUID, acc.customer_id, acc.acc, acc.alternate_acc, acc.account_type_id, 
  typ.account_type, acc.Principal, ai.Interest, 
  acc.principal - acc.credit + acc.debit balprin,
  COALESCE(ai.interest,0) - COALESCE(ai.credit,0) + COALESCE(ai.debit,0) balint,
  stat.Title Status, tr.Frequency, n, paid_n, Date_Start, Maturity,
  sc.system_date - (date_part('dow'::text, sc.system_date) - grp.meeting_day::double precision - 1)::integer AS meetingdate
FROM Account acc
INNER JOIN Account_Type typ on acc.account_type_id = typ.id
INNER JOIN Reference stat on acc.status_code = stat.code and lower(stat.Ref_Type) = 'accountstatus'
INNER JOIN Customer cus on cus.Id = acc.Customer_Id
INNER JOIN Customer_Group grp on cus.customer_group_id = grp.id
INNER JOIN Office unit on grp.office_id = unit.id
INNER JOIN Office ar on unit.parent_id = ar.id
INNER JOIN system_config sc on ar.id = sc.office_id 
LEFT JOIN Account_Interest ai on acc.id = ai.Account_Id
LEFT JOIN Account_Term tr on acc.id = tr.Account_Id
`

func populateAccountStat(
	q *QueriesAccount, ctx context.Context,
	sql string, param ...interface{}) (map[string]AccountStat, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make(map[string]AccountStat)
	for rows.Next() {
		var i AccountStat
		err := rows.Scan(
			&i.AccountId,
			&i.Uuid,
			&i.CustomerId,
			&i.Acc,
			&i.AlternateAcc,
			&i.AccountTypeId,
			&i.AccountType,
			&i.Principal,
			&i.Interest,
			&i.BalPrin,
			&i.BalInt,
			&i.Status,
			&i.Frequency,
			&i.N,
			&i.PaidN,
			&i.DateStart,
			&i.Maturity,
			&i.MeetingDate,
		)
		if err != nil {
			return items, err
		}
		items[i.Acc] = i
	}
	if err := rows.Close(); err != nil {
		return items, err
	}
	if err := rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}

func (q *QueriesAccount) GetAccountStat(
	ctx context.Context, accList []string) (map[string]AccountStat, error) {
	accs := util.String2SqlList(accList)
	sql := fmt.Sprintf(`%v WHERE acc.Acc in %s `, getAccountStat, accs)
	return populateAccountStat(q, ctx, sql)
}

func populateAccount(q *QueriesAccount, ctx context.Context, sql string, param ...interface{}) ([]AccountInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountInfo{}
	for rows.Next() {
		var i AccountInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CustomerId,
			&i.Acc,
			&i.AlternateAcc,
			&i.AccountName,
			&i.Balance,
			&i.NonCurrent,
			&i.ContractDate,
			&i.Credit,
			&i.Debit,
			&i.Isbudget,
			&i.LastActivityDate,
			&i.OpenDate,
			&i.PassbookLine,
			&i.PendingTrnAmt,
			&i.Principal,
			&i.ClassId,
			&i.AccountTypeId,
			&i.BudgetAccountId,
			&i.CategoryId,
			&i.Currency,
			&i.OfficeId,
			&i.ReferredbyId,
			&i.StatusCode,
			&i.Remarks,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
			return items, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return items, err
	}
	if err := rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}

func (q *QueriesAccount) GetAccount(ctx context.Context, id int64) (AccountInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.id = $1`, getAccount)
	log.Printf("sql: %v", script)
	items, err := populateAccount(q, ctx, script, id)
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return AccountInfo{}, fmt.Errorf("account ID:%v not found", id)
	}
}

func (q *QueriesAccount) GetAccountbyAcc(ctx context.Context, accList []string) ([]AccountInfo, error) {
	accs := util.String2SqlList(accList)
	sql := fmt.Sprintf(`%v WHERE d.Acc in %s `, getAccount, accs)
	return populateAccount(q, ctx, sql)
}

func (q *QueriesAccount) GetAccountbyAltAcc(ctx context.Context, altAcc string) (AccountInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.alternate_acc = $1`, getAccount)
	log.Printf("sql: %v", script)
	items, err := populateAccount(q, ctx, script, altAcc)
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return AccountInfo{}, fmt.Errorf("account ID:%v not found", altAcc)
	}
}

func (q *QueriesAccount) GetAccountbyUuid(ctx context.Context, uuid uuid.UUID) (AccountInfo, error) {
	script := fmt.Sprintf(`%v WHERE mr.UUID = $1`, getAccount)
	log.Printf("sql: %v", script)
	items, err := populateAccount(q, ctx, script, uuid)
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return AccountInfo{}, fmt.Errorf("account UUID:%v not found", uuid)
	}
}

func (q *QueriesAccount) ListAccount(ctx context.Context, arg ListAccountParams) ([]AccountInfo, error) {
	script := fmt.Sprintf(
		`%v WHERE Customer_Id = $1
			ORDER BY id LIMIT $2 OFFSET $3`, getAccount)
	log.Printf("sql: %v", script)
	items, err := populateAccount(q, ctx, script, arg.CustomerId, arg.Limit, arg.Offset)
	log.Printf("items: %v", items)
	if len(items) > 0 {
		return items, err
	} else {
		return []AccountInfo{}, fmt.Errorf("account Customer_ID:%v not found", arg)
	}
}

type ListAccountParams struct {
	CustomerId int64 `json:"customerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET 
  Customer_Id        = $2,
  Acc                = $3,
  Alternate_Acc      = $4,
  Account_Name       = $5,
  Balance            = $6,
  Non_Current        = $7,
  Contract_Date      = $8,
  Credit             = $9,
  Debit              = $10,
  isBudget           = $11,
  Last_Activity_Date = $12,
  Open_Date          = $13,
  Passbook_Line      = $14,
  Pending_Trn_Amt    = $15,
  Principal          = $16,
  Class_Id           = $17,
  Account_Type_Id    = $18,
  Budget_Account_Id  = $19,
  Category_Id        = $20,
  Currency           = $21,
  Office_Id          = $22,
  Referredby_Id      = $23,
  Status_Code          = $24,
  Remarks            = $25,
  Other_Info         = $26
WHERE id = $1
RETURNING 
  id, uuid, customer_id, acc, alternate_acc, account_name, balance, non_current, 
  contract_date, credit, debit, isbudget, last_activity_date, open_date, 
  passbook_line, pending_trn_amt, principal, class_id, Account_Type_Id,
   budget_account_id, Category_Id, currency, office_id, referredby_id, status_Code, remarks, other_info
`

func (q *QueriesAccount) UpdateAccount(ctx context.Context, arg AccountRequest) (model.Account, error) {
	row_ := q.db.QueryRowContext(ctx, updateAccount,
		arg.Id,
		arg.CustomerId,
		arg.Acc,
		arg.AlternateAcc,
		arg.AccountName,
		arg.Balance,
		arg.NonCurrent,
		arg.ContractDate,
		arg.Credit,
		arg.Debit,
		arg.Isbudget,
		arg.LastActivityDate,
		arg.OpenDate,
		arg.PassbookLine,
		arg.PendingTrnAmt,
		arg.Principal,
		arg.ClassId,
		arg.AccountTypeId,
		arg.BudgetAccountId,
		arg.CategoryId,
		arg.Currency,
		arg.OfficeId,
		arg.ReferredbyId,
		arg.StatusCode,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Account
	err := row_.Scan(
		&i.Id,
		&i.Uuid,
		&i.CustomerId,
		&i.Acc,
		&i.AlternateAcc,
		&i.AccountName,
		&i.Balance,
		&i.NonCurrent,
		&i.ContractDate,
		&i.Credit,
		&i.Debit,
		&i.Isbudget,
		&i.LastActivityDate,
		&i.OpenDate,
		&i.PassbookLine,
		&i.PendingTrnAmt,
		&i.Principal,
		&i.ClassId,
		&i.AccountTypeId,
		&i.BudgetAccountId,
		&i.CategoryId,
		&i.Currency,
		&i.OfficeId,
		&i.ReferredbyId,
		&i.StatusCode,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

func AccountModel2Info(m model.Account) AccountInfo {
	return AccountInfo{
		Id:               m.Id,
		Uuid:             m.Uuid,
		CustomerId:       m.CustomerId,
		Acc:              m.Acc,
		AlternateAcc:     m.AlternateAcc,
		AccountName:      m.AccountName,
		Balance:          m.Balance,
		NonCurrent:       m.NonCurrent,
		ContractDate:     m.ContractDate,
		Credit:           m.Credit,
		Debit:            m.Debit,
		Isbudget:         m.Isbudget,
		LastActivityDate: m.LastActivityDate,
		OpenDate:         m.OpenDate,
		PassbookLine:     m.PassbookLine,
		PendingTrnAmt:    m.PendingTrnAmt,
		Principal:        m.Principal,
		ClassId:          m.ClassId,
		AccountTypeId:    m.AccountTypeId,
		BudgetAccountId:  m.BudgetAccountId,
		Currency:         m.Currency,
		OfficeId:         m.OfficeId,
		ReferredbyId:     m.ReferredbyId,
		StatusCode:       m.StatusCode,
		Remarks:          m.Remarks,
		OtherInfo:        m.OtherInfo,
	}
}

func AccountInfo2Model(m AccountInfo) model.Account {
	return model.Account{
		Id:               m.Id,
		Uuid:             m.Uuid,
		CustomerId:       m.CustomerId,
		Acc:              m.Acc,
		AlternateAcc:     m.AlternateAcc,
		AccountName:      m.AccountName,
		Balance:          m.Balance,
		NonCurrent:       m.NonCurrent,
		ContractDate:     m.ContractDate,
		Credit:           m.Credit,
		Debit:            m.Debit,
		Isbudget:         m.Isbudget,
		LastActivityDate: m.LastActivityDate,
		OpenDate:         m.OpenDate,
		PassbookLine:     m.PassbookLine,
		PendingTrnAmt:    m.PendingTrnAmt,
		Principal:        m.Principal,
		ClassId:          m.ClassId,
		AccountTypeId:    m.AccountTypeId,
		BudgetAccountId:  m.BudgetAccountId,
		CategoryId:       m.CategoryId,
		Currency:         m.Currency,
		OfficeId:         m.OfficeId,
		ReferredbyId:     m.ReferredbyId,
		StatusCode:       m.StatusCode,
		Remarks:          m.Remarks,
		OtherInfo:        m.OtherInfo,
	}
}

func AccountModel2Request(m model.Account) AccountRequest {
	return AccountRequest{
		Id:               m.Id,
		Uuid:             m.Uuid,
		CustomerId:       m.CustomerId,
		Acc:              m.Acc,
		AlternateAcc:     m.AlternateAcc,
		AccountName:      m.AccountName,
		Balance:          m.Balance,
		NonCurrent:       m.NonCurrent,
		ContractDate:     m.ContractDate,
		Credit:           m.Credit,
		Debit:            m.Debit,
		Isbudget:         m.Isbudget,
		LastActivityDate: m.LastActivityDate,
		OpenDate:         m.OpenDate,
		PassbookLine:     m.PassbookLine,
		PendingTrnAmt:    m.PendingTrnAmt,
		Principal:        m.Principal,
		ClassId:          m.ClassId,
		AccountTypeId:    m.AccountTypeId,
		BudgetAccountId:  m.BudgetAccountId,
		CategoryId:       m.CategoryId,
		Currency:         m.Currency,
		OfficeId:         m.OfficeId,
		ReferredbyId:     m.ReferredbyId,
		StatusCode:       m.StatusCode,
		Remarks:          m.Remarks,
		OtherInfo:        m.OtherInfo,
	}
}
