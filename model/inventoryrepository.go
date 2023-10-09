package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type InventoryRepository struct {
	Id                  int64          `json:"id"`
	Uuid                uuid.UUID      `json:"uuid"`
	CentralOfficeId     int64          `json:"centralOfficeId"`
	RepositoryCode      string         `json:"repositoryCode"`
	Repository          string         `json:"repository"`
	OfficeId            int64          `json:"officeId"`
	CustodianId         sql.NullInt64  `json:"custodianId"`
	GeographyId         sql.NullInt64  `json:"geographyId"`
	LocationDescription sql.NullString `json:"locationDescription"`
	Remarks             sql.NullString `json:"remarks"`
	OtherInfo           sql.NullString `json:"otherInfo"`
}
