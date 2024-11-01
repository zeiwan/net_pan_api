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
	Id           string `json:"id"`
	Name         string `json:"name"`
	ParentId     string `json:"parentId"`
	MD5          string `json:"md5"`
}
type folderList struct {
	Id       string `json:"id"`
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

//type createShareLink struct {
//	ResCode       int    `json:"res_code"`
//	ResMessage    string `json:"res_message"`
//	ShareLinkList []struct {
//		AccessCode string `json:"accessCode"`
//		AccessUrl  string `json:"accessUrl"`
//		FileId     string `json:"fileId"`
//		ShareId    int64  `json:"shareId"`
//		Url        string `json:"url"`
//	} `json:"shareLinkList"`
//}
