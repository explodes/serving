package statusz_test

import (
	"github.com/explodes/serving/statusz"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterStatuszWebpage(t *testing.T) {
	mux := http.NewServeMux()
	statusz.RegisterStatuszWebpage("foo:65535", mux)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/statusz", nil)

	handler, pattern := mux.Handler(r)

	assert.NotNil(t, handler)
	assert.Equal(t, "/statusz", pattern)

	handler.ServeHTTP(w, r)
	responseBody := readBody(t, w)

	assert.Equal(t, 200, w.Result().StatusCode)
	assert.True(t, strings.Contains(responseBody, "http://foo:65535/GetStatus?ts="))
}

func TestRegisterStatuszWebpage_brokenPipe(t *testing.T) {
	mux := http.NewServeMux()
	statusz.RegisterStatuszWebpage("foo:65535", mux)
	w := &brokenPipe{}
	r := httptest.NewRequest("GET", "/statusz", nil)

	handler, pattern := mux.Handler(r)

	assert.NotNil(t, handler)
	assert.Equal(t, "/statusz", pattern)

	handler.ServeHTTP(w, r)

	assert.Equal(t, 200, w.statusCode)
}
