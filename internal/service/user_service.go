// Package service provides the business logic for the microservice.
package service

import (
	"FACEIT-coding-test/internal/infrastructure/publisher"
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"FACEIT-coding-test/internal/domain"
)

type UserService struct {
	userRepo  domain.UserRepository
	publisher *publisher.RabbitMQPublisher
}

func NewUserService(userRepo domain.UserRepository, pub *publisher.RabbitMQPublisher) *UserService {
	return &UserService{userRepo, pub}
}

func (us *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := us.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := us.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	if err := us.userRepo.Update(ctx, user); err != nil {
		return err
	}

	// Publish event
	err = us.publisher.Publish("user", "user.updated", publisher.UserUpdatedEvent{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Timestamp: 0,
	})
	if err != nil {
		log.Println("Failed to publish event:", err)
	}

	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, id string) error {
	if _, err := us.userRepo.FindByID(ctx, id); err != nil {
		return err
	}

	if err := us.userRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := us.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) ListUsers(ctx context.Context, query string, offset, limit int) ([]*domain.User, error) {
	users, err := us.userRepo.List(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
