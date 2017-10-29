package simple

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Foxcapades/Go-ChainRequest/request/method"
)

func TestRequestBuilder_Build(t *testing.T) {
	expUrl := "test.com"
	expMeth := method.PATCH
	expBody := "body text"

	test1 := RequestBuilder(http.NewRequest)
	test2 := newRequest(expUrl, expMeth)

	test2.SetBody([]byte(expBody))

	req, err := test1.Build(test2)

	if err != nil {
		t.Error(err)
	}

	if req.URL.String() != expUrl {
		t.Error("Expected", expUrl, "got", req.URL.String())
	}

	if req.Method != string(expMeth) {
		t.Error("Expected", expMeth, "got", req.Method)
	}

	body, err := req.GetBody()

	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		t.Error(err)
	}

	if expBody != string(bytes) {
		t.Error("Expected", expBody, "got", string(bytes))
	}
}

func TestGetBodyReader_ReturnsReaderIfBodyIsNotEmpty(t *testing.T) {
	exp1 := "test data string"
	data := []byte(exp1)
	actu := make([]byte, len(data))

	read := getBodyReader(data)

	read.Read(actu)

	if exp1 != string(actu) {
		t.Error()
	}
}

func TestGetBodyReader_ReturnsNilIfBodyIsEmpty(t *testing.T) {
	if getBodyReader([]byte{}) != nil {
		t.Error()
	}
}
