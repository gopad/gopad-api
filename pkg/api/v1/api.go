package v1

import (
	"context"
	"net/http"
	"path"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/gopad/gopad-api/pkg/api/v1/models"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations"
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/gopad/gopad-api/pkg/token"
	"github.com/gopad/gopad-api/pkg/upload"
	"github.com/rs/zerolog/log"
)

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	Handler http.Handler
}

// New creates a new API that adds the custom Handler implementations.
func New(
	cfg *config.Config,
	uploads upload.Upload,
	usersService users.Service,
	teamsService teams.Service,
) *API {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to analyze openapi")

		return nil
	}

	spec.Spec().Host = cfg.Server.Host
	spec.Spec().BasePath = path.Join(
		cfg.Server.Root,
		spec.Spec().BasePath,
	)

	api := operations.NewGopadAPI(spec)

	api.AuthLoginUserHandler = LoginUserHandler(cfg, usersService)
	api.AuthRefreshAuthHandler = RefreshAuthHandler(cfg)
	api.AuthVerifyAuthHandler = VerifyAuthHandler()

	api.ProfileTokenProfileHandler = TokenProfileHandler(cfg)
	api.ProfileUpdateProfileHandler = UpdateProfileHandler(usersService)
	api.ProfileShowProfileHandler = ShowProfileHandler(usersService)

	api.TeamListTeamsHandler = ListTeamsHandler(teamsService)
	api.TeamShowTeamHandler = ShowTeamHandler(teamsService)
	api.TeamCreateTeamHandler = CreateTeamHandler(teamsService)
	api.TeamUpdateTeamHandler = UpdateTeamHandler(teamsService)
	api.TeamDeleteTeamHandler = DeleteTeamHandler(teamsService)
	api.TeamListTeamUsersHandler = ListTeamUsersHandler(teamsService)
	api.TeamAppendTeamToUserHandler = AppendTeamToUserHandler(teamsService, usersService)
	api.TeamPermitTeamUserHandler = PermitTeamUserHandler(teamsService, usersService)
	api.TeamDeleteTeamFromUserHandler = DeleteTeamFromUserHandler(teamsService, usersService)

	api.UserListUsersHandler = ListUsersHandler(usersService)
	api.UserShowUserHandler = ShowUserHandler(usersService)
	api.UserCreateUserHandler = CreateUserHandler(usersService)
	api.UserUpdateUserHandler = UpdateUserHandler(usersService)
	api.UserDeleteUserHandler = DeleteUserHandler(usersService)
	api.UserListUserTeamsHandler = ListUserTeamsHandler(usersService)
	api.UserAppendUserToTeamHandler = AppendUserToTeamHandler(usersService, teamsService)
	api.UserPermitUserTeamHandler = PermitUserTeamHandler(usersService, teamsService)
	api.UserDeleteUserFromTeamHandler = DeleteUserFromTeamHandler(usersService, teamsService)

	// TODO: needs context for request id logging?
	api.HeaderAuth = func(val string) (*models.User, error) {
		t, err := token.Parse(val, cfg.Session.Secret)

		if err != nil {
			log.Warn().
				Err(err).
				Str("token", val).
				Msg("failed to parse token")

			return nil, errors.New(401, "incorrect auth")
		}

		user, err := usersService.Show(
			context.Background(),
			t.Text,
		)

		if err != nil {
			log.Warn().
				Err(err).
				Str("token", val).
				Msg("failed to fetch user")

			return nil, errors.New(401, "incorrect auth")
		}

		return convertUser(user), nil
	}

	// TODO: needs context for request id logging?
	api.BasicAuth = func(username, password string) (*models.User, error) {
		user, err := usersService.ByBasicAuth(
			context.Background(),
			username,
			password,
		)

		if err != nil {
			log.Warn().
				Err(err).
				Str("username", username).
				Msg("failed to auth user")

			return nil, errors.New(401, "incorrect auth")
		}

		return convertUser(user), nil
	}

	return &API{
		Handler: api.Serve(nil),
	}
}

func convertAuthToken(record *token.Result) *models.AuthToken {
	if record.ExpiresAt.IsZero() {
		return &models.AuthToken{
			Token: record.Token,
		}
	}

	expiresAt := strfmt.DateTime(record.ExpiresAt)

	return &models.AuthToken{
		Token:     record.Token,
		ExpiresAt: &expiresAt,
	}
}

func convertAuthVerify(record *models.User) *models.AuthVerify {
	createdAt := strfmt.DateTime(record.CreatedAt)

	return &models.AuthVerify{
		Username:  *record.Username,
		CreatedAt: &createdAt,
	}
}

// convertProfile is a simple helper to convert between different model formats.
func convertProfile(record *model.User) *models.Profile {
	teams := make([]*models.TeamUser, 0)

	for _, team := range record.Teams {
		teams = append(teams, convertTeamUser(team))
	}

	return &models.Profile{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Username:  &record.Username,
		Password:  nil,
		Email:     &record.Email,
		Active:    &record.Active,
		Admin:     &record.Admin,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Teams:     teams,
	}
}

// convertTeam is a simple helper to convert between different model formats.
func convertTeam(record *model.Team) *models.Team {
	users := make([]*models.TeamUser, 0)

	for _, user := range record.Users {
		users = append(users, convertTeamUser(user))
	}

	return &models.Team{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Name:      &record.Name,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Users:     users,
	}
}

// convertUser is a simple helper to convert between different model formats.
func convertUser(record *model.User) *models.User {
	teams := make([]*models.TeamUser, 0)

	for _, team := range record.Teams {
		teams = append(teams, convertTeamUser(team))
	}

	return &models.User{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Username:  &record.Username,
		Password:  nil,
		Email:     &record.Email,
		Active:    &record.Active,
		Admin:     &record.Admin,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
		Teams:     teams,
	}
}

// convertTeamUser is a simple helper to convert between different model formats.
func convertTeamUser(record *model.TeamUser) *models.TeamUser {
	userID := strfmt.UUID(record.UserID)
	teamID := strfmt.UUID(record.TeamID)

	return &models.TeamUser{
		TeamID:    &teamID,
		Team:      convertTeam(record.Team),
		UserID:    &userID,
		User:      convertUser(record.User),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
