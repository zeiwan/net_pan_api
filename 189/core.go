package _189

import (
	"errors"
	"fmt"
	"github.com/ZeiWan/NetPanSDK/model"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"time"
)

func (c core) login(account model.Account) (r *req.Client, err error) {
	client := req.C()
	tempUrl := "https://cloud.189.cn/api/portal/loginUrl.action?redirectURL=https%3A%2F%2Fcloud.189.cn%2Fmain.action"
	var lt, reqId string

	resp, err := client.
		SetRedirectPolicy(req.RedirectPolicy(func(req *http.Request, via []*http.Request) error {
			if req.URL.Query().Get("lt") != "" {
				lt = req.URL.Query().Get("lt")
			}
			if req.URL.Query().Get("reqId") != "" {
				reqId = req.URL.Query().Get("reqId")
			}
			return nil
		})).R().Get(tempUrl)

	if err != nil {
		return
	}

	cookies := resp.Cookies()
	client.SetRedirectPolicy(req.MaxRedirectPolicy(10))

	appConfResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/74.0",
			"Referer":    resp.Request.URL.String(),
			"Cookie":     cookiesToString(cookies),
			"lt":         lt,
			"reqId":      reqId,
		}).
		SetFormData(map[string]string{
			"appKey":  "cloud",
			"version": "2.0",
		}).
		Post("https://open.e.189.cn/api/logbox/oauth2/appConf.do")
	if err != nil {
		return
	}

	data := jsoniter.Get(appConfResp.Bytes(), "data")
	accountType := data.Get("accountType").ToString()
	clientType := data.Get("clientType").ToString()
	paramId := data.Get("paramId").ToString()
	mailSuffix := data.Get("mailSuffix").ToString()
	returnUrl := data.Get("returnUrl").ToString()

	encryptConfResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/74.0",
			"Referer":    "https://open.e.189.cn/api/logbox/separate/web/index.html",
			"Cookie":     cookiesToString(cookies),
		}).
		SetFormData(map[string]string{
			"appId": "cloud",
		}).
		Post("https://open.e.189.cn/api/logbox/config/encryptConf.do")
	if err != nil {
		return
	}

	if resCode := jsoniter.Get(encryptConfResp.Bytes(), "result").ToInt(); resCode != 0 {
		err = fmt.Errorf("Failed to get encrypt config")
		return
	}

	encryptData := jsoniter.Get(encryptConfResp.Bytes(), "data")
	pubKey := encryptData.Get("pubKey").ToString()
	pre := encryptData.Get("pre").ToString()
	vCodeRS := ""
	userRsa := rsaEncode([]byte(account.Username), pubKey)
	passwordRsa := rsaEncode([]byte(account.Password), pubKey)

	loginResp, err := client.R().
		SetHeaders(map[string]string{
			"lt":         lt,
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/76.0",
			"Referer":    "https://open.e.189.cn/",
		}).
		SetFormData(map[string]string{
			"version":      "v2.0",
			"appKey":       "cloud",
			"accountType":  accountType,
			"userName":     pre + userRsa,
			"epd":          pre + passwordRsa,
			"validateCode": vCodeRS,
			"captchaToken": "",
			"returnUrl":    returnUrl,
			"mailSuffix":   mailSuffix,
			"paramId":      paramId,
			"clientType":   clientType,
			"dynamicCheck": "FALSE",
			"cb_SaveName":  "1",
			"isOauth2":     "false",
		}).
		Post("https://open.e.189.cn/api/logbox/oauth2/loginSubmit.do")
	if err != nil {
		return
	}

	if restCode := jsoniter.Get(loginResp.Bytes(), "result").ToInt(); restCode == 0 {
		toUrl := jsoniter.Get(loginResp.Bytes(), "toUrl").ToString()

		resp, err = client.SetRedirectPolicy(req.MaxRedirectPolicy(10)).R().Get(toUrl)
		if err != nil {
			return
		}

		resp, err = client.R().Get("https://cloud.189.cn/v2/getUserBriefInfo.action?noCache=" + random())
		if err != nil {
			return
		}
		return client, err
	} else if restCode == -2 {
		err = errors.New(jsoniter.Get(loginResp.Bytes(), "msg").ToString())
		return
	} else {
		err = errors.New(jsoniter.Get(loginResp.Bytes(), "msg").ToString())
		return
	}
}
func (c core) userInfo() (resp model.UserInfo, err error) {
	values := url.Values{}
	path := "/open/user/getUserInfoForPortal.action"
	err = c.invoker.Get(path, values, &resp)
	return
}
func (c core) getShareInfoByCodeV2(code string) (resp model.ShareInfoResp, err error) {
	path := "/open/share/getShareInfoByCodeV2.action"
	values := url.Values{}
	values.Set("shareCode", code)
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}
func (c core) checkAccessCode(code, pwd string) (resp checkAccessCode, err error) {
	path := "/open/share/checkAccessCode.action"
	values := url.Values{}
	values.Set("shareCode", code)
	values.Set("accessCode", pwd)
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}
func (c core) shareFolderList(req model.ShareInfoResp) (resp listShareDirResp, err error) {
	path := "/open/share/listShareDir.action"
	values := url.Values{}
	values.Set("pageNum", "1")
	values.Set("pageSize", "60")
	values.Set("fileId", req.FileId)
	values.Set("shareDirFileId", req.FileId)
	values.Set("isFolder", "true")
	values.Set("shareId", cast.ToString(req.ShareId))
	values.Set("shareMode", cast.ToString(req.ShareMode))
	values.Set("iconOption", "5")
	values.Set("orderBy", "lastOpTime")
	values.Set("descending", "true")
	values.Set("accessCode", req.Pwd)
	err = c.invoker.Get(path, values, &resp)

	if err != nil {
		return
	}
	return
}
func (c core) createFolder(parentFolderId, folderName string) (resp model.CreateFolderResp, err error) {
	path := "/open/file/createFolder.action"
	values := url.Values{}
	values.Set("parentFolderId", parentFolderId)
	values.Set("folderName", folderName)
	var r createFolderResp
	err = c.invoker.Post(path, values, &r)
	resp.Id = cast.ToString(r.Id)
	resp.Name = r.Name

	if err != nil {
		return
	}
	return
}
func (c core) getMyFolder(id string) (resp []model.MyFolderListResp, err error) {
	path := "/portal/getObjectFolderNodes.action"

	values := url.Values{}
	if id == "" {
		id = "-11"
	}
	values.Set("id", id)
	values.Set("orderBy", "1")
	values.Set("order", "ASC")
	err = c.invoker.Post(path, values, &resp)
	if err != nil {
		return
	}
	return
}
func (c core) rename(folderId, newFolderName string) (bool, error) {
	values := url.Values{}
	values.Set("folderId", folderId)
	values.Set("destFolderName", newFolderName)

	path := "/open/file/renameFolder.action"
	err := c.invoker.Post(path, values, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (c core) createBatchTask(types, targetFolderId, shareId string, taskInfos []model.TaskInfosReq) (taskId taskInfoResp, err error) {
	marshal, err := jsoniter.Marshal(taskInfos)
	if err != nil {
		return
	}
	values := url.Values{}
	values.Set("type", types)
	values.Set("taskInfos", string(marshal))
	values.Set("targetFolderId", targetFolderId)
	if shareId != "" {
		values.Set("shareId", shareId)
	}
	path := "/open/batch/createBatchTask.action"
	err = c.invoker.Post(path, values, &taskId)
	return
}
func (c core) checkBatchTask(types, taskId string, maxRetries uint8) (err error) {
	values := url.Values{}
	values.Set("type", types)
	values.Set("taskId", taskId)
	path := "/open/batch/checkBatchTask.action"
	var resp checkBatchTaskResp
	err = c.invoker.Post(path, values, &resp)
	if err != nil {
		return
	}
	switch resp.TaskStatus {
	case -1:
		err = errors.New(resp.ResMessage)
		return
	case 2:
		return
	case 4:
		return
	}
	if resp.SubTaskCount != resp.SuccessedCount && maxRetries > 0 {
		time.Sleep(1 * time.Second)
		return c.checkBatchTask(types, taskId, maxRetries-1)
	}
	return
}

//	func (c core) shareLink(fileId string, expireTime, shareType uint8) (string, error) {
//		values := url.Values{}
//		values.Set("fileId", fileId)
//		values.Set("expireTime", cast.ToString(expireTime))
//		values.Set("shareType", cast.ToString(shareType))
//
//		path := "/open/share/createShareLink.action"
//		c.invoker.Get(path, values, nil)
//	}
func (c core) getMyFileAll(folderId string) (resp model.MyFileListResp, err error) {
	path := "/open/file/listFiles.action"
	values := url.Values{}
	values.Set("pageNum", "1")
	values.Set("pageSize", "60")
	values.Set("mediaType", "0")
	values.Set("folderId", folderId)
	values.Set("iconOption", "5")
	values.Set("orderBy", "lastOpTime")
	values.Set("descending", "true")
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}
