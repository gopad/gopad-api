package members

import (
	"context"
	"time"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

type loggingService struct {
	service Service
	logger  zerolog.Logger
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(s Service) Service {
	return &loggingService{
		service: s,
		logger:  log.With().Str("service", "members").Logger(),
	}
}

// List implements the Service interface for logging.
func (s *loggingService) List(ctx context.Context, params model.MemberParams) ([]*model.Member, int64, error) {
	start := time.Now()
	records, counter, err := s.service.List(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("user", params.UserID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch members")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, counter, err
}

// Attach implements the Service interface for logging.
func (s *loggingService) Attach(ctx context.Context, params model.MemberParams) error {
	start := time.Now()
	err := s.service.Attach(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "attach").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("user", params.UserID).
		Str("perm", params.Perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to attach member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Permit implements the Service interface for logging.
func (s *loggingService) Permit(ctx context.Context, params model.MemberParams) error {
	start := time.Now()
	err := s.service.Permit(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permit").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("user", params.UserID).
		Str("perm", params.Perm).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to permit member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

// Drop implements the Service interface for logging.
func (s *loggingService) Drop(ctx context.Context, params model.MemberParams) error {
	start := time.Now()
	err := s.service.Drop(ctx, params)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "drop").
		Dur("duration", time.Since(start)).
		Str("team", params.TeamID).
		Str("user", params.UserID).
		Logger()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to drop member")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) requestID(ctx context.Context) string {
	id, ok := hlog.IDFromCtx(ctx)

	if ok {
		return id.String()
	}

	return ""
}
