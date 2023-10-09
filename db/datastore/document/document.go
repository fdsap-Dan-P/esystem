package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
	"simplebank/util"
	"simplebank/util/images"
)

const createDocument = `-- name: CreateDocument: one
INSERT INTO Documents (
  Uuid, Code, Server_Id, File_Path, Doc_Date, Thumbnail, DocType_Id, 
  Description, Other_Info) 
SELECT $1, $2, s.ID, $4, $5, $6, $7, $8, $9
FROM Server s WHERE lower(s.Code) = lower($3)
ON CONFLICT (Code) DO UPDATE SET
  Code = EXCLUDED.Code, 
  Server_Id = EXCLUDED.Server_Id, 
  File_Path = EXCLUDED.File_Path, 
  Doc_Date = EXCLUDED.Doc_Date, 
  Thumbnail = EXCLUDED.Thumbnail, 
  DocType_Id = EXCLUDED.DocType_Id, 
  Description = EXCLUDED.Description, 
  Other_Info = EXCLUDED.Other_Info
RETURNING 
  Id, UUID, Code, Server_Id, File_Path, Doc_Date, Thumbnail, DocType_Id, 
  Description, Other_Info
`

type DocumentRequest struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	Code        string         `json:"code"`
	ServerCode  string         `json:"serverCode"`
	FilePath    string         `json:"filePath"`
	DocDate     sql.NullTime   `json:"docDate"`
	Thumbnail   []byte         `json:"thumbnail"`
	DoctypeId   int64          `json:"doctypeId"`
	Description sql.NullString `json:"description"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesDocument) AddDocument(
	ctx context.Context, arg DocumentRequest, sql string) (model.Document, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid = util.UUID()
	}
	row := q.db.QueryRowContext(ctx, sql,
		arg.Uuid,
		arg.Code,
		arg.ServerCode,
		arg.FilePath,
		arg.DocDate,
		arg.Thumbnail,
		arg.DoctypeId,
		arg.Description,
		arg.OtherInfo,
	)
	var i model.Document
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ServerId,
		&i.FilePath,
		&i.DocDate,
		&i.Thumbnail,
		&i.DoctypeId,
		&i.Description,
		&i.OtherInfo,
	)
	return i, err
}

func (q *QueriesDocument) CreateDocumentImageFromURL(
	ctx context.Context, arg DocumentRequest, url string) (model.Document, error) {

	dc, dcErr := q.GetDocumentbyCode(ctx, arg.Code)
	if dcErr == nil {
		arg.Uuid = dc.Uuid
	} else {
		if arg.Uuid == uuid.Nil {
			arg.Uuid = util.UUID()
		}
	}

	// targetPath, _, _ := file.FileSpecs(arg.FilePath)

	svr, svrErr := q.GetServerbyCode(ctx, arg.ServerCode)
	if svrErr != nil {
		return model.Document{}, svrErr
	}

	img := images.NewImageDocumentFromURL(url, svr.HomePath, arg.Uuid, arg.FilePath)
	imgErr := img.CreateThumbnail()
	if imgErr != nil {
		return model.Document{}, imgErr
	}
	arg.FilePath = img.ImageData().NewFile.FullPath()

	return q.AddDocument(ctx, arg, createDocument)
}

func (q *QueriesDocument) CreateDocument(ctx context.Context, arg DocumentRequest) (model.Document, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid = util.UUID()
	}
	return q.AddDocument(ctx, arg, createDocument)
}

const deleteDocument = `-- name: DeleteDocument :exec
DELETE FROM Documents
WHERE id = $1
`

func (q *QueriesDocument) DeleteDocument(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteDocument, id)
	return err
}

type DocumentInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	Code         string         `json:"code"`
	ServerId     int64          `json:"serverId"`
	ServerCode   string         `json:"serverCode"`
	FilePath     string         `json:"filePath"`
	HomePath     string         `json:"homePath"`
	DocDate      sql.NullTime   `json:"docDate"`
	Thumbnail    []byte         `json:"thumbnail"`
	DoctypeId    int64          `json:"doctypeId"`
	DoctypeCode  string         `json:"doctypeCode"`
	DoctypeTitle string         `json:"doctypeTitle"`
	Description  sql.NullString `json:"description"`
	OtherInfo    sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getDocument = `-- name: GetDocument :one
