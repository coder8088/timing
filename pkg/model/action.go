package model

type Action struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	GroupId int64  `json:"group_id"`
}
