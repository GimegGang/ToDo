package handlers

import (
	"ToDo/internal/storage/sqlite"
	"html/template"
	"log/slog"
	"net/http"
)

type GetTodo interface {
	GetTodo() ([]sqlite.Task, error)
}

func GetTodoHandler(log *slog.Logger, s GetTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := s.GetTodo()
		if err != nil {
			log.Error("GetTodoHandler", "err", err)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		if err = tmpl.ExecuteTemplate(w, "index.html", tasks); err != nil {
			log.Error("GetTodoHandler", "err", err)
			return
		}
	}
}
