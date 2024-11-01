package aliyun

import (
	"github.com/imroc/req/v3"
)

type core struct {
	invoker invoker
}
type CloudALi struct {
	core core
}

type invoker struct {
	client *req.Client
}
