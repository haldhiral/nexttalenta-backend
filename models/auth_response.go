package models

type AuthResponse struct {
	Username  string `json:"username"`
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
}
