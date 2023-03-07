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

func TestSavePlacements(t *testing.T) {
	cases := []struct {
		desc       string
		placement  models.Placement
		request    string
		mockReturn error
		expected   string
	}{
		{
			"Placement must be saved",
			models.Placement{
				ProductCode: "992afd25-09bd-49e6-82de-c873923d8d09",
				WarehouseId: 1,
				QTY:         1,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SavePlacements",
				"params": [
					{
						"placements": [
							{
								"product_code": "992afd25-09bd-49e6-82de-c873923d8d09",
								"warehouse_id": 1,
								"qty": 1
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
			"Placement must be not saved",
			models.Placement{
				Id:          1,
				ProductCode: "992afd25-09bd-49e6-82de-c873923d8d09",
				WarehouseId: 1,
				QTY:         1,
			},
			`{
				"jsonrpc": "2.0",
				"method": "warehouse.SavePlacements",
				"params": [
					{
						"placements": [
							{
								"id": 1,
								"product_code": "992afd25-09bd-49e6-82de-c873923d8d09",
								"warehouse_id": 1,
								"qty": 1
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
		mockPlacementRepo := new(mocks.MockPlacementRepository)
		mockPlacementRepo.On("InTransaction", context.Background(), mock.Anything).Return(tc.mockReturn)

		service := services.NewService(nil, nil, mockPlacementRepo, testLogger)
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
		mockPlacementRepo.AssertExpectations(t)
	}
}
