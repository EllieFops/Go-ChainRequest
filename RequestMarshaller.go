package creq

type RequestMarshaller interface {
  Marshal(interface{}) ([]byte, error)
}
