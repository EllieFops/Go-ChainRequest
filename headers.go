package creq

type RequestHeader string

// Standard Request Fields
const (
  // Content-Types that are acceptable for the response.
  // See Content negotiation.
  ACCEPT RequestHeader = "Accept"

  // Character sets that are acceptable.
  ACCEPT_CHARSET RequestHeader = "Accept-Charset"

  // List of acceptable encodings. See HTTP compression.
  ACCEPT_ENCODING RequestHeader = "Accept-Encoding"

  // List of acceptable human languages for response. See Content negotiation.
  ACCEPT_LANGUAGE RequestHeader = "Accept-Language"

  // Acceptable version in time.
  ACCEPT_DATETIME RequestHeader = "Accept-Datetime"

  // Initiates a request for cross-origin resource sharing with Origin (below).
  ACCESS_CONTROL_REQUEST_METHOD RequestHeader = "Access-Control-Request-Method"

  //Initiates a request for cross-origin resource sharing with Origin (below).
  ACCESS_CONTROL_REQUEST_HEADERS RequestHeader = "Access-Control-Request-Headers"

  // Authentication credentials for HTTP authentication.
  AUTHORIZATION RequestHeader = "Authorization"

  // Used to specify directives that must be obeyed by all caching mechanisms
  // along the request-response chain.
  CACHE_CONTROL RequestHeader = "Cache-Control"

  // Control options for the current connection and list of hop-by-hop request
  // fields.
  CONNECTION RequestHeader = "Connection"

  // An HTTP cookie previously sent by the server with Set-Cookie (below).
  COOKIE RequestHeader = "Cookie"

  // The length of the request body in octets (8-bit bytes).
  CONTENT_LENGTH RequestHeader = "Content-Length"

  // A Base64-encoded binary MD5 sum of the content of the request body.
  CONTENT_MD5 RequestHeader = "Content-MD5"

  // The MIME type of the body of the request (used with POST and PUT requests).
  CONTENT_TYPE RequestHeader = "Content-Type"

  // The date and time that the message was originated (in "HTTP-date" format as
  // defined by RFC 7231 Date/Time Formats).
  DATE RequestHeader = "Date"

  // Indicates that particular server behaviors are required by the client.
  EXPECT RequestHeader = "Expect"

  // Disclose original information of a client connecting to a web server
  // through an HTTP proxy.
  FORWARDED RequestHeader = "Forwarded"

  // The email address of the user making the request.
  FROM RequestHeader = "From"

  // The domain name of the server (for virtual hosting), and the TCP port
  // number on which the server is listening. The port number may be omitted if
  // the port is the standard port for the service requested
  HOST RequestHeader = "Host"

  // Only perform the action if the client supplied entity matches the same
  // entity on the server. This is mainly for methods like PUT to only update a
  // resource if it has not been modified since the user last updated it.
  IF_MATCH RequestHeader = "If-Match"

  // Allows a 304 Not Modified to be returned if content is unchanged.
  IF_MODIFIED_SINCE RequestHeader = "If-Modified-Since"

  // Allows a 304 Not Modified to be returned if content is unchanged, see HTTP
  // ETag.
  IF_NONE_MATCH RequestHeader = "If-None-Match"

  // If the entity is unchanged, send me the part(s) that I am missing;
  // otherwise, send me the entire new entity.
  IF_RANGE RequestHeader = "If-Range"

  // Only send the response if the entity has not been modified since a specific
  // time.
  IF_UNMODIFIED_SINCE RequestHeader = "If-Unmodified-Since"

  // Limit the number of times the message can be forwarded through proxies or
  // gateways.
  MAX_FORWARDS RequestHeader = "Max-Forwards"

  // Initiates a request for cross-origin resource sharing (asks server for
  // Access-Control-* response fields).
  ORIGIN RequestHeader = "Origin"

  // Implementation-specific fields that may have various effects anywhere
  // along the request-response chain.
  PRAGMA RequestHeader = "Pragma"

  // Authorization credentials for connecting to a proxy.
  PROXY_AUTHORIZATION RequestHeader = "Proxy-Authorization"

  // Request only part of an entity. Bytes are numbered from 0. See Byte
  // serving.
  RANGE RequestHeader = "Range"

  // This is the address of the previous web page from which a link to the
  // currently requested page was followed. (The word “referrer” has been
  // misspelled in the RFC as well as in most implementations to the point that
  // it has become standard usage and is considered correct terminology)
  REFERER RequestHeader = "Referer"

  // The transfer encodings the user agent is willing to accept: the same values
  // as for the response header field Transfer-Encoding can be used, plus the
  // "trailers" value (related to the "chunked" transfer method) to notify the
  // server it expects to receive additional fields in the trailer after the
  // last, zero-sized, chunk
  TE RequestHeader = "TE"

  // The user agent string of the user agent.
  USER_AGENT RequestHeader = "User-Agent"

  // Ask the server to upgrade to another protocol
  UPGRADE RequestHeader = "Upgrade"

  // Informs the server of proxies through which the request was sent.
  VIA RequestHeader = "Via"

  // A general warning about possible problems with the entity body.
  WARNING RequestHeader = "Warning"
)

