-- +goose Up
-- +goose StatementBegin
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (1, 1, 100);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (2, 1, 200);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (3, 1, 300);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (1, 2, 400);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (2, 2, 500);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (3, 3, 600);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (1, 4, 700);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (2, 5, 800);
INSERT INTO placements (warehouse_id, product_id, qty) VALUES (3, 5, 900);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE FROM placements;
-- +goose StatementEnd
