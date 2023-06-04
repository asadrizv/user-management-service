package postgresql

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// Client represents a PostgreSQL database client.
type Client struct {
	db *sql.DB
}

// NewClient creates a new PostgreSQL client with the given configuration.
func NewClient(dsn string) (*Client, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}

// Close closes the database connection.
func (c *Client) Close() error {
	return c.db.Close()
}

// Exec executes the given SQL query with the given arguments.
func (c *Client) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// QueryRow executes the given SQL query with the given arguments and returns a single row.
func (c *Client) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}

// Query executes the given SQL query with the given arguments and returns multiple rows.
func (c *Client) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
