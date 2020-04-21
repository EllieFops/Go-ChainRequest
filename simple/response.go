package simple

import (
	"io/ioutil"
	"net/http"

	res "github.com/Foxcapades/Go-ChainRequest/response"
)

type response struct {
	raw *http.Response
	err error

	body    []byte
	cookies map[string]*http.Cookie
}

func (r *response) GetResponseCode() (uint16, error) {
	if r.err != nil {
		return 0, r.err
	}
	return uint16(r.raw.StatusCode), nil
}

func (r *response) MustGetResponseCode() uint16 {
	if r.err != nil {
		panic(r.err)
	}
	return uint16(r.raw.StatusCode)
}

func (r *response) GetHeader(key string) (string, error) {
	if r.err != nil {
		return "", r.err
	}
	return r.raw.Header.Get(key), nil
}

func (r *response) MustGetHeader(key string) string {
	if r.err != nil {
		panic(r.err)
	}
	return r.raw.Header.Get(key)
}

func (r *response) LookupHeader(key string) (val string, ok bool, err error) {
	if r.err != nil {
		return "", false, r.err
	}

	if v, exists := r.raw.Header[key]; exists {
		return v[0], exists, nil
	}

	return "", false, nil
}

func (r *response) MustLookupHeader(key string) (val string, ok bool) {
	if r.err != nil {
		panic(r.err)
	}

	if v, exists := r.raw.Header[key]; exists {
		return v[0], exists
	}

	return "", false
}

func (r *response) GetError() error {
	return r.err
}

func (r *response) GetRawResponse() (*http.Response, error) {
	return r.raw, r.err
}

func (r *response) GetBody() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}

	if r.body == nil {
		body, err := ioutil.ReadAll(r.raw.Body)
		r.body = body
		return body, err
	}

	return r.body, nil
}

func (r *response) MustGetBody() []byte {
	bod, err := r.GetBody()
	if err != nil {
		panic(err)
	}
	return bod
}

func (r *response) GetCookie(name string) (*http.Cookie, error) {
	if r.err != nil {
		return nil, r.err
	}

	// populate cookie map on demand
	if r.cookies == nil {
		r.cookies = make(map[string]*http.Cookie, len(r.raw.Cookies()))
		for _, c := range r.raw.Cookies() {
			r.cookies[c.Name] = c
		}
	}

	return r.cookies[name], nil
}

func (r *response) MustGetCookie(name string) *http.Cookie {
	if c, e := r.GetCookie(name); e != nil {
		panic(e)
	} else {
		return c
	}
}

func (r *response) UnmarshalBody(in interface{}, un res.Unmarshaller) error {
	dat, err := r.GetBody()

	if err != nil {
		return err
	}

	return un.Unmarshal(dat, in)
}

func (r *response) MustUnmarshalBody(in interface{}, un res.Unmarshaller) {
	if dat, err := r.GetBody(); err != nil {
		panic(err)
	} else {
		if err = un.Unmarshal(dat, in); err != nil {
			panic(err)
		}
	}
}

func (r *response) Close() error {
	if r.raw == nil {
		return nil
	}

	return r.raw.Body.Close()
}
