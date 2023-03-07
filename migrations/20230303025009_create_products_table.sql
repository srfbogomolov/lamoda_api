-- +goose Up
-- +goose StatementBegin
CREATE TABLE "products" (
    "code" uuid DEFAULT gen_random_uuid(),
    "name" text UNIQUE NOT NULL,
    "size" int NOT NULL,
    "qty" int NOT NULL,
    PRIMARY KEY ("code")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "products" CASCADE;
-- +goose StatementEnd
