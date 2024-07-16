package handlers

import (
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
	"strconv"
)

type UpdateTodo interface {
	UpdateTodo(id int, title string, description string) error
}

func EditHandler(log *slog.Logger, updateTodo UpdateTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Error("Request parsing error", "error", err)
			return
		}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("Request parsing error", "error", err)
			return
		}
		title, desc := r.Form.Get("title"), r.Form.Get("desc")

		if err = updateTodo.UpdateTodo(id, title, desc); err != nil {
			log.Error("Update todo error", "error", err)
			return
		}

		log.Info("Update todo success", "id", id, "title", title, "desc", desc)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
