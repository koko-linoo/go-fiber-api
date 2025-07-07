package models

type LoginRequest struct {
	Username string `json:"username" validator:"required"`
	Password string `json:"password" validator:"required"`
}
