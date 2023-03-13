-- +goose Up
-- +goose StatementBegin
CREATE TABLE reservations (
  id bigserial,
  placement_id bigserial UNIQUE NOT NULL,
  qty int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (placement_id)
    REFERENCES placements(id)
    ON DELETE CASCADE,
  CONSTRAINT reservations_qty_gt_zero
    CHECK (qty > 0)
);

CREATE INDEX ux_reservations_placement_id ON reservations(placement_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX ux_reservations_placement_id CASCADE;

DROP TABLE reservations CASCADE;
-- +goose StatementEnd
