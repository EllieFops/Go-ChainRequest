package simple

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/Foxcapades/Go-ChainRequest/request/header"
	"github.com/Foxcapades/Go-ChainRequest/request/method"
)

func TestRequest_AddHeaderCreatesNewHeaderWhenKeyDoesNotYetExist(t *testing.T) {
	test := newRequest("test", method.GET)
	head := header.ACCEPT
	value := "test value"
	test.AddHeader(head, value)

	if _, ok := test.headers[head]; !ok {
		t.Error()
	}
}

func TestRequest_AddHeaderAppendsToHeaderWhenKeyAlreadyExists(t *testing.T) {
	test := newRequest("test", method.GET)
	head := header.ACCEPT
	value1 := "test value"
	value2 := "other value"
	expected := "test value; other value"
	test.AddHeader(head, value1).AddHeader(head, value2)

	if val := test.headers[head]; val != expected {
		t.Error()
	}
}

func TestRequest_SetBody(t *testing.T) {
	test := newRequest("test", method.GET)
	expe := "test value"
	body := []byte(expe)

	test.SetBody(body)

	if string(test.body) != expe {
		t.Error()
	}
}

func TestRequest_GetBody_WhenNilReturnsEmptySliceAndFalse(t *testing.T) {
	test := newRequest("test", method.GET)

	if len(test.GetBody()) != 0 {
		t.Error()
	}
}

func TestRequest_GetBody_WhenNotNilReturnsBodyAndTrue(t *testing.T) {
	test := newRequest("test", method.GET)
	expe := "test value"
	body := []byte(expe)
	test.body = body
	if string(test.GetBody()) != expe {
		t.Error()
	}
}

func TestRequest_SetHeaderCreatesNewHeaderWhenKeyDoesNotYetExist(t *testing.T) {
	test := newRequest("test", method.GET)
	head := header.ACCEPT
	value := "test value"
	test.SetHeader(head, value)

	if _, ok := test.headers[head]; !ok {
		t.Error()
	}
}

func TestRequest_SetHeaderOverwritesHeaderWhenKeyAlreadyExists(t *testing.T) {
	test := newRequest("test", method.GET)
	head := header.ACCEPT
	value1 := "test value"
	value2 := "other value"
	test.SetHeader(head, value1).SetHeader(head, value2)

	if val := test.headers[head]; val != value2 {
		t.Error()
	}
}

func TestRequest_GetHeaderReturnsFalseIfKeyWasNotSet(t *testing.T) {
	test := newRequest("test", method.GET)
	if _, ok := test.GetHeader(header.CONTENT_TYPE); ok {
		t.Error()
	}
}

func TestRequest_GetHeaderReturnsValueAndTrueIfKeyWasSet(t *testing.T) {
	test := newRequest("test", method.GET)
	head := header.ACCEPT
	value1 := "test value"

	test.headers[head] = value1

	if val, ok := test.GetHeader(head); val != value1 || !ok {
		t.Error()
	}
}

func TestRequest_SetMethod(t *testing.T) {
	test := newRequest("test", method.GET)
	expe := method.CONNECT

	test.SetMethod(expe)

	if test.method != expe {
		t.Error()
	}
}

func TestRequest_GetMethod(t *testing.T) {
	test := newRequest("test", method.GET)
	expe := method.CONNECT

	test.method = expe

	if test.GetMethod() != expe {
		t.Error()
	}
}

func TestRequest_MarshalBody_StoresErrorsReturnedByGivenMarshaller(t *testing.T) {
	test := newRequest("test", method.GET)
	err := errors.New("test error")
	fun := func(interface{}) ([]byte, error) { return nil, err }

	test.MarshalBody(nil, MarshallerFunc(fun))

	if test.err != err {
		t.Error()
	}
}

func TestRequest_MarshalBody_StoresBodyReturnedByGivenMarshaller(t *testing.T) {
	test := newRequest("test", method.GET)
	data := "test body"
	exp := []byte(data)
	fun := func(interface{}) ([]byte, error) { return exp, nil }

	test.MarshalBody(nil, MarshallerFunc(fun))

	if string(test.body) != data {
		t.Error()
	}
}

func TestRequest_assignHeaders(t *testing.T) {
	test := newRequest("test", method.GET)
	data := http.Header{}
	exp1 := "application/json"
	exp2 := "foo bar"

	test.AddHeader(header.CONTENT_TYPE, exp1).
		AddHeader(header.AUTHORIZATION, exp2)

	test.assignHeaders(data)

	if a, b := data[string(header.CONTENT_TYPE)]; a[0] != exp1 || !b {
		t.Error()
	}

	if a, b := data[string(header.AUTHORIZATION)]; a[0] != exp2 || !b {
		t.Error()
	}
}

type dummyClient struct {
	Response  *http.Response
	Error     error
	CallCount uint
}

func (d *dummyClient) Do(*http.Request) (*http.Response, error) {
	d.CallCount++
	return d.Response, d.Error
}

func TestRequest_SubmitReturnsResponseOnSuccess(t *testing.T) {
	test := newRequest("test.test", method.GET)
	client := dummyClient{}
	expect := http.Response{Status: "Test Response"}

	client.Response = &expect
	test.client = &client

	res, err := test.Submit().GetRawResponse()

	if err != nil || !reflect.DeepEqual(res, &expect) {
		t.Error("Expected (", &expect, ", <nil> ) got (", res, ",", err, ")")
	}
}

func TestRequest_SubmitReturnsErrorOnFail(t *testing.T) {
	test := newRequest("test.test", method.GET)
	client := dummyClient{}
	expect := errors.New("test error")

	client.Error = expect
	test.client = &client

	res, err := test.Submit().GetRawResponse()

	if res != nil || !reflect.DeepEqual(err, expect) {
		t.Error("Expected (", expect, ", <nil> ) got (", res, ",", err, ")")
	}
}

func TestRequest_SubmitReturnsErrorIfExistsBeforeRequestSend(t *testing.T) {
	test := newRequest("test.test", method.GET)
	expect := errors.New("test error")
	test.err = expect

	err := test.Submit().GetError()

	if !reflect.DeepEqual(err, expect) {
		t.Error("Expected", expect, "got", err)
	}
}
