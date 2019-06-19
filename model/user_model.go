package model

type User struct {
	UserID string `json:user_id`
	Password string `json:password`
	AccessToken string
}