package creq

type UnmarshallerFunc func([]byte, interface{}) error

func (r UnmarshallerFunc) Unmarshal(dat []byte, in interface{}) error {
  return r(dat, in)
}
