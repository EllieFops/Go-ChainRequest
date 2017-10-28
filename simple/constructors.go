package simple

import (
	"github.com/Foxcapades/Go-ChainRequest"
	req "github.com/Foxcapades/Go-ChainRequest/request"
	"github.com/Foxcapades/Go-ChainRequest/request/method"
)

func GetRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.GET,
	}
}

func PostRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.POST,
	}
}

func PutRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.PUT,
	}
}

func DeleteRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.DELETE,
	}
}

func PatchRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.PATCH,
	}
}

func HeadRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.HEAD,
	}
}

func OptionsRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.OPTIONS,
	}
}

func ConnectRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.CONNECT,
	}
}

func InfoRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.INFO,
	}
}

func TraceRequest(url string) creq.Request {
	return &request{
		url:     url,
		headers: map[req.Header]string{},
		method:  method.TRACE,
	}
}
