package types

type PasswordDB struct {
	Passwords map[string]string `json:"passwords"`
}