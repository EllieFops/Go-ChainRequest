package creq

import "net/http"

type Response interface {
  GetError() error

  GetRawResponse() (*http.Response, error)

  GetBody() ([]byte, error)

  UnmarshalBody(interface{}, ResponseUnmarshaller) error
}
