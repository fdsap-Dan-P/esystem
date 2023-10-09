// Package SocialMedia implements common functionality needed for social media web applications.
package model

import (
	"database/sql"

	"github.com/google/uuid"
)

//go:generate stringer -type=MoodState
type MoodState int

// All possible mood states.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

// Post represents a Social Media Post type.
// swagger:model Post
type Post struct {
	Uuid         uuid.UUID      `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"String"`
	UserId       int64          `json:"userId" example:"name@mail.com" format:"String"`
	Caption      string         `json:"caption"`
	MessageBody  string         `json:"messageBody"`
	Url          string         `json:"url"`
	ImageUri     string         `json:"imageUri"`
	ThumbnailUri string         `json:"thumbnailUri"`
	Keywords     []string       `json:"keywords"`
	Likers       []int64        `json:"likers"`
	Mood         MoodState      `json:"mood"`
	MoodEmoji    string         `json:"moodEmoji"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood states.
var Moods map[string]MoodState
var MoodsEmoji map[MoodState]string

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}

	MoodsEmoji = map[MoodState]string{MoodStateNeutral: "\xF0\x9F\x98\x90", MoodStateHappy: "\xF0\x9F\x98\x8A", MoodStateSad: "\xF0\x9F\x98\x9E", MoodStateAngry: "\xF0\x9F\x98\xA0", MoodStateHopeful: "\xF0\x9F\x98\x8C", MoodStateThrilled: "\xF0\x9F\x98\x81", MoodStateBored: "\xF0\x9F\x98\xB4", MoodStateShy: "\xF0\x9F\x98\xB3", MoodStateComical: "\xF0\x9F\x98\x9C", MoodStateOnCloudNine: "\xF0\x9F\x98\x82"}

}
