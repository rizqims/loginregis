package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}
