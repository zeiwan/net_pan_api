package model

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cookie   string `json:"cookie"`
	Token    string `json:"token"`
}
type UserInfo struct {
	NickName string `json:"nickName"`
	Account  string `json:"account"`
}
