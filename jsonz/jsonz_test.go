package jsonz

import (
	"bytes"
	"encoding/json"
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/statusz"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *http.Server {
	addr := &spb.Address{
		Host: "",
		Port: 1,
	}
	statuszServer := statusz.NewStatuszServer()
	return createServer(addr, statuszServer)
}

func TestServeStatuszWebpage(t *testing.T) {
	server := mockServer()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/statusz", nil)
	server.Handler.ServeHTTP(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}

func TestServe404(t *testing.T) {
	server := mockServer()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/nope", nil)
	server.Handler.ServeHTTP(w, req)
	res := w.Result()

	assert.Equal(t, 404, res.StatusCode)
}

func TestServeStatuszJson(t *testing.T) {
	server := mockServer()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/GetStatus", bytes.NewBufferString(`{}`))
	server.Handler.ServeHTTP(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("content-type"))

	bodyJson := struct {
		Status struct {
			Timestamp struct {
				Nanoseconds uint64
			}
			Groups []interface{}
		}
	}{}
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&bodyJson))

	assert.NotNil(t, bodyJson.Status.Groups)
	if bodyJson.Status.Timestamp.Nanoseconds == 0 {
		assert.Fail(t, "timestamp not in JSON response")
	}
}
