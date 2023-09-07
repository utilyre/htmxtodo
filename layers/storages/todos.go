package storages

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Todo struct {
	ID        uint64    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Title string `db:"title"`
	Body  string `db:"body"`
	Done  bool   `db:"done"`
}

type TodosStorage struct {
	db *sqlx.DB
}

func NewTodosStorage(db *sqlx.DB) TodosStorage {
	return TodosStorage{db: db}
}

func (s TodosStorage) Create(todo *Todo) error {
	query := `
	INSERT INTO "todos"
	("title", "body")
	VALUES ($1, $2)
	RETURNING *;
	`

	return s.db.Get(todo, query, todo.Title, todo.Body)
}

func (s TodosStorage) Read(todo *Todo) error {
	query := `
	SELECT *
	FROM "todos"
	WHERE id = $1;
	`

	return s.db.Get(todo, query, todo.ID)
}

func (s TodosStorage) ReadAll(todos *[]Todo) error {
	query := `
	SELECT *
	FROM "todos"
	ORDER BY "id";
	`

	return s.db.Select(todos, query)
}

func (s TodosStorage) Update(todo *Todo) error {
	query := `
	UPDATE "todos"
	SET "title" = $1, "body" = $2
	WHERE "id" = $3
	RETURNING *;
	`

	return s.db.Get(todo, query, todo.Title, todo.Body, todo.ID)
}

func (s TodosStorage) ToggleDone(todo *Todo) error {
	query := `
	UPDATE "todos"
	SET "done" = NOT "done"
	WHERE "id" = $1
	RETURNING *;
	`

	return s.db.Get(todo, query, todo.ID)
}

func (s TodosStorage) Delete(todo *Todo) error {
	query := `
	DELETE
	FROM "todos"
	WHERE "id" = $1
	RETURNING *;
	`

	return s.db.Get(todo, query, todo.ID)
}
