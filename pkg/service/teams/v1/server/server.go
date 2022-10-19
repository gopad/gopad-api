package serverv1

import (
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/metrics"
	"github.com/gopad/gopad-api/pkg/service/teams/repository"
	"github.com/gopad/gopad-api/pkg/upload"
)

// NewTeamsServer initializes the team server.
func NewTeamsServer(
	cfg *config.Config,
	uploads upload.Upload,
	metricz *metrics.Metrics,
	repository repository.TeamsRepository,
) *TeamsServer {
	return &TeamsServer{
		config:     cfg,
		uploads:    uploads,
		metrics:    metricz,
		repository: repository,
	}
}

// TeamsServer provides all handlers for teams API.
type TeamsServer struct {
	config     *config.Config
	uploads    upload.Upload
	metrics    *metrics.Metrics
	repository repository.TeamsRepository
}
