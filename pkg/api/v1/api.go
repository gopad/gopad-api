package v1

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations"
	"github.com/rs/zerolog/log"
)

//go:generate gorunpkg github.com/go-swagger/go-swagger/cmd/swagger generate server --target . --name Gopad --spec ../../../openapi/v1.yml --exclude-main --regenerate-configureapi

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	Handler http.Handler
}

// New creates a new API that adds the custom Handler implementations.
func New() *API {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to analyze openapi")

		return nil
	}

	api := operations.NewGopadAPI(spec)

	api.Middleware = func(b middleware.Builder) http.Handler {
		return middleware.Spec("", nil, api.Context().RoutesHandler(b))
	}

	return &API{
		Handler: api.Serve(nil),
	}
}
