package expense

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(e *Expense) error {
	e.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	query := `
	INSERT OR IGNORE INTO expenses 
	(id, amount, category, description, date, created_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		e.ID,
		e.Amount,
		e.Category,
		e.Description,
		e.Date,
		e.CreatedAt,
	)
	return err
}

func (r *Repository) GetByID(id string) (*Expense, error) {
	row := r.db.QueryRow(`
	SELECT id, amount, category, description, date, created_at
	FROM expenses WHERE id = ?`, id)

	var e Expense
	err := row.Scan(
		&e.ID,
		&e.Amount,
		&e.Category,
		&e.Description,
		&e.Date,
		&e.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *Repository) List(category string) ([]Expense, error) {
	query := `
	SELECT id, amount, category, description, date, created_at
	FROM expenses
	WHERE (? = '' OR category = ?)
	ORDER BY date DESC
	`

	rows, err := r.db.Query(query, category, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		rows.Scan(
			&e.ID,
			&e.Amount,
			&e.Category,
			&e.Description,
			&e.Date,
			&e.CreatedAt,
		)
		expenses = append(expenses, e)
	}
	return expenses, nil
}
