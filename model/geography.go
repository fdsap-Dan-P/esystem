package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Geography struct {
	Id                    int64               `json:"id"`
	Uuid                  uuid.UUID           `json:"uuid"`
	Code                  int64               `json:"code"`
	ShortName             sql.NullString      `json:"shortName"`
	Location              string              `json:"location"`
	TypeId                int64               `json:"typeId"`
	ParentId              sql.NullInt64       `json:"parentId"`
	ZipCode               sql.NullString      `json:"zipCode"`
	Latitude              decimal.NullDecimal `json:"latitude"`
	Longitude             decimal.NullDecimal `json:"longitude"`
	AddressUrl            sql.NullString      `json:"addressUrl"`
	SimpleLocation        sql.NullString      `json:"simpleLocation"`
	FullLocation          sql.NullString      `json:"fullLocation"`
	VecSimpleLocation     sql.NullString      `json:"vecSimpleLocation"`
	VecFullLocation       sql.NullString      `json:"vecFullLocation"`
	VecFullSimpleLocation sql.NullString      `json:"vecFullSimpleLocation"`
	OtherInfo             sql.NullString      `json:"otherInfo"`
}
