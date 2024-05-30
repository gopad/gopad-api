package members

import (
	"context"
	"errors"

	"github.com/gopad/gopad-api/pkg/model"
)

var (
	// ErrInvalidListParams defines the error if list receives invalid params.
	ErrInvalidListParams = errors.New("invalid parameters for list")

	// ErrNotFound defines the error if a member could not be found.
	ErrNotFound = errors.New("team or user not found")

	// ErrAlreadyAssigned defines the error if a member is already assigned.
	ErrAlreadyAssigned = errors.New("membership already exists")

	// ErrNotAssigned defines the error if a member is not assigned.
	ErrNotAssigned = errors.New("membership is not defined")
)

// Service handles all interactions with members.
type Service interface {
	List(context.Context, model.MemberParams) ([]*model.Member, int64, error)
	Attach(context.Context, model.MemberParams) error
	Permit(context.Context, model.MemberParams) error
	Drop(context.Context, model.MemberParams) error
}

type service struct {
	members Service
}

// NewService returns a Service that handles all interactions with members.
func NewService(members Service) Service {
	return &service{
		members: members,
	}
}

// List implements the Service interface.
func (s *service) List(ctx context.Context, params model.MemberParams) ([]*model.Member, int64, error) {
	return s.members.List(ctx, params)
}

// Attach implements the Service interface.
func (s *service) Attach(ctx context.Context, params model.MemberParams) error {
	return s.members.Attach(ctx, params)
}

// Permit implements the Service interface.
func (s *service) Permit(ctx context.Context, params model.MemberParams) error {
	return s.members.Permit(ctx, params)
}

// Drop implements the Service interface.
func (s *service) Drop(ctx context.Context, params model.MemberParams) error {
	return s.members.Drop(ctx, params)
}
