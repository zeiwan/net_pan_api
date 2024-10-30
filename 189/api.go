package _189

import (
	"core/model"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
)

func (c *Cloud189) GetMyFolder(id string) (resp []model.MyFolderListResp, err error) {
	resp, err = c.core.getMyFolder(id)
	return
}
func (c *Cloud189) CreateFolder(parentFolderId, folderName string) (resp model.CreateFolderResp, err error) {
	resp, err = c.core.createFolder(parentFolderId, folderName)
	return
}
func (c *Cloud189) GetShareNoteFileList(req model.ShareInfoResp) (list []model.SharePageFileListResp, err error) {
	resp, err := c.core.shareFolderList(req)
	for _, f := range resp.FileListAO.FileList {
		list = append(list, model.SharePageFileListResp{
			Id:   cast.ToString(f.Id),
			Name: f.Name,
		})
	}
	return
}
func (c *Cloud189) GetSharePageFolderList(req model.ShareInfoResp) (list []model.SharePageFolderListResp, err error) {
	resp, err := c.core.shareFolderList(req)
	if err != nil {
		return
	}
	for _, f := range resp.FileListAO.FolderList {
		list = append(list, model.SharePageFolderListResp{
			Id:   cast.ToString(f.Id),
			Name: f.Name,
		})
	}

	return
}
func (c *Cloud189) GetShareInfo(url, pwd string) (info model.ShareInfoResp, err error) {
	code := ParseShareCode(url)
	info, err = c.core.getShareInfoByCodeV2(code)
	if err != nil {
		return
	}

	info.Pwd = pwd

	if info.Pwd != "" {
		resp, err := c.core.checkAccessCode(code, pwd)
		if err != nil {
			return info, err
		}
		info.ShareId = resp.ShareId
	}
	return
}
func (c *Cloud189) UserInfo() (resp model.UserInfo, err error) {
	resp, err = c.core.userInfo()
	return
}
func (c *Cloud189) AuthLogin(account model.Account) (client *req.Client, err error) {
	client, err = c.core.login(account)
	return
}
func (c *Cloud189) NewClient(r *req.Client) {
	c.core.invoker = invoker{client: r}
}
