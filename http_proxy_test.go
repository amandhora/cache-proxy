package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleReqGet(t *testing.T) {

	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/proxy", nil)
	if err != nil {
		t.Fatal(err)

	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleReq)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)

	}

}

func TestHandleReqNoQParams(t *testing.T) {

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/proxy", nil)
	if err != nil {
		t.Fatal(err)

	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleReq)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)

	}

}

func TestHandleReqInvalidQParam(t *testing.T) {

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/proxy", nil)
	if err != nil {
		t.Fatal(err)

	}
	q := req.URL.Query()
	q.Add("unknown", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleReq)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code: got %v want %v",
			status, http.StatusOK)

	}

}
