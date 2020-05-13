package models

type GetTokenRequest struct {
	GrantType string `json:"grant_type"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}