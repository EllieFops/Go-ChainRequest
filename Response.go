package creq

import "net/http"

type Response interface {
  GetRawResponse() *http.Response

  GetStatusCode() int

  GetStatus() string

  GetBody() ([]byte, error)

  GetHeader(string) string

  GetHeaders() http.Header

  UnmarshalBody(interface{}, ResponseUnmarshaller) error
}
