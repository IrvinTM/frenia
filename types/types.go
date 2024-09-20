package types

type PasswordDB struct {
	Passwords []Password `json:"passwords"`
}

type Password struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}