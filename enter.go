package core

import (
	_189 "core/189"
	"core/model"
	"github.com/imroc/req/v3"
)

type NetPan interface {
	// AuthLogin 获取注入Client
	AuthLogin(account model.Account) (*req.Client, error)
	// NewClient  注入Client
	NewClient(client *req.Client)
	//	UserInfo 获取用户信息
	UserInfo() (model.UserInfo, error)
	//	GetShareInfo 获取分享链接详情
	GetShareInfo(url, pwd string) (model.ShareInfoResp, error)
	//	GetSharePageFolderList  获取分享页面文件夹
	GetSharePageFolderList(model.ShareInfoResp) ([]model.SharePageFolderListResp, error)
	//	GetShareNoteFolderList 	获取分享页面文件
	GetShareNoteFileList(model.ShareInfoResp) ([]model.SharePageFileListResp, error)
	// CreateFolder 创建文件夹
	CreateFolder(parentFolderId, folderName string) (model.CreateFolderResp, error)
	// GetMyFolder	获取我的目录
	GetMyFolder(id string) ([]model.MyFolderListResp, error)
}

func NewCloud189() NetPan {
	return &_189.Cloud189{}
}
