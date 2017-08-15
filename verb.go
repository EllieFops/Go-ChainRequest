package creq

type HttpVerb string

const (
  CONNECT HttpVerb = "CONNECT"
  DELETE  HttpVerb = "DELETE"
  GET     HttpVerb = "GET"
  HEAD    HttpVerb = "HEAD"
  INFO    HttpVerb = "INFO"
  OPTIONS HttpVerb = "OPTIONS"
  PATCH   HttpVerb = "PATCH"
  POST    HttpVerb = "POST"
  PUT     HttpVerb = "PUT"
  TRACE   HttpVerb = "TRACE"
)
