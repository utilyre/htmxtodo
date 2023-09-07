-- +goose Up
-- +goose StatementBegin
ALTER TABLE "todos"
ADD "done" boolean NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "todos"
DROP COLUMN "done";
-- +goose StatementEnd
