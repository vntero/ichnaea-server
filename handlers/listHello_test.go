package handlers_test

import (
	"ichnaea-server/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListHello(t *testing.T) {
	expected := "ichnaea server is live!\n"
	
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.ListHello)
	handler.ServeHTTP(rr, req)

	// assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
