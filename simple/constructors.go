package simple

import (
	"github.com/Foxcapades/Go-ChainRequest"
	req "github.com/Foxcapades/Go-ChainRequest/request"
	"github.com/Foxcapades/Go-ChainRequest/request/method"
	"net/http"
)

func GetRequest(url string) creq.Request {
	return newRequest(url, method.GET)
}

func PostRequest(url string) creq.Request {
	return newRequest(url, method.POST)
}

func PutRequest(url string) creq.Request {
	return newRequest(url, method.PUT)
}

func DeleteRequest(url string) creq.Request {
	return newRequest(url, method.DELETE)
}

func PatchRequest(url string) creq.Request {
	return newRequest(url, method.PATCH)
}

func HeadRequest(url string) creq.Request {
	return newRequest(url, method.HEAD)
}

func OptionsRequest(url string) creq.Request {
	return newRequest(url, method.OPTIONS)
}

func ConnectRequest(url string) creq.Request {
	return newRequest(url, method.CONNECT)
}

func InfoRequest(url string) creq.Request {
	return newRequest(url, method.INFO)
}

func TraceRequest(url string) creq.Request {
	return newRequest(url, method.TRACE)
}

func newRequest(url string, meth req.Method) *request {
	return &request{
		url:     url,
		method:  meth,
		headers: map[req.Header]string{},
		client:  http.DefaultClient,
		factory: RequestBuilder(http.NewRequest),
	}
}
