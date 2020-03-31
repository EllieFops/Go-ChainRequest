package creq

import (
	"github.com/Foxcapades/Go-ChainRequest/response"
	"net/http"
)

type Response interface {
	GetError() error

	GetResponseCode() uint16

	GetRawResponse() (*http.Response, error)

	GetBody() ([]byte, error)

	UnmarshalBody(interface{}, response.Unmarshaller) error
}
