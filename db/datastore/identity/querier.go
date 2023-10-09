package db

import (
	"context"
	// "database/sql"
	"simplebank/model"

	"github.com/google/uuid"
)

type QuerierIdentity interface {
	CreateIdentityInfo(ctx context.Context, arg IdentityInfoRequest) (model.IdentityInfo, error)
	GetIdentityInfo(ctx context.Context, id int64) (IdentityInfoInfo, error)
	GetIdentityInfobyUuId(ctx context.Context, uuid uuid.UUID) (IdentityInfoInfo, error)
	ListIdentityInfo(ctx context.Context, arg ListIdentityInfoParams) ([]IdentityInfoInfo, error)
	UpdateIdentityInfo(ctx context.Context, arg IdentityInfoRequest) (model.IdentityInfo, error)
	DeleteIdentityInfo(ctx context.Context, id int64) error

	CreatePersonalInfo(ctx context.Context, arg PersonalInfoRequest) (model.PersonalInfo, error)
	GetPersonalInfo(ctx context.Context, id int64) (PersonalInfoInfo, error)
	GetPersonalInfobyUuId(ctx context.Context, uuid uuid.UUID) (PersonalInfoInfo, error)
	ListPersonalInfo(ctx context.Context, arg ListPersonalInfoParams) ([]PersonalInfoInfo, error)
	UpdatePersonalInfo(ctx context.Context, arg PersonalInfoRequest) (model.PersonalInfo, error)
	DeletePersonalInfo(ctx context.Context, id int64) error

	CreateAddressList(ctx context.Context, arg AddressListRequest) (model.AddressList, error)
	GetAddressList(ctx context.Context, uuid uuid.UUID) (AddressListInfo, error)
	GetAddressListbyUuId(ctx context.Context, uuid uuid.UUID) (AddressListInfo, error)
	ListAddressList(ctx context.Context, arg ListAddressListParams) ([]AddressListInfo, error)
	UpdateAddressList(ctx context.Context, arg AddressListRequest) (model.AddressList, error)
	DeleteAddressList(ctx context.Context, uuid uuid.UUID) error

	CreateIds(ctx context.Context, arg IdsRequest) (model.Ids, error)
	GetIds(ctx context.Context, uuid uuid.UUID) (IdsInfo, error)
	GetIdsbyUuId(ctx context.Context, uuid uuid.UUID) (IdsInfo, error)
	GetbyIds(ctx context.Context, arg GetbyIdsParams) ([]IdsInfo, error)
	ListIds(ctx context.Context, arg ListIdsParams) ([]IdsInfo, error)
	UpdateIds(ctx context.Context, arg IdsRequest) (model.Ids, error)
	DeleteIds(ctx context.Context, uuid uuid.UUID) error

	CreateEducational(ctx context.Context, arg EducationalRequest) (model.Educational, error)
	GetEducational(ctx context.Context, uuid uuid.UUID) (EducationalInfo, error)
	GetEducationalbyUuId(ctx context.Context, uuid uuid.UUID) (EducationalInfo, error)
	ListEducational(ctx context.Context, arg ListEducationalParams) ([]EducationalInfo, error)
	UpdateEducational(ctx context.Context, arg EducationalRequest) (model.Educational, error)
	DeleteEducational(ctx context.Context, uuid uuid.UUID) error

	CreateEmployment(ctx context.Context, arg EmploymentRequest) (model.Employment, error)
	GetEmployment(ctx context.Context, uuid uuid.UUID) (EmploymentInfo, error)
	GetEmploymentbyUuId(ctx context.Context, uuid uuid.UUID) (EmploymentInfo, error)
	ListEmployment(ctx context.Context, arg ListEmploymentParams) ([]EmploymentInfo, error)
	UpdateEmployment(ctx context.Context, arg EmploymentRequest) (model.Employment, error)
	DeleteEmployment(ctx context.Context, uuid uuid.UUID) error

	CreateOffice(ctx context.Context, arg OfficeRequest) (model.Office, error)
	GetOffice(ctx context.Context, id int64) (OfficeInfo, error)
	GetOfficebyUuId(ctx context.Context, uuid uuid.UUID) (OfficeInfo, error)
	ListOffice(ctx context.Context, arg ListOfficeParams) ([]OfficeInfo, error)
	UpdateOffice(ctx context.Context, arg OfficeRequest) (model.Office, error)
	DeleteOffice(ctx context.Context, id int64) error

	CreateOfficer(ctx context.Context, arg OfficerRequest) (model.Officer, error)
	GetOfficer(ctx context.Context, id int64) (OfficerInfo, error)
	GetOfficerbyUuid(ctx context.Context, uuid uuid.UUID) (OfficerInfo, error)
	ListOfficer(ctx context.Context, arg ListOfficerParams) ([]OfficerInfo, error)
	UpdateOfficer(ctx context.Context, arg OfficerRequest) (model.Officer, error)
	DeleteOfficer(ctx context.Context, uuid uuid.UUID) error

	CreateEmployee(ctx context.Context, arg EmployeeRequest) (model.Employee, error)
	GetEmployee(ctx context.Context, id int64) (EmployeeInfo, error)
	GetEmployeebyUuId(ctx context.Context, uuid uuid.UUID) (EmployeeInfo, error)
	ListEmployee(ctx context.Context, arg ListEmployeeParams) ([]EmployeeInfo, error)
	UpdateEmployee(ctx context.Context, arg EmployeeRequest) (model.Employee, error)
	DeleteEmployee(ctx context.Context, id int64) error

	CreateIncomeSource(ctx context.Context, arg IncomeSourceRequest) (model.IncomeSource, error)
	GetIncomeSource(ctx context.Context, uuid uuid.UUID) (IncomeSourceInfo, error)
	GetIncomeSourcebyUuId(ctx context.Context, uuid uuid.UUID) (IncomeSourceInfo, error)
	ListIncomeSource(ctx context.Context, arg ListIncomeSourceParams) ([]IncomeSourceInfo, error)
	UpdateIncomeSource(ctx context.Context, arg IncomeSourceRequest) (model.IncomeSource, error)
	DeleteIncomeSource(ctx context.Context, uuid uuid.UUID) error

	CreateContact(ctx context.Context, arg ContactRequest) (model.Contact, error)
	GetContact(ctx context.Context, uuid uuid.UUID) (ContactInfo, error)
	GetContactbyUuId(ctx context.Context, uuid uuid.UUID) (ContactInfo, error)
	ListContact(ctx context.Context, arg ListContactParams) ([]ContactInfo, error)
	UpdateContact(ctx context.Context, arg ContactRequest) (model.Contact, error)
	DeleteContact(ctx context.Context, uuid uuid.UUID) error

	CreateRelation(ctx context.Context, arg RelationRequest) (model.Relation, error)
	GetRelation(ctx context.Context, uuid uuid.UUID) (RelationInfo, error)
	GetRelationbyUuId(ctx context.Context, uuid uuid.UUID) (RelationInfo, error)
	ListRelation(ctx context.Context, arg ListRelationParams) ([]RelationInfo, error)
	UpdateRelation(ctx context.Context, arg RelationRequest) (model.Relation, error)
	DeleteRelation(ctx context.Context, uuid uuid.UUID) error

	CreateIdentitySpecsDate(ctx context.Context, arg IdentitySpecsDateRequest) (model.IdentitySpecsDate, error)
	GetIdentitySpecsDate(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsDateInfo, error)
	GetIdentitySpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsDateInfo, error)
	ListIdentitySpecsDate(ctx context.Context, arg ListIdentitySpecsDateParams) ([]IdentitySpecsDateInfo, error)
	UpdateIdentitySpecsDate(ctx context.Context, arg IdentitySpecsDateRequest) (model.IdentitySpecsDate, error)
	DeleteIdentitySpecsDate(ctx context.Context, uuid uuid.UUID) error

	CreateIdentitySpecsString(ctx context.Context, arg IdentitySpecsStringRequest) (model.IdentitySpecsString, error)
	GetIdentitySpecsString(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsStringInfo, error)
	GetIdentitySpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsStringInfo, error)
	ListIdentitySpecsString(ctx context.Context, arg ListIdentitySpecsStringParams) ([]IdentitySpecsStringInfo, error)
	UpdateIdentitySpecsString(ctx context.Context, arg IdentitySpecsStringRequest) (model.IdentitySpecsString, error)
	DeleteIdentitySpecsString(ctx context.Context, uuid uuid.UUID) error

	CreateIdentitySpecsNumber(ctx context.Context, arg IdentitySpecsNumberRequest) (model.IdentitySpecsNumber, error)
	GetIdentitySpecsNumber(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsNumberInfo, error)
	GetIdentitySpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsNumberInfo, error)
	ListIdentitySpecsNumber(ctx context.Context, arg ListIdentitySpecsNumberParams) ([]IdentitySpecsNumberInfo, error)
	UpdateIdentitySpecsNumber(ctx context.Context, arg IdentitySpecsNumberRequest) (model.IdentitySpecsNumber, error)
	DeleteIdentitySpecsNumber(ctx context.Context, uuid uuid.UUID) error

	CreateIdentitySpecsRef(ctx context.Context, arg IdentitySpecsRefRequest) (model.IdentitySpecsRef, error)
	GetIdentitySpecsRef(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsRefInfo, error)
	GetIdentitySpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsRefInfo, error)
	ListIdentitySpecsRef(ctx context.Context, arg ListIdentitySpecsRefParams) ([]IdentitySpecsRefInfo, error)
	UpdateIdentitySpecsRef(ctx context.Context, arg IdentitySpecsRefRequest) (model.IdentitySpecsRef, error)
	DeleteIdentitySpecsRef(ctx context.Context, uuid uuid.UUID) error
}

var _ QuerierIdentity = (*QueriesIdentity)(nil)
