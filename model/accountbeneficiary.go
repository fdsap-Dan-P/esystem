package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type AccountBeneficiary struct {
	Uuid               uuid.UUID      `json:"uuid"`
	AccountId          int64          `json:"accountId"`
	Iiid               int64          `json:"iiid"`
	Series             int16          `json:"series"`
	BeneficiaryTypeId  int64          `json:"beneficiaryTypeId"`
	RelationshipTypeId int64          `json:"relationshipTypeId"`
	OtherInfo          sql.NullString `json:"otherInfo"`
}
