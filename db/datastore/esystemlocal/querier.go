package db

import (
	"context"
	"time"
)

type QuerierLocal interface {
	Sql2Csv(ctx context.Context, sql string, filenamePath string) error
	GetModifiedTable(ctx context.Context, modCtr int64) (ModifiedTableInfo, error)
	ListModifiedTable(ctx context.Context) ([]ModifiedTable, error)
	UpdateModifiedTableUploaded(ctx context.Context, modCtr []int64, updated bool) error

	//  Area
	CreateArea(ctx context.Context, arg AreaRequest) error
	GetArea(ctx context.Context, areaCode int64) (AreaInfo, error)
	ListArea(ctx context.Context) ([]AreaInfo, error)
	UpdateArea(ctx context.Context, arg AreaRequest) error
	DeleteArea(ctx context.Context, areaCode int64) error
	AreaCSV(ctx context.Context, filenamePath string) error
	//  Unit
	CreateUnit(ctx context.Context, arg UnitRequest) error
	GetUnit(ctx context.Context, unitCode int64) (UnitInfo, error)
	ListUnit(ctx context.Context) ([]UnitInfo, error)
	UpdateUnit(ctx context.Context, arg UnitRequest) error
	DeleteUnit(ctx context.Context, unitCode int64) error
	UnitCSV(ctx context.Context, filenamePath string) error
	//  Center
	CreateCenter(ctx context.Context, arg CenterRequest) error
	GetCenter(ctx context.Context, centerCode string) (CenterInfo, error)
	ListCenter(ctx context.Context) ([]CenterInfo, error)
	UpdateCenter(ctx context.Context, arg CenterRequest) error
	DeleteCenter(ctx context.Context, centerCode string) error
	CenterCSV(ctx context.Context, filenamePath string) error
	//  Customer
	CreateCustomer(ctx context.Context, arg CustomerRequest) error
	GetCustomer(ctx context.Context, cid int64) (CustomerInfo, error)
	ListCustomer(ctx context.Context) ([]CustomerInfo, error)
	UpdateCustomer(ctx context.Context, arg CustomerRequest) error
	DeleteCustomer(ctx context.Context, cid int64) error
	CustomerCSV(ctx context.Context, filenamePath string) error
	//  Addresses
	CreateAddresses(ctx context.Context, arg AddressesRequest) (int64, error)
	GetAddresses(ctx context.Context, seqNum int64) (AddressesInfo, error)
	ListAddresses(ctx context.Context) ([]AddressesInfo, error)
	UpdateAddresses(ctx context.Context, arg AddressesRequest) error
	DeleteAddresses(ctx context.Context, seqNum int64) error
	AddressesCSV(ctx context.Context, filenamePath string) error
	//  LnMaster
	CreateLnMaster(ctx context.Context, arg LnMasterRequest) error
	GetLnMaster(ctx context.Context, acc string) (LnMasterInfo, error)
	ListLnMaster(ctx context.Context) ([]LnMasterInfo, error)
	UpdateLnMaster(ctx context.Context, arg LnMasterRequest) error
	DeleteLnMaster(ctx context.Context, acc string) error
	LnMasterCSV(ctx context.Context, filenamePath string) error
	//  SaMaster
	CreateSaMaster(ctx context.Context, arg SaMasterRequest) error
	GetSaMaster(ctx context.Context, acc string) (SaMasterInfo, error)
	ListSaMaster(ctx context.Context) ([]SaMasterInfo, error)
	UpdateSaMaster(ctx context.Context, arg SaMasterRequest) error
	DeleteSaMaster(ctx context.Context, acc string) error
	SaMasterCSV(ctx context.Context, filenamePath string) error
	//  TrnMaster
	CreateTrnMaster(ctx context.Context, arg TrnMasterRequest) error
	GetTrnMaster(ctx context.Context, trnDate time.Time, trn int64) (TrnMasterInfo, error)
	ListTrnMaster(ctx context.Context) ([]TrnMasterInfo, error)
	UpdateTrnMaster(ctx context.Context, arg TrnMasterRequest) error
	DeleteTrnMaster(ctx context.Context, trnDate time.Time, trn int64) error
	TrnMasterCSV(ctx context.Context, filenamePath string) error
	//  SaTrnMaster
	CreateSaTrnMaster(ctx context.Context, arg SaTrnMasterRequest) error
	GetSaTrnMaster(ctx context.Context, trnDate time.Time, trn int64) (SaTrnMasterInfo, error)
	ListSaTrnMaster(ctx context.Context) ([]SaTrnMasterInfo, error)
	UpdateSaTrnMaster(ctx context.Context, arg SaTrnMasterRequest) error
	DeleteSaTrnMaster(ctx context.Context, trnDate time.Time, trn int64) error
	SaTrnMasterCSV(ctx context.Context, filenamePath string) error
	//  LoanInst
	CreateLoanInst(ctx context.Context, arg LoanInstRequest) error
	GetLoanInst(ctx context.Context, acc string, dnum int64) (LoanInstInfo, error)
	ListLoanInst(ctx context.Context) ([]LoanInstInfo, error)
	UpdateLoanInst(ctx context.Context, arg LoanInstRequest) error
	DeleteLoanInst(ctx context.Context, acc string, dnum int64) error
	LoanInstCSV(ctx context.Context, filenamePath string) error
	//  LnChrgData
	CreateLnChrgData(ctx context.Context, arg LnChrgDataRequest) error
	GetLnChrgData(ctx context.Context, acc string, chrgCode int64, refAcc string) (LnChrgDataInfo, error)
	ListLnChrgData(ctx context.Context) ([]LnChrgDataInfo, error)
	UpdateLnChrgData(ctx context.Context, arg LnChrgDataRequest) error
	DeleteLnChrgData(ctx context.Context, acc string, chrgCode int64, refAcc string) error
	LnChrgDataCSV(ctx context.Context, filenamePath string) error
	//  CustAddInfoList
	CreateCustAddInfoList(ctx context.Context, arg CustAddInfoListRequest) error
	GetCustAddInfoList(ctx context.Context, infoCode int64) (CustAddInfoListInfo, error)
	ListCustAddInfoList(ctx context.Context) ([]CustAddInfoListInfo, error)
	UpdateCustAddInfoList(ctx context.Context, arg CustAddInfoListRequest) error
	DeleteCustAddInfoList(ctx context.Context, infoCode int64) error
	CustAddInfoListCSV(ctx context.Context, filenamePath string) error
	//  CustAddInfoGroup
	CreateCustAddInfoGroup(ctx context.Context, arg CustAddInfoGroupRequest) error
	GetCustAddInfoGroup(ctx context.Context, infoGroup int64) (CustAddInfoGroupInfo, error)
	ListCustAddInfoGroup(ctx context.Context) ([]CustAddInfoGroupInfo, error)
	UpdateCustAddInfoGroup(ctx context.Context, arg CustAddInfoGroupRequest) error
	DeleteCustAddInfoGroup(ctx context.Context, infoGroup int64) error
	CustAddInfoGroupCSV(ctx context.Context, filenamePath string) error
	//  CustAddInfoGroupNeed
	CreateCustAddInfoGroupNeed(ctx context.Context, arg CustAddInfoGroupNeedRequest) error
	GetCustAddInfoGroupNeed(ctx context.Context, infoGroup int64, infoCode int64) (CustAddInfoGroupNeedInfo, error)
	ListCustAddInfoGroupNeed(ctx context.Context) ([]CustAddInfoGroupNeedInfo, error)
	UpdateCustAddInfoGroupNeed(ctx context.Context, arg CustAddInfoGroupNeedRequest) error
	DeleteCustAddInfoGroupNeed(ctx context.Context, infoGroup int64, infoCode int64) error
	CustAddInfoGroupNeedCSV(ctx context.Context, filenamePath string) error
	//  CustAddInfo
	CreateCustAddInfo(ctx context.Context, arg CustAddInfoRequest) error
	GetCustAddInfo(ctx context.Context, cid int64, infoDate time.Time, infoCode int64) (CustAddInfoInfo, error)
	ListCustAddInfo(ctx context.Context) ([]CustAddInfoInfo, error)
	UpdateCustAddInfo(ctx context.Context, arg CustAddInfoRequest) error
	DeleteCustAddInfo(ctx context.Context, cid int64, infoDate time.Time, infoCode int64) error
	CustAddInfoCSV(ctx context.Context, filenamePath string) error
	//  MutualFund
	CreateMutualFund(ctx context.Context, arg MutualFundRequest) error
	GetMutualFund(ctx context.Context, cid int64, orNo int64, trnDate time.Time) (MutualFundInfo, error)
	ListMutualFund(ctx context.Context) ([]MutualFundInfo, error)
	UpdateMutualFund(ctx context.Context, arg MutualFundRequest) error
	DeleteMutualFund(ctx context.Context, cid int64, orNo int64, trnDate time.Time) error
	MutualFundCSV(ctx context.Context, filenamePath string) error
	//  ReferencesDetails
	CreateReferencesDetails(ctx context.Context, arg ReferencesDetailsRequest) error
	GetReferencesDetails(ctx context.Context, id int64) (ReferencesDetailsInfo, error)
	ListReferencesDetails(ctx context.Context) ([]ReferencesDetailsInfo, error)
	UpdateReferencesDetails(ctx context.Context, arg ReferencesDetailsRequest) error
	DeleteReferencesDetails(ctx context.Context, id int64) error
	ReferencesDetailsCSV(ctx context.Context, filenamePath string) error
	//  CenterWorker
	CreateCenterWorker(ctx context.Context, arg CenterWorkerRequest) error
	GetCenterWorker(ctx context.Context, aoID int64) (CenterWorkerInfo, error)
	ListCenterWorker(ctx context.Context) ([]CenterWorkerInfo, error)
	UpdateCenterWorker(ctx context.Context, arg CenterWorkerRequest) error
	DeleteCenterWorker(ctx context.Context, aoID int64) error
	CenterWorkerCSV(ctx context.Context, filenamePath string) error
	//  Writeoff
	CreateWriteoff(ctx context.Context, arg WriteoffRequest) error
	GetWriteoff(ctx context.Context, acc string) (WriteoffInfo, error)
	ListWriteoff(ctx context.Context) ([]WriteoffInfo, error)
	UpdateWriteoff(ctx context.Context, arg WriteoffRequest) error
	DeleteWriteoff(ctx context.Context, acc string) error
	WriteoffCSV(ctx context.Context, filenamePath string) error
	//  Accounts
	CreateAccounts(ctx context.Context, arg AccountsRequest) error
	GetAccounts(ctx context.Context, acc string) (AccountsInfo, error)
	ListAccounts(ctx context.Context) ([]AccountsInfo, error)
	UpdateAccounts(ctx context.Context, arg AccountsRequest) error
	DeleteAccounts(ctx context.Context, acc string) error
	AccountsCSV(ctx context.Context, filenamePath string) error
	//  JnlHeaders
	CreateJnlHeaders(ctx context.Context, arg JnlHeadersRequest) error
	GetJnlHeaders(ctx context.Context, trn string) (JnlHeadersInfo, error)
	ListJnlHeaders(ctx context.Context) ([]JnlHeadersInfo, error)
	UpdateJnlHeaders(ctx context.Context, arg JnlHeadersRequest) error
	DeleteJnlHeaders(ctx context.Context, trn string) error
	JnlHeadersCSV(ctx context.Context, filenamePath string) error
	//  JnlDetails
	CreateJnlDetails(ctx context.Context, arg JnlDetailsRequest) error
	GetJnlDetails(ctx context.Context, trn string, acc string) (JnlDetailsInfo, error)
	ListJnlDetails(ctx context.Context) ([]JnlDetailsInfo, error)
	UpdateJnlDetails(ctx context.Context, arg JnlDetailsRequest) error
	DeleteJnlDetails(ctx context.Context, trn string, acc string) error
	JnlDetailsCSV(ctx context.Context, filenamePath string) error
	//  LedgerDetails
	CreateLedgerDetails(ctx context.Context, arg LedgerDetailsRequest) error
	GetLedgerDetails(ctx context.Context, trnDate time.Time, trn string) (LedgerDetailsInfo, error)
	ListLedgerDetails(ctx context.Context) ([]LedgerDetailsInfo, error)
	UpdateLedgerDetails(ctx context.Context, arg LedgerDetailsRequest) error
	DeleteLedgerDetails(ctx context.Context, trnDate time.Time, trn string) error
	LedgerDetailsCSV(ctx context.Context, filenamePath string) error

	CreateMultiplePaymentReceipt(ctx context.Context, arg MultiplePaymentReceiptRequest) error
	GetMultiplePaymentReceipt(ctx context.Context, orNo int64) (MultiplePaymentReceipt, error)
	ListMultiplePaymentReceipt(ctx context.Context) ([]MultiplePaymentReceipt, error)
	UpdateMultiplePaymentReceipt(ctx context.Context, arg MultiplePaymentReceiptRequest) error
	DeleteMultiplePaymentReceipt(ctx context.Context, orNo int64) error
	MultiplePaymentReceiptCSV(ctx context.Context, filenamePath string) error

	//  CreateInActiveCID(ctx context.Context, arg InActiveCIDRequest) error

	CreateInActiveCID(ctx context.Context, arg InActiveCIDRequest) error
	GetInActiveCID(ctx context.Context, cid int64, dateStart time.Time) (InActiveCID, error)
	ListInActiveCID(ctx context.Context) ([]InActiveCID, error)
	UpdateInActiveCID(ctx context.Context, arg InActiveCIDRequest) error
	DeleteInActiveCID(ctx context.Context, cid int64, dateStart time.Time) error
	InActiveCIDCSV(ctx context.Context, filenamePath string) error

	CreateLnBeneficiary(ctx context.Context, arg LnBeneficiaryRequest) error
	GetLnBeneficiary(ctx context.Context, acc string) (LnBeneficiary, error)
	ListLnBeneficiary(ctx context.Context) ([]LnBeneficiary, error)
	UpdateLnBeneficiary(ctx context.Context, arg LnBeneficiaryRequest) error
	DeleteLnBeneficiary(ctx context.Context, acc string) error
	LnBeneficiaryCSV(ctx context.Context, filenamePath string) error

	CreateReactivateWriteoff(ctx context.Context, arg ReactivateWriteoffRequest) error
	GetReactivateWriteoff(ctx context.Context, id int64) (ReactivateWriteoff, error)
	GetReactivateWriteoffbyCID(ctx context.Context, cid int64) (ReactivateWriteoff, error)
	ListReactivateWriteoff(ctx context.Context) ([]ReactivateWriteoff, error)
	UpdateReactivateWriteoff(ctx context.Context, arg ReactivateWriteoffRequest) error
	DeleteReactivateWriteoff(ctx context.Context, id int64) error
	ReactivateWriteoffCSV(ctx context.Context, filenamePath string) error

	GetColSht(ctx context.Context) ([]ColSht, error)
	GetColShtPerCID(ctx context.Context, cid int64) ([]ColSht, error)
	ColShtCSV(ctx context.Context, filenamePath string) error
}

var _ QuerierLocal = (*QueriesLocal)(nil)
