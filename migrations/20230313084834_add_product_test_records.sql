-- +goose Up
-- +goose StatementBegin
INSERT INTO products (name, size) VALUES ('test_product_1', 10);
INSERT INTO products (name, size) VALUES ('test_product_2', 20);
INSERT INTO products (name, size) VALUES ('test_product_3', 30);
INSERT INTO products (name, size) VALUES ('test_product_4', 40);
INSERT INTO products (name, size) VALUES ('test_product_5', 50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE FROM products;
-- +goose StatementEnd
