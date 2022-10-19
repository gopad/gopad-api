package repository

import (
	"context"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/opentracing/opentracing-go"
)

// TracingRequestID returns the request ID as string for tracing
type TracingRequestID func(context.Context) string

// TracingRepository implements MembersRepository interface.
type TracingRepository struct {
	upstream  MembersRepository
	requestID TracingRequestID
}

// NewTracingRepository wraps the MembersRepository and provides tracing for its methods.
func NewTracingRepository(repository MembersRepository, requestID TracingRequestID) MembersRepository {
	return &TracingRepository{
		upstream:  repository,
		requestID: requestID,
	}
}

// List implements the MembersRepository interface.
func (r *TracingRepository) List(ctx context.Context, teamID, userID string) ([]*model.Member, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MembersRepository.List")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("team", teamID)
	span.SetTag("user", userID)
	defer span.Finish()

	return r.upstream.List(ctx, teamID, userID)
}

// Append implements the MembersRepository interface.
func (r *TracingRepository) Append(ctx context.Context, teamID, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MembersRepository.Append")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("team", teamID)
	span.SetTag("user", userID)
	defer span.Finish()

	return r.upstream.Append(ctx, teamID, userID)
}

// Drop implements the MembersRepository interface.
func (r *TracingRepository) Drop(ctx context.Context, teamID, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MembersRepository.Drop")
	span.SetTag("request", r.requestID(ctx))
	span.SetTag("team", teamID)
	span.SetTag("user", userID)
	defer span.Finish()

	return r.upstream.Drop(ctx, teamID, userID)
}
