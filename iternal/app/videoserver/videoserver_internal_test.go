package videoserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideoServer_HandlHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	//assert.Equal(t, req.Body.String(), "Hello world!")
	assert.Equal(t, http.StatusOK, rec.Code)
}
