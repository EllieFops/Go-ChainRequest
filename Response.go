package creq

import (
	"github.com/Foxcapades/Go-ChainRequest/response"
	"net/http"
)

type Response interface {
	GetError() error

	GetResponseCode() (uint16, error)

	MustGetResponseCode() uint16

	GetRawResponse() (*http.Response, error)

	GetBody() ([]byte, error)

	MustGetBody() []byte

	GetHeader(key string) (string, error)

	MustGetHeader(key string) string

	LookupHeader(key string) (val string, ok bool, err error)

	MustLookupHeader(key string) (val string, ok bool)

	UnmarshalBody(interface{}, response.Unmarshaller) error
}
