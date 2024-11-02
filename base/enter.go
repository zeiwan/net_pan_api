package base

import (
	_189 "github.com/ZeiWan/NetPanSDK/189"
	"github.com/ZeiWan/NetPanSDK/aliyun"
	"github.com/ZeiWan/NetPanSDK/module"
	"github.com/ZeiWan/NetPanSDK/quark"

	"github.com/imroc/req/v3"
)

type NetPan interface {
	// AuthLogin 获取注入Client
	AuthLogin(account module.Account) (*req.Client, error)
	// NewClient  注入Client
	NewClient(client *req.Client)
	//	UserInfo 获取用户信息
	UserInfo() (module.UserInfo, error)
	//	GetShareInfo 获取分享链接详情
	GetShareInfo(url, pwd string) (module.ShareInfoResp, error)
	//	GetSharePageFolderList  获取分享页面文件夹
	GetSharePageFolderList(module.ShareInfoResp) ([]module.SharePageFolderListResp, error)
	//	GetShareNoteFolderList 	获取分享页面文件
	GetSharePageFileList(module.ShareInfoResp) ([]module.SharePageFileListResp, error)
	// GetSharePageAll  获取分享页面所有内容
	GetSharePageAll(module.ShareInfoResp) (module.SharePageALL, error)
	// CreateFolder 创建文件夹
	CreateFolder(parentFolderId, folderName string) (module.CreateFolderResp, error)
	// GetMyFolder	获取我的目录
	GetMyFolder(id string) ([]module.MyFolderListResp, error)
	GetMyFileAll(id string) (module.MyDirAll, error)
	//	GetMyFolderAll 获取我的目录所有文件
	//GetMyFolderAll(id string) (model.MyFolderALL, error)
	////	ShareLink 生成分享链接
	//ShareLink(fileId string, expireTime, shareType uint8) (string, error)
	//	Rename 重命名文件
	Rename(folderId, newFolderName string) (bool, error)
	// Delete 删除文件夹
	Delete(taskInfos []module.TaskInfosReq) error
	// Move 移动文件
	Move(targetFolderId string, taskInfos []module.TaskInfosReq) error
	// Copy 复制文件
	Copy(targetFolderId string, taskInfos []module.TaskInfosReq) error
}

func NewCloud189() NetPan {
	return &_189.Cloud189{}
}
func NewCloudALi() NetPan {
	return &aliyun.CloudALi{}
}

func NewCloudQuark() NetPan {
	return &quark.CloudQuark{}
}
