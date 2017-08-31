package creq

type ResponseUnmarshaller interface {
  Unmarshal([]byte, interface{}) error
}
