package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddNumber(ctx context.Context, num int) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO numbers (value) VALUES ($1)", num)
	return err
}

func (r *Repository) GetAllSorted(ctx context.Context) ([]int, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT value FROM numbers ORDER BY value ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []int
	for rows.Next() {
		var num int
		if err := rows.Scan(&num); err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, rows.Err()
}
