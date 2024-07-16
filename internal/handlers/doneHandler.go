package handlers

import (
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
	"strconv"
)

type DeleteHandler interface {
	DeleteTodo(id int) error
}

func DeleteTodoHandler(log *slog.Logger, deleteHandler DeleteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("Invalid id parameter")
			return
		}

		if err = deleteHandler.DeleteTodo(id); err != nil {
			log.Error("Error deleting todo")
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
