package services_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"net/rpc"
	"testing"

	"github.com/gorilla/mux"
	grlrpc "github.com/gorilla/rpc"

	grljson "github.com/gorilla/rpc/json"
	"github.com/srfbogomolov/warehouse_api/internal/app"
	"github.com/srfbogomolov/warehouse_api/internal/domain"
	"github.com/srfbogomolov/warehouse_api/internal/services"
	repomock "github.com/srfbogomolov/warehouse_api/internal/services/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

type testService struct {
	item      domain.ItemRepository
	product   domain.ProductRepository
	warehouse domain.WarehouseRepository
}

func init() {
	testLogger = createTestLogger()
}

func createTestLogger() *zap.SugaredLogger {
	cfgJSON := []byte(`{"level": "info", "encoding": "json"}`)
	var cfg zap.Config
	if err := json.Unmarshal(cfgJSON, &cfg); err != nil {
		panic(err)
	}
	return zap.Must(cfg.Build()).Sugar()
}

func createTestHandler(service *services.LamodaService) *mux.Router {
	server := grlrpc.NewServer()
	server.RegisterCodec(grljson.NewCodec(), app.CodecContentType)
	server.RegisterService(service, app.ServiceName)

	router := mux.NewRouter()
	router.Handle(app.ServicePath, server)

	return router
}

func TestLamodaService_GetInStock(t *testing.T) {
	type getByIdResult struct {
		warehouse domain.Warehouse
		err       error
	}
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, nil)

	type testCase struct {
		test       string
		request    string
		mockReturn getByIdResult
		expected   string
	}

	testCases := []testCase{
		{
			test: "Empty args",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.GetInStock",
				"params": [],
				"id": 3
			}`,
			mockReturn: getByIdResult{},
			expected:   "warehouse id is not specified",
		},
		{
			test: "Warehouse not exists",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.GetInStock",
				"params": [
					{
						"warehouse_id": 999
					}
				],
				"id": 3
			}`,
			mockReturn: getByIdResult{
				warehouse: domain.Warehouse{},
				err:       domain.ErrWarehouseNotFound,
			},
			expected: "warehouse not found",
		},
		{
			test: "Warehouse not exists",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.GetInStock",
				"params": [
					{
						"warehouse_id": 1
					}
				],
				"id": 3
			}`,
			mockReturn: getByIdResult{
				warehouse: domain.Warehouse{},
				err:       domain.ErrWarehouseNotFound,
			},
			expected: "warehouse not found",
		},
		{
			test: "Warehouse is not available",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.GetInStock",
				"params": [
					{
						"warehouse_id": 1
					}
				],
				"id": 3
			}`,
			mockReturn: getByIdResult{
				warehouse: notAvailableWarehouse,
				err:       nil,
			},
			expected: "warehouse is not available",
		},
	}

	for _, tc := range testCases {
		mockWarehouseRepository := new(repomock.MockWarehouseRepository)
		mockWarehouseRepository.On("GetById", context.Background(), mock.Anything).Return(tc.mockReturn.warehouse, tc.mockReturn.err)

		testService := services.LamodaService{Warehouse: mockWarehouseRepository}

		handler := createTestHandler(&testService)

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("POST", app.ServicePath, bytes.NewBufferString(tc.request))
		req.Header.Set("Content-Type", app.CodecContentType)
		handler.ServeHTTP(recorder, req)

		resp := rpc.Response{}
		if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			panic(err)
		}

		assert.Equal(t, tc.expected, resp.Error)
	}
}

