package _189

type checkAccessCode struct {
	ShareId int64 `json:"shareId"`
}
type listShareDirResp struct {
	FileListAO struct {
		Count        int          `json:"count"`
		FileList     []fileList   `json:"fileList"`
		FileListSize int          `json:"fileListSize"`
		FolderList   []folderList `json:"folderList"`
	} `json:"fileListAO"`
}
type fileList struct {
	FileListSize int    `json:"fileListSize"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
}
type folderList struct {
	FileListSize int    `json:"fileListSize"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ParentId     int64  `json:"parentId"`
}
