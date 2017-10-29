package creq

import (
	"net/http"
)

type RequestBuilder interface {
	Build(request Request) (*http.Request, error)
}
