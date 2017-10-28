package simple

type MarshallerFunc func(interface{}) ([]byte, error)

func (m MarshallerFunc) Marshal(dat interface{}) ([]byte, error) {
	return m(dat)
}
