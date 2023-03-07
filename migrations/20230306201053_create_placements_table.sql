-- +goose Up
-- +goose StatementBegin
CREATE TABLE placements (
  id bigserial,
  warehouse_id bigserial,
  product_code uuid,
  qty int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (warehouse_id)
    REFERENCES warehouses(id)
    ON DELETE CASCADE,
  FOREIGN KEY (product_code)
    REFERENCES products(code)
    ON DELETE CASCADE,
  CONSTRAINT uk_placements_warehouse_id_product_code
    UNIQUE (warehouse_id, product_code)
  CONSTRAINT placements_qty_greater_zero
    CHECK (qty > 0)
);


CREATE INDEX ux_placements_warehouse_id ON placements(warehouse_id);
CREATE INDEX ux_placements_product_code ON placements(product_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX ux_placements_product_code CASCADE;
DROP INDEX ux_placements_warehouse_id CASCADE;

DROP TABLE placements CASCADE;
-- +goose StatementEnd
