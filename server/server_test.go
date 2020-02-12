package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartSession(t *testing.T) {
	server := httptest.NewServer(StartSession())
	defer server.Close()

	req, err := http.Get(server.URL)

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	output := string(result)
	expected := "hello session"

	if output != expected {
		t.Errorf("output %v not expected", output)
	}
}
