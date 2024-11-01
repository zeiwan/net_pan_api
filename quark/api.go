package quark

import (
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
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp = append(resp, model.SharePageFolderListResp{
				Id:       s.Fid,
				Name:     s.FileName,
				IsFolder: cast.ToUint8(s.FileType),
				ParentId: s.Fid,
				SToken:   s.ShareFidToken,
			})
		}

	}

	return resp, nil
}

func (c *CloudQuark) GetSharePageFileList(info model.ShareInfoResp) (resp []model.SharePageFileListResp, err error) {
	m, err := c.core.shareDetail(info)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if !s.Dir {
			resp = append(resp, model.SharePageFileListResp{
				Id:       s.Fid,
				Name:     s.FileName,
				IsFolder: cast.ToUint8(s.FileType),
				//ParentId: s.Fid,
				//SToken:   s.ShareFidToken,
			})
		}

	}

	return resp, nil
}

func (c *CloudQuark) GetSharePageAll(info model.ShareInfoResp) (resp model.SharePageALL, err error) {
	m, err := c.core.shareDetail(info)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp.FolderList = append(resp.FolderList, model.SharePageFolderListResp{
				Id:       s.Fid,
				Name:     s.FileName,
				IsFolder: cast.ToUint8(s.FileType),
				ParentId: s.Fid,
				SToken:   s.ShareFidToken,
			})
		} else {
			resp.FileList = append(resp.FileList, model.SharePageFileListResp{
				Id:       s.Fid,
				Name:     s.FileName,
				IsFolder: cast.ToUint8(s.FileType),
				//ParentId: s.Fid,
				//SToken:   s.ShareFidToken,
			})
		}

	}
	return
}

func (c *CloudQuark) CreateFolder(parentFolderId, folderName string) (model.CreateFolderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) GetMyFolder(id string) (resp []model.MyFolderListResp, err error) {
	m, err := c.core.getMyFolderNodes(id)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp = append(resp, model.MyFolderListResp{
				Id:   s.Fid,
				Name: s.FileName,
			})
		}

	}
	return
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
func (c *CloudQuark) GetMyFileAll(id string) (resp model.MyDirAll, err error) {
	m, err := c.core.getMyFolderNodes(id)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp.FolderList = append(resp.FolderList, model.MyFolderListResp{
				Id:   s.Fid,
				Name: s.FileName,
			})
		} else {
			resp.FileList = append(resp.FileList, model.MyFileListResp{
				Id:   s.Fid,
				Name: s.FileName,
			})
		}
	}
	return
}
