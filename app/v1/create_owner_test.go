package v1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mesa-mestre/app/v1/mocks"
	"mesa-mestre/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "mesa-mestre/app/v1"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateOwnerHandler(t *testing.T) {
	tests := []struct {
		name           string
		mockError      error
		expectedStatus int
	}{
		{
			name:           "successful creation",
			mockError:      nil,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "owner already exists",
			mockError:      domain.ErrConflict,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "internal error",
			mockError:      domain.ErrUnexpected,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// router + huma
			r := chi.NewRouter()
			cfg := huma.DefaultConfig("Mesa Mestre API", "1.0.0")
			cfg.Transformers = nil
			api := humachi.New(r, cfg)

			// mock
			ownerCreator := setupOwnerCreatorMock(tt.mockError)

			// register handler
			handler := v1.NewOwnerHandler(ownerCreator)
			huma.Post(api, "/api/v1/owners", handler.CreateOwnerHandler)

			// test server
			server := httptest.NewServer(r)
			defer server.Close()

			// request body
			body := map[string]interface{}{
				"name":  "Michael Scott",
				"email": "michael.scott@dundermifflin.com",
			}

			payload, _ := json.Marshal(body)

			req, err := http.NewRequest(
				http.MethodPost,
				server.URL+"/api/v1/owners",
				bytes.NewBuffer(payload),
			)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			// execute
			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)

			// assert
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

		})
	}
}

func setupOwnerCreatorMock(err error) *mocks.OnwerCreatorMock {
	return &mocks.OnwerCreatorMock{
		CreateOwnerFunc: func(ctx context.Context, name string, email string) error {
			return err
		},
	}
}
