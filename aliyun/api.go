package aliyun

import (
	"github.com/imroc/req/v3"
	"github.com/zeiwan/net_pan_api/module"
)

func (c CloudALi) SaveFile(taskReq module.TaskShareReq, taskInfos []module.TaskInfosReq) error {
	panic("implement me")
}
func (c CloudALi) AuthLogin(account module.Account) (*req.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) NewClient(client *req.Client) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) UserInfo() (module.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetShareInfo(url, pwd string) (module.ShareInfoResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageFolderList(resp module.ShareInfoResp) ([]module.SharePageFolderListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageFileList(resp module.ShareInfoResp) ([]module.SharePageFileListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetSharePageAll(resp module.ShareInfoResp) (module.SharePageALL, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) CreateFolder(parentFolderId, folderName string) (module.CreateFolderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetMyFolder(id string) ([]module.MyFolderListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Rename(folderId, newFolderName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Delete(taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Move(targetFolderId string, taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) Copy(targetFolderId string, taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CloudALi) GetMyFileAll(id string) (module.MyFolderAll, error) {
	//TODO implement me
	panic("implement me")
}
