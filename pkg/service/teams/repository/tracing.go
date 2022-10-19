package repository

import (
	"context"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/opentracing/opentracing-go"
)

// TracingRequestID returns the request ID as string for tracing
type TracingRequestID func(context.Context) string

// TracingRepository implements TeamsRepository interface.
type TracingRepository struct {
	upstream  TeamsRepository
	requestID TracingRequestID
}

// NewTracingRepository wraps the TeamsRepository and provides tracing for its methods.
func NewTracingRepository(repository TeamsRepository, requestID TracingRequestID) TeamsRepository {
	return &TracingRepository{
		upstream:  repository,
		requestID: requestID,
	}
}

// List implements the TeamsRepository interface.
func (r *TracingRepository) List(ctx context.Context) ([]*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.List")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	return r.upstream.List(ctx)
}

// Create implements the TeamsRepository interface.
func (r *TracingRepository) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.Create")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	record, err := r.upstream.Create(ctx, team)
	span.SetTag("name", r.extractIdentifier(record))

	return record, err
}

// Update implements the TeamsRepository interface.
func (r *TracingRepository) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.Update")
	span.SetTag("request", r.requestID(ctx))
	defer span.Finish()

	record, err := r.upstream.Update(ctx, team)
	span.SetTag("name", r.extractIdentifier(record))

	return record, err
}

// Show implements the TeamsRepository interface.
func (r *TracingRepository) Show(ctx context.Context, name string) (*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.Show")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Show(ctx, name)
}

// Delete implements the TeamsRepository interface.
func (r *TracingRepository) Delete(ctx context.Context, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.Delete")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Delete(ctx, name)
}

// Exists implements the TeamsRepository interface.
func (r *TracingRepository) Exists(ctx context.Context, name string) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TeamsRepository.Exists")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return r.upstream.Exists(ctx, name)
}

func (r *TracingRepository) extractIdentifier(record *model.Team) string {
	if record == nil {
		return ""
	}

	if record.ID != "" {
		return record.ID
	}

	if record.Slug != "" {
		return record.Slug
	}

	if record.Name != "" {
		return record.Name
	}

	return ""
}
