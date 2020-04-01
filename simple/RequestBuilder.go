package simple

import (
	"bytes"
	"io"

	"net/http"

	"github.com/Foxcapades/Go-ChainRequest"
)

type RequestBuilder func(string, string, io.Reader) (*http.Request, error)

func (r RequestBuilder) Build(request creq.Request) (*http.Request, error) {
	return http.NewRequest(
		string(request.GetMethod()),
		request.GetUrl(),
		getBodyReader(request.GetBody()),
	)
}

func getBodyReader(body []byte) io.Reader {
	if len(body) > 0 {
		return bytes.NewReader(body)
	}

	return nil
}
