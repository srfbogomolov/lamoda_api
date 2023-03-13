-- +goose Up
-- +goose StatementBegin
INSERT INTO warehouses (name, is_available) VALUES ('test_warehouse_1', true);
INSERT INTO warehouses (name, is_available) VALUES ('test_warehouse_2', true);
INSERT INTO warehouses (name, is_available) VALUES ('test_warehouse_3', false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE warehouses;
-- +goose StatementEnd
