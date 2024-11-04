package module

type TaskInfosReq struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	IsFolder int    `json:"isFolder"`
}
type TaskShareReq struct {
	ShareId        string `json:"shareId"`
	TargetFolderId string `json:"targetFolderId"`
}
