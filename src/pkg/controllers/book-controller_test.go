package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

}
