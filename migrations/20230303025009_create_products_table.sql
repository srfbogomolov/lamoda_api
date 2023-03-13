-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  id bigserial,
  code uuid DEFAULT gen_random_uuid() NOT NULL,
  name text NOT NULL,
  size int NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT products_size_gte_zero
    CHECK (size >= 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products CASCADE;
-- +goose StatementEnd
