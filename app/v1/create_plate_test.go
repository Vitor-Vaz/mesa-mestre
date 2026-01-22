package v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	v1 "mesa-mestre/app/v1"
	"mesa-mestre/app/v1/mocks"
	"mesa-mestre/domain"
	"mesa-mestre/extension/testhelpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlateHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedResp   string
		mockError      error
		expectedStatus int
	}{
		{
			name:           "successful creation",
			expectedResp:   "",
			mockError:      nil,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "internal error",
			expectedResp:   `{"title":"Internal Server Error","status":500,"detail":"Internal server error"}`,
			mockError:      domain.ErrUnexpected,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			plateCreator := setupPlateCreatorMock(tt.mockError)

			handler := v1.NewPlateHandler(plateCreator)

			r := testhelpers.CreatePostApiRouter("/api/v1/plates", handler.CreatePlateHandler)

			server := httptest.NewServer(r.C)
			defer server.Close()

			body := map[string]interface{}{
				"name":              "Spaghetti Carbonara",
				"plate_description": "Classic Italian pasta dish with eggs, cheese, pancetta, and pepper.",
				"price":             25.50,
			}
			bodyBytes, _ := json.Marshal(body)

			resp, err := http.Post(server.URL+"/api/v1/plates", "application/json", bytes.NewReader(bodyBytes))
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus != http.StatusNoContent {
				respBody, _ := io.ReadAll(resp.Body)
				assert.JSONEq(t, tt.expectedResp, string(respBody))
			}

		})
	}
}

func setupPlateCreatorMock(mockError error) *mocks.PlateCreatorMock {
	return &mocks.PlateCreatorMock{
		CreatePlateFunc: func(ctx context.Context, name, description string, price float64) error {
			return mockError
		},
	}
}
