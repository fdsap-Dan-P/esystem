package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Connectivity int16

const (
	Local   Connectivity = 0
	Network Connectivity = 1
	Service Connectivity = 2
)

func (s Connectivity) String() string {
	switch s {
	case Local:
		return "Local"
	case Network:
		return "Network"
	case Service:
		return "Service"
	}
	return "unknown"
}

type Server struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	Code         string         `json:"code"`
	Connectivity Connectivity   `json:"connectivity"`
	NetAddress   string         `json:"netAddress"`
	Certificate  sql.NullString `json:"certificate"`
	HomePath     string         `json:"homePath"`
	Description  sql.NullString `json:"description"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}
