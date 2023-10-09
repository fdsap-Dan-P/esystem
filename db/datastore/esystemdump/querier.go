package db

import (
	"context"
	model "simplebank/db/datastore/esystemlocal"
	"time"
)

type QuerierDump interface {
	//  ModifiedTable
	ListModifiedTable(ctx context.Context, brCode string) ([]ModifiedTable, error)
	//  BranchList
	CreateBranchList(ctx context.Context, arg BranchList) error
	GetBranchList(ctx context.Context, brCode string) (BranchList, error)
	ListBranchList(ctx context.Context) ([]BranchList, error)
	UpdateBranchList(ctx context.Context, arg BranchListLight) error
	DeleteBranchList(ctx context.Context, brCode string) error
	//  Area
	CreateArea(ctx context.Context, arg model.Area) error
	GetArea(ctx context.Context, brCode string, areaCode int64) (model.Area, error)
	ListArea(ctx context.Context, modCtr int64) ([]model.Area, error)
	UpdateArea(ctx context.Context, arg model.Area) error
	DeleteArea(ctx context.Context, brCode string, areaCode int64) error
	//  Unit
	CreateUnit(ctx context.Context, arg model.Unit) error
	GetUnit(ctx context.Context, brCode string, unitCode int64) (model.Unit, error)
	ListUnit(ctx context.Context, modCtr int64) ([]model.Unit, error)
	UpdateUnit(ctx context.Context, arg model.Unit) error
	DeleteUnit(ctx context.Context, brCode string, unitCode int64) error
	//  Center
	CreateCenter(ctx context.Context, arg model.Center) error
	GetCenter(ctx context.Context, brCode string, centerCode string) (model.Center, error)
	ListCenter(ctx context.Context, modCtr int64) ([]model.Center, error)
	UpdateCenter(ctx context.Context, arg model.Center) error
	DeleteCenter(ctx context.Context, brCode string, centerCode string) error
	//  Customer
	CreateCustomer(ctx context.Context, arg model.Customer) error
	GetCustomer(ctx context.Context, brCode string, cid int64) (model.Customer, error)
	ListCustomer(ctx context.Context, modCtr int64) ([]model.Customer, error)
	UpdateCustomer(ctx context.Context, arg model.Customer) error
	DeleteCustomer(ctx context.Context, brCode string, cid int64) error
	//  Addresses
	CreateAddresses(ctx context.Context, arg model.Addresses) error
	GetAddresses(ctx context.Context, brCode string, seqNum int64) (model.Addresses, error)
	ListAddresses(ctx context.Context, modCtr int64) ([]model.Addresses, error)
	UpdateAddresses(ctx context.Context, arg model.Addresses) error
	DeleteAddresses(ctx context.Context, brCode string, seqNum int64) error
	//  LnMaster
	CreateLnMaster(ctx context.Context, arg model.LnMaster) error
	GetLnMaster(ctx context.Context, brCode string, acc string) (model.LnMaster, error)
	ListLnMaster(ctx context.Context, modCtr int64) ([]model.LnMaster, error)
	UpdateLnMaster(ctx context.Context, arg model.LnMaster) error
	DeleteLnMaster(ctx context.Context, brCode string, acc string) error
	//  SaMaster
	CreateSaMaster(ctx context.Context, arg model.SaMaster) error
	GetSaMaster(ctx context.Context, brCode string, acc string) (model.SaMaster, error)
	ListSaMaster(ctx context.Context, modCtr int64) ([]model.SaMaster, error)
	UpdateSaMaster(ctx context.Context, arg model.SaMaster) error
	DeleteSaMaster(ctx context.Context, brCode string, acc string) error
	//  LoanInst
	CreateTrnMaster(ctx context.Context, arg model.TrnMaster) error
	GetTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) (model.TrnMaster, error)
	ListTrnMaster(ctx context.Context, modCtr int64) ([]model.TrnMaster, error)
	UpdateTrnMaster(ctx context.Context, arg model.TrnMaster) error
	DeleteTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) error
	//  SaTrnMaster
	CreateSaTrnMaster(ctx context.Context, arg model.SaTrnMaster) error
	GetSaTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) (model.SaTrnMaster, error)
	ListSaTrnMaster(ctx context.Context, modCtr int64) ([]model.SaTrnMaster, error)
	UpdateSaTrnMaster(ctx context.Context, arg model.SaTrnMaster) error
	DeleteSaTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) error
	//  LoanInst
	CreateLoanInst(ctx context.Context, arg model.LoanInst) error
	GetLoanInst(ctx context.Context, brCode string, acc string, dnum int64) (model.LoanInst, error)
	ListLoanInst(ctx context.Context, modCtr int64) ([]model.LoanInst, error)
	UpdateLoanInst(ctx context.Context, arg model.LoanInst) error
	DeleteLoanInst(ctx context.Context, brCode string, acc string, dnum int64) error
	//  LnChrgData
	CreateLnChrgData(ctx context.Context, arg model.LnChrgData) error
	GetLnChrgData(ctx context.Context, brCode string, acc string, chrgCode int64, refAcc string) (model.LnChrgData, error)
	ListLnChrgData(ctx context.Context, modCtr int64) ([]model.LnChrgData, error)
	UpdateLnChrgData(ctx context.Context, arg model.LnChrgData) error
	DeleteLnChrgData(ctx context.Context, brCode string, acc string, chrgCode int64, refAcc string) error
	//  CustAddInfoList
	CreateCustAddInfoList(ctx context.Context, arg model.CustAddInfoList) error
	GetCustAddInfoList(ctx context.Context, brCode string, infoCode int64) (model.CustAddInfoList, error)
	ListCustAddInfoList(ctx context.Context, modCtr int64) ([]model.CustAddInfoList, error)
	UpdateCustAddInfoList(ctx context.Context, arg model.CustAddInfoList) error
	DeleteCustAddInfoList(ctx context.Context, brCode string, infoCode int64) error
	//  CustAddInfoGroup
	CreateCustAddInfoGroup(ctx context.Context, arg model.CustAddInfoGroup) error
	GetCustAddInfoGroup(ctx context.Context, brCode string, infoGroup int64) (model.CustAddInfoGroup, error)
	ListCustAddInfoGroup(ctx context.Context, modCtr int64) ([]model.CustAddInfoGroup, error)
	UpdateCustAddInfoGroup(ctx context.Context, arg model.CustAddInfoGroup) error
	DeleteCustAddInfoGroup(ctx context.Context, brCode string, infoGroup int64) error
	//  CustAddInfoGroupNeed
	CreateCustAddInfoGroupNeed(ctx context.Context, arg model.CustAddInfoGroupNeed) error
	GetCustAddInfoGroupNeed(ctx context.Context, brCode string, infoGroup int64, infoCode int64) (model.CustAddInfoGroupNeed, error)
	ListCustAddInfoGroupNeed(ctx context.Context, modCtr int64) ([]model.CustAddInfoGroupNeed, error)
	UpdateCustAddInfoGroupNeed(ctx context.Context, arg model.CustAddInfoGroupNeed) error
	DeleteCustAddInfoGroupNeed(ctx context.Context, brCode string, infoGroup int64, infoCode int64) error
	//  CustAddInfo
	CreateCustAddInfo(ctx context.Context, arg model.CustAddInfo) error
	GetCustAddInfo(ctx context.Context, brCode string, cid int64, infoDate time.Time, infoCode int64) (model.CustAddInfo, error)
	ListCustAddInfo(ctx context.Context, modCtr int64) ([]model.CustAddInfo, error)
	UpdateCustAddInfo(ctx context.Context, arg model.CustAddInfo) error
	DeleteCustAddInfo(ctx context.Context, brCode string, cid int64, infoDate time.Time, infoCode int64) error
	BulkInsertCustAddInfo(ctx context.Context, rows []model.CustAddInfo) error
	//  MutualFund
	CreateMutualFund(ctx context.Context, arg model.MutualFund) error
	GetMutualFund(ctx context.Context, brCode string, cid int64, orNo int64) (model.MutualFund, error)
	ListMutualFund(ctx context.Context, modCtr int64) ([]model.MutualFund, error)
	UpdateMutualFund(ctx context.Context, arg model.MutualFund) error
	DeleteMutualFund(ctx context.Context, brCode string, cid int64, orNo int64) error
	//  ReferencesDetails
	CreateReferencesDetails(ctx context.Context, arg model.ReferencesDetails) error
	GetReferencesDetails(ctx context.Context, brCode string, id int64) (model.ReferencesDetails, error)
	ListReferencesDetails(ctx context.Context, modCtr int64) ([]model.ReferencesDetails, error)
	UpdateReferencesDetails(ctx context.Context, arg model.ReferencesDetails) error
	DeleteReferencesDetails(ctx context.Context, brCode string, id int64) error
	//  CenterWorker
	CreateCenterWorker(ctx context.Context, arg model.CenterWorker) error
	GetCenterWorker(ctx context.Context, brCode string, aoID int64) (model.CenterWorker, error)
	ListCenterWorker(ctx context.Context, modCtr int64) ([]model.CenterWorker, error)
	UpdateCenterWorker(ctx context.Context, arg model.CenterWorker) error
	DeleteCenterWorker(ctx context.Context, brCode string, aoID int64) error
	//  Writeoff
	CreateWriteoff(ctx context.Context, arg model.Writeoff) error
	GetWriteoff(ctx context.Context, brCode string, acc string) (model.Writeoff, error)
	ListWriteoff(ctx context.Context, modCtr int64) ([]model.Writeoff, error)
	UpdateWriteoff(ctx context.Context, arg model.Writeoff) error
	DeleteWriteoff(ctx context.Context, brCode string, acc string) error
	//  Accounts
	CreateAccounts(ctx context.Context, arg model.Accounts) error
	GetAccounts(ctx context.Context, brCode string, acc string) (model.Accounts, error)
	ListAccounts(ctx context.Context, modCtr int64) ([]model.Accounts, error)
	UpdateAccounts(ctx context.Context, arg model.Accounts) error
	DeleteAccounts(ctx context.Context, brCode string, acc string) error
	//  JnlHeaders
	CreateJnlHeaders(ctx context.Context, arg model.JnlHeaders) error
	GetJnlHeaders(ctx context.Context, brCode string, trn string) (model.JnlHeaders, error)
	ListJnlHeaders(ctx context.Context, modCtr int64) ([]model.JnlHeaders, error)
	UpdateJnlHeaders(ctx context.Context, arg model.JnlHeaders) error
	DeleteJnlHeaders(ctx context.Context, brCode string, trn string) error
	//  JnlDetails
	CreateJnlDetails(ctx context.Context, arg model.JnlDetails) error
	GetJnlDetails(ctx context.Context, brCode string, trn string, acc string) (model.JnlDetails, error)
	ListJnlDetails(ctx context.Context, modCtr int64) ([]model.JnlDetails, error)
	UpdateJnlDetails(ctx context.Context, arg model.JnlDetails) error
	DeleteJnlDetails(ctx context.Context, brCode string, trn string, acc string) error
	//  LedgerDetails
	CreateLedgerDetails(ctx context.Context, arg model.LedgerDetails) error
	GetLedgerDetails(ctx context.Context, brCode string, trnDate time.Time, trn string) (model.LedgerDetails, error)
	ListLedgerDetails(ctx context.Context, modCtr int64) ([]model.LedgerDetails, error)
	UpdateLedgerDetails(ctx context.Context, arg model.LedgerDetails) error
	DeleteLedgerDetails(ctx context.Context, brCode string, trnDate time.Time, trn string) error
	//  LedgerDetails
	CreateUsersList(ctx context.Context, arg model.UsersList) error
	GetUsersList(ctx context.Context, brCode string, userId string) (model.UsersList, error)
	ListUsersList(ctx context.Context, modCtr int64) ([]model.UsersList, error)
	UpdateUsersList(ctx context.Context, arg model.UsersList) error
	DeleteUsersList(ctx context.Context, brCode string, serId string) error

	CreateMultiplePaymentReceipt(ctx context.Context, arg model.MultiplePaymentReceipt) error
	GetMultiplePaymentReceipt(ctx context.Context, brCode string, orNo int64) (model.MultiplePaymentReceipt, error)
	ListMultiplePaymentReceipt(ctx context.Context, modCtr int64) ([]model.MultiplePaymentReceipt, error)
	UpdateMultiplePaymentReceipt(ctx context.Context, arg model.MultiplePaymentReceipt) error
	DeleteMultiplePaymentReceipt(ctx context.Context, brCode string, orNo int64) error

	CreateInActiveCID(ctx context.Context, arg model.InActiveCID) error
	GetInActiveCID(ctx context.Context, brCode string, cid int64, dataStart time.Time) (model.InActiveCID, error)
	ListInActiveCID(ctx context.Context, modCtr int64) ([]model.InActiveCID, error)
	UpdateInActiveCID(ctx context.Context, arg model.InActiveCID) error
	DeleteInActiveCID(ctx context.Context, brCode string, cid int64, dataStart time.Time) error

	CreateLnBeneficiary(ctx context.Context, arg model.LnBeneficiary) error
	GetLnBeneficiary(ctx context.Context, brCode string, acc string) (model.LnBeneficiary, error)
	ListLnBeneficiary(ctx context.Context, modCtr int64) ([]model.LnBeneficiary, error)
	UpdateLnBeneficiary(ctx context.Context, arg model.LnBeneficiary) error
	DeleteLnBeneficiary(ctx context.Context, brCode string, acc string) error

	CreateReactivateWriteoff(ctx context.Context, arg model.ReactivateWriteoff) error
	GetReactivateWriteoff(ctx context.Context, brCode string, cid int64) (model.ReactivateWriteoff, error)
	ListReactivateWriteoff(ctx context.Context, modCtr int64) ([]model.ReactivateWriteoff, error)
	UpdateReactivateWriteoff(ctx context.Context, arg model.ReactivateWriteoff) error
	DeleteReactivateWriteoff(ctx context.Context, brCode string, cid int64) error

	CreateColSht(ctx context.Context, arg model.ColSht) error
	DeleteColSht(ctx context.Context, brCode string) error
	GetColSht(ctx context.Context, acc string) (model.ColSht, error)
	ColShtPerBranch(ctx context.Context, brCode string) ([]model.ColSht, error)
	ColShtPerCID(ctx context.Context, cid int64) ([]model.ColSht, error)
	ColShtPerCenter(ctx context.Context, cenCode string) ([]model.ColSht, error)
	UpdateColSht(ctx context.Context, arg model.ColSht) error
}

var _ QuerierDump = (*QueriesDump)(nil)
