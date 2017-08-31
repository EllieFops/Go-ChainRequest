package creq

import (
  "net/http"
  "io/ioutil"
)

type response struct {
  raw *http.Response

  body []byte
}

func (r *response) GetRawResponse() *http.Response {
  return r.raw
}

func (r *response) GetStatusCode() int {
  return r.raw.StatusCode
}

func (r *response) GetStatus() string {
  return r.raw.Status
}

func (r *response) GetBody() ([]byte, error) {
  if r.body == nil {
    body, err := ioutil.ReadAll(r.raw.Body)
    r.body = body
    return body, err
  }

  return r.body, nil
}

func (r *response) GetHeader(in string) string {
  return r.raw.Header.Get(in)
}

func (r *response) GetHeaders() http.Header {
  return r.raw.Header
}

func (r *response) UnmarshalBody(in interface{}, un ResponseUnmarshaller) error {
  dat, err := r.GetBody()

  if err != nil { return err }

  return un.Unmarshal(dat, in)
}
