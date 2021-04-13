package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHandler)

	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"message": "Doing Things"}`
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
