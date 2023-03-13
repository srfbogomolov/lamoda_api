-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouses (
  id bigserial,
  name text NOT NULL,
  is_available boolean NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE warehouses CASCADE;
-- +goose StatementEnd
