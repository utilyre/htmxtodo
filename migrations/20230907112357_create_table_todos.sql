-- +goose Up
-- +goose StatementBegin
CREATE TABLE "todos" (
    "id" serial NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    "title" varchar(100) NOT NULL,
    "body" text,
    PRIMARY KEY ("id")
);

CREATE FUNCTION "update_todos_updated_at"()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER "update_todos_updated_at"
    BEFORE UPDATE
    ON "todos"
    FOR EACH ROW
    EXECUTE PROCEDURE update_todos_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER "update_todos_updated_at" ON "todos";
DROP FUNCTION "update_todos_updated_at";
DROP TABLE "todos";
-- +goose StatementEnd
