package mock

import (
	"context"

	"FACEIT-coding-test/internal/domain"
)

// UserRepositoryMock is a mock implementation of UserRepository.
type UserRepositoryMock struct {
	users []*domain.User
}

// NewUserRepositoryMock creates a new instance of UserRepositoryMock.
func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		users: []*domain.User{},
	}
}

// Create adds a new user to the mock repository.
func (ur *UserRepositoryMock) Create(ctx context.Context, user *domain.User) error {
	ur.users = append(ur.users, user)
	return nil
}

// Update updates an existing user in the mock repository.
func (ur *UserRepositoryMock) Update(ctx context.Context, user *domain.User) error {
	for i, u := range ur.users {
		if u.ID == user.ID {
			ur.users[i] = user
			return nil
		}
	}
	return domain.ErrUserNotFound
}

// Delete removes a user from the mock repository.
func (ur *UserRepositoryMock) Delete(ctx context.Context, id string) error {
	for i, user := range ur.users {
		if user.ID == id {
			ur.users = append(ur.users[:i], ur.users[i+1:]...)
			return nil
		}
	}
	return domain.ErrUserNotFound
}

// FindByID finds a user in the mock repository by their ID.
func (ur *UserRepositoryMock) FindByID(ctx context.Context, id string) (*domain.User, error) {
	for _, user := range ur.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

// List retrieves a list of users from the mock repository, with optional filtering and pagination.
func (ur *UserRepositoryMock) List(ctx context.Context, query string, offset, limit int) ([]*domain.User, error) {
	// In this mock implementation, we don't support filtering or pagination.
	return ur.users, nil
}

// Count returns the number of users in the mock repository.
func (ur *UserRepositoryMock) Count(ctx context.Context, query string) (int, error) {
	return len(ur.users), nil
}
