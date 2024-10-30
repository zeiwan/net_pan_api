package test

import (
	"core"
	"core/model"
	"fmt"
	"testing"
)

var c189 = core.NewCloud189()

func login() {
	var obj model.Account
	obj.Username = "15315496321"
	obj.Password = "mogen123"

	client, err := c189.AuthLogin(obj)
	if err != nil {
		panic(err)
	}
	c189.NewClient(client)

	return
}
func TestLogin(t *testing.T) {
	login()
}
func TestUserInfo(t *testing.T) {
	login()
	resp, err := c189.UserInfo()

	fmt.Println(resp, err)
}
func TestShare(t *testing.T) {
	login()
	resp, err := c189.GetShareInfo("https://cloud.189.cn/web/share?code=eUb2IfYnUbYb", "s5ir")
	fmt.Println(resp, err)
	// 获取当前页面文件夹
	folders, err := c189.GetSharePageFolderList(resp)
	fmt.Println(folders, err)
	// 获取分享目录节点内容
	resp.FileId = "724901134144036403"
	files, err := c189.GetShareNoteFileList(resp)
	fmt.Println(files, err)
}
