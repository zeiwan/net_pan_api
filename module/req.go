package module

type TaskInfosReq struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	IsFolder int    `json:"isFolder"`
}
