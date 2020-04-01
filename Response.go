package creq

import (
	"github.com/Foxcapades/Go-ChainRequest/response"
	"net/http"
)

type Response interface {
	GetError() error

	GetResponseCode() (uint16, error)

	GetRawResponse() (*http.Response, error)

	GetBody() ([]byte, error)

	GetHeader(key string) (string error)

	MustGetHeader(key string) string

	LookupHeader(key string) (val string, ok bool, err error)

	MustLookupHeader(key string) (val string, ok bool)

	UnmarshalBody(interface{}, response.Unmarshaller) error
}
