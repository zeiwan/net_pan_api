package quark

import (
	"errors"
	"github.com/ZeiWan/NetPanSDK/model"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

const (
	baseUrl = "https://pan.quark.cn"
	driveH  = "https://drive-h.quark.cn"
)

func (c core) login(account model.Account) (*req.Client, error) {
	values := url.Values{}
	values.Set("fr", "pc")
	values.Set("platform", "pc")

	// 创建一个 HTTP 客户端
	client := req.C()
	client.QueryParams = values
	client.SetCommonHeader("Cookie", account.Cookie)

	// 发送 GET 请求获取账户信息
	path := "/account/info"
	resp, err := client.R().Get(baseUrl + path) // 替换为实际的 baseUrl

	if resp.StatusCode != 200 {
		err = errors.New(jsoniter.Get(resp.Bytes(), "message").ToString())
		return nil, err
	}

	return client, err
}
func (c core) userInfo() (resp model.UserInfo, err error) {

	values := url.Values{}

	path := "/account/info"
	var obj userInfo
	err = c.invoker.Get(baseUrl, path, values, &obj)
	resp.LoginName = obj.Data.Nickname
	return
}
func (c core) shareInfo(code, pwd string) (info model.ShareInfoResp, err error) {
	values := url.Values{}

	values.Set("uc_param_str", "")
	values.Set("__t", GetTimestamp())

	body := make(map[string]any)
	body["pwd_id"] = code
	body["passcode"] = pwd
	//
	path := "/1/clouddrive/share/sharepage/token"
	var obj any
	err = c.invoker.Post(driveH, path, values, body, &obj)

	bytes, err := jsoniter.Marshal(obj)

	status := jsoniter.Get(bytes, "status").ToInt()

	if status == 404 {
		err = errors.New(jsoniter.Get(bytes, "message").ToString())
		return
	}

	info.SToken = jsoniter.Get(bytes, "data", "stoken").ToString()
	info.FileName = jsoniter.Get(bytes, "data", "title").ToString()
	info.ShareMode = jsoniter.Get(bytes, "data", "expired_type").ToInt32()
	info.FileId = "0"
	info.Code = code
	info.Pwd = pwd
	return info, err
}
func (c core) shareDetail(info model.ShareInfoResp) (resp sharePageFolderListResp, err error) {
	values := url.Values{}

	values.Add("uc_param_str", "")
	values.Add("pwd_id", info.Code)
	values.Add("stoken", info.SToken)
	values.Add("pdir_fid", info.FileId)
	values.Add("force", "0")
	values.Add("_page", "1")
	values.Add("_size", "50")
	values.Add("_fetch_banner", "1")
	values.Add("_fetch_share", "1")
	values.Add("_fetch_total", "1")
	values.Add("_sort", "file_type:asc,updated_at:desc")
	values.Add("__t", GetTimestamp())

	path := "/1/clouddrive/share/sharepage/detail"

	err = c.invoker.Get(driveH, path, values, &resp)
	return
}
