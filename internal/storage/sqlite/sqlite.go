package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	*sql.DB
}

type Task struct {
	Id          int64
	Title       string
	Description string
}

func New(storagePath string) (*Storage, error) {
	const op = "internal/storage/mysql.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS todo (
		    id INTEGER PRIMARY KEY,
		    title TEXT NOT NULL,
		    description TEXT
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err = stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db}, nil
}

func (s *Storage) CreateTodo(title string, description string) error {
	const op = "internal/storage/mysql.CreateTodo"

	if title == "" {
		return fmt.Errorf("%s: %w", op, "title is required")
	}

	stmt, err := s.DB.Prepare("INSERT INTO todo (title, description) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if _, err = stmt.Exec(title, description); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetTodo() ([]Task, error) {
	const op = "internal/storage/mysql.GetTodo"

	var tasks []Task
	rows, err := s.DB.Query("SELECT id, title, description FROM todo")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var task Task
		if err = rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

func (s *Storage) DeleteTodo(id int) error {
	const op = "internal/storage/mysql.DeleteTodo"

	stmt, err := s.DB.Prepare("DELETE FROM todo WHERE id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err = stmt.Exec(id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdateTodo(id int, title string, description string) error {
	const op = "internal/storage/mysql.UpdateTodo"

	if title == "" || description == "" {
		return fmt.Errorf("%s: %w", op, "title or description is required")
	}

	stmt, err := s.DB.Prepare("UPDATE todo SET title=?, description=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if _, err = stmt.Exec(title, description, id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
