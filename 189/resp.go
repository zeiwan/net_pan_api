package _189

type checkAccessCode struct {
	ShareId int64 `json:"shareId"`
}
type listShareDirResp struct {
	FileListAO struct {
		FileList   []fileList   `json:"fileList"`
		FolderList []folderList `json:"folderList"`
	} `json:"fileListAO"`
}
type fileList struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
	MD5      string `json:"md5"`
}
type folderList struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
	MD5      string `json:"md5"`
}
type taskInfoResp struct {
	TaskId string `json:"taskId"`
}
type checkBatchTaskResp struct {
	ResMessage     string `json:"res_message"`
	FailedCount    int    `json:"failedCount"`
	Process        int    `json:"process"`
	SkipCount      int    `json:"skipCount"`
	SubTaskCount   int    `json:"subTaskCount"`
	SuccessedCount int    `json:"successedCount"`
	TaskId         string `json:"taskId"`
	TaskStatus     int    `json:"taskStatus"`
}
type createFolderResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
