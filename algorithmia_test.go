package algorithmia

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	client = NewClient("test-token")
)

func TestQuerySuccessful(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testQuerySuccessfulHandler))
	defer server.Close()
	algorithmiaURI = server.URL + "/%s/%s"

	res, err := client.Query("user", "project", 123)
	if err != nil {
		t.Error(err)
	}

	if res.Result != "some-result" {
		t.Error("unexpected result", res.Result)
	}
}

func testQuerySuccessfulHandler(rw http.ResponseWriter, req *http.Request) {
	res := new(Response)
	res.Result = "some-result"

	encoder := json.NewEncoder(rw)
	encoder.Encode(&res)

	rw.WriteHeader(http.StatusOK)
}
