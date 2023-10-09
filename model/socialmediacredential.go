package model

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type SocialProviderType string

const (
	SocialProviderTypeFacebook SocialProviderType = "facebook"
	SocialProviderTypeTwitter  SocialProviderType = "twitter"
	SocialProviderTypeGoogle   SocialProviderType = "google"
)

func (e *SocialProviderType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SocialProviderType(s)
	case string:
		*e = SocialProviderType(s)
	default:
		return fmt.Errorf("unsupported scan type for SocialProviderType: %T", src)
	}
	return nil
}

type SocialMediaCredential struct {
	Uuid         uuid.UUID          `json:"uuid"`
	UserId       int64              `json:"userId"`
	ProviderKey  string             `json:"providerKey"`
	ProviderType SocialProviderType `json:"providerType"`
	OtherInfo    sql.NullString     `json:"otherInfo"`
}
