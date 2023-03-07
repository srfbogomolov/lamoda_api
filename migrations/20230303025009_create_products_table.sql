-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  code uuid DEFAULT gen_random_uuid(),
  name text UNIQUE NOT NULL,
  size int NOT NULL,
  qty int NOT NULL,
  PRIMARY KEY (code),
  CONSTRAINT products_size_greater_or_equal_zero
    CHECK (size >= 0),
  CONSTRAINT products_qty_greater_zero
    CHECK (qty > 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products CASCADE;
-- +goose StatementEnd
