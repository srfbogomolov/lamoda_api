-- +goose Up
-- +goose StatementBegin
CREATE TABLE "replacements" (
  "id" bigserial,
  "product_id" bigserial,
  "warehouse_id" bigserial,
  "qty" bigint NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE,
  FOREIGN KEY (warehouse_id)
        REFERENCES warehouses(id)
        ON DELETE CASCADE,
  CONSTRAINT uk_replacement_product_id_warehouse_id UNIQUE (product_id, warehouse_id)
);

CREATE INDEX "ux_replacement_product_id" ON replacements(product_id);

CREATE INDEX "ux_replacement_warehouse_id" ON replacements(warehouse_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "ux_replacement_warehouse_id" CASCADE;
DROP INDEX "ux_replacement_product_id" CASCADE;

DROP TABLE "replacements" CASCADE;
-- +goose StatementEnd
