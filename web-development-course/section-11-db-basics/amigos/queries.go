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
