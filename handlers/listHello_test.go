package handlers_test

import (
	"ichnaea-server/handlers"
	"log"
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

	log.Printf("handler: %v\n", handler)
	log.Printf("rr.Code: %v\n", rr.Code)
	log.Printf("rr.Body.String(): %v\n", rr.Body.String())

	// assertions
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
