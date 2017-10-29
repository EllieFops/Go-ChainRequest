package response

type Unmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
