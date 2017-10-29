package simple

import (
	"bytes"
	"github.com/Foxcapades/Go-ChainRequest"
	"io"
	"net/http"
)

type RequestBuilder func(string, string, io.Reader) (*http.Request, error)

func (r RequestBuilder) Build(request creq.Request) (*http.Request, error) {
	return http.NewRequest(
		string(request.GetMethod()),
		string(request.GetUrl()),
		getBodyReader(request.GetBody()),
	)
}

func getBodyReader(body []byte) io.Reader {
	if len(body) > 0 {
		return bytes.NewReader(body)
	}

	return nil
}
