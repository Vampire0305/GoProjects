package auth

import (
	"database/sql"
)

type AuthRepository interface {
	CreateUser(username string, passwordHash string) (int64, error)
	FindByUsername(username string) (*User, error)
}

type PostgresAuthRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) AuthRepository {
	return &PostgresAuthRepository{DB: db}
}

func (r *PostgresAuthRepository) CreateUser(username string, passwordHash string) (int64, error) {
	var id int64
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id;`
	err := r.DB.QueryRow(query, username, passwordHash).Scan(&id)
	return id, err
}

func (r *PostgresAuthRepository) FindByUsername(username string) (*User, error) {
	query := `SELECT id, username, password_hash, created_at FROM users WHERE username = $1`

	user := &User{}
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, nil
}
