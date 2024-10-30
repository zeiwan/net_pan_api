package test

import (
	"core"
	"core/model"
	"fmt"
	"testing"
)

var (
	pan core.NetPan
	//client *req.Client
)

func login() {
	var obj model.Account
	obj.Username = "15315496321"
	obj.Password = "mogen123"

	pan, _ = core.GetPan("189")

	client, _ := pan.AuthLogin(obj)
	resp, err := pan.UserInfo(client)
	fmt.Println(resp, err)
	return
}
func TestLogin(t *testing.T) {
	login()

}
func TestUserInfo(t *testing.T) {
	//login()
	//fmt.Println(client)
	//resp, err := pan.UserInfo(client)
	//fmt.Println(resp, err)
}
