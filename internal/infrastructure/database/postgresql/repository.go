package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"FACEIT-coding-test/internal/domain"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// UserRepositoryPostgres is a concrete implementation of the UserRepository interface for PostgreSQL.
type UserRepositoryPostgres struct {
	db *sql.DB
}

// NewUserRepositoryPostgres creates a new UserRepositoryPostgres with the given database connection string.
func NewUserRepositoryPostgres(connStr string) (*UserRepositoryPostgres, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &UserRepositoryPostgres{db}, nil
}

// Create creates a new user in the database.
func (ur *UserRepositoryPostgres) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (id, name, email, password, country, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := ur.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password, user.Country, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

// Update updates an existing user in the database.
func (ur *UserRepositoryPostgres) Update(ctx context.Context, user *domain.User) error {
	query := `
        UPDATE users
        SET "name"=$1, "email"=$2, "country"=$3, "updated_at"=$4
        WHERE id=$5
    `

	_, err := ur.db.ExecContext(ctx, query, user.Name, user.Email, user.Country, time.Now().UTC(), user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

// Delete deletes a user with the given ID from the database.
func (ur *UserRepositoryPostgres) Delete(ctx context.Context, id string) error {
	query := `
        DELETE FROM users WHERE id=$1
    `

	_, err := ur.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

// FindByID retrieves a user with the given ID from the database.
func (ur *UserRepositoryPostgres) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
        SELECT id, name, email, password, country, created_at, updated_at
        FROM users
        WHERE id=$1
    `

	var user domain.User
	err := ur.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Country, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(domain.ErrUserNotFound, "user not found")
		}
		return nil, errors.Wrap(err, "failed to find user")
	}

	return &user, nil
}

// List retrieves a list of users from the database, with optional filtering and pagination.
func (ur *UserRepositoryPostgres) List(ctx context.Context, query string, offset, limit int) ([]*domain.User, error) {
	// Create the SQL query for filtering.
	if query != "" {
		query = strings.TrimSpace(query)
		query = strings.ToLower(query)
		query = fmt.Sprintf("WHERE %s", query)
	}

	// Create the SQL query for pagination.
	if limit <= 0 {
		return nil, errors.Wrap(domain.ErrInvalidLimit, "invalid limit")
	}

	if offset < 0 {
		return nil, errors.Wrap(domain.ErrInvalidOffset, "invalid offset")
	}

	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)

	// Execute the SQL query.
	query = `
        SELECT id, name, email, password, country, created_at, updated_at
        FROM users
        ` + query

	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}
	defer rows.Close()

	// Parse the query results into a list of User entities.
	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Country, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	return users, nil
}

func (r *UserRepositoryPostgres) Count(ctx context.Context, query string) (int, error) {
	var count int

	sql := `SELECT COUNT(*) FROM users`

	if query != "" {
		sql = sql + " WHERE " + query
	}

	err := r.db.QueryRowContext(ctx, sql).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count users")
	}

	return count, nil
}
