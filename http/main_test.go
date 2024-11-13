package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	app := setup()

	tests := []struct {
		name           string
		reqBody        User
		expectedStatus int
	}{

		{"Valid user", User{Email: "test@example.com", Fullname: "John Doe", Age: 25}, 200},
		{"Invalid email", User{Email: "invalid-email", Fullname: "John Doe", Age: 25}, 400},
		{"Invalid fullname", User{Email: "test@example.com", Fullname: "", Age: 25}, 400},
		{"Invalid age", User{Email: "test@example.com", Fullname: "John Doe", Age: 0}, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.reqBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectedStatus, resp.StatusCode)
		})
	}
}
