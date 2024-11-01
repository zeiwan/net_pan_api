package model

type ShareInfoResp struct {
	ShareId        int64  `json:"shareId"` // 189
	FileId         string `json:"fileId"`
	Pwd            string `json:"pwd"`
	ShareMode      int32  `json:"shareMode"`
	ShareDirFileId string `json:"shareDirFileId"`
	FileName       string `json:"fileName"`
	SToken         string `json:"stoken"` // quark
	Code           string `json:"code"`   // quark
}

type SharePageFolderListResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsFolder uint8  `json:"isFolder"`
	ParentId string `json:"parentId"`
	SToken   string `json:"SToken"` // quark
}
type SharePageFileListResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsFolder uint8  `json:"isFolder"`
}
type SharePageALL struct {
	FolderList []SharePageFolderListResp
	FileList   []SharePageFileListResp
}

type CreateFolderResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MyFolderListResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"` // 用来识别
}
type MyFileListResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"` // 用来识别
}

type MyDirAll struct {
	FolderList []MyFolderListResp
	FileList   []MyFileListResp
}
