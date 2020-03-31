package creq

import (
	"net/http"
)

type RequestBuilder interface {
	Build(request Request) (*http.Request, error)
}

type RequestBuilderFunc func(request Request) (*http.Request, error)

func (r RequestBuilderFunc) Build(request Request) (*http.Request, error) {
	return r(request)
}
