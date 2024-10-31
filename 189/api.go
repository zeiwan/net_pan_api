package _189

import (
	"github.com/ZeiWan/NetPanSDK/model"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
)

const (
	DELETE = "DELETE"
	Move   = "Move"
	COPY   = "COPY"
)

// func (c *Cloud189) ShareLink(fileId string, expireTime, shareType uint8) (string, error) {
//
// }
func (c *Cloud189) Copy(targetFolderId string, taskInfos []model.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(COPY, targetFolderId, "", taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(COPY, taskInfo.TaskId, 3)
}

func (c *Cloud189) Move(targetFolderId string, taskInfos []model.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(Move, targetFolderId, "", taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(Move, taskInfo.TaskId, 3)
}

func (c *Cloud189) Delete(taskInfos []model.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(DELETE, "", "", taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(DELETE, taskInfo.TaskId, 3)
}

func (c *Cloud189) Rename(folderId, newFolderName string) (ok bool, err error) {
	ok, err = c.core.rename(folderId, newFolderName)
	return
}
func (c *Cloud189) GetMyFolder(id string) (resp []model.MyFolderListResp, err error) {
	resp, err = c.core.getMyFolder(id)
	return
}
func (c *Cloud189) CreateFolder(parentFolderId, folderName string) (resp model.CreateFolderResp, err error) {
	resp, err = c.core.createFolder(parentFolderId, folderName)
	return
}
func (c *Cloud189) GetSharePageFileList(req model.ShareInfoResp) (list []model.SharePageFileListResp, err error) {
	resp, err := c.core.shareFolderList(req)
	for _, f := range resp.FileListAO.FileList {
		list = append(list, model.SharePageFileListResp{
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 0,
		})
	}
	return
}
func (c *Cloud189) GetSharePageAll(req model.ShareInfoResp) (list model.SharePageALL, err error) {
	resp, err := c.core.shareFolderList(req)
	var fileLists []model.SharePageFileListResp
	var folderLists []model.SharePageFolderListResp
	for _, f := range resp.FileListAO.FileList {
		fileLists = append(fileLists, model.SharePageFileListResp{
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 0,
		})
	}
	for _, f := range resp.FileListAO.FileList {
		folderLists = append(folderLists, model.SharePageFolderListResp{
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 0,
		})
	}
	list = model.SharePageALL{
		FileList:   fileLists,
		FolderList: folderLists,
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
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 1,
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
