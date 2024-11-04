package quark

import (
	"github.com/imroc/req/v3"
)

type core struct {
	invoker invoker
}
type CloudQuark struct {
	core core
}

type invoker struct {
	client *req.Client
}
