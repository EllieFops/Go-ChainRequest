package simple

import (
	res "github.com/Foxcapades/Go-ChainRequest/response"
	"io/ioutil"
	"net/http"
)

type response struct {
	raw *http.Response
	err error

	body []byte
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
