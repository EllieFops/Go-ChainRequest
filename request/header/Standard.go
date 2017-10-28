package header

import "github.com/Foxcapades/Go-ChainRequest/request"

// Standard Request Fields
const (
	// Content-Types that are acceptable for the response.
	// See Content negotiation.
	ACCEPT request.Header = "Accept"

	// Character sets that are acceptable.
	ACCEPT_CHARSET request.Header = "Accept-Charset"

	// List of acceptable encodings. See HTTP compression.
	ACCEPT_ENCODING request.Header = "Accept-Encoding"

	// List of acceptable human languages for response. See Content negotiation.
	ACCEPT_LANGUAGE request.Header = "Accept-Language"

	// Acceptable version in time.
	ACCEPT_DATETIME request.Header = "Accept-Datetime"

	// Initiates a request for cross-origin resource sharing with Origin (below).
	ACCESS_CONTROL_REQUEST_METHOD request.Header = "Access-Control-Request-Method"

	//Initiates a request for cross-origin resource sharing with Origin (below).
	ACCESS_CONTROL_REQUEST_HEADERS request.Header = "Access-Control-Request-request.Headers"

	// Authentication credentials for HTTP authentication.
	AUTHORIZATION request.Header = "Authorization"

	// Used to specify directives that must be obeyed by all caching mechanisms
	// along the request-response chain.
	CACHE_CONTROL request.Header = "Cache-Control"

	// Control options for the current connection and list of hop-by-hop request
	// fields.
	CONNECTION request.Header = "Connection"

	// An HTTP cookie previously sent by the server with Set-Cookie (below).
	COOKIE request.Header = "Cookie"

	// The length of the request body in octets (8-bit bytes).
	CONTENT_LENGTH request.Header = "Content-Length"

	// A Base64-encoded binary MD5 sum of the content of the request body.
	CONTENT_MD5 request.Header = "Content-MD5"

	// The MIME type of the body of the request (used with POST and PUT requests).
	CONTENT_TYPE request.Header = "Content-Type"

	// The date and time that the message was originated (in "HTTP-date" format as
	// defined by RFC 7231 Date/Time Formats).
	DATE request.Header = "Date"

	// Indicates that particular server behaviors are required by the client.
	EXPECT request.Header = "Expect"

	// Disclose original information of a client connecting to a web server
	// through an HTTP proxy.
	FORWARDED request.Header = "Forwarded"

	// The email address of the user making the request.
	FROM request.Header = "From"

	// The domain name of the server (for virtual hosting), and the TCP port
	// number on which the server is listening. The port number may be omitted if
	// the port is the standard port for the service requested
	HOST request.Header = "Host"

	// Only perform the action if the client supplied entity matches the same
	// entity on the server. This is mainly for methods like PUT to only update a
	// resource if it has not been modified since the user last updated it.
	IF_MATCH request.Header = "If-Match"

	// Allows a 304 Not Modified to be returned if content is unchanged.
	IF_MODIFIED_SINCE request.Header = "If-Modified-Since"

	// Allows a 304 Not Modified to be returned if content is unchanged, see HTTP
	// ETag.
	IF_NONE_MATCH request.Header = "If-None-Match"

	// If the entity is unchanged, send me the part(s) that I am missing;
	// otherwise, send me the entire new entity.
	IF_RANGE request.Header = "If-Range"

	// Only send the response if the entity has not been modified since a specific
	// time.
	IF_UNMODIFIED_SINCE request.Header = "If-Unmodified-Since"

	// Limit the number of times the message can be forwarded through proxies or
	// gateways.
	MAX_FORWARDS request.Header = "Max-Forwards"

	// Initiates a request for cross-origin resource sharing (asks server for
	// Access-Control-* response fields).
	ORIGIN request.Header = "Origin"

	// Implementation-specific fields that may have various effects anywhere
	// along the request-response chain.
	PRAGMA request.Header = "Pragma"

	// Authorization credentials for connecting to a proxy.
	PROXY_AUTHORIZATION request.Header = "Proxy-Authorization"

	// Request only part of an entity. Bytes are numbered from 0. See Byte
	// serving.
	RANGE request.Header = "Range"

	// This is the address of the previous web page from which a link to the
	// currently requested page was followed. (The word “referrer” has been
	// misspelled in the RFC as well as in most implementations to the point that
	// it has become standard usage and is considered correct terminology)
	REFERER request.Header = "Referer"

	// The transfer encodings the user agent is willing to accept: the same values
	// as for the response header field Transfer-Encoding can be used, plus the
	// "trailers" value (related to the "chunked" transfer method) to notify the
	// server it expects to receive additional fields in the trailer after the
	// last, zero-sized, chunk
	TE request.Header = "TE"

	// The user agent string of the user agent.
	USER_AGENT request.Header = "User-Agent"

	// Ask the server to upgrade to another protocol
	UPGRADE request.Header = "Upgrade"

	// Informs the server of proxies through which the request was sent.
	VIA request.Header = "Via"

	// A general warning about possible problems with the entity body.
	WARNING request.Header = "Warning"
)
