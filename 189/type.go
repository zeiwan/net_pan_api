package _189

import (
	"github.com/imroc/req/v3"
)

type core struct {
	invoker invoker
}
type Cloud189 struct {
	core core
}

type invoker struct {
	client *req.Client
}
