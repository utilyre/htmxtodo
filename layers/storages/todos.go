package storages

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Todo struct {
	ID        uint64    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Title string `db:"title"`
	Body  string `db:"body"`
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
	RETURNING "id", "created_at", "updated_at";
	`

	if err := s.db.QueryRow(
		query,
		todo.Title,
		todo.Body,
	).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (s TodosStorage) Read(todo *Todo) error {
	query := `
	SELECT *
	FROM "todos"
	WHERE id = ?;
	`

	if err := s.db.Get(todo, query, todo.ID); err != nil {
		return err
	}

	return nil
}

func (s TodosStorage) ReadAll() ([]Todo, error) {
	query := `
	SELECT *
	FROM "todos";
	`

	todos := []Todo{}
	if err := s.db.Select(&todos, query); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s TodosStorage) Update(todo *Todo) error {
	query := `
	UPDATE "todos"
	SET "title" = $1, "body" = $2
	WHERE "id" = $3
	RETURNING "updated_at";
	`

	if err := s.db.QueryRow(
		query,
		todo.Title,
		todo.Body,
		todo.ID,
	).Scan(&todo.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (s TodosStorage) Delete(todo *Todo) error {
	query := `
	DELETE
	FROM "todos"
	WHERE "id" = $1
	RETURNING *;
	`

	if err := s.db.Get(todo, query, todo.ID); err != nil {
		return err
	}

	return nil
}
