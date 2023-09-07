package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/utilyre/htmxtodo/layers/storages"
)

type Todo struct {
	ID    uint64 `form:"id" validate:"isdefault"`
	Title string `form:"title" validate:"required,min=3"`
	Body  string `form:"body" validate:"omitempty,min=10"`
}

func (t *Todo) Storage() *storages.Todo {
	return &storages.Todo{
		ID:    t.ID,
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

	buf := new(bytes.Buffer)
	if err := h.tmpl.ExecuteTemplate(buf, "todo", todo); err != nil {
		return err
	}

	return c.HTML(http.StatusCreated, buf.String())
}
