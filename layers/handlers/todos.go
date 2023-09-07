package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/utilyre/htmxtodo/layers/storages"
)

type Todo struct {
	ID    uint64 `param:"id" validate:"omitempty,min=1"`
	Title string `form:"title" validate:"omitempty,min=3"`
	Body  string `form:"body" validate:"omitempty,min=10"`
	Done  bool   `validate:"-"`
}

func (t *Todo) Storage() *storages.Todo {
	return &storages.Todo{
		Title: t.Title,
		Body:  t.Body,
	}
}

type todosHandler struct {
	storage storages.TodosStorage
	tmpl    *template.Template
}

func Todos(e *echo.Echo, storage storages.TodosStorage, tmpl *template.Template) {
	g := e.Group("/todos")
	h := todosHandler{storage: storage, tmpl: tmpl}

	g.POST("", h.create)
	g.GET("", h.readAll)
	g.PUT(":id/toggle", h.toggle)
}

func (h todosHandler) create(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	if err := c.Validate(todo); err != nil {
		return err
	}

	sTodo := todo.Storage()
	if err := h.storage.Create(sTodo); err != nil {
		return err
	}

	todo.Done = sTodo.Done

	buf := new(bytes.Buffer)
	if err := h.tmpl.ExecuteTemplate(buf, "todo", todo); err != nil {
		return err
	}

	return c.HTML(http.StatusCreated, buf.String())
}

func (h todosHandler) readAll(c echo.Context) error {
	sTodos := []storages.Todo{}
	if err := h.storage.ReadAll(&sTodos); err != nil {
		return err
	}

	todos := make([]Todo, 0, len(sTodos))
	for _, sTodo := range sTodos {
		todos = append(todos, Todo{
			ID:    sTodo.ID,
			Title: sTodo.Title,
			Body:  sTodo.Body,
			Done:  sTodo.Done,
		})
	}

	buf := new(bytes.Buffer)
	if err := h.tmpl.ExecuteTemplate(buf, "todos", todos); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, buf.String())
}

func (h todosHandler) toggle(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	if err := c.Validate(todo); err != nil {
		return err
	}

	sTodo := todo.Storage()
	if err := h.storage.ToggleDone(sTodo); err != nil {
		return err
	}

	todo.Title = sTodo.Title
	todo.Body = sTodo.Body
	todo.Done = sTodo.Done

	buf := new(bytes.Buffer)
	if err := h.tmpl.ExecuteTemplate(buf, "todo", todo); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, buf.String())
}
