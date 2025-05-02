package model

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type RegisterRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
