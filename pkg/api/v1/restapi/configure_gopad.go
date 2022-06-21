// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gopad/gopad-api/pkg/api/v1/models"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/profile"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/team"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/user"
)

//go:generate swagger generate server --target ../../v1 --name Gopad --spec ../../../../openapi/v1.yml --principal models.User --default-scheme https --exclude-main

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

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuth == nil {
		api.BasicAuth = func(user string, pass string) (*models.User, error) {
			return nil, errors.NotImplemented("basic auth  (Basic) has not yet been implemented")
		}
	}
	// Applies when the "X-API-Key" header is set
	if api.HeaderAuth == nil {
		api.HeaderAuth = func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (Header) X-API-Key from header param [X-API-Key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.TeamAppendTeamToUserHandler == nil {
		api.TeamAppendTeamToUserHandler = team.AppendTeamToUserHandlerFunc(func(params team.AppendTeamToUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.AppendTeamToUser has not yet been implemented")
		})
	}
	if api.UserAppendUserToTeamHandler == nil {
		api.UserAppendUserToTeamHandler = user.AppendUserToTeamHandlerFunc(func(params user.AppendUserToTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.AppendUserToTeam has not yet been implemented")
		})
	}
	if api.TeamCreateTeamHandler == nil {
		api.TeamCreateTeamHandler = team.CreateTeamHandlerFunc(func(params team.CreateTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.CreateTeam has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamHandler == nil {
		api.TeamDeleteTeamHandler = team.DeleteTeamHandlerFunc(func(params team.DeleteTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeam has not yet been implemented")
		})
	}
	if api.TeamDeleteTeamFromUserHandler == nil {
		api.TeamDeleteTeamFromUserHandler = team.DeleteTeamFromUserHandlerFunc(func(params team.DeleteTeamFromUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.DeleteTeamFromUser has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.UserDeleteUserFromTeamHandler == nil {
		api.UserDeleteUserFromTeamHandler = user.DeleteUserFromTeamHandlerFunc(func(params user.DeleteUserFromTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserFromTeam has not yet been implemented")
		})
	}
	if api.TeamListTeamUsersHandler == nil {
		api.TeamListTeamUsersHandler = team.ListTeamUsersHandlerFunc(func(params team.ListTeamUsersParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeamUsers has not yet been implemented")
		})
	}
	if api.TeamListTeamsHandler == nil {
		api.TeamListTeamsHandler = team.ListTeamsHandlerFunc(func(params team.ListTeamsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.ListTeams has not yet been implemented")
		})
	}
	if api.UserListUserTeamsHandler == nil {
		api.UserListUserTeamsHandler = user.ListUserTeamsHandlerFunc(func(params user.ListUserTeamsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUserTeams has not yet been implemented")
		})
	}
	if api.UserListUsersHandler == nil {
		api.UserListUsersHandler = user.ListUsersHandlerFunc(func(params user.ListUsersParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.ListUsers has not yet been implemented")
		})
	}
	if api.AuthLoginUserHandler == nil {
		api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginUser has not yet been implemented")
		})
	}
	if api.TeamPermitTeamUserHandler == nil {
		api.TeamPermitTeamUserHandler = team.PermitTeamUserHandlerFunc(func(params team.PermitTeamUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.PermitTeamUser has not yet been implemented")
		})
	}
	if api.UserPermitUserTeamHandler == nil {
		api.UserPermitUserTeamHandler = user.PermitUserTeamHandlerFunc(func(params user.PermitUserTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.PermitUserTeam has not yet been implemented")
		})
	}
	if api.AuthRefreshAuthHandler == nil {
		api.AuthRefreshAuthHandler = auth.RefreshAuthHandlerFunc(func(params auth.RefreshAuthParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation auth.RefreshAuth has not yet been implemented")
		})
	}
	if api.ProfileShowProfileHandler == nil {
		api.ProfileShowProfileHandler = profile.ShowProfileHandlerFunc(func(params profile.ShowProfileParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation profile.ShowProfile has not yet been implemented")
		})
	}
	if api.TeamShowTeamHandler == nil {
		api.TeamShowTeamHandler = team.ShowTeamHandlerFunc(func(params team.ShowTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.ShowTeam has not yet been implemented")
		})
	}
	if api.UserShowUserHandler == nil {
		api.UserShowUserHandler = user.ShowUserHandlerFunc(func(params user.ShowUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.ShowUser has not yet been implemented")
		})
	}
	if api.ProfileTokenProfileHandler == nil {
		api.ProfileTokenProfileHandler = profile.TokenProfileHandlerFunc(func(params profile.TokenProfileParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation profile.TokenProfile has not yet been implemented")
		})
	}
	if api.ProfileUpdateProfileHandler == nil {
		api.ProfileUpdateProfileHandler = profile.UpdateProfileHandlerFunc(func(params profile.UpdateProfileParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation profile.UpdateProfile has not yet been implemented")
		})
	}
	if api.TeamUpdateTeamHandler == nil {
		api.TeamUpdateTeamHandler = team.UpdateTeamHandlerFunc(func(params team.UpdateTeamParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation team.UpdateTeam has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}
	if api.AuthVerifyAuthHandler == nil {
		api.AuthVerifyAuthHandler = auth.VerifyAuthHandlerFunc(func(params auth.VerifyAuthParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation auth.VerifyAuth has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
