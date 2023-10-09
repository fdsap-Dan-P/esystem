package db

import (
	"context"
	"database/sql"
	"fmt"
	dsRef "simplebank/db/datastore/reference"
	"simplebank/model"

	"github.com/google/uuid"
)

type QuerierAccount interface {
	CreateAccount(ctx context.Context, arg AccountRequest) (model.Account, error)
	GetAccount(ctx context.Context, id int64) (AccountInfo, error)
	GetAccountbyAcc(ctx context.Context, acc []string) ([]AccountInfo, error)
	GetAccountbyAltAcc(ctx context.Context, altAcc string) (AccountInfo, error)
	GetAccountStat(ctx context.Context, acc []string) (map[string]AccountStat, error)
	GetAccountbyUuid(ctx context.Context, uuid uuid.UUID) (AccountInfo, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]AccountInfo, error)
	UpdateAccount(ctx context.Context, arg AccountRequest) (model.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (model.Account, error)
	// AccountModel2Info(m model.Account) AccountInfo
	// AccountInfo2Model(m AccountInfo) model.Account
	// AccountModel2Request(m model.Account) AccountRequest

	CreateProduct(ctx context.Context, arg ProductRequest) (model.Product, error)
	GetProduct(ctx context.Context, id int64) (ProductInfo, error)
	GetProductbyUuid(ctx context.Context, uuid uuid.UUID) (ProductInfo, error)
	GetProductbyName(ctx context.Context, name string) (ProductInfo, error)
	ListProduct(ctx context.Context, arg ListProductParams) ([]ProductInfo, error)
	UpdateProduct(ctx context.Context, arg ProductRequest) (model.Product, error)
	DeleteProduct(ctx context.Context, id int64) error

	CreateAccountType(ctx context.Context, arg AccountTypeRequest) (model.AccountType, error)
	GetAccountType(ctx context.Context, id int64) (AccountTypeInfo, error)
	GetAccountTypebyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeInfo, error)
	GetAccountTypebyName(ctx context.Context, name string) (AccountTypeInfo, error)
	ListAccountType(ctx context.Context, arg ListAccountTypeParams) ([]AccountTypeInfo, error)
	UpdateAccountType(ctx context.Context, arg AccountTypeRequest) (model.AccountType, error)
	DeleteAccountType(ctx context.Context, id int64) error

	CreateAccountTypeFilter(ctx context.Context, arg AccountTypeFilterRequest) (model.AccountTypeFilter, error)
	GetAccountTypeFilter(ctx context.Context, officeId int64, acctType int64) (AccountTypeFilterInfo, error)
	GetAccountTypeFilterbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeFilterInfo, error)
	ListAccountTypeFilter(ctx context.Context, arg ListAccountTypeFilterParams) ([]AccountTypeFilterInfo, error)
	UpdateAccountTypeFilter(ctx context.Context, arg AccountTypeFilterRequest) (model.AccountTypeFilter, error)
	DeleteAccountTypeFilter(ctx context.Context, uuid uuid.UUID) error

	CreateAccountTypeGroup(ctx context.Context, arg AccountTypeGroupRequest) (model.AccountTypeGroup, error)
	GetAccountTypeGroup(ctx context.Context, id int64) (AccountTypeGroupInfo, error)
	GetAccountTypeGroupbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeGroupInfo, error)
	ListAccountTypeGroup(ctx context.Context, arg ListAccountTypeGroupParams) ([]AccountTypeGroupInfo, error)
	UpdateAccountTypeGroup(ctx context.Context, arg AccountTypeGroupRequest) (model.AccountTypeGroup, error)
	DeleteAccountTypeGroup(ctx context.Context, uuid uuid.UUID) error

	CreateOfficeAccountType(ctx context.Context, arg OfficeAccountTypeRequest) (model.OfficeAccountType, error)
	GetOfficeAccountType(ctx context.Context, id int64) (OfficeAccountTypeInfo, error)
	GetOfficeAccountTypebyName(ctx context.Context, name string) (OfficeAccountTypeInfo, error)
	GetOfficeAccountTypebyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountTypeInfo, error)
	ListOfficeAccountType(ctx context.Context, arg ListOfficeAccountTypeParams) ([]OfficeAccountTypeInfo, error)
	UpdateOfficeAccountType(ctx context.Context, arg OfficeAccountTypeRequest) (model.OfficeAccountType, error)
	DeleteOfficeAccountType(ctx context.Context, id int64) error

	CreateAccountClass(ctx context.Context, arg AccountClassRequest) (model.AccountClass, error)
	GetAccountClass(ctx context.Context, id int64) (AccountClassInfo, error)
	GetAccountClassbyUuid(ctx context.Context, uuid uuid.UUID) (AccountClassInfo, error)
	GetAccountClassbyKeys(ctx context.Context, productID int64, groupID int64, classID int64) (AccountClassInfo, error)
	ListAccountClass(ctx context.Context, arg ListAccountClassParams) ([]AccountClassInfo, error)
	UpdateAccountClass(ctx context.Context, arg AccountClassRequest) (model.AccountClass, error)
	DeleteAccountClass(ctx context.Context, id int64) error

	CreateChargeType(ctx context.Context, arg ChargeTypeRequest) (model.ChargeType, error)
	GetChargeType(ctx context.Context, id int64) (ChargeTypeInfo, error)
	GetChargeTypebyUuid(ctx context.Context, uuid uuid.UUID) (ChargeTypeInfo, error)
	GetChargeTypebyName(ctx context.Context, name string) (ChargeTypeInfo, error)
	ListChargeType(ctx context.Context, arg ListChargeTypeParams) ([]ChargeTypeInfo, error)
	UpdateChargeType(ctx context.Context, arg ChargeTypeRequest) (model.ChargeType, error)
	DeleteChargeType(ctx context.Context, id int64) error

	CreateAccountInterest(ctx context.Context, arg AccountInterestRequest) (model.AccountInterest, error)
	GetAccountInterest(ctx context.Context, id int64) (AccountInterestInfo, error)
	GetAccountInterestbyUuid(ctx context.Context, uuid uuid.UUID) (AccountInterestInfo, error)
	ListAccountInterest(ctx context.Context, arg ListAccountInterestParams) ([]AccountInterestInfo, error)
	UpdateAccountInterest(ctx context.Context, arg AccountInterestRequest) (model.AccountInterest, error)
	DeleteAccountInterest(ctx context.Context, id int64) error

	CreateAccountTerm(ctx context.Context, arg AccountTermRequest) (model.AccountTerm, error)
	GetAccountTerm(ctx context.Context, id int64) (AccountTermInfo, error)
	GetAccountTermbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTermInfo, error)
	ListAccountTerm(ctx context.Context, arg ListAccountTermParams) ([]AccountTermInfo, error)
	UpdateAccountTerm(ctx context.Context, arg AccountTermRequest) (model.AccountTerm, error)
	DeleteAccountTerm(ctx context.Context, id int64) error

	CreateAccountInventory(ctx context.Context, arg AccountInventoryRequest) (model.AccountInventory, error)
	GetAccountInventory(ctx context.Context, id int64) (AccountInventoryInfo, error)
	GetAccountInventorybyUuid(ctx context.Context, uuid uuid.UUID) (AccountInventoryInfo, error)
	ListAccountInventory(ctx context.Context, arg ListAccountInventoryParams) ([]AccountInventoryInfo, error)
	UpdateAccountInventory(ctx context.Context, arg AccountInventoryRequest) (model.AccountInventory, error)
	DeleteAccountInventory(ctx context.Context, id int64) error

	CreateInventoryDetail(ctx context.Context, arg InventoryDetailRequest) (model.InventoryDetail, error)
	GetInventoryDetail(ctx context.Context, id int64) (InventoryDetailInfo, error)
	GetInventoryDetailbyUuid(ctx context.Context, uuid uuid.UUID) (InventoryDetailInfo, error)
	ListInventoryDetail(ctx context.Context, arg ListInventoryDetailParams) ([]InventoryDetailInfo, error)
	UpdateInventoryDetail(ctx context.Context, arg InventoryDetailRequest) (model.InventoryDetail, error)
	DeleteInventoryDetail(ctx context.Context, id int64) error

	CreateInventoryItem(ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error)
	GetInventoryItem(ctx context.Context, id int64) (InventoryItemInfo, error)
	GetInventoryItembyUuid(ctx context.Context, uuid uuid.UUID) (InventoryItemInfo, error)
	ListInventoryItembyBrand(ctx context.Context, arg ListInventoryItembyBrandParams) ([]InventoryItemInfo, error)
	ListInventoryItembyGeneric(ctx context.Context, arg ListInventoryItembyGenericParams) ([]InventoryItemInfo, error)
	UpdateInventoryItem(ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error)
	DeleteInventoryItem(ctx context.Context, id int64) error
	InventoryItemFilter(ctx context.Context, arg InventoryItemFilterParams) ([]InventoryItemInfo, error)
	InventoryItemSearch(ctx context.Context, arg InventoryItemSearchParams) ([]InventoryItemInfo, error)

	CreateAccountSpecsDate(ctx context.Context, arg AccountSpecsDateRequest) (model.AccountSpecsDate, error)
	GetAccountSpecsDate(ctx context.Context, AccountItemId int64, specsId int64) (AccountSpecsDateInfo, error)
	GetAccountSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsDateInfo, error)
	ListAccountSpecsDate(ctx context.Context, arg ListAccountSpecsDateParams) ([]AccountSpecsDateInfo, error)
	UpdateAccountSpecsDate(ctx context.Context, arg AccountSpecsDateRequest) (model.AccountSpecsDate, error)
	DeleteAccountSpecsDate(ctx context.Context, uuid uuid.UUID) error

	CreateAccountSpecsString(ctx context.Context, arg AccountSpecsStringRequest) (model.AccountSpecsString, error)
	GetAccountSpecsString(ctx context.Context, AccountItemId int64, specsId int64) (AccountSpecsStringInfo, error)
	GetAccountSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsStringInfo, error)
	ListAccountSpecsString(ctx context.Context, arg ListAccountSpecsStringParams) ([]AccountSpecsStringInfo, error)
	UpdateAccountSpecsString(ctx context.Context, arg AccountSpecsStringRequest) (model.AccountSpecsString, error)
	DeleteAccountSpecsString(ctx context.Context, uuid uuid.UUID) error

	CreateAccountSpecsNumber(ctx context.Context, arg AccountSpecsNumberRequest) (model.AccountSpecsNumber, error)
	GetAccountSpecsNumber(ctx context.Context, AccountItemId int64, specsId int64) (AccountSpecsNumberInfo, error)
	GetAccountSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsNumberInfo, error)
	ListAccountSpecsNumber(ctx context.Context, arg ListAccountSpecsNumberParams) ([]AccountSpecsNumberInfo, error)
	UpdateAccountSpecsNumber(ctx context.Context, arg AccountSpecsNumberRequest) (model.AccountSpecsNumber, error)
	DeleteAccountSpecsNumber(ctx context.Context, uuid uuid.UUID) error

	CreateAccountSpecsRef(ctx context.Context, arg AccountSpecsRefRequest) (model.AccountSpecsRef, error)
	GetAccountSpecsRef(ctx context.Context, AccountItemId int64, specsId int64) (AccountSpecsRefInfo, error)
	GetAccountSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsRefInfo, error)
	ListAccountSpecsRef(ctx context.Context, arg ListAccountSpecsRefParams) ([]AccountSpecsRefInfo, error)
	UpdateAccountSpecsRef(ctx context.Context, arg AccountSpecsRefRequest) (model.AccountSpecsRef, error)
	DeleteAccountSpecsRef(ctx context.Context, uuid uuid.UUID) error

	CreateInventorySpecsDate(ctx context.Context, arg InventorySpecsDateRequest) (model.InventorySpecsDate, error)
	GetInventorySpecsDate(ctx context.Context, InventoryItemId int64, specsId int64) (InventorySpecsDateInfo, error)
	GetInventorySpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsDateInfo, error)
	ListInventorySpecsDate(ctx context.Context, arg ListInventorySpecsDateParams) ([]InventorySpecsDateInfo, error)
	UpdateInventorySpecsDate(ctx context.Context, arg InventorySpecsDateRequest) (model.InventorySpecsDate, error)
	DeleteInventorySpecsDate(ctx context.Context, uuid uuid.UUID) error

	CreateInventorySpecsString(ctx context.Context, arg InventorySpecsStringRequest) (model.InventorySpecsString, error)
	GetInventorySpecsString(ctx context.Context, InventoryItemId int64, specsId int64) (InventorySpecsStringInfo, error)
	GetInventorySpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsStringInfo, error)
	ListInventorySpecsString(ctx context.Context, arg ListInventorySpecsStringParams) ([]InventorySpecsStringInfo, error)
	UpdateInventorySpecsString(ctx context.Context, arg InventorySpecsStringRequest) (model.InventorySpecsString, error)
	DeleteInventorySpecsString(ctx context.Context, uuid uuid.UUID) error

	CreateInventorySpecsNumber(ctx context.Context, arg InventorySpecsNumberRequest) (model.InventorySpecsNumber, error)
	GetInventorySpecsNumber(ctx context.Context, InventoryItemId int64, specsId int64) (InventorySpecsNumberInfo, error)
	GetInventorySpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsNumberInfo, error)
	ListInventorySpecsNumber(ctx context.Context, arg ListInventorySpecsNumberParams) ([]InventorySpecsNumberInfo, error)
	UpdateInventorySpecsNumber(ctx context.Context, arg InventorySpecsNumberRequest) (model.InventorySpecsNumber, error)
	DeleteInventorySpecsNumber(ctx context.Context, uuid uuid.UUID) error

	CreateInventorySpecsRef(ctx context.Context, arg InventorySpecsRefRequest) (model.InventorySpecsRef, error)
	GetInventorySpecsRef(ctx context.Context, InventoryItemId int64, specsId int64) (InventorySpecsRefInfo, error)
	GetInventorySpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsRefInfo, error)
	ListInventorySpecsRef(ctx context.Context, arg ListInventorySpecsRefParams) ([]InventorySpecsRefInfo, error)
	UpdateInventorySpecsRef(ctx context.Context, arg InventorySpecsRefRequest) (model.InventorySpecsRef, error)
	DeleteInventorySpecsRef(ctx context.Context, uuid uuid.UUID) error

	CreateAccountParamDate(ctx context.Context, arg AccountParamDateRequest) (model.AccountParamDate, error)
	GetAccountParamDate(ctx context.Context, InventoryItemId int64, specsId int64) (AccountParamDateInfo, error)
	GetAccountParamDatebyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamDateInfo, error)
	ListAccountParamDate(ctx context.Context, arg ListAccountParamDateParams) ([]AccountParamDateInfo, error)
	UpdateAccountParamDate(ctx context.Context, arg AccountParamDateRequest) (model.AccountParamDate, error)
	DeleteAccountParamDate(ctx context.Context, uuid uuid.UUID) error

	CreateAccountParamString(ctx context.Context, arg AccountParamStringRequest) (model.AccountParamString, error)
	GetAccountParamString(ctx context.Context, InventoryItemId int64, specsId int64) (AccountParamStringInfo, error)
	GetAccountParamStringbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamStringInfo, error)
	ListAccountParamString(ctx context.Context, arg ListAccountParamStringParams) ([]AccountParamStringInfo, error)
	UpdateAccountParamString(ctx context.Context, arg AccountParamStringRequest) (model.AccountParamString, error)
	DeleteAccountParamString(ctx context.Context, uuid uuid.UUID) error

	CreateAccountParamNumber(ctx context.Context, arg AccountParamNumberRequest) (model.AccountParamNumber, error)
	GetAccountParamNumber(ctx context.Context, InventoryItemId int64, specsId int64) (AccountParamNumberInfo, error)
	GetAccountParamNumberbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamNumberInfo, error)
	ListAccountParamNumber(ctx context.Context, arg ListAccountParamNumberParams) ([]AccountParamNumberInfo, error)
	UpdateAccountParamNumber(ctx context.Context, arg AccountParamNumberRequest) (model.AccountParamNumber, error)
	DeleteAccountParamNumber(ctx context.Context, uuid uuid.UUID) error

	CreateAccountParamRef(ctx context.Context, arg AccountParamRefRequest) (model.AccountParamRef, error)
	GetAccountParamRef(ctx context.Context, InventoryItemId int64, specsId int64) (AccountParamRefInfo, error)
	GetAccountParamRefbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamRefInfo, error)
	ListAccountParamRef(ctx context.Context, arg ListAccountParamRefParams) ([]AccountParamRefInfo, error)
	UpdateAccountParamRef(ctx context.Context, arg AccountParamRefRequest) (model.AccountParamRef, error)
	DeleteAccountParamRef(ctx context.Context, uuid uuid.UUID) error

	CreateSchedule(ctx context.Context, arg ScheduleRequest) (model.Schedule, error)
	GetSchedule(ctx context.Context, id int64) (ScheduleInfo, error)
	GetSchedulebyUuid(ctx context.Context, uuid uuid.UUID) (ScheduleInfo, error)
	GetSchedulebyAccId(ctx context.Context, accountId int64) (map[int16]ScheduleInfo, error)
	GetSchedulebyAcc(ctx context.Context, acc []string) (map[string]map[int16]ScheduleInfo, error)
	UpdateSchedule(ctx context.Context, arg ScheduleRequest) (model.Schedule, error)
	DeleteSchedule(ctx context.Context, id int64) error

	CreateOtherSchedule(ctx context.Context, arg OtherScheduleRequest) (model.OtherSchedule, error)
	GetOtherSchedule(ctx context.Context, uuid uuid.UUID) (OtherScheduleInfo, error)
	GetOtherSchedulebyUuid(ctx context.Context, uuid uuid.UUID) (OtherScheduleInfo, error)
	ListOtherSchedule(ctx context.Context, arg ListOtherScheduleParams) ([]OtherScheduleInfo, error)
	UpdateOtherSchedule(ctx context.Context, arg OtherScheduleRequest) (model.OtherSchedule, error)
	DeleteOtherSchedule(ctx context.Context, uuid uuid.UUID) error

	CreateAccountBeneficiary(ctx context.Context, arg AccountBeneficiaryRequest) (model.AccountBeneficiary, error)
	GetAccountBeneficiary(ctx context.Context, uuid uuid.UUID) (AccountBeneficiaryInfo, error)
	GetAccountBeneficiarybyUuid(ctx context.Context, uuid uuid.UUID) (AccountBeneficiaryInfo, error)
	ListAccountBeneficiary(ctx context.Context, arg ListAccountBeneficiaryParams) ([]AccountBeneficiaryInfo, error)
	UpdateAccountBeneficiary(ctx context.Context, arg AccountBeneficiaryRequest) (model.AccountBeneficiary, error)
	DeleteAccountBeneficiary(ctx context.Context, uuid uuid.UUID) error

	CreateOfficeAccount(ctx context.Context, arg OfficeAccountRequest) (model.OfficeAccount, error)
	GetOfficeAccount(ctx context.Context, id int64) (OfficeAccountInfo, error)
	GetOfficeAccountbyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountInfo, error)
	ListOfficeAccount(ctx context.Context, arg ListOfficeAccountParams) ([]OfficeAccountInfo, error)
	UpdateOfficeAccount(ctx context.Context, arg OfficeAccountRequest) (model.OfficeAccount, error)
	DeleteOfficeAccount(ctx context.Context, id int64) error

	CreateAccountParam(ctx context.Context, arg AccountParamRequest) (model.AccountParam, error)
	GetAccountParam(ctx context.Context, id int64) (AccountParamInfo, error)
	GetAccountParambyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamInfo, error)
	ListAccountParam(ctx context.Context, arg ListAccountParamParams) ([]AccountParamInfo, error)
	UpdateAccountParam(ctx context.Context, arg AccountParamRequest) (model.AccountParam, error)
	DeleteAccountParam(ctx context.Context, id int64) error

	CreateGlAccount(ctx context.Context, arg GlAccountRequest) (model.GlAccount, error)
	GetGlAccount(ctx context.Context, id int64) (GlAccountInfo, error)
	GetGlAccountbyUuid(ctx context.Context, uuid uuid.UUID) (GlAccountInfo, error)
	ListGlAccount(ctx context.Context, arg ListGlAccountParams) ([]GlAccountInfo, error)
	UpdateGlAccount(ctx context.Context, arg GlAccountRequest) (model.GlAccount, error)
	DeleteGlAccount(ctx context.Context, id int64) error

	CreateInventoryRepository(ctx context.Context, arg InventoryRepositoryRequest) (model.InventoryRepository, error)
	GetInventoryRepository(ctx context.Context, id int64) (InventoryRepositoryInfo, error)
	GetInventoryRepositorybyUuid(ctx context.Context, uuid uuid.UUID) (InventoryRepositoryInfo, error)
	GetInventoryRepositorybyName(ctx context.Context, name string) (InventoryRepositoryInfo, error)
	ListInventoryRepository(ctx context.Context, arg ListInventoryRepositoryParams) ([]InventoryRepositoryInfo, error)
	UpdateInventoryRepository(ctx context.Context, arg InventoryRepositoryRequest) (model.InventoryRepository, error)
	DeleteInventoryRepository(ctx context.Context, id int64) error
}

var _ QuerierAccount = (*QueriesAccount)(nil)

var _ StoreAccount = (*SQLStoreAccount)(nil)

func (q *QueriesAccount) WithTx(tx *sql.Tx) *QueriesAccount {
	return &QueriesAccount{
		db: tx,
	}
}

// SQLStore provides all functions to execute SQL queriesUser and users
type SQLStoreAccount struct {
	db *sql.DB
	*QueriesAccount
	*dsRef.QueriesReference
}

// Store defines all functions to execute db queriesUser and users
type StoreAccount interface {
	QuerierAccount
	dsRef.QuerierReference
	CreateInventoryItemFull(ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error)
	CreateInventoryDetailFull(ctx context.Context, arg InventoryDetailFullRequest) (model.InventoryDetail, error) // TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// NewStore creates a new store
func NewStoreAccount(db *sql.DB) StoreAccount {
	return &SQLStoreAccount{
		QueriesAccount:   New(db),
		QueriesReference: dsRef.New(db),
		db:               db,
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStoreAccount) ExecTx(ctx context.Context,
	fn func(*QueriesAccount) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