func TestLamodaService_Reserve(t *testing.T) {
	type getByIdReturn struct {
		warehouse domain.Warehouse
		err       error
	}
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, nil)
	type testCase struct {
		test          string
		request       string
		getByIdReturn getByIdReturn
		reserveReturn error
		expected      string
	}

	testCases := []testCase{
		{
			test: "Warehouse id is not specified",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Reserve",
				"params": [],
				"id": 3
			}`,
			expected: "warehouse id is not specified",
		},
		{
			test: "Product codes are not specified",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Reserve",
				"params": [{
					"warehouse_id": 1
				}],
				"id": 3
			}`,
			expected: "product codes are not specified",
		},
		{
			test: "Warehouse not exists",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Reserve",
				"params": [
					{
						"warehouse_id": 1,
						"product_codes": ["b29a07c5-472f-4778-bb04-dab16e9502bb"]
					}
				],
				"id": 3
			}`,
			getByIdReturn: getByIdReturn{
				warehouse: domain.Warehouse{},
				err:       domain.ErrWarehouseNotFound,
			},
			expected: "warehouse not found",
		},
		{
			test: "Warehouse is not available",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Reserve",
				"params": [
					{
						"warehouse_id": 1,
						"product_codes": ["b29a07c5-472f-4778-bb04-dab16e9502bb"]
					}
				],
				"id": 3
			}`,
			getByIdReturn: getByIdReturn{
				warehouse: notAvailableWarehouse,
				err:       nil,
			},
			expected: "warehouse is not available",
		},
	}

	for _, tc := range testCases {
		mockWarehouseRepository := new(repomock.MockWarehouseRepository)
		mockWarehouseRepository.On("GetById", context.Background(), mock.Anything).
			Return(tc.getByIdReturn.warehouse, tc.getByIdReturn.err)
		mockWarehouseRepository.On("Reserve", context.Background(), mock.Anything).
			Return(tc.reserveReturn)

		testService := services.LamodaService{
			Warehouse: mockWarehouseRepository,
		}

		handler := createTestHandler(&testService)

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("POST", app.ServicePath, bytes.NewBufferString(tc.request))
		req.Header.Set("Content-Type", app.CodecContentType)
		handler.ServeHTTP(recorder, req)

		resp := rpc.Response{}
		if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			panic(err)
		}

		assert.Equal(t, tc.expected, resp.Error)
	}
}

func TestLamodaService_Release(t *testing.T) {
	type getByIdReturn struct {
		warehouse domain.Warehouse
		err       error
	}
	notAvailableWarehouse, _ := domain.NewWarehouse("test", false, nil)
	type testCase struct {
		test          string
		request       string
		getByIdReturn getByIdReturn
		reserveReturn error
		expected      string
	}

	testCases := []testCase{
		{
			test: "Warehouse id is not specified",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Release",
				"params": [],
				"id": 3
			}`,
			expected: "warehouse id is not specified",
		},
		{
			test: "Product codes are not specified",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Release",
				"params": [{
					"warehouse_id": 1
				}],
				"id": 3
			}`,
			expected: "product codes are not specified",
		},
		{
			test: "Warehouse not exists",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Release",
				"params": [
					{
						"warehouse_id": 1,
						"product_codes": ["b29a07c5-472f-4778-bb04-dab16e9502bb"]
					}
				],
				"id": 3
			}`,
			getByIdReturn: getByIdReturn{
				warehouse: domain.Warehouse{},
				err:       domain.ErrWarehouseNotFound,
			},
			expected: "warehouse not found",
		},
		{
			test: "Warehouse is not available",
			request: `{
				"jsonrpc": "2.0",
				"method": "lamoda.Release",
				"params": [
					{
						"warehouse_id": 1,
						"product_codes": ["b29a07c5-472f-4778-bb04-dab16e9502bb"]
					}
				],
				"id": 3
			}`,
			getByIdReturn: getByIdReturn{
				warehouse: notAvailableWarehouse,
				err:       nil,
			},
			expected: "warehouse is not available",
		},
	}

	for _, tc := range testCases {
		mockWarehouseRepository := new(repomock.MockWarehouseRepository)
		mockWarehouseRepository.On("GetById", context.Background(), mock.Anything).
			Return(tc.getByIdReturn.warehouse, tc.getByIdReturn.err)
		mockWarehouseRepository.On("Release", context.Background(), mock.Anything).
			Return(tc.reserveReturn)

		testService := services.LamodaService{
			Warehouse: mockWarehouseRepository,
		}

		handler := createTestHandler(&testService)

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("POST", app.ServicePath, bytes.NewBufferString(tc.request))
		req.Header.Set("Content-Type", app.CodecContentType)
		handler.ServeHTTP(recorder, req)

		resp := rpc.Response{}
		if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			panic(err)
		}

		assert.Equal(t, tc.expected, resp.Error)
	}
}