SELECT 
  d.Id, mr.UUId, d.Code, s.Id Server_Id, s.Code Server_Code, 
  d.File_Path, s.HomePath, d.Doc_Date, d.Thumbnail, d.DocType_Id, 
  typ.code DocType_Code, typ.title DocType_Title, d.Description, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Documents d 
INNER JOIN Server s on d.Server_Id = s.Id
INNER JOIN Reference typ on d.DocType_Id = typ.Id
INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateDocument(q *QueriesDocument, ctx context.Context, sql string, param ...interface{}) ([]DocumentInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DocumentInfo{}
	for rows.Next() {
		var i DocumentInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.ServerId,
			&i.ServerCode,
			&i.FilePath,
			&i.HomePath,
			&i.DocDate,
			&i.Thumbnail,
			&i.DoctypeId,
			&i.DoctypeCode,
			&i.DoctypeTitle,
			&i.Description,
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

func (q *QueriesDocument) ListDocument(ctx context.Context, arg ListDocumentParams) ([]DocumentInfo, error) {
	sql := fmt.Sprintf(`%v WHERE DocType_Id = $1 ORDER BY id LIMIT $2 OFFSET $3 `, getDocument)
	return populateDocument(q, ctx, sql, arg.DocTypeId, arg.Limit, arg.Offset)
}

func (q *QueriesDocument) GetDocument(ctx context.Context, id int64) (DocumentInfo, error) {
	sql := fmt.Sprintf(`%v WHERE d.id = $1 `, getDocument)
	items, err := populateDocument(q, ctx, sql, id)

	for _, val := range items {
		return val, err
	}
	return DocumentInfo{}, fmt.Errorf("schedule UUID:%v not found", id)
}

func (q *QueriesDocument) GetDocumentbyUuid(ctx context.Context, uuid uuid.UUID) (DocumentInfo, error) {
	sql := fmt.Sprintf(`%v WHERE mr.UUID = $1 `, getDocument)
	items, err := populateDocument(q, ctx, sql, uuid)

	for _, val := range items {
		return val, err
	}
	return DocumentInfo{}, fmt.Errorf("schedule UUID:%v not found", uuid)
}

func (q *QueriesDocument) GetDocumentbyCode(ctx context.Context, code string) (DocumentInfo, error) {
	sql := fmt.Sprintf(`%v WHERE d.Code = $1 `, getDocument)
	items, err := populateDocument(q, ctx, sql, code)

	for _, val := range items {
		return val, err
	}
	return DocumentInfo{}, fmt.Errorf("schedule UUID:%v not found", code)
}

type ListDocumentParams struct {
	DocTypeId int64 `json:"docTypeId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

const updateDocument = `-- name: UpdateDocument :one
UPDATE Documents SET 
  Code = $2,
  Server_Id = s.Id,
  File_Path = $4,
  Doc_Date = $5,
  Thumbnail = $6,
  DocType_Id = $7,
  Description = $8,
  Other_Info = $9
FROM Server s 
WHERE Documents.uuid = $1 and s.Code = $3
RETURNING 
  Documents.Id, Documents.UUID, Documents.Code, Documents.Server_Id, 
  Documents.File_Path, Documents.Doc_Date, Documents.Thumbnail, Documents.DocType_Id, 
  Documents.Description, Documents.Other_Info
`

func (q *QueriesDocument) UpdateDocument(ctx context.Context, arg DocumentRequest) (model.Document, error) {
	return q.AddDocument(ctx, arg, updateDocument)
}

func (q *QueriesDocument) LoadData(ctx context.Context, tableName string, filepathname string, brCode string) error {
	_, err := q.db.ExecContext(ctx,
		`select loaddata ($1, $2, $3)`, tableName, filepathname, brCode)
	return err
}
