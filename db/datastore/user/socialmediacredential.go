package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createSocialMediaCredential = `-- name: CreateSocialMediaCredential: one
INSERT INTO Social_Media_Credential (
User_Id, Provider_Key, Provider_Type, Other_Info
) VALUES (
$1, $2, $3, $4
) RETURNING UUId, User_Id, Provider_Key, Provider_Type, Other_Info
`

type SocialMediaCredentialRequest struct {
	Uuid         uuid.UUID                `json:"uuid"`
	UserId       int64                    `json:"userId"`
	ProviderKey  string                   `json:"ProviderKey"`
	ProviderType model.SocialProviderType `json:"ProviderType"`
	OtherInfo    sql.NullString           `json:"otherInfo"`
}

func (q *QueriesUser) CreateSocialMediaCredential(ctx context.Context, arg SocialMediaCredentialRequest) (model.SocialMediaCredential, error) {
	row := q.db.QueryRowContext(ctx, createSocialMediaCredential,
		arg.UserId,
		arg.ProviderKey,
		arg.ProviderType,
		arg.OtherInfo,
	)
	var i model.SocialMediaCredential
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProviderKey,
		&i.ProviderType,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSocialMediaCredential = `-- name: DeleteSocialMediaCredential :exec
DELETE FROM Social_Media_Credential
WHERE uuid = $1
`

func (q *QueriesUser) DeleteSocialMediaCredential(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSocialMediaCredential, uuid)
	return err
}

type SocialMediaCredentialInfo struct {
	Uuid         uuid.UUID                `json:"uuid"`
	UserId       int64                    `json:"userId"`
	ProviderKey  string                   `json:"ProviderKey"`
	ProviderType model.SocialProviderType `json:"ProviderType"`
	OtherInfo    sql.NullString           `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getSocialMediaCredential = `-- name: GetSocialMediaCredential :one
SELECT 
mr.UUId, User_Id, Provider_Key, Provider_Type, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Social_Media_Credential d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesUser) GetSocialMediaCredential(ctx context.Context, uuid uuid.UUID) (SocialMediaCredentialInfo, error) {
	row := q.db.QueryRowContext(ctx, getSocialMediaCredential, uuid)
	var i SocialMediaCredentialInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProviderKey,
		&i.ProviderType,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getSocialMediaCredentialbyUuId = `-- name: GetSocialMediaCredentialbyUuId :one
SELECT 
mr.UUId, User_Id, Provider_Key, Provider_Type, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Social_Media_Credential d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetSocialMediaCredentialbyUuId(ctx context.Context, uuid uuid.UUID) (SocialMediaCredentialInfo, error) {
	row := q.db.QueryRowContext(ctx, getSocialMediaCredentialbyUuId, uuid)
	var i SocialMediaCredentialInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProviderKey,
		&i.ProviderType,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listSocialMediaCredential = `-- name: ListSocialMediaCredential:many
SELECT 
mr.UUId, User_Id, Provider_Key, Provider_Type, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Social_Media_Credential d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListSocialMediaCredentialParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListSocialMediaCredential(ctx context.Context, arg ListSocialMediaCredentialParams) ([]SocialMediaCredentialInfo, error) {
	rows, err := q.db.QueryContext(ctx, listSocialMediaCredential, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SocialMediaCredentialInfo{}
	for rows.Next() {
		var i SocialMediaCredentialInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.ProviderKey,
			&i.ProviderType,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSocialMediaCredential = `-- name: UpdateSocialMediaCredential :one
UPDATE Social_Media_Credential SET 
User_Id = $2,
Provider_Key = $3,
Provider_Type = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, User_Id, Provider_Key, Provider_Type, Other_Info
`

func (q *QueriesUser) UpdateSocialMediaCredential(ctx context.Context, arg SocialMediaCredentialRequest) (model.SocialMediaCredential, error) {
	row := q.db.QueryRowContext(ctx, updateSocialMediaCredential,

		arg.Uuid,
		arg.UserId,
		arg.ProviderKey,
		arg.ProviderType,
		arg.OtherInfo,
	)
	var i model.SocialMediaCredential
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ProviderKey,
		&i.ProviderType,
		&i.OtherInfo,
	)
	return i, err
}
