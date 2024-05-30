package v1

import (
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/metrics"
	"github.com/gopad/gopad-api/pkg/service/members"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/gopad/gopad-api/pkg/session"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/gopad/gopad-api/pkg/upload"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yml ../../../openapi/v1.yml

var (
	_ StrictServerInterface = (*API)(nil)
)

// New creates a new API that adds the handler implementations.
func New(
	cfg *config.Config,
	registry *metrics.Metrics,
	sess *session.Session,
	uploads upload.Upload,
	storage store.Store,
	teamsService teams.Service,
	usersService users.Service,
	membersService members.Service,
) *API {
	return &API{
		config:   cfg,
		registry: registry,
		session:  sess,
		uploads:  uploads,
		storage:  storage,
		teams:    teamsService,
		users:    usersService,
		members:  membersService,
	}
}

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	config   *config.Config
	registry *metrics.Metrics
	session  *session.Session
	uploads  upload.Upload
	storage  store.Store
	teams    teams.Service
	users    users.Service
	members  members.Service
}
