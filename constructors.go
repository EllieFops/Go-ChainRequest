package creq

func NewGetRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  GET,
  }
}

func NewPostRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  POST,
  }
}

func NewPutRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  PUT,
  }
}

func NewDeleteRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  DELETE,
  }
}

func NewPatchRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  PATCH,
  }
}

func NewHeadRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  HEAD,
  }
}

func NewOptionsRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  OPTIONS,
  }
}

func NewConnectRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  CONNECT,
  }
}

func NewInfoRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  INFO,
  }
}

func NewTraceRequest(url string) Request {
  return &request{
    url:     url,
    headers: map[RequestHeader]string{},
    method:  TRACE,
  }
}
