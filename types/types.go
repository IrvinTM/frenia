package types

type PasswordDB struct {
	Passwords map[string]string `json:"passwords"`
}

type Configuration struct {
	Salt string `json:"salt"`
	Mode string `json:"mode"`
}
