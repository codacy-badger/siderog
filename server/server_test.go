package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestStartSession(t *testing.T) {
	sess := Session{}
	mockConn := redigomock.NewConn()
	mockConn.Clear()

	server := httptest.NewServer(sess.Start(mockConn))
	defer server.Close()

	// register redigomock connection
	mockConn.Command("SET", "029384028095203892", "true").Expect("029384028095203892")

	expectedToken := "029384028095203892"
	urlQuery := fmt.Sprintf("%v/start_session?token=%v", server.URL, expectedToken)
	req, err := http.Get(urlQuery)

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	output := string(result)

	if output != expectedToken {
		t.Errorf("output %v not expected", output)
	}
}

func TestCheckSession(t *testing.T) {
	sess := Session{}
	mockConn := redigomock.NewConn()
	token := "029384028095203892"
	expectResult := "true"

	mockConn.Command("GET", "029384028095203892").Expect([]uint8{'t', 'r', 'u', 'e'})

	server := httptest.NewServer(sess.Check(mockConn))
	defer server.Close()

	urlQuery := fmt.Sprintf("%v/check_session?token=%v", server.URL, token)
	req, err := http.Get(urlQuery)

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	output := string(result)

	if output != expectResult {
		t.Errorf("%v is not what I expect to read", output)
	}
}
