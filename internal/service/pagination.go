// Package service provides the business logic for the microservice.
package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"net/url"
	"strconv"
	"strings"

	"FACEIT-coding-test/internal/domain"
)

type PaginationOptions struct {
	Query  string
	Offset int
	Limit  int
}

func DefaultPaginationOptions() PaginationOptions {
	return PaginationOptions{
		Query:  "",
		Offset: 0,
		Limit:  10,
	}
}

func (us *UserService) PaginateUsers(options PaginationOptions) ([]*domain.User, error) {
	var query string
	var offset int
	var limit int

	if options.Query != "" {
		query = strings.TrimSpace(options.Query)
		query = strings.ToLower(query)
	}

	if options.Offset < 0 {
		return nil, domain.ErrInvalidOffset
	}

	if options.Limit <= 0 {
		return nil, domain.ErrInvalidLimit
	}

	offset = options.Offset
	limit = options.Limit

	users, err := us.userRepo.List(context.Background(), query, offset, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) PaginateUsersWithCount(options PaginationOptions) ([]*domain.User, int, error) {
	var query string
	var offset int
	var limit int

	if options.Query != "" {
		query = strings.TrimSpace(options.Query)
		query = strings.ToLower(query)
		query = fmt.Sprintf("WHERE %s", query)
	}

	if options.Offset < 0 {
		return nil, 0, domain.ErrInvalidOffset
	}

	if options.Limit <= 0 {
		return nil, 0, domain.ErrInvalidLimit
	}

	offset = options.Offset
	limit = options.Limit

	users, err := us.userRepo.List(context.Background(), query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	count, err := us.userRepo.Count(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (us *UserService) PaginateUsersWithPage(options PaginationOptions) ([]*domain.User, int, int, error) {
	var query string
	var page int
	var limit int

	if options.Query != "" {
		query = strings.TrimSpace(options.Query)
		query = strings.ToLower(query)
		query = fmt.Sprintf("WHERE %s", query)
	}

	if options.Offset < 0 {
		return nil, 0, 0, domain.ErrInvalidOffset
	}

	if options.Limit <= 0 {
		return nil, 0, 0, domain.ErrInvalidLimit
	}

	page = options.Offset/options.Limit + 1
	limit = options.Limit

	users, err := us.userRepo.List(context.Background(), query, (page-1)*limit, limit)
	if err != nil {
		return nil, 0, 0, err
	}

	count, err := us.userRepo.Count(context.Background(), query)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	return users, page, totalPages, nil
}

func (us *UserService) ParsePaginationParams(params url.Values) (PaginationOptions, error) {
	var pagination PaginationOptions

	if query := params.Get("query"); query != "" {
		pagination.Query = query
	}

	if offsetStr := params.Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return pagination, errors.New("invalid offset parameter")
		}
		pagination.Offset = offset
	}

	if limitStr := params.Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return pagination, errors.New("invalid limit parameter")
		}
		pagination.Limit = limit
	}

	return pagination, nil
}
