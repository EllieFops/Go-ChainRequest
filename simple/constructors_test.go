package simple

import (
	"testing"

	"github.com/Foxcapades/Go-ChainRequest/request/method"
)

func TestConnectRequest(t *testing.T) {
	if a := ConnectRequest("test"); a.GetMethod() != method.CONNECT {
		t.Error("Expected", method.CONNECT, "got", a.GetMethod())
	}
}

func TestDeleteRequest(t *testing.T) {
	if a := DeleteRequest("test"); a.GetMethod() != method.DELETE {
		t.Error("Expected", method.DELETE, "got", a.GetMethod())
	}
}

func TestPostRequest(t *testing.T) {
	if a := PostRequest("test"); a.GetMethod() != method.POST {
		t.Error("Expected", method.POST, "got", a.GetMethod())
	}
}

func TestPatchRequest(t *testing.T) {
	if a := PatchRequest("test"); a.GetMethod() != method.PATCH {
		t.Error("Expected", method.PATCH, "got", a.GetMethod())
	}
}

func TestPutRequest(t *testing.T) {
	if a := PutRequest("test"); a.GetMethod() != method.PUT {
		t.Error("Expected", method.PUT, "got", a.GetMethod())
	}
}

func TestGetRequest(t *testing.T) {
	if a := GetRequest("test"); a.GetMethod() != method.GET {
		t.Error("Expected", method.GET, "got", a.GetMethod())
	}
}

func TestHeadRequest(t *testing.T) {
	if a := HeadRequest("test"); a.GetMethod() != method.HEAD {
		t.Error("Expected", method.HEAD, "got", a.GetMethod())
	}
}

func TestOptionsRequest(t *testing.T) {
	if a := OptionsRequest("test"); a.GetMethod() != method.OPTIONS {
		t.Error("Expected", method.OPTIONS, "got", a.GetMethod())
	}
}

func TestInfoRequest(t *testing.T) {
	if a := InfoRequest("test"); a.GetMethod() != method.INFO {
		t.Error("Expected", method.INFO, "got", a.GetMethod())
	}
}

func TestTraceRequest(t *testing.T) {
	if a := TraceRequest("test"); a.GetMethod() != method.TRACE {
		t.Error("Expected", method.TRACE, "got", a.GetMethod())
	}
}
