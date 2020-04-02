package creq

import (
	"github.com/Foxcapades/Go-ChainRequest/response"
	"net/http"
)

type Response interface {
	// GetError returns the last error from either the request or another method
	// on Response itself.
	GetError() error

	// GetResponseCode returns the HTTP response code for the request.
	//
	// If an error occurred during the request then the returned value will be 0
	// and that error will be returned.
	GetResponseCode() (uint16, error)

	// MustGetResponseCode returns the HTTP response code for the request or
	// panics if that request failed.
	MustGetResponseCode() uint16

	// GetRawResponse returns the raw Golang http.Response object backing this
	// instance.
	GetRawResponse() (*http.Response, error)

	// GetBody returns the response body as a byte array.
	//
	// If the request failed, then the returned body will be empty, and the
	// request error will be returned.
	GetBody() ([]byte, error)

	// MustGetBody returns the response body as a byte array or panics if the
	// request failed.
	MustGetBody() []byte

	// GetHeader retrieves the header value (if set) at the specified key from the
	// response.
	//
	// If the response does not have that header, or the request failed, the
	// return value will be ""
	//
	// The returned error will be nil unless the request failed.
	GetHeader(key string) (string, error)

	// MustGetHeader retrieves the header value (if set) at the specified key from
	// the response or panics if the request failed.
	//
	// If the header does not exist in the response, "" is returned.
	MustGetHeader(key string) string

	// LookupHeader retrieves the header value at the specified key from the
	// response along with a boolean flag indicating whether that header existed.
	//
	// If the header did not exist or the request errored, val will be "", and
	// ok will be false.
	//
	// The returned error will be nil unless the request failed.
	LookupHeader(key string) (val string, ok bool, err error)

	// MustLookupHeader retrieves the header value at the specified key from the
	// response along with a boolean flag indicating whether that header existed.
	//
	// If the header did not exist val will be "", and ok will be false.
	//
	// If the request errored, this method panics
	MustLookupHeader(key string) (val string, ok bool)

	// UnmarshalBody attempts to unmarshal the raw response body into the given
	// pointer using the given Unmarshaller.
	//
	// If the request failed the unmarshaller will not be called and the error
	// will be returned.
	//
	// If unmashalling fails, the returned error will be the unmarshalling error.
	UnmarshalBody(interface{}, response.Unmarshaller) error

	// MustUnmarshalBody attempts to unmarshal the raw response body into the
	// given pointer using the given Unmarshaller.
	//
	// If the request failed, or unmarshalling fails, this method panics.
	MustUnmarshalBody(interface{}, response.Unmarshaller)
}
