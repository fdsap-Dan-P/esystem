package db

import (
	"context"
	"simplebank/model"

	"github.com/google/uuid"
)

// var QueriesAccount *account.QueriesAccount = account.New(testDB)

type QuerierDocument interface {
	LoadData(ctx context.Context, tableName string, filepathname string, brCode string) error

	CreateServer(ctx context.Context, arg ServerRequest) (model.Server, error)
	GetServer(ctx context.Context, id int64) (ServerInfo, error)
	GetServerbyCode(ctx context.Context, code string) (ServerInfo, error)
	GetServerbyUuId(ctx context.Context, uuid uuid.UUID) (ServerInfo, error)
	ListServer(ctx context.Context, arg ListServerParams) ([]ServerInfo, error)
	UpdateServer(ctx context.Context, arg ServerRequest) (model.Server, error)
	DeleteServer(ctx context.Context, id int64) error

	CreateDocumentImageFromURL(ctx context.Context, arg DocumentRequest, url string) (model.Document, error)
	GetDocumentbyCode(ctx context.Context, code string) (DocumentInfo, error)
	CreateDocument(ctx context.Context, arg DocumentRequest) (model.Document, error)
	GetDocument(ctx context.Context, id int64) (DocumentInfo, error)
	GetDocumentbyUuid(ctx context.Context, uuid uuid.UUID) (DocumentInfo, error)
	ListDocument(ctx context.Context, arg ListDocumentParams) ([]DocumentInfo, error)
	UpdateDocument(ctx context.Context, arg DocumentRequest) (model.Document, error)
	DeleteDocument(ctx context.Context, id int64) error
}

var _ QuerierDocument = (*QueriesDocument)(nil)
