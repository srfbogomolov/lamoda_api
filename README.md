# lamoda_api

## Инструкцию по запуску сервиса
Для запуска приложения необходимо выполний команду ```make service/up```, а затем применить миграции с помощью команды ```make migrations/up```.

## Инструкцию по запуску тестов
Запус тестов выполняется с помощью команды ```make tests/run```.

## Описание API методов с запросом и ответом

### host: `http://0.0.0.0:8080/lamoda`

* [Резервирование товара на складе для доставки](#резервирование-товара-на-складе-для-доставки)
* [Освобождение резерва товаров](#освобождение-резерва-товаров)
* [Получение количества оставшихся товаров на складе](#получение-количества-оставшихся-товаров-на-складе)

### Резервирование товара на складе для доставки
---------------------

На вход принимает идентификатор склада и массив уникальных кодов товаров (коды могут повторяться).

```go
type ReleaseReserveArgs struct {
	WarehouseId  uint        `json:"warehouse_id"`
	ProductCodes []uuid.UUID `json:"product_codes"`
}
```

Возвращает статус выполнения операции.

```go
type ReleaseReserveReply struct {
	Status string `json:"status"`
}
```

Запрос:

```bash
curl -X POST \
http://localhost:8080/lamoda \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
    "jsonrpc": "2.0",
    "method": "lamoda.Reserve",
    "params": [
        {
            "warehouse_id": 1,
            "product_codes": [
                "90f6481d-97b3-4f07-a412-ec5d6b13aaa0",
                "b29a07c5-472f-4778-bb04-dab16e9502bb",
                "9e02ae0e-eac1-4f70-a743-cb8442351bbf"
            ]
        }
    ],
    "id": 1
}'
```

Ответ:

```json
{
    "result": {
        "status": "OK"
    },
    "error": null,
    "id": 1
}
```

### Освобождение резерва товаров
----------------------

На вход принимает идентификатор склада и массив уникальных кодов товаров (коды могут повторяться).

```go
type ReleaseReserveArgs struct {
	WarehouseId  uint        `json:"warehouse_id"`
	ProductCodes []uuid.UUID `json:"product_codes"`
}
```

Возвращает статус выполнения операции.

```go
type ReleaseReserveReply struct {
	Status string `json:"status"`
}
```

Запрос:

```bash
curl -X POST \
http://localhost:8080/lamoda \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
    "jsonrpc": "2.0",
    "method": "lamoda.Release",
    "params": [
        {
            "warehouse_id": 1,
            "product_codes": [
                "90f6481d-97b3-4f07-a412-ec5d6b13aaa0",
                "b29a07c5-472f-4778-bb04-dab16e9502bb",
                "9e02ae0e-eac1-4f70-a743-cb8442351bbf"
            ]
        }
    ],
    "id": 2
}'
```

Ответ:

```json
{
    "result": {
        "status": "OK"
    },
    "error": null,
    "id": 2
}
```

### Получение количества оставшихся товаров на складе
-----------------------------------

На вход принимает идентификатор склада.

```go
type GetInStockArgs struct {
	WarehouseId uint `json:"warehouse_id"`
}
```

Возвращает массив оставшихся на складе товаров и статус выполнения операции.

```go
type GetInStockReply struct {
	AvailableProducts []AvailableProduct `json:"available_products"`
	Status            string             `json:"status"`
}
```

Запрос:

```bash
curl -X POST \
http://localhost:8080/lamoda \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
    "jsonrpc": "2.0",
    "method": "lamoda.GetInStock",
    "params": [
        {
            "warehouse_id": 1
        }
    ],
    "id": 3
}'
```

Ответ:

```json
{
    "result": {
        "available_products": [
            {
                "id": 1,
                "code": "90f6481d-97b3-4f07-a412-ec5d6b13aaa0",
                "name": "test_product_1",
                "size": 10,
                "available_qty": 100
            },
            {
                "id": 2,
                "code": "b29a07c5-472f-4778-bb04-dab16e9502bb",
                "name": "test_product_2",
                "size": 20,
                "available_qty": 400
            },
            {
                "id": 4,
                "code": "9e02ae0e-eac1-4f70-a743-cb8442351bbf",
                "name": "test_product_4",
                "size": 40,
                "available_qty": 700
            }
        ],
        "status": "OK"
    },
    "error": null,
    "id": 3
}
```
