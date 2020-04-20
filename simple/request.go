package simple

import (
	"net/http"
	"time"

	"github.com/Foxcapades/Go-ChainRequest"
	req "github.com/Foxcapades/Go-ChainRequest/request"
)

type request struct {
	headers map[req.Header]string
	method  req.Method
	url     string
	body    []byte
	cookies []*http.Cookie
	err     error
	client  *http.Client
	factory creq.RequestBuilder
}

func (r *request) AddCookie(cookie *http.Cookie) creq.Request {
	r.cookies = append(r.cookies, cookie)
	return r
}

func (r *request) SetHttpClient(client *http.Client) creq.Request {
	r.client = client
	return r
}

func (r *request) DisableRedirects() creq.Request {
	r.client.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return r
}

func (r *request) SetRequestBuilder(builder creq.RequestBuilder) creq.Request {
	r.factory = builder
	return r
}

func (r *request) MarshalBody(body interface{}, marshaller req.Marshaller) creq.Request {
	r.body, r.err = marshaller.Marshal(body)
	return r
}

func (r *request) SetBody(body []byte) creq.Request {
	r.body = body
	return r
}

func (r *request) GetBody() []byte {
	if r.body == nil {
		return make([]byte, 0)
	}
	return r.body
}

func (r *request) AddHeader(key req.Header, value string) creq.Request {
	if _, ok := r.headers[key]; ok {
		r.headers[key] += "; " + value
	} else {
		r.headers[key] = value
	}

	return r
}

func (r *request) SetHeader(key req.Header, value string) creq.Request {
	r.headers[key] = value
	return r
}

func (r *request) GetHeader(key req.Header) (string, bool) {
	a, b := r.headers[key]
	return a, b
}

func (r *request) SetMethod(method req.Method) creq.Request {
	r.method = method
	return r
}

func (r *request) GetMethod() req.Method {
	return r.method
}

func (r *request) GetUrl() string {
	return r.url
}

func (r *request) Submit() creq.Response {
	if r.err != nil {
		return &response{err: r.err}
	}

	quest, err := r.factory.Build(r)

	if nil != err {
		return &response{err: err}
	}

	r.assignHeaders(quest.Header)

	for _, c := range r.cookies {
		quest.AddCookie(c)
	}

	res, err := r.client.Do(quest)

	return &response{raw: res, err: err}
}

func (r *request) assignHeaders(head http.Header) {
	for header := range r.headers {
		head.Set(string(header), r.headers[header])
	}
}
