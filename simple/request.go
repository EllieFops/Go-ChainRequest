package simple

import "net/http"
import "bytes"
import (
	"github.com/Foxcapades/Go-ChainRequest"
	req "github.com/Foxcapades/Go-ChainRequest/request"
	"io"
)

type request struct {
	headers map[req.Header]string
	method  req.Method
	url     string
	body    []byte
	err     error
}

func (r request) MarshalBody(
	body interface{},
	marshaller req.Marshaller,
) creq.Request {
	r.body, r.err = marshaller.Marshal(body)

	return &r
}

func (r request) SetBody(body []byte) creq.Request {
	r.body = body

	return &r
}

func (r request) GetBody() ([]byte, bool) {
	if r.body == nil {
		return make([]byte, 0), false
	}
	return r.body, true
}

func (r request) AddHeader(key req.Header, value string) creq.Request {
	if _, ok := r.headers[key]; ok {
		r.headers[key] += "; " + value
	} else {
		r.headers[key] = value
	}

	return &r
}

func (r request) SetHeader(key req.Header, value string) creq.Request {
	r.headers[key] = value

	return &r
}

func (r request) GetHeader(key req.Header) (string, bool) {
	a, b := r.headers[key]

	return a, b
}

func (r request) SetMethod(method req.Method) creq.Request {
	r.method = method

	return &r
}

func (r request) GetMethod() req.Method {
	return r.method
}

func (r request) Submit() creq.Response {
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
