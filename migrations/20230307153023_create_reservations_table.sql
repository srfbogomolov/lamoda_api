-- +goose Up
-- +goose StatementBegin
CREATE TABLE reservations (
  id bigserial,
  placement_id bigserial UNIQUE,
  qty int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (placement_id)
    REFERENCES placements(id)
    ON DELETE CASCADE
  CONSTRAINT reservations_qty_greater_zero
    CHECK (qty > 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reservations CASCADE;
-- +goose StatementEnd
