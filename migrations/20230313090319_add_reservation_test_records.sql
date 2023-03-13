-- +goose Up
-- +goose StatementBegin
INSERT INTO reservations (placement_id, qty) VALUES (1, 50);
INSERT INTO reservations (placement_id, qty) VALUES (2, 100);
INSERT INTO reservations (placement_id, qty) VALUES (3, 150);
INSERT INTO reservations (placement_id, qty) VALUES (4, 100);
INSERT INTO reservations (placement_id, qty) VALUES (5, 400);
INSERT INTO reservations (placement_id, qty) VALUES (6, 300);
INSERT INTO reservations (placement_id, qty) VALUES (7, 200);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE FROM reservations;
-- +goose StatementEnd
