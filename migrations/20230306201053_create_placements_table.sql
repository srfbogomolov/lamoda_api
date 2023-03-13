-- +goose Up
-- +goose StatementBegin
CREATE TABLE placements (
  id bigserial,
  warehouse_id bigserial,
  product_id bigserial,
  qty int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (warehouse_id)
    REFERENCES warehouses(id)
    ON DELETE CASCADE,
  FOREIGN KEY (product_id)
    REFERENCES products(id)
    ON DELETE CASCADE,
  CONSTRAINT uk_placements_warehouse_id_product_id
    UNIQUE (warehouse_id, product_id),
  CONSTRAINT placements_qty_gte_zero
    CHECK (qty >= 0)
);

CREATE INDEX ux_placements_warehouse_id ON placements(warehouse_id);
CREATE INDEX ux_placements_product_id ON placements(product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX ux_placements_product_id CASCADE;
DROP INDEX ux_placements_warehouse_id CASCADE;

DROP TABLE placements CASCADE;
-- +goose StatementEnd
