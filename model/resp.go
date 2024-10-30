package model

type ShareInfoResp struct {
	ShareId        int64  `json:"shareId"`
	FileId         string `json:"fileId"`
	Pwd            string `json:"pwd"`
	ShareMode      int32  `json:"shareMode"`
	ShareDirFileId string `json:"shareDirFileId"`
	FileName       string `json:"fileName"`
}

type SharePageFolderListResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type SharePageFileListResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CreateFolderResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type MyFolderListResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
