package main

import (
	"ToDo/internal/config"
	"ToDo/internal/handlers"
	"ToDo/internal/logger"
	"ToDo/internal/storage/sqlite"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")
	log := logger.New(cfg.Env)

	log.Info("Config and Logger being loaded")

	db, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to initialize database", "error", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)

	router.Get("/", handlers.GetTodoHandler(log, db))
	router.Post("/", handlers.CreateTodoHandler(log, db))
	router.Get("/done/{id}", handlers.DeleteTodoHandler(log, db))
	router.Post("/edit/{id}", handlers.EditHandler(log, db))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	fs := http.FileServer(http.Dir("static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("server error", "error", err)
	}
}
