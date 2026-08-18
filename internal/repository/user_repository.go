package repository

import (
	"database/sql"
	"fmt"
	"go-crud/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create user
func (r *UserRepository) Create(user *models.User) error {
	query := `
        INSERT INTO users (first_name, last_name, email, phone, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	return err
}

// Get all users
func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, created_at, updated_at FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Get user by ID
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, created_at, updated_at 
              FROM users WHERE id = $1`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// Update user
func (r *UserRepository) Update(id int, user *models.User) error {
	query := `
        UPDATE users 
        SET first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = $5
        WHERE id = $6
    `

	result, err := r.db.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.UpdatedAt,
		id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

// Delete user
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
