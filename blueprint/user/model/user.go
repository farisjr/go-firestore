package model

type Users struct {
	UserId int    `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
}
