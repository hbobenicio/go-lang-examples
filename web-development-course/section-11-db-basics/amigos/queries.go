package amigos

import (
	"context"
	"database/sql"
	"fmt"
)

// List lists all amigos
func List(ctx context.Context, db *sql.DB) ([]Amigo, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM amigos")
	if err != nil {
		return nil, fmt.Errorf("amigos list query: %v", err)
	}

	var result []Amigo
	for rows.Next() {
		var amigo Amigo
		if err := rows.Scan(&amigo.ID, &amigo.Name); err != nil {
			return nil, fmt.Errorf("amigos list query: %v", err)
		}

		result = append(result, amigo)
	}

	return result, nil
}

// Create persists a new Amigo
func Create(ctx context.Context, db *sql.DB, amigo Amigo) (int64, error) {
	row := db.QueryRowContext(ctx, "INSERT INTO amigos (name) VALUES ($1) RETURNING id", amigo.Name)

	var newAmigoID int64
	if err := row.Scan(&newAmigoID); err != nil {
		return 0, fmt.Errorf("amigos create query: %v", err)
	}

	return newAmigoID, nil
}

// Get retrieves an amigo by its ID
func Get(ctx context.Context, db *sql.DB, id int64) (*Amigo, error) {
	row := db.QueryRowContext(ctx, "SELECT id, name FROM amigos WHERE id = $1", id)

	var amigo Amigo
	if err := row.Scan(&amigo.ID, &amigo.Name); err != nil {
		return nil, err
	}

	return &amigo, nil
}

// Delete remover an amigo by its ID
func Delete(ctx context.Context, db *sql.DB, id int64) error {
	result, err := db.ExecContext(ctx, "DELETE FROM amigos WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("amigos delete query: %v", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("amigos delete query: %v", err)
	}

	return nil
}
