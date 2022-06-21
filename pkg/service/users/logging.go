package users

import (
	"context"
	"time"

	"github.com/gopad/gopad-api/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggingRequestID returns the request ID as string for logging
type LoggingRequestID func(context.Context) string

type loggingService struct {
	service   Service
	requestID LoggingRequestID
	logger    zerolog.Logger
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(s Service, requestID LoggingRequestID) Service {
	return &loggingService{
		service:   s,
		requestID: requestID,
		logger:    log.With().Str("service", "users").Logger(),
	}
}

func (s *loggingService) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	return s.service.ByBasicAuth(ctx, username, password)
}

func (s *loggingService) List(ctx context.Context) ([]*model.User, error) {
	start := time.Now()
	records, err := s.service.List(ctx)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "list").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all users")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) Show(ctx context.Context, name string) (*model.User, error) {
	start := time.Now()
	record, err := s.service.Show(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "show").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to find user by name")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	start := time.Now()
	record, err := s.service.Create(ctx, user)

	name := ""

	if record != nil {
		name = record.Username
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "create").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to create user")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	start := time.Now()
	record, err := s.service.Update(ctx, user)

	id := ""
	name := ""

	if record != nil {
		id = record.ID
		name = record.Username
	}

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "update").
		Dur("duration", time.Since(start)).
		Str("id", id).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to update user")
	} else {
		logger.Debug().
			Msg("")
	}

	return record, err
}

func (s *loggingService) Delete(ctx context.Context, name string) error {
	start := time.Now()
	err := s.service.Delete(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "delete").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil && err != ErrNotFound {
		logger.Warn().
			Err(err).
			Msg("failed to delete user")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) ListTeams(ctx context.Context, name string) ([]*model.TeamUser, error) {
	start := time.Now()
	records, err := s.service.ListTeams(ctx, name)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "listTeams").
		Dur("duration", time.Since(start)).
		Str("name", name).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to find all user teams")
	} else {
		logger.Debug().
			Msg("")
	}

	return records, err
}

func (s *loggingService) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	start := time.Now()
	err := s.service.AppendTeam(ctx, userID, teamID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "appendTeam").
		Dur("duration", time.Since(start)).
		Str("user", userID).
		Str("team", teamID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to append user to team")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	start := time.Now()
	err := s.service.PermitTeam(ctx, userID, teamID, perm)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "permitTeam").
		Dur("duration", time.Since(start)).
		Str("user", userID).
		Str("team", teamID).
		Str("perm", perm).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to update team perms")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}

func (s *loggingService) DropTeam(ctx context.Context, userID, teamID string) error {
	start := time.Now()
	err := s.service.DropTeam(ctx, userID, teamID)

	logger := s.logger.With().
		Str("request", s.requestID(ctx)).
		Str("method", "dropTeam").
		Dur("duration", time.Since(start)).
		Str("user", userID).
		Str("team", teamID).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("failed to drop user from team")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
