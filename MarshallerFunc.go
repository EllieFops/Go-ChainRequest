package creq

type MarshallerFunc func(interface{}) ([]byte, error)

func (self MarshallerFunc) Marshal(dat interface{}) ([]byte, error) {
  return self(dat)
}
