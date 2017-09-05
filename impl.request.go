package creq

import "net/http"
import "bytes"
import "io"

type request struct {
  headers map[RequestHeader]string
  method  HttpVerb
  url     string
  body    []byte
  err     error
}

func (r request) MarshalBody(
  body interface{},
  marshaller RequestMarshaller,
) Request {
  r.body, r.err = marshaller.Marshal(body)

  return &r
}

func (r request) SetBody(body []byte) Request {
  r.body = body

  return &r
}

func (r request) GetBody() ([]byte, bool) {
  if r.body == nil {
    return make([]byte, 0), false
  }
  return r.body, true
}

func (r request) AddHeader(key RequestHeader, value string) Request {
  if _, ok := r.headers[key]; ok {
    r.headers[key] += "; " + value
  } else {
    r.headers[key] = value
  }

  return &r
}

func (r request) SetHeader(key RequestHeader, value string) Request {
  r.headers[key] = value

  return &r
}

func (r request) GetHeader(key RequestHeader) (string, bool) {
  a, b := r.headers[key]

  return a, b
}

func (r request) SetMethod(method HttpVerb) Request {
  r.method = method

  return &r
}

func (r request) GetMethod() HttpVerb {
  return r.method
}

func (r request) Submit() Response {
  var res *http.Response

  if r.err != nil {
    return &response{err: r.err}
  }

  req, err := http.NewRequest(string(r.method), r.url, r.getBodyReader())
  if nil != err {
    return &response{err: err}
  }

  for header := range r.headers {
    req.Header.Set(string(header), r.headers[header])
  }

  res, err = http.DefaultClient.Do(req)
  return &response{raw: res, err: err}
}

func (r request) getBodyReader() io.Reader {
  if len(r.body) > 0 {
    return bytes.NewReader(r.body)
  }

  return nil
}
