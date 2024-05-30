package v1

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/markbates/goth"
	"github.com/rs/zerolog/log"
)

// ExternalInitialize implements the v1.ServerInterface.
func (a *API) ExternalInitialize(ctx context.Context, request ExternalInitializeRequestObject) (ExternalInitializeResponseObject, error) {
	provider, err := goth.GetProvider(
		request.Provider,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect provider")

		return ExternalInitialize404JSONResponse{
			Message: ToPtr("Failed to detect provider"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	session, err := provider.BeginAuth(
		setAuthState(request.Params.State),
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to init provider")

		return ExternalInitialize412JSONResponse{
			Message: ToPtr("Failed to init provider"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	url, err := session.GetAuthURL()

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect auth url")

		return ExternalInitialize404JSONResponse{
			Message: ToPtr("Failed to detect auth url"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	a.session.Put(
		ctx,
		request.Provider,
		session.Marshal(),
	)

	return ExternalInitializeRedirectResponse{
		url: url,
	}, nil
}

// ExternalInitializeRedirectResponse defines the response to redirect to a defined URL.
type ExternalInitializeRedirectResponse struct {
	url string
}

// VisitExternalInitializeResponse implements the middleware.Responder interface for redirects.
func (r ExternalInitializeRedirectResponse) VisitExternalInitializeResponse(w http.ResponseWriter) error {
	w.Header().Set(
		"Location",
		r.url,
	)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.WriteHeader(
		http.StatusTemporaryRedirect,
	)

	return nil
}

// ExternalCallback implements the v1.ServerInterface.
func (a *API) ExternalCallback(ctx context.Context, request ExternalCallbackRequestObject) (ExternalCallbackResponseObject, error) {
	provider, err := goth.GetProvider(
		request.Provider,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to detect provider")

		return ExternalCallback404JSONResponse{
			Message: ToPtr("Failed to detect provider"),
			Status:  ToPtr(http.StatusNotFound),
		}, nil
	}

	session, err := provider.UnmarshalSession(
		a.session.Get(
			ctx,
			request.Provider,
		),
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to parse session")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to parse session"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	// defer gothic.Logout(w, r)

	if err := verifyAuthState(
		request.Params.State,
		session,
	); err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to verify state")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to verify state"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	external, err := provider.FetchUser(
		session,
	)

	log.Debug().
		Str("provider", external.Provider).
		Str("email", external.Email).
		Str("name", external.Name).
		Str("firstname", external.FirstName).
		Str("lastname", external.LastName).
		Str("nickname", external.NickName).
		Str("user_id", external.UserID).
		Msg("requested auth")

	if err == nil {
		nickname := slugify.Slugify(external.NickName)

		user, err := a.users.External(
			ctx,
			external.Provider,
			external.UserID,
			nickname,
			external.Email,
			external.Name,
		)

		if err != nil {
			log.Error().
				Err(err).
				Str("provider", request.Provider).
				Str("username", nickname).
				Msg("Failed to create user")

			return ExternalCallback412JSONResponse{
				Message: ToPtr("Failed to create user"),
				Status:  ToPtr(http.StatusPreconditionFailed),
			}, nil
		}

		a.session.Put(
			ctx,
			"user",
			user.ID,
		)

		log.Debug().
			Str("provider", request.Provider).
			Str("username", user.Username).
			Str("email", user.Email).
			Str("external", external.UserID).
			Msg("authenticated")

		return ExternalCallbackRedirectResponse{
			url: strings.Join(
				[]string{
					a.config.Server.Host,
					a.config.Server.Root,
				},
				"",
			),
		}, nil
	}

	authValues := url.Values{}

	if request.Params.Code != nil {
		authValues.Set("code", FromPtr(request.Params.Code))
	}

	if request.Params.State != nil {
		authValues.Set("state", FromPtr(request.Params.State))
	}

	if _, err = session.Authorize(
		provider,
		authValues,
	); err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to authorize session")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to authorize session"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	a.session.Put(
		ctx,
		request.Provider,
		session.Marshal(),
	)

	external, err = provider.FetchUser(session)

	log.Debug().
		Str("provider", external.Provider).
		Str("email", external.Email).
		Str("name", external.Name).
		Str("firstname", external.FirstName).
		Str("lastname", external.LastName).
		Str("nickname", external.NickName).
		Str("user_id", external.UserID).
		Msg("requested auth")

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Msg("Failed to fetch user")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to fetch user"),
			Status:  ToPtr(http.StatusInternalServerError),
		}, nil
	}

	nickname := slugify.Slugify(external.NickName)

	user, err := a.users.External(
		ctx,
		external.Provider,
		external.UserID,
		nickname,
		external.Email,
		external.Name,
	)

	if err != nil {
		log.Error().
			Err(err).
			Str("provider", request.Provider).
			Str("username", nickname).
			Msg("Failed to create user")

		return ExternalCallback412JSONResponse{
			Message: ToPtr("Failed to create user"),
			Status:  ToPtr(http.StatusPreconditionFailed),
		}, nil
	}

	a.session.Put(
		ctx,
		"user",
		user.ID,
	)

	log.Debug().
		Str("provider", request.Provider).
		Str("username", user.Username).
		Str("email", user.Email).
		Str("external", external.UserID).
		Msg("authenticated")

	return ExternalCallbackRedirectResponse{
		url: strings.Join(
			[]string{
				a.config.Server.Host,
				a.config.Server.Root,
			},
			"",
		),
	}, nil
}

// ExternalCallbackRedirectResponse defines the response to redirect to a defined URL.
type ExternalCallbackRedirectResponse struct {
	url string
}

// VisitExternalCallbackResponse implements the middleware.Responder interface for redirects.
func (r ExternalCallbackRedirectResponse) VisitExternalCallbackResponse(w http.ResponseWriter) error {
	w.Header().Set(
		"Location",
		r.url,
	)

	w.Header().Set(
		"Content-Type",
		"text/html",
	)

	w.WriteHeader(
		http.StatusTemporaryRedirect,
	)

	return nil
}

func setAuthState(state *string) string {
	if state != nil && len(FromPtr(state)) > 0 {
		return FromPtr(state)
	}

	nonceBytes := make([]byte, 64)

	if _, err := io.ReadFull(
		rand.Reader,
		nonceBytes,
	); err != nil {
		log.Error().
			Err(err).
			Msg("Source of randomness unavailable")

		panic("source of randomness unavailable")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func verifyAuthState(state *string, sess goth.Session) error {
	rawAuth, err := sess.GetAuthURL()

	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuth)

	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")

	if originalState != "" && (state == nil || originalState != FromPtr(state)) {
		return fmt.Errorf("state token mismatch")
	}

	return nil
}
