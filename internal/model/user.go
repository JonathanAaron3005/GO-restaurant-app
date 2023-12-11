package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"-"` //ga akan kasih field ini ketika struct User digunakan sebagai response
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
