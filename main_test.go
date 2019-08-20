package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()

	res, err := http.Get(testServer.URL)
	defer res.Body.Close()
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ioutil read all error")
	}
	if string(body) != "Hello, World" {
		t.Errorf("got:%+v, wont:%+v", string(body), "Hello, World")
	}
}
