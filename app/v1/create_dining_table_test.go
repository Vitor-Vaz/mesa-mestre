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

func TestCreateDiningTableHandler(t *testing.T) {
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
			name:           "dining table already exists",
			expectedResp:   `{"title": "Conflict", "status": 409, "detail": "Dining table already exists"}`,
			mockError:      domain.ErrConflict,
			expectedStatus: http.StatusConflict,
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

			diningTableCreator := setupdiningTableCreatorMock(tt.mockError)

			handler := v1.NewDiningTableHandler(diningTableCreator)

			r := testhelpers.CreatePostApiRouter("/api/v1/dining-table", handler.CreateDiningTableHandler)

			server := httptest.NewServer(r.C)
			defer server.Close()

			body := map[string]interface{}{
				"number": 1,
				"seats":  4,
			}
			bodyBytes, _ := json.Marshal(body)

			req, _ := http.NewRequest(http.MethodPost, server.URL+"/api/v1/dining-table", bytes.NewReader(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			respBody, _ := io.ReadAll(resp.Body)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			if tt.expectedResp != "" {
				assert.JSONEq(t, tt.expectedResp, string(respBody))
			}
		})
	}
}

func setupdiningTableCreatorMock(mockError error) *mocks.DiningTableCreatorMock {
	return &mocks.DiningTableCreatorMock{
		CreateDiningTableFunc: func(ctx context.Context, number int32, seats int32) error {
			return mockError
		},
	}
}
