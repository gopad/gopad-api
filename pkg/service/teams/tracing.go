package teams

import (
	"context"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/opentracing/opentracing-go"
)

// TracingRequestID returns the request ID as string for tracing
type TracingRequestID func(context.Context) string

type tracingService struct {
	service   Service
	requestID TracingRequestID
}

// NewTracingService wraps the Service and provides tracing for its methods.
func NewTracingService(s Service, requestID TracingRequestID) Service {
	return &tracingService{
		service:   s,
		requestID: requestID,
	}
}

func (s *tracingService) List(ctx context.Context) ([]*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.List")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.List(ctx)
}

func (s *tracingService) Show(ctx context.Context, id string) (*model.Team, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.Show")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	defer span.Finish()

	return s.service.Show(ctx, id)
}

func (s *tracingService) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	name := ""

	if team != nil {
		name = team.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.Create")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Create(ctx, team)
}

func (s *tracingService) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	id := ""
	name := ""

	if team != nil {
		id = team.ID
		name = team.Name
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.Update")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("id", id)
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Update(ctx, team)
}

func (s *tracingService) Delete(ctx context.Context, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.Delete")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.Delete(ctx, name)
}

func (s *tracingService) ListUsers(ctx context.Context, name string) ([]*model.TeamUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.ListUsers")
	span.SetTag("request", s.requestID(ctx))
	span.SetTag("name", name)
	defer span.Finish()

	return s.service.ListUsers(ctx, name)
}

func (s *tracingService) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.AppendUser")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.AppendUser(ctx, teamID, userID, perm)
}

func (s *tracingService) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.PermitUser")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.PermitUser(ctx, teamID, userID, perm)
}

func (s *tracingService) DropUser(ctx context.Context, teamID, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "teams.Service.DropUser")
	span.SetTag("request", s.requestID(ctx))
	defer span.Finish()

	return s.service.DropUser(ctx, teamID, userID)
}
