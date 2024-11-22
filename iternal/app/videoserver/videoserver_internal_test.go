package videoserver

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"video/iternal/app/store/teststore"
)

func TestServer_HandlUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	tc := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@user.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"email":    "userr.com",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tt.payload)
			req, _ := http.NewRequest("POST", "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
		})
	}
}
