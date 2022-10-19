package repository

import (
	"context"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/opentracing/opentracing-go"
)

// TracingRequestID returns the request ID as string for tracing
type TracingRequestID func(context.Context) string

// TracingRepository implements UsersRepository interface.
type TracingRepository struct {
	upstream  UsersRepository
	requestID TracingRequestID
}

// NewTracingRepository wraps the UsersRepository and provides tracing for its methods.
func NewTracingRepository(repository UsersRepository, requestID TracingRequestID) UsersRepository {
	return &TracingRepository{
		upstream:  repository,
		requestID: requestID,
	}
}

// List implements the UsersRepository interface.
func (r *TracingRepository) List(ctx context.Context) ([]*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.List")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	return r.upstream.List(ctx)
}

// Create implements the UsersRepository interface.
func (r *TracingRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.Create")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	record, err := r.upstream.Create(ctx, user)
	span.SetTag("name", r.extractIdentifier(record))

	return record, err
}

// Update implements the UsersRepository interface.
func (r *TracingRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.Update")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	record, err := r.upstream.Update(ctx, user)
	span.SetTag("name", r.extractIdentifier(record))

	return record, err
}

// Show implements the UsersRepository interface.
func (r *TracingRepository) Show(ctx context.Context, name string) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.Show")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Show(ctx, name)
}

// Delete implements the UsersRepository interface.
func (r *TracingRepository) Delete(ctx context.Context, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.Delete")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Delete(ctx, name)
}

// Exists implements the UsersRepository interface.
func (r *TracingRepository) Exists(ctx context.Context, name string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRepository.Exists")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Exists(ctx, name)
}

func (r *TracingRepository) extractIdentifier(record *model.User) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	if record.Slug != "" {
		return record.Slug
	}

	if record.Username != "" {
		return record.Username
	}

	return ""
}
