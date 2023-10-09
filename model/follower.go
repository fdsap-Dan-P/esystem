package model

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	Uuid         uuid.UUID `json:"uuid"`
	UserId       int64     `json:"userId"`
	FollowerId   int64     `json:"followerId"`
	DateFollowed time.Time `json:"dateFollowed"`
	IsFollower   bool      `json:"isFollower"`
}
