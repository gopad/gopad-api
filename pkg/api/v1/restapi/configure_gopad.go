// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/profile"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/team"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/user"
)

//go:generate gorunpkg github.com/go-swagger/go-swagger/cmd/swagger generate server --target ../../v1 --name Gopad --spec ../../../../assets/apiv1.yml --exclude-main

func configureFlags(api *operations.GopadAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GopadAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AuthAuthLoginHandler == nil {
		api.AuthAuthLoginHandler = auth.AuthLoginHandlerFunc(func(params auth.AuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthLogin has not yet been implemented")
		})
	}
	if api.AuthAuthRefreshHandler == nil {
		api.AuthAuthRefreshHandler = auth.AuthRefreshHandlerFunc(func(params auth.AuthRefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthRefresh has not yet been implemented")
		})
	}
	if api.AuthAuthVerifyHandler == nil {
		api.AuthAuthVerifyHandler = auth.AuthVerifyHandlerFunc(func(params auth.AuthVerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.AuthVerify has not yet been implemented")
		})
	}
	if api.ProfileProfileShowHandler == nil {
		api.ProfileProfileShowHandler = profile.ProfileShowHandlerFunc(func(params profile.ProfileShowParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileShow has not yet been implemented")
		})
	}
	if api.ProfileProfileTokenHandler == nil {
		api.ProfileProfileTokenHandler = profile.ProfileTokenHandlerFunc(func(params profile.ProfileTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileToken has not yet been implemented")
		})
	}
	if api.ProfileProfileUpdateHandler == nil {
		api.ProfileProfileUpdateHandler = profile.ProfileUpdateHandlerFunc(func(params profile.ProfileUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.ProfileUpdate has not yet been implemented")
		})
	}
	if api.TeamTeamCreateHandler == nil {
		api.TeamTeamCreateHandler = team.TeamCreateHandlerFunc(func(params team.TeamCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamCreate has not yet been implemented")
		})
	}
	if api.TeamTeamDeleteHandler == nil {
		api.TeamTeamDeleteHandler = team.TeamDeleteHandlerFunc(func(params team.TeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamDelete has not yet been implemented")
		})
	}
	if api.TeamTeamIndexHandler == nil {
		api.TeamTeamIndexHandler = team.TeamIndexHandlerFunc(func(params team.TeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamIndex has not yet been implemented")
		})
	}
	if api.TeamTeamShowHandler == nil {
		api.TeamTeamShowHandler = team.TeamShowHandlerFunc(func(params team.TeamShowParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamShow has not yet been implemented")
		})
	}
	if api.TeamTeamUpdateHandler == nil {
		api.TeamTeamUpdateHandler = team.TeamUpdateHandlerFunc(func(params team.TeamUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUpdate has not yet been implemented")
		})
	}
	if api.TeamTeamUserAppendHandler == nil {
		api.TeamTeamUserAppendHandler = team.TeamUserAppendHandlerFunc(func(params team.TeamUserAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserAppend has not yet been implemented")
		})
	}
	if api.TeamTeamUserDeleteHandler == nil {
		api.TeamTeamUserDeleteHandler = team.TeamUserDeleteHandlerFunc(func(params team.TeamUserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserDelete has not yet been implemented")
		})
	}
	if api.TeamTeamUserIndexHandler == nil {
		api.TeamTeamUserIndexHandler = team.TeamUserIndexHandlerFunc(func(params team.TeamUserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserIndex has not yet been implemented")
		})
	}
	if api.TeamTeamUserPermHandler == nil {
		api.TeamTeamUserPermHandler = team.TeamUserPermHandlerFunc(func(params team.TeamUserPermParams) middleware.Responder {
			return middleware.NotImplemented("operation team.TeamUserPerm has not yet been implemented")
		})
	}
	if api.UserUserCreateHandler == nil {
		api.UserUserCreateHandler = user.UserCreateHandlerFunc(func(params user.UserCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserCreate has not yet been implemented")
		})
	}
	if api.UserUserDeleteHandler == nil {
		api.UserUserDeleteHandler = user.UserDeleteHandlerFunc(func(params user.UserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserDelete has not yet been implemented")
		})
	}
	if api.UserUserIndexHandler == nil {
		api.UserUserIndexHandler = user.UserIndexHandlerFunc(func(params user.UserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserIndex has not yet been implemented")
		})
	}
	if api.UserUserShowHandler == nil {
		api.UserUserShowHandler = user.UserShowHandlerFunc(func(params user.UserShowParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserShow has not yet been implemented")
		})
	}
	if api.UserUserTeamAppendHandler == nil {
		api.UserUserTeamAppendHandler = user.UserTeamAppendHandlerFunc(func(params user.UserTeamAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamAppend has not yet been implemented")
		})
	}
	if api.UserUserTeamDeleteHandler == nil {
		api.UserUserTeamDeleteHandler = user.UserTeamDeleteHandlerFunc(func(params user.UserTeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamDelete has not yet been implemented")
		})
	}
	if api.UserUserTeamIndexHandler == nil {
		api.UserUserTeamIndexHandler = user.UserTeamIndexHandlerFunc(func(params user.UserTeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamIndex has not yet been implemented")
		})
	}
	if api.UserUserTeamPermHandler == nil {
		api.UserUserTeamPermHandler = user.UserTeamPermHandlerFunc(func(params user.UserTeamPermParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserTeamPerm has not yet been implemented")
		})
	}
	if api.UserUserUpdateHandler == nil {
		api.UserUserUpdateHandler = user.UserUpdateHandlerFunc(func(params user.UserUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserUpdate has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
