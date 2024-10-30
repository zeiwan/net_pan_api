package core

import (
	"core/model"
	"github.com/imroc/req/v3"
)

type NetPan interface {
	// AuthLogin 获取连接池
	AuthLogin(account model.Account) (*req.Client, error)
	// UserInfo 获取用户信息
	UserInfo(r *req.Client) (model.UserInfo, error)
}

var PanMap = map[string]NetPan{}

// RegisterPan 注册连接池
func RegisterPan(mode string, pan NetPan) {
	PanMap[mode] = pan
}
func GetPan(mode string) (pan NetPan, ok bool) {
	pan, ok = PanMap[mode]
	return
}
