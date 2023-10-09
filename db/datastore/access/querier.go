package db

import (
	"context"

	"github.com/google/uuid"

	"simplebank/model"
)

type QuerierAccess interface {
	CreateAccessRole(ctx context.Context, arg AccessRoleRequest) (model.AccessRole, error)
	GetAccessRole(ctx context.Context, id int64) (AccessRoleInfo, error)
	GetAccessRolebyUuId(ctx context.Context, uuid uuid.UUID) (AccessRoleInfo, error)
	GetAccessRolebyName(ctx context.Context, name string) (AccessRoleInfo, error)
	ListAccessRole(ctx context.Context, arg ListAccessRoleParams) ([]AccessRoleInfo, error)
	UpdateAccessRole(ctx context.Context, arg AccessRoleRequest) (model.AccessRole, error)
	DeleteAccessRole(ctx context.Context, id int64) error

	CreateAccessProduct(ctx context.Context, arg AccessProductRequest) (model.AccessProduct, error)
	GetAccessProduct(ctx context.Context, uuid uuid.UUID) (AccessProductInfo, error)
	GetAccessProductbyUuId(ctx context.Context, uuid uuid.UUID) (AccessProductInfo, error)
	ListAccessProduct(ctx context.Context, arg ListAccessProductParams) ([]AccessProductInfo, error)
	UpdateAccessProduct(ctx context.Context, arg AccessProductRequest) (model.AccessProduct, error)
	DeleteAccessProduct(ctx context.Context, uuid uuid.UUID) error

	CreateAccessAccountType(ctx context.Context, arg AccessAccountTypeRequest) (model.AccessAccountType, error)
	GetAccessAccountType(ctx context.Context, uuid uuid.UUID) (AccessAccountTypeInfo, error)
	GetAccessAccountTypebyUuId(ctx context.Context, uuid uuid.UUID) (AccessAccountTypeInfo, error)
	ListAccessAccountType(ctx context.Context, arg ListAccessAccountTypeParams) ([]AccessAccountTypeInfo, error)
	UpdateAccessAccountType(ctx context.Context, arg AccessAccountTypeRequest) (model.AccessAccountType, error)
	DeleteAccessAccountType(ctx context.Context, uuid uuid.UUID) error

	CreateAccessObject(ctx context.Context, arg AccessObjectRequest) (model.AccessObject, error)
	GetAccessObject(ctx context.Context, uuid uuid.UUID) (AccessObjectInfo, error)
	GetAccessObjectbyUuId(ctx context.Context, uuid uuid.UUID) (AccessObjectInfo, error)
	ListAccessObject(ctx context.Context, arg ListAccessObjectParams) ([]AccessObjectInfo, error)
	UpdateAccessObject(ctx context.Context, arg AccessObjectRequest) (model.AccessObject, error)
	DeleteAccessObject(ctx context.Context, uuid uuid.UUID) error
}

var _ QuerierAccess = (*QueriesAccess)(nil)
