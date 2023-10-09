package model

type ReferenceGroup struct {
	GroupId int64  `json:"groupId"`
	RefId   int64  `json:"refId"`
	Title   string `json:"title"`
	Active  bool   `json:"active"`
}
