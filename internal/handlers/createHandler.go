package handlers

import (
	"log/slog"
	"net/http"
)

type CreateTodo interface {
	CreateTodo(title string, description string) error
}

func CreateTodoHandler(log *slog.Logger, createHandler CreateTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Error("Request parsing error", "error", err)
			return
		}

		title, description := r.Form.Get("title"), r.Form.Get("desc")

		log.Info("Create todo", "title", title, "description", description)

		if err := createHandler.CreateTodo(title, description); err != nil {
			log.Error("Error creating todo", "error", err)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
