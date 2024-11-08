package _189

import (
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"github.com/zeiwan/net_pan_api/module"
)

const (
	DELETE     = "DELETE"
	Move       = "Move"
	COPY       = "COPY"
	SHARE_SAVE = "SHARE_SAVE"
)

func (c *Cloud189) SaveFile(taskReq module.TaskShareReq, taskInfos []module.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(SHARE_SAVE, taskReq.TargetFolderId, taskReq.ShareId, taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(SHARE_SAVE, taskInfo.TaskId, 3)
}

func (c *Cloud189) GetMyFileAll(targetFolderId string) (resp module.MyFolderAll, err error) {
	m, err := c.core.getMyFileAll(targetFolderId)

	if err != nil {
		return
	}
	for _, r := range m.FileListAO.FileList {
		resp.FileList = append(resp.FileList, module.MyFileListResp{
			Id:   cast.ToString(r.Id),
			Name: r.Name,
			Tag:  r.MD5,
		})
	}
	for _, r := range m.FileListAO.FolderList {
		resp.FolderList = append(resp.FolderList, module.MyFolderListResp{
			Id:   cast.ToString(r.Id),
			Name: r.Name,
			Tag:  r.MD5,
		})
	}
	return
}

func (c *Cloud189) Copy(targetFolderId string, taskInfos []module.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(COPY, targetFolderId, "", taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(COPY, taskInfo.TaskId, 3)
}

func (c *Cloud189) Move(targetFolderId string, taskInfos []module.TaskInfosReq) (err error) {
	taskInfo, err := c.core.createBatchTask(Move, targetFolderId, "", taskInfos)
	if err == nil && taskInfo.TaskId != "" {
		return
	}
	return c.core.checkBatchTask(Move, taskInfo.TaskId, 3)
}

func (c *Cloud189) Delete(taskInfos []module.TaskInfosReq) (err error) {
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
func (c *Cloud189) GetMyFolder(id string) (resp []module.MyFolderListResp, err error) {
	resp, err = c.core.getMyFolder(id)
	return
}
func (c *Cloud189) CreateFolder(parentFolderId, folderName string) (resp module.CreateFolderResp, err error) {
	resp, err = c.core.createFolder(parentFolderId, folderName)
	return
}
func (c *Cloud189) GetSharePageFileList(req module.ShareInfoResp) (list []module.SharePageFileListResp, err error) {
	resp, err := c.core.shareFolderList(req)
	for _, f := range resp.FileListAO.FileList {
		list = append(list, module.SharePageFileListResp{
			Id:   cast.ToString(f.Id),
			Name: f.Name,
			Tag:  f.MD5,
		})
	}
	return
}
func (c *Cloud189) GetSharePageAll(req module.ShareInfoResp) (list module.SharePageALL, err error) {
	resp, err := c.core.shareFolderList(req)
	var fileLists []module.SharePageFileListResp
	var folderLists []module.SharePageFolderListResp
	for _, f := range resp.FileListAO.FileList {
		fileLists = append(fileLists, module.SharePageFileListResp{
			Id:   cast.ToString(f.Id),
			Name: f.Name,
			Tag:  f.MD5,
		})
	}
	for _, f := range resp.FileListAO.FolderList {
		folderLists = append(folderLists, module.SharePageFolderListResp{
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 0,
			Tag:      f.MD5,
		})
	}
	list = module.SharePageALL{
		FileList:   fileLists,
		FolderList: folderLists,
	}
	return
}
func (c *Cloud189) GetSharePageFolderList(req module.ShareInfoResp) (list []module.SharePageFolderListResp, err error) {
	resp, err := c.core.shareFolderList(req)
	if err != nil {
		return
	}
	for _, f := range resp.FileListAO.FolderList {
		list = append(list, module.SharePageFolderListResp{
			Id:       cast.ToString(f.Id),
			Name:     f.Name,
			IsFolder: 1,
			Tag:      f.MD5,
		})
	}

	return
}
func (c *Cloud189) GetShareInfo(url, pwd string) (info module.ShareInfoResp, err error) {
	code := parseShareCode(url)
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
func (c *Cloud189) UserInfo() (resp module.UserInfo, err error) {
	resp, err = c.core.userInfo()
	return
}
func (c *Cloud189) AuthLogin(account module.Account) (client *req.Client, err error) {
	client, err = c.core.login(account)
	return
}
func (c *Cloud189) NewClient(r *req.Client) {
	c.core.invoker = invoker{client: r}
}
