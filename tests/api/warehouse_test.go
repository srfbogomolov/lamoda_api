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

func TestSaveWarehouses(t *testing.T) {
	cases := []struct {
		desc       string
		warehouse  models.Warehouse
		request    string
		mockReturn error
		expected   any
	}{
		{
			"Warehouse must be saved",
			models.Warehouse{
				ID:         1,
				Name:       "test",
				IsAvalible: true,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SaveWarehouses",
				"params": [
					{
						"warehouses": [
							{
								"id": 1,
								"name": "test",
								"is_available": true
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
			"Warehouse must be not saved",
			models.Warehouse{
				ID:         1,
				Name:       "test",
				IsAvalible: true,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SaveWarehouses",
				"params": [
					{
						"warehouses": [
							{
								"id": 1,
								"name": "test",
								"is_available": true
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
		mockWarehouseRepo := new(mocks.MockWarehouseRepository)
		mockWarehouseRepo.On("InTransaction", context.Background(), mock.Anything).Return(tc.mockReturn)

		service := services.NewService(mockWarehouseRepo, nil, testLogger)
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
		mockWarehouseRepo.AssertExpectations(t)
	}
}
