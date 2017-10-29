package header

import "github.com/Foxcapades/Go-ChainRequest/request"

// Nonstandard Common Request Fields
const (
	// Mainly used to identify Ajax requests. Most JavaScript frameworks send this
	// field with value of XMLHttpRequest
	X_REQUESTED_WITH request.Header = "X-Requested-With"

	// Requests a web application to disable their tracking of a user. This is
	// Mozilla's version of the X-Do-Not-Track header field (since Firefox 4.0
	// Beta 11). Safari and IE9 also have support for this field. On March 7,
	// 2011, a draft proposal was submitted to IETF. The W3C Tracking Protection
	// Working Group is producing a specification.
	DNT request.Header = "DNT"

	// A de facto standard for identifying the originating IP address of a client
	// connecting to a web server through an HTTP proxy or load balancer
	X_FORWARDED_FOR request.Header = "X-Forwarded-For"

	// A de facto standard for identifying the original host requested by the
	// client in the Host HTTP request header, since the host name and/or port of
	// the reverse proxy (load balancer) may differ from the origin server
	// handling the request.
	X_FORWARDED_HOST request.Header = "X-Forwarded-Host"

	// A de facto standard for identifying the originating protocol of an HTTP
	// request, since a reverse proxy (or a load balancer) may communicate with a
	// web server using HTTP even if the request to the reverse proxy is HTTPS.
	// An alternative form of the header (X-ProxyUser-Ip) is used by Google
	// clients talking to Google servers.
	X_FORWARDED_PROTO request.Header = "X-Forwarded-Proto"

	// Non-standard header field used by Microsoft applications and load-balancers
	FRONT_END_HTTPS request.Header = "Front-End-Https"

	// Requests a web application to override the method specified in the request
	// (typically POST) with the method given in the header field (typically PUT
	// or DELETE). This can be used when a user agent or firewall prevents PUT or
	// DELETE methods from being sent directly (note that this is either a bug in
	// the software component, which ought to be fixed, or an intentional
	// configuration, in which case bypassing it may be the wrong thing to do).
	X_HTTP_METHOD_OVERRIDE request.Header = "X-Http-Method-Override"

	// Allows easier parsing of the MakeModel/Firmware that is usually found in
	// the User-Agent String of AT&T Devices
	X_ATT_DEVICEID request.Header = "X-ATT-DeviceId"

	// Links to an XML file on the Internet with a full description and details
	// about the device currently connecting. In the example to the right is an
	// XML file for an AT&T Samsung Galaxy S2.
	X_WAP_PROFILE request.Header = "X-Wap-Profile"

	// Implemented as a misunderstanding of the HTTP specifications. Common
	// because of mistakes in implementations of early HTTP versions. Has exactly
	// the same functionality as standard Connection field
	PROXY_CONNECTION request.Header = "Proxy-Connection"

	// Server-side deep packet insertion of a unique ID identifying customers of
	// Verizon Wireless; also known as "perma-cookie" or "supercookie"
	X_UIDH request.Header = "X-UIDH"

	// Used to prevent cross-site request forgery.
	X_CSRF_TOKEN request.Header = "X-Csrf-Token"

	// Used to prevent cross-site request forgery.
	X_XSRF_TOKEN request.Header = "X-XSRF-TOKEN"

	// Used to prevent cross-site request forgery.
	X_CSRFToken request.Header = "X-CSRFToken"

	// Correlates HTTP requests between a client and server.
	X_REQUEST_ID request.Header = "X-Request-ID"

	// Correlates HTTP requests between a client and server.
	X_CORRELATION_ID request.Header = "X-Correlation-ID"
)
