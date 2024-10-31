package _189

import (
	"errors"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

const (
	baseUrl = "https://cloud.189.cn/api"
)

func (i *invoker) Get(path string, params url.Values, data interface{}) error {
	client := i.client.SetBaseURL(baseUrl).DevMode().R()

	client.QueryParams = params
	client.SetQueryParam("noCache", random())
	client.SetHeader("Accept", "application/json;charset=UTF-8")

	err := i.do(client, "GET", path, &data)
	return err
}
func (i *invoker) do(client *req.Request, method string, path string, data interface{}) (err error) {
	resp, err := client.Send(method, path)
	if resp.StatusCode == 400 {
		resMessage := jsoniter.Get(resp.Bytes(), "res_message").ToString()
		errorCode := jsoniter.Get(resp.Bytes(), "errorCode").ToString()
		if errorCode == "InvalidSessionKey" {
			_ = i.resetClient()
			_ = i.do(client, method, path, &data)
		}
		return errors.New(resMessage)
	}

	if err != nil {
		return
	}
	err = resp.Into(&data)
	if err != nil {
		return
	}
	return
}

func (i *invoker) Post(path string, params url.Values, data interface{}) error {
	client := i.client.SetBaseURL(baseUrl).DevMode().R()

	client.FormData = params
	client.SetQueryParam("noCache", random())
	client.SetHeader("Accept", "application/json;charset=UTF-8")
	client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	err := i.do(client, "POST", path, &data)
	return err
}

func (i *invoker) resetClient() error {
	// 设置查询参数
	values := url.Values{}
	values.Set("redirectURL", "https://cloud.189.cn/web/redirect.html")
	values.Set("defaultSaveName", "3")
	values.Set("defaultSaveNameCheck", "uncheck")
	values.Set("browserId", "48d76ec45fec4cf27901a171759c289e")

	// 创建客户端
	client := i.client.R()

	client.QueryParams = values

	_, err := client.Get("https://cloud.189.cn/api/portal/loginUrl.action")
	if err != nil {
		return err
	}
	i.client = client.GetClient()
	return nil
}