// Nonstandard Common Request Fields
const (
  // Mainly used to identify Ajax requests. Most JavaScript frameworks send this
  // field with value of XMLHttpRequest
  X_REQUESTED_WITH RequestHeader = "X-Requested-With"

  // Requests a web application to disable their tracking of a user. This is
  // Mozilla's version of the X-Do-Not-Track header field (since Firefox 4.0
  // Beta 11). Safari and IE9 also have support for this field. On March 7,
  // 2011, a draft proposal was submitted to IETF. The W3C Tracking Protection
  // Working Group is producing a specification.
  DNT RequestHeader = "DNT"

  // A de facto standard for identifying the originating IP address of a client
  // connecting to a web server through an HTTP proxy or load balancer
  X_FORWARDED_FOR RequestHeader = "X-Forwarded-For"

  // A de facto standard for identifying the original host requested by the
  // client in the Host HTTP request header, since the host name and/or port of
  // the reverse proxy (load balancer) may differ from the origin server
  // handling the request.
  X_FORWARDED_HOST RequestHeader = "X-Forwarded-Host"

  // A de facto standard for identifying the originating protocol of an HTTP
  // request, since a reverse proxy (or a load balancer) may communicate with a
  // web server using HTTP even if the request to the reverse proxy is HTTPS.
  // An alternative form of the header (X-ProxyUser-Ip) is used by Google
  // clients talking to Google servers.
  X_FORWARDED_PROTO RequestHeader = "X-Forwarded-Proto"

  // Non-standard header field used by Microsoft applications and load-balancers
  FRONT_END_HTTPS RequestHeader = "Front-End-Https"

  // Requests a web application to override the method specified in the request
  // (typically POST) with the method given in the header field (typically PUT
  // or DELETE). This can be used when a user agent or firewall prevents PUT or
  // DELETE methods from being sent directly (note that this is either a bug in
  // the software component, which ought to be fixed, or an intentional
  // configuration, in which case bypassing it may be the wrong thing to do).
  X_HTTP_METHOD_OVERRIDE RequestHeader = "X-Http-Method-Override"

  // Allows easier parsing of the MakeModel/Firmware that is usually found in
  // the User-Agent String of AT&T Devices
  X_ATT_DEVICEID RequestHeader = "X-ATT-DeviceId"

  // Links to an XML file on the Internet with a full description and details
  // about the device currently connecting. In the example to the right is an
  // XML file for an AT&T Samsung Galaxy S2.
  X_WAP_PROFILE RequestHeader = "X-Wap-Profile"

  // Implemented as a misunderstanding of the HTTP specifications. Common
  // because of mistakes in implementations of early HTTP versions. Has exactly
  // the same functionality as standard Connection field
  PROXY_CONNECTION RequestHeader = "Proxy-Connection"

  // Server-side deep packet insertion of a unique ID identifying customers of
  // Verizon Wireless; also known as "perma-cookie" or "supercookie"
  X_UIDH RequestHeader = "X-UIDH"

  // Used to prevent cross-site request forgery.
  X_CSRF_TOKEN RequestHeader = "X-Csrf-Token"

  // Used to prevent cross-site request forgery.
  X_XSRF_TOKEN RequestHeader = "X-XSRF-TOKEN"

  // Used to prevent cross-site request forgery.
  X_CSRFToken  RequestHeader = "X-CSRFToken"

  // Correlates HTTP requests between a client and server.
  X_REQUEST_ID RequestHeader = "X-Request-ID"

  // Correlates HTTP requests between a client and server.
  X_CORRELATION_ID RequestHeader = "X-Correlation-ID"
)
