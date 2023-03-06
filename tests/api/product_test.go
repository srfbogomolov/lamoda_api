package api_test

import (
	"bytes"
	"context"
	jsonparse "encoding/json"
	"errors"
	"net/http/httptest"
	"net/rpc"
	"testing"

	"github.com/srfbogomolov/warehouse_api/internal/app"
	"github.com/srfbogomolov/warehouse_api/internal/models"
	"github.com/srfbogomolov/warehouse_api/internal/services"
	"github.com/srfbogomolov/warehouse_api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveProducts(t *testing.T) {
	cases := []struct {
		desc       string
		product    models.Product
		request    string
		mockReturn error
		expected   any
	}{
		{
			"Product must be saved",
			models.Product{
				ID:   0,
				Name: "test",
				Size: 0,
				Code: "0667da7a-5c13-4be3-8aba-b5005914f38c",
				QTY:  0,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SaveProducts",
				"params": [
					{
						"products": [
							{
								"id": 0,
								"name": "test",
								"size": 0,
								"code": "0667da7a-5c13-4be3-8aba-b5005914f38c",
								"qty": 0
							}
						]
					}
				],
				"id": 1
			}`,
			nil,
			"",
		},
		{
			"Product must be not saved",
			models.Product{
				ID:   0,
				Name: "test",
				Size: 0,
				Code: "0667da7a-5c13-4be3-8aba-b5005914f38c",
				QTY:  0,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SaveProducts",
				"params": [
					{
						"products": [
							{
								"id": 0,
								"name": "test",
								"size": 0,
								"code": "0667da7a-5c13-4be3-8aba-b5005914f38c",
								"qty": 0
							}
						]
					}
				],
				"id": 1
			}`,
			errors.New("error"),
			"transaction execution error",
		},
	}

	for _, tc := range cases {
		mockProductRepo := new(mocks.MockProductRepository)
		mockProductRepo.On("InTransaction", context.Background(), mock.Anything).Return(tc.mockReturn)

		service := services.NewService(nil, mockProductRepo, testLogger)
		handler := app.NewHandler(service)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/warehouse", bytes.NewBufferString(tc.request))
		req.Header.Set("Content-Type", "application/json")
		handler.ServeHTTP(recorder, req)

		resp := rpc.Response{}
		if err := jsonparse.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			panic(err)
		}

		assert.Equal(t, tc.expected, resp.Error)
		mockProductRepo.AssertExpectations(t)
	}
}
