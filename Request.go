package creq

import (
	"github.com/Foxcapades/Go-ChainRequest/request"
	"net/http"
)

type Request interface {

	// Marshals the request body into an array of bytes
	//
	// Parameters:
	//   body: interface{}
	//     Raw data that should be marshalled into an array of bytes and set
	//     as the request body.
	//   marshaller: request.Marshaller
	//     The marshaller to turn the request body into an array of bytes.
	//
	// Returns:
	//   Request: The current instance of this Request
	MarshalBody(body interface{}, marshaller request.Marshaller) Request

	// Sets the Request body
	//
	// Parameters:
	//   body: []byte
	//     Raw data that should be sent with the HTTP request on submit.
	//     NOTE: This data will be ignored for HTTP verbs that do not support a
	//     request body.
	//
	// Returns:
	//   Request: The current instance of this Request
	SetBody(body []byte) Request

	// Gets the Request Body
	//
	// Returns:
	//   []byte: The body data (empty slice if not present)
	GetBody() []byte

	// Adds a Header to this Request
	//
	// AddHeader does not overwrite existing headers, if the header specified by
	// `key` is already set on this request, the value passed by the parameter
	// `value` will be appended to the existing value, separated by a semicolon.
	//
	// If the desired action is overwriting previous data use SetHeader.
	//
	// Parameters:
	//   key:   request.Header
	//     The header that should be appended to this Request
	//   value: string
	//     The value that should be assigned to the given header for this request.
	//
	// Returns:
	//   Request: the current instance of this Request
	AddHeader(key request.Header, value string) Request

	// Sets a Header to this Request
	//
	// SetHeader overwrites existing headers, if the header specified by `key` is
	// already set on this request, the value passed by the parameter `value` will
	// replace the existing value.
	//
	// If the desired action is appending to the previous data use AddHeader.
	//
	// Parameters:
	//   key:   request.Header
	//     The header that should be appended to this Request
	//   value: string
	//     The value that should be assigned to the given header for this request.
	//
	// Returns:
	//   Request: the current instance of this Request
	SetHeader(key request.Header, value string) Request

	// Gets a header, if set, from this Request
	//
	// Parameters:
	//   key: request.Header
	//     The header value that should be returned.
	//
	// Returns:
	//   string: the value assigned to that key, empty string if not present
	//   bool:   whether or not that key was set on this Request
	GetHeader(key request.Header) (string, bool)

	// Sets the request method to use for this Request.
	//
	// Custom HTTP verbs are allowed, but must be wrapped in the HttpVerb type to
	// make it explicit that something abnormal is being done in those cases.
	//
	// Parameters:
	//   method: request.Method
	//     The request method or HTTP verb to send this request as.
	//
	// Returns:
	//   Request: the current instance of this Request
	SetMethod(method request.Method) Request

	// Gets the request method set on this Request
	//
	// Returns:
	//   request.Method: the method that is currently set on this Request
	GetMethod() request.Method

	// Submit this Request
	//
	// Returns:
	//   Response A wrapped http response
	//   error    the error returned from submitting this request
	Submit() Response

	GetUrl() string

	SetHttpClient(client *http.Client) Request

	SetRequestBuilder(builder RequestBuilder) Request

	DisableRedirects() Request
}
