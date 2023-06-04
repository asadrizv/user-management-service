package unit_test

import (
	"FACEIT-coding-test/internal/infrastructure/database/mock"
	"context"
	"os"
	"testing"

	"FACEIT-coding-test/internal/domain"
	"github.com/stretchr/testify/assert"
)

var userRepo *mock.UserRepositoryMock

func TestMain(m *testing.M) {
	// Set up user repository
	userRepo = mock.NewUserRepositoryMock()

	// Run tests
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestUserRepository(t *testing.T) {
	// Create a test user
	user := &domain.User{
		ID:      "123",
		Name:    "John",
		Email:   "johndoe@example.com",
		Country: "US",
	}

	ctx := context.Background()
	// Save the user to the database
	err := userRepo.Create(context.Background(), user)
	assert.NoError(t, err)

	// Find the user by ID
	foundUser, err := userRepo.FindByID(ctx, "123")
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	// Update the user
	user.Name = "Jane"
	err = userRepo.Update(ctx, user)
	assert.NoError(t, err)

	// Find the user by ID again
	foundUser, err = userRepo.FindByID(ctx, "123")
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	// Delete the user
	err = userRepo.Delete(ctx, "123")
	assert.NoError(t, err)

	// Try to find the user by ID again (should fail)
	_, err = userRepo.FindByID(ctx, "123")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
}

func TestUserRepositoryList(t *testing.T) {
	// Create some test users
	users := []*domain.User{
		{ID: "1", Name: "John", Email: "johndoe@example.com", Password: "password", Country: "US"},
		{ID: "2", Name: "Jane", Email: "janedoe@example.com", Password: "password", Country: "UK"},
		{ID: "3", Name: "Bob", Email: "bobsmith@example.com", Password: "password", Country: "US"},
	}

	// Save the users to the database
	for _, user := range users {
		err := userRepo.Create(context.Background(), user)
		assert.NoError(t, err)
	}

	// Test listing all users
	listAllUsers(t, "", 0, 0, users)

	// Test filtering by country
	listAllUsers(t, "", 0, 2, users)

}

func TestFindByID(t *testing.T) {
	// Create some test users
	users := []*domain.User{
		{ID: "1", Name: "John", Email: "johndoe@example.com", Password: "password", Country: "US"},
		{ID: "2", Name: "Jane", Email: "janedoe@example.com", Password: "password", Country: "UK"},
		{ID: "3", Name: "Bob", Email: "bobsmith@example.com", Password: "password", Country: "US"},
	}

	// Save the users to the database
	for _, user := range users {
		err := userRepo.Create(context.Background(), user)
		assert.NoError(t, err)
	}

	usr, err := userRepo.FindByID(context.Background(), "1")
	if err != nil {
		return
	}

	assert.Equal(t, "1", usr.ID, "Id of fetched user must be equal to the one passed in the function")

}

func listAllUsers(t *testing.T, query string, offset, limit int, expectedUsers []*domain.User) {
	// List all users
	users, err := userRepo.List(context.Background(), query, offset, limit)
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedUsers, users)
}
