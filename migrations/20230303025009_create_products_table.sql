-- +goose Up
-- +goose StatementBegin
CREATE TABLE "products" (
    "id" bigserial,
    "name" text NOT NULL,
    "size" smallint NOT NULL,
    "code" uuid NOT NULL DEFAULT gen_random_uuid(),
    "qty" bigint NOT NULL,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "products" CASCADE;
-- +goose StatementEnd
