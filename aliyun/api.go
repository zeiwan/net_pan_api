package aliyun

import (
	"github.com/ZeiWan/NetPanSDK/model"
	"github.com/imroc/req/v3"
)

func (c CloudALi) AuthLogin(account model.Account) (*req.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) NewClient(client *req.Client) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) UserInfo() (model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetShareInfo(url, pwd string) (model.ShareInfoResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageFolderList(resp model.ShareInfoResp) ([]model.SharePageFolderListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageFileList(resp model.ShareInfoResp) ([]model.SharePageFileListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageAll(resp model.ShareInfoResp) (model.SharePageALL, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) CreateFolder(parentFolderId, folderName string) (model.CreateFolderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetMyFolder(id string) ([]model.MyFolderListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Rename(folderId, newFolderName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Delete(taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Move(targetFolderId string, taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Copy(targetFolderId string, taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetMyFileAll(id string) (model.MyDirAll, error) {
	//TODO implement me
	panic("implement me")
}
