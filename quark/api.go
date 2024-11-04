package quark

import (
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"github.com/zeiwan/net_pan_api/module"
)

func (c *CloudQuark) SaveFile(taskReq module.TaskShareReq, taskInfos []module.TaskInfosReq) error {
	panic("method SaveFile not implemented")
}

func (c *CloudQuark) AuthLogin(account module.Account) (*req.Client, error) {
	return c.core.login(account)
}

func (c *CloudQuark) NewClient(client *req.Client) {
	c.core.invoker = invoker{client: client}
}

func (c *CloudQuark) UserInfo() (module.UserInfo, error) {
	return c.core.userInfo()
}

func (c *CloudQuark) GetShareInfo(url, pwd string) (module.ShareInfoResp, error) {
	code := parseShareCode(url)
	return c.core.shareInfo(code, pwd)
}

func (c *CloudQuark) GetSharePageFolderList(info module.ShareInfoResp) (resp []module.SharePageFolderListResp, err error) {
	m, err := c.core.shareDetail(info)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp = append(resp, module.SharePageFolderListResp{
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

func (c *CloudQuark) GetSharePageFileList(info module.ShareInfoResp) (resp []module.SharePageFileListResp, err error) {
	m, err := c.core.shareDetail(info)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if !s.Dir {
			resp = append(resp, module.SharePageFileListResp{
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

func (c *CloudQuark) GetSharePageAll(info module.ShareInfoResp) (resp module.SharePageALL, err error) {
	m, err := c.core.shareDetail(info)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp.FolderList = append(resp.FolderList, module.SharePageFolderListResp{
				Id:       s.Fid,
				Name:     s.FileName,
				IsFolder: cast.ToUint8(s.FileType),
				ParentId: s.Fid,
				SToken:   s.ShareFidToken,
			})
		} else {
			resp.FileList = append(resp.FileList, module.SharePageFileListResp{
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

func (c *CloudQuark) CreateFolder(parentFolderId, folderName string) (module.CreateFolderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) GetMyFolder(id string) (resp []module.MyFolderListResp, err error) {
	m, err := c.core.getMyFolderNodes(id)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp = append(resp, module.MyFolderListResp{
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

func (c *CloudQuark) Delete(taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Move(targetFolderId string, taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}

func (c *CloudQuark) Copy(targetFolderId string, taskInfos []module.TaskInfosReq) error {
	//TODO implement me
	panic("implement me")
}
func (c *CloudQuark) GetMyFileAll(id string) (resp module.MyFolderAll, err error) {
	m, err := c.core.getMyFolderNodes(id)
	if err != nil {
		return
	}
	for _, s := range m.Data.List {
		if s.Dir {
			resp.FolderList = append(resp.FolderList, module.MyFolderListResp{
				Id:   s.Fid,
				Name: s.FileName,
			})
		} else {
			resp.FileList = append(resp.FileList, module.MyFileListResp{
				Id:   s.Fid,
				Name: s.FileName,
			})
		}
	}
	return
}
