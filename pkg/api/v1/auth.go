package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gopad/gopad-api/pkg/api/v1/models"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/gopad/gopad-api/pkg/token"
)

// LoginUserHandler implements the handler for the AuthLoginUser operation.
func LoginUserHandler(cfg *config.Config, usersService users.Service) auth.LoginUserHandlerFunc {
	return func(params auth.LoginUserParams) middleware.Responder {
		user, err := usersService.ByBasicAuth(
			params.HTTPRequest.Context(),
			*params.AuthLogin.Username,
			params.AuthLogin.Password.String(),
		)

		if err != nil {
			if err == users.ErrNotFound {
				message := "wrong username or password"

				return auth.NewLoginUserUnauthorized().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if err == users.ErrWrongAuth {
				message := "wrong username or password"

				return auth.NewLoginUserUnauthorized().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return auth.NewLoginUserDefault(http.StatusInternalServerError)
		}

		result, err := token.New(user.Username).Expiring(cfg.Session.Secret, cfg.Session.Expire)

		if err != nil {
			return auth.NewLoginUserDefault(http.StatusInternalServerError)
		}

		return auth.NewLoginUserOK().WithPayload(convertAuthToken(result))
	}
}

// RefreshAuthHandler implements the handler for the AuthRefreshAuth operation.
func RefreshAuthHandler(cfg *config.Config) auth.RefreshAuthHandlerFunc {
	return func(params auth.RefreshAuthParams, principal *models.User) middleware.Responder {
		result, err := token.New(*principal.Username).Expiring(cfg.Session.Secret, cfg.Session.Expire)

		if err != nil {
			return auth.NewRefreshAuthDefault(http.StatusInternalServerError)
		}

		return auth.NewRefreshAuthOK().WithPayload(convertAuthToken(result))
	}
}

// VerifyAuthHandler implements the handler for the AuthVerifyAuth operation.
func VerifyAuthHandler() auth.VerifyAuthHandlerFunc {
	return func(params auth.VerifyAuthParams, principal *models.User) middleware.Responder {
		return auth.NewVerifyAuthOK().WithPayload(convertAuthVerify(principal))
	}
}
