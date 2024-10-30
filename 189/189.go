package _189

import (
	"core"
	"core/model"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func init() {
	fmt.Println("Cloud189 init")
	core.RegisterPan("189", &Cloud189{})
}
func (c Cloud189) UserInfo(r *req.Client) (resp model.UserInfo, err error) {

	fmt.Println("Cloud189 UserInfo", r)
	//values := url.Values{}
	path := "/open/user/getUserInfoForPortal.action"

	request, err := r.R().Get(baseUrl + path)
	fmt.Println(request)
	//err = c.Invoker.Get(path, values, &resp)
	return resp, nil
}

func (c Cloud189) AuthLogin(account model.Account) (r *req.Client, err error) {

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
	userRsa := RsaEncode([]byte(account.Username), pubKey)
	passwordRsa := RsaEncode([]byte(account.Password), pubKey)

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

		resp, err = client.R().Get("https://cloud.189.cn/v2/getUserBriefInfo.action?noCache=" + Random())
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
