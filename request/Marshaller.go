package request

type Marshaller interface {
	Marshal(interface{}) ([]byte, error)
}
