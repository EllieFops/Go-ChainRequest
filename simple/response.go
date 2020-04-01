package simple

import (
	"io/ioutil"
	"net/http"

	res "github.com/Foxcapades/Go-ChainRequest/response"
)

type response struct {
	raw *http.Response
	err error

	body []byte
}

func (r *response) GetResponseCode() (uint16, error) {
	if r.err != nil {
		return 0, r.err
	}
	return uint16(r.raw.StatusCode), nil
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

func (r *response) UnmarshalBody(in interface{}, un res.Unmarshaller) error {
	dat, err := r.GetBody()

	if err != nil {
		return err
	}

	return un.Unmarshal(dat, in)
}

func (r *response) MustGetResponseCode() uint16 {
	if r.err != nil {
		panic(r.err)
	}
	return uint16(r.raw.StatusCode)
}
