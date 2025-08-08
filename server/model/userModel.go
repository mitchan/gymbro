package model

type CreateUser struct {
	Username string
	Email    string
	Password string
}

type ResponseLoginUser struct {
	AccessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}
