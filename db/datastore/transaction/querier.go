package db

import (
	"context"
	"simplebank/model"

	"github.com/google/uuid"
)

// var QueriesAccount *account.QueriesAccount = account.New(testDB)

type QuerierTransaction interface {
	// GetEntry(ctx context.Context, id int64) (model.Entry, error)
	// ListEntries(ctx context.Context, arg ListEntriesParams) ([]model.Entry, error)
	// CreateEntry(ctx context.Context, arg CreateEntryParams) (model.Entry, error)

	CreateTransfer(ctx context.Context, arg CreateTransferParams) (model.Transfer, error)
	GetTransfer(ctx context.Context, id int64) (model.Transfer, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]model.Transfer, error)

	CreateInventoryTran(ctx context.Context, arg InventoryTranRequest) (model.InventoryTran, error)
	// GetInventoryTran(ctx context.Context, uuid uuid.UUID) (InventoryTranInfo, error)
	GetInventoryTranbyUuid(ctx context.Context, uuid uuid.UUID) (InventoryTranInfo, error)
	ListInventoryTran(ctx context.Context, arg ListInventoryTranParams) ([]InventoryTranInfo, error)
	UpdateInventoryTran(ctx context.Context, arg InventoryTranRequest) (model.InventoryTran, error)
	DeleteInventoryTran(ctx context.Context, uuid uuid.UUID) error

	CreateAccountTran(ctx context.Context, arg AccountTranRequest) (model.AccountTran, error)
	GetAccountTran(ctx context.Context, trnTypeCode int64, series int64) (AccountTranInfo, error)
	GetAccountTranbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTranInfo, error)
	ListAccountTran(ctx context.Context, arg ListAccountTranParams) ([]AccountTranInfo, error)
	UpdateAccountTran(ctx context.Context, arg AccountTranRequest) (model.AccountTran, error)
	DeleteAccountTran(ctx context.Context, uuid uuid.UUID) error

	CreateTicketTypeAction(ctx context.Context, arg TicketTypeActionRequest) (model.TicketTypeAction, error)
	GetTicketTypeAction(ctx context.Context, id int64) (TicketTypeActionInfo, error)
	GetTicketTypeActionbyUuid(ctx context.Context, uuid uuid.UUID) (TicketTypeActionInfo, error)
	ListTicketTypeAction(ctx context.Context, arg ListTicketTypeActionParams) ([]TicketTypeActionInfo, error)
	UpdateTicketTypeAction(ctx context.Context, arg TicketTypeActionRequest) (model.TicketTypeAction, error)
	DeleteTicketTypeAction(ctx context.Context, uuid uuid.UUID) error

	CreateTicketTypeStatus(ctx context.Context, arg TicketTypeStatusRequest) (model.TicketTypeStatus, error)
	GetTicketTypeStatus(ctx context.Context, int int64) (TicketTypeStatusInfo, error)
	GetTicketTypeStatusbyUuid(ctx context.Context, uuid uuid.UUID) (TicketTypeStatusInfo, error)
	ListTicketTypeStatus(ctx context.Context, arg ListTicketTypeStatusParams) ([]TicketTypeStatusInfo, error)
	UpdateTicketTypeStatus(ctx context.Context, arg TicketTypeStatusRequest) (model.TicketTypeStatus, error)
	DeleteTicketTypeStatus(ctx context.Context, uuid uuid.UUID) error

	CreateCustomerEvent(ctx context.Context, arg CustomerEventRequest) (model.CustomerEvent, error)
	GetCustomerEvent(ctx context.Context, uuid uuid.UUID) (CustomerEventInfo, error)
	GetCustomerEventbyUuid(ctx context.Context, uuid uuid.UUID) (CustomerEventInfo, error)
	ListCustomerEvent(ctx context.Context, arg ListCustomerEventParams) ([]CustomerEventInfo, error)
	UpdateCustomerEvent(ctx context.Context, arg CustomerEventRequest) (model.CustomerEvent, error)
	DeleteCustomerEvent(ctx context.Context, uuid uuid.UUID) error

	CreateEmployeeEvent(ctx context.Context, arg EmployeeEventRequest) (model.EmployeeEvent, error)
	GetEmployeeEvent(ctx context.Context, uuid uuid.UUID) (EmployeeEventInfo, error)
	GetEmployeeEventbyUuid(ctx context.Context, uuid uuid.UUID) (EmployeeEventInfo, error)
	ListEmployeeEvent(ctx context.Context, arg ListEmployeeEventParams) ([]EmployeeEventInfo, error)
	UpdateEmployeeEvent(ctx context.Context, arg EmployeeEventRequest) (model.EmployeeEvent, error)
	DeleteEmployeeEvent(ctx context.Context, uuid uuid.UUID) error

	CreateJournalDetail(ctx context.Context, arg JournalDetailRequest) (model.JournalDetail, error)
	GetJournalDetail(ctx context.Context, uuid uuid.UUID) (JournalDetailInfo, error)
	GetJournalDetailbyUuid(ctx context.Context, uuid uuid.UUID) (JournalDetailInfo, error)
	ListJournalDetail(ctx context.Context, arg ListJournalDetailParams) ([]JournalDetailInfo, error)
	UpdateJournalDetail(ctx context.Context, arg JournalDetailRequest) (model.JournalDetail, error)
	DeleteJournalDetail(ctx context.Context, uuid uuid.UUID) error

	CreateOfficeAccountTran(ctx context.Context, arg OfficeAccountTranRequest) (model.OfficeAccountTran, error)
	GetOfficeAccountTran(ctx context.Context, uuid uuid.UUID) (OfficeAccountTranInfo, error)
	GetOfficeAccountTranbyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountTranInfo, error)
	ListOfficeAccountTran(ctx context.Context, arg ListOfficeAccountTranParams) ([]OfficeAccountTranInfo, error)
	UpdateOfficeAccountTran(ctx context.Context, arg OfficeAccountTranRequest) (model.OfficeAccountTran, error)
	DeleteOfficeAccountTran(ctx context.Context, uuid uuid.UUID) error

	NewDailyTicket(ctx context.Context, arg NewDailyTicketRequest) (model.Ticket, error)
	CreateTicket(ctx context.Context, arg TicketRequest) (model.Ticket, error)
	GetTicket(ctx context.Context, id int64) (TicketInfo, error)
	GetTicketbyUuid(ctx context.Context, uuid uuid.UUID) (TicketInfo, error)
	ListTicket(ctx context.Context, arg ListTicketParams) ([]TicketInfo, error)
	UpdateTicket(ctx context.Context, arg TicketRequest) (model.Ticket, error)
	DeleteTicket(ctx context.Context, uuid uuid.UUID) error

	// CreateTrnAction(ctx context.Context, arg TrnActionRequest) (model.TrnAction, error)
	// GetTrnAction(ctx context.Context, uuid uuid.UUID) (TrnActionInfo, error)
	// GetTrnActionbyUuid(ctx context.Context, uuid uuid.UUID) (TrnActionInfo, error)
	// ListTrnAction(ctx context.Context, arg ListTrnActionParams) ([]TrnActionInfo, error)
	// UpdateTrnAction(ctx context.Context, arg TrnActionRequest) (model.TrnAction, error)
	// DeleteTrnAction(ctx context.Context, uuid uuid.UUID) error

	CreateTrnHeadRelation(ctx context.Context, arg TrnHeadRelationRequest) (model.TrnHeadRelation, error)
	GetTrnHeadRelation(ctx context.Context, uuid uuid.UUID) (TrnHeadRelationInfo, error)
	GetTrnHeadRelationbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadRelationInfo, error)
	ListTrnHeadRelation(ctx context.Context, arg ListTrnHeadRelationParams) ([]TrnHeadRelationInfo, error)
	UpdateTrnHeadRelation(ctx context.Context, arg TrnHeadRelationRequest) (model.TrnHeadRelation, error)
	DeleteTrnHeadRelation(ctx context.Context, uuid uuid.UUID) error

	CreateTrnHeadSpecsDate(ctx context.Context, arg TrnHeadSpecsDateRequest) (model.TrnHeadSpecsDate, error)
	GetTrnHeadSpecsDate(ctx context.Context, TrnHeadItemId int64, specsId int64) (TrnHeadSpecsDateInfo, error)
	GetTrnHeadSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsDateInfo, error)
	ListTrnHeadSpecsDate(ctx context.Context, arg ListTrnHeadSpecsDateParams) ([]TrnHeadSpecsDateInfo, error)
	UpdateTrnHeadSpecsDate(ctx context.Context, arg TrnHeadSpecsDateRequest) (model.TrnHeadSpecsDate, error)
	DeleteTrnHeadSpecsDate(ctx context.Context, uuid uuid.UUID) error

	CreateTrnHeadSpecsString(ctx context.Context, arg TrnHeadSpecsStringRequest) (model.TrnHeadSpecsString, error)
	GetTrnHeadSpecsString(ctx context.Context, TrnHeadItemId int64, specsId int64) (TrnHeadSpecsStringInfo, error)
	GetTrnHeadSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsStringInfo, error)
	ListTrnHeadSpecsString(ctx context.Context, arg ListTrnHeadSpecsStringParams) ([]TrnHeadSpecsStringInfo, error)
	UpdateTrnHeadSpecsString(ctx context.Context, arg TrnHeadSpecsStringRequest) (model.TrnHeadSpecsString, error)
	DeleteTrnHeadSpecsString(ctx context.Context, uuid uuid.UUID) error

	CreateTrnHeadSpecsNumber(ctx context.Context, arg TrnHeadSpecsNumberRequest) (model.TrnHeadSpecsNumber, error)
	GetTrnHeadSpecsNumber(ctx context.Context, TrnHeadItemId int64, specsId int64) (TrnHeadSpecsNumberInfo, error)
	GetTrnHeadSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsNumberInfo, error)
	ListTrnHeadSpecsNumber(ctx context.Context, arg ListTrnHeadSpecsNumberParams) ([]TrnHeadSpecsNumberInfo, error)
	UpdateTrnHeadSpecsNumber(ctx context.Context, arg TrnHeadSpecsNumberRequest) (model.TrnHeadSpecsNumber, error)
	DeleteTrnHeadSpecsNumber(ctx context.Context, uuid uuid.UUID) error

	CreateTrnHeadSpecsRef(ctx context.Context, arg TrnHeadSpecsRefRequest) (model.TrnHeadSpecsRef, error)
	GetTrnHeadSpecsRef(ctx context.Context, TrnHeadItemId int64, specsId int64) (TrnHeadSpecsRefInfo, error)
	GetTrnHeadSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsRefInfo, error)
	ListTrnHeadSpecsRef(ctx context.Context, arg ListTrnHeadSpecsRefParams) ([]TrnHeadSpecsRefInfo, error)
	UpdateTrnHeadSpecsRef(ctx context.Context, arg TrnHeadSpecsRefRequest) (model.TrnHeadSpecsRef, error)
	DeleteTrnHeadSpecsRef(ctx context.Context, uuid uuid.UUID) error

	CreateTicketActionConditionDate(ctx context.Context, arg TicketActionConditionDateRequest) (model.TicketActionConditionDate, error)
	GetTicketActionConditionDate(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionDateInfo, error)
	GetTicketActionConditionDatebyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionDateInfo, error)
	ListTicketActionConditionDate(ctx context.Context, arg ListTicketActionConditionDateParams) ([]TicketActionConditionDateInfo, error)
	UpdateTicketActionConditionDate(ctx context.Context, arg TicketActionConditionDateRequest) (model.TicketActionConditionDate, error)
	DeleteTicketActionConditionDate(ctx context.Context, uuid uuid.UUID) error

	CreateTicketActionConditionString(ctx context.Context, arg TicketActionConditionStringRequest) (model.TicketActionConditionString, error)
	GetTicketActionConditionString(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionStringInfo, error)
	GetTicketActionConditionStringbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionStringInfo, error)
	ListTicketActionConditionString(ctx context.Context, arg ListTicketActionConditionStringParams) ([]TicketActionConditionStringInfo, error)
	UpdateTicketActionConditionString(ctx context.Context, arg TicketActionConditionStringRequest) (model.TicketActionConditionString, error)
	DeleteTicketActionConditionString(ctx context.Context, uuid uuid.UUID) error

	CreateTicketActionConditionNumber(ctx context.Context, arg TicketActionConditionNumberRequest) (model.TicketActionConditionNumber, error)
	GetTicketActionConditionNumber(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionNumberInfo, error)
	GetTicketActionConditionNumberbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionNumberInfo, error)
	ListTicketActionConditionNumber(ctx context.Context, arg ListTicketActionConditionNumberParams) ([]TicketActionConditionNumberInfo, error)
	UpdateTicketActionConditionNumber(ctx context.Context, arg TicketActionConditionNumberRequest) (model.TicketActionConditionNumber, error)
	DeleteTicketActionConditionNumber(ctx context.Context, uuid uuid.UUID) error

	CreateTicketActionConditionRef(ctx context.Context, arg TicketActionConditionRefRequest) (model.TicketActionConditionRef, error)
	GetTicketActionConditionRef(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionRefInfo, error)
	GetTicketActionConditionRefbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionRefInfo, error)
	ListTicketActionConditionRef(ctx context.Context, arg ListTicketActionConditionRefParams) ([]TicketActionConditionRefInfo, error)
	UpdateTicketActionConditionRef(ctx context.Context, arg TicketActionConditionRefRequest) (model.TicketActionConditionRef, error)
	DeleteTicketActionConditionRef(ctx context.Context, uuid uuid.UUID) error

	CreateTicketItemAssigned(ctx context.Context, arg TicketItemAssignedRequest) (model.TicketItemAssigned, error)
	// GetTicketItemAssigned(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketItemAssignedInfo, error)
	GetTicketItemAssignedbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemAssignedInfo, error)
	ListTicketItemAssigned(ctx context.Context, arg ListTicketItemAssignedParams) ([]TicketItemAssignedInfo, error)
	UpdateTicketItemAssigned(ctx context.Context, arg TicketItemAssignedRequest) (model.TicketItemAssigned, error)
	DeleteTicketItemAssigned(ctx context.Context, uuid uuid.UUID) error
}

var _ QuerierTransaction = (*QueriesTransaction)(nil)
