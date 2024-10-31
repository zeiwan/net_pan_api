package quark

import (
	"errors"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

func (i *invoker) Get(url, path string, params url.Values, data interface{}) error {

	client := i.client.SetBaseURL(url).DevMode().R()

	params.Add("pr", "ucpro")
	params.Add("fr", "pc")

	client.QueryParams = params

	err := i.do(client, "GET", path, &data)
	return err
}
func (i *invoker) do(client *req.Request, method string, path string, data interface{}) (err error) {
	resp, err := client.Send(method, path)

	if resp.StatusCode == 400 {
		resMessage := jsoniter.Get(resp.Bytes(), "res_message").ToString()
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

func (i *invoker) Post(url, path string, params url.Values, body map[string]interface{}, data interface{}) error {

	client := i.client.DevMode().SetBaseURL(url).R()

	params.Add("pr", "ucpro")
	params.Add("fr", "pc")

	client.QueryParams = params

	client.SetBody(body)

	client.SetHeader("Accept", "application/json, text/plain, */*")
	client.SetHeader("Content-Type", "application/json")
	err := i.do(client, "POST", path, &data)
	return err
}
