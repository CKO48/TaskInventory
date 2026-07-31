package repo

import (
	"context"
	"database/sql"
	"taskmanager/domain"
	"taskmanager/repo"

	_ "modernc.org/sqlite" // Driver para SQLite
)

type SQLiteRepo struct {
	db *sql.DB
}

// InitSQLiteDB crea/abre la base de datos, asegura la creación de la tabla e inicializa la estructura.
func InitSQLiteDB() (repo.Database, error) {
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		title TEXT PRIMARY KEY,
		description TEXT,
		status BOOLEAN NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteRepo{db: db}, nil
}

// SaveTask inserta una nueva tarea o actualiza su descripción/estado si ya existe (UPSERT).
func (r *SQLiteRepo) SaveTask(name string, description string, status bool) (sql.Result, error) {
	query := `
	INSERT INTO tasks (title, description, status) 
	VALUES (?, ?, ?)
	ON CONFLICT(title) DO UPDATE SET
		description = excluded.description,
		status = excluded.status;`

	return r.db.Exec(query, name, description, status)
}

// DeleteTask elimina una tarea por su título.
func (r *SQLiteRepo) DeleteTask(name string) (sql.Result, error) {
	query := `DELETE FROM tasks WHERE title = ?;`
	return r.db.Exec(query, name)
}

// Load recupera todas las tareas registradas de la base de datos usando el contexto.
func (r *SQLiteRepo) Load(ctx context.Context) (map[string]domain.Task, error) {
	query := `SELECT title, description, status FROM tasks;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]domain.Task)

	for rows.Next() {
		var title, description string
		var status bool

		if err := rows.Scan(&title, &description, &status); err != nil {
			return nil, err
		}

		// Se asume la función constructora domain.NewStandartTask(title, description, status)
		task := domain.NewStandartTask(title, description, status)
		tasks[title] = task
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Close cierra la conexión activa con SQLite.
func (r *SQLiteRepo) Close() error {
	return r.db.Close()
}
