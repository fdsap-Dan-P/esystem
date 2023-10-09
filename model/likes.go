package model

import (
	"time"

	"github.com/google/uuid"
)

type Likes struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserId    int64     `json:"userId"`
	Mood      int32     `json:"mood"`
	DateLiked time.Time `json:"dateLiked"`
}
