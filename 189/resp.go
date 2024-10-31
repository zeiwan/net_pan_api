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
	ParentId     int64  `json:"parentId"`
}
type folderList struct {
	FileListSize int    `json:"fileListSize"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	ParentId     int64  `json:"parentId"`
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
