package users

import (
	"context"
	"errors"

	"github.com/gopad/gopad-api/pkg/model"
)

var (
	// ErrNotFound is returned when a user was not found.
	ErrNotFound = errors.New("user not found")

	// ErrAlreadyAssigned is returned when a user is already assigned.
	ErrAlreadyAssigned = errors.New("user is already assigned")

	// ErrNotAssigned is returned when a user is not assigned.
	ErrNotAssigned = errors.New("user is not assigned")
)

// Service handles all interactions with users.
type Service interface {
	External(context.Context, string, string, string, string, string) (*model.User, error)
	List(context.Context, model.ListParams) ([]*model.User, int64, error)
	Show(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) error
	Update(context.Context, *model.User) error
	Delete(context.Context, string) error
	Exists(context.Context, string) (bool, error)
}

type service struct {
	users Service
}

// NewService returns a Service that handles all interactions with users.
func NewService(users Service) Service {
	return &service{
		users: users,
	}
}

// External implements the Service interface.
func (s *service) External(ctx context.Context, provider, ref, username, email, fullname string) (*model.User, error) {
	return s.users.External(ctx, provider, ref, username, email, fullname)
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.ListParams) ([]*model.User, int64, error) {
	return s.users.List(ctx, params)
}

// Show implements the Service interface.
func (s *service) Show(ctx context.Context, id string) (*model.User, error) {
	return s.users.Show(ctx, id)
}

// Create implements the Service interface.
func (s *service) Create(ctx context.Context, user *model.User) error {
	return s.users.Create(ctx, user)
}

// Update implements the Service interface.
func (s *service) Update(ctx context.Context, user *model.User) error {
	return s.users.Update(ctx, user)
}

// Delete implements the Service interface.
func (s *service) Delete(ctx context.Context, name string) error {
	return s.users.Delete(ctx, name)
}

// Exists implements the Service interface.
func (s *service) Exists(ctx context.Context, name string) (bool, error) {
	return s.users.Exists(ctx, name)
}
