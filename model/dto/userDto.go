package dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Age         string `json:"age"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PassConfirm string `json:"passConfirm"`
}
