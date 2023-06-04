package domain

import "context"

// UserRepository is an interface that defines methods for interacting with user data.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*User, error)
	List(ctx context.Context, query string, offset, limit int) ([]*User, error)
	Count(ctx context.Context, query string) (int, error)
}
