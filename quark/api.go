package quark

import (
	"fmt"
	"github.com/ZeiWan/NetPanSDK/model"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
)

func (c *CloudQuark) AuthLogin(account model.Account) (*req.Client, error) {
	return c.core.login(account)
}

func (c *CloudQuark) NewClient(client *req.Client) {
	c.core.invoker = invoker{client: client}
}

func (c *CloudQuark) UserInfo() (model.UserInfo, error) {
	return c.core.userInfo()
}

func (c *CloudQuark) GetShareInfo(url, pwd string) (model.ShareInfoResp, error) {
	code := parseShareCode(url)
	return c.core.shareInfo(code, pwd)
}

func (c *CloudQuark) GetSharePageFolderList(info model.ShareInfoResp) (resp []model.SharePageFolderListResp, err error) {
	m, err := c.core.shareDetail(info)
	fmt.Println(resp, err)

	for _, s := range m.Data.List {
		resp = append(resp, model.SharePageFolderListResp{
			Id:       s.Fid,
			Name:     s.FileName,
			IsFolder: cast.ToUint8(s.FileType),
			ParentId: s.Fid,
			SToken:   s.ShareFidToken,
		})
	}

	return resp, nil
}

func (c *CloudQuark) GetSharePageFileList(resp model.ShareInfoResp) ([]model.SharePageFileListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) GetSharePageAll(resp model.ShareInfoResp) (model.SharePageALL, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) CreateFolder(parentFolderId, folderName string) (model.CreateFolderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) GetMyFolder(id string) ([]model.MyFolderListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Rename(folderId, newFolderName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Delete(taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Move(targetFolderId string, taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Copy(targetFolderId string, taskInfos []model.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}
func (c *CloudQuark) GetMyFileAll(id string) (model.MyFileListResp, error) {
	//TODO implement me
	panic("implement me")
}
