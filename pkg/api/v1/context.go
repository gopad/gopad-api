package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/rs/zerolog/log"
)

const (
	groupContext contextKey = "group"
	userContext  contextKey = "user"
)

// GroupToContext is used to put the requested group into the context.
func (a *API) GroupToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "group_id")

		record, err := a.storage.Groups.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrGroupNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find group"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "GroupToContext").
				Str("group", id).
				Msg("Failed to load group")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load group"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			groupContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GroupFromContext is used to get the requested group from the context.
func (a *API) GroupFromContext(ctx context.Context) *model.Group {
	record, ok := ctx.Value(groupContext).(*model.Group)

	if !ok {
		return nil
	}

	return record
}

// UserToContext is used to put the requested user into the context.
func (a *API) UserToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "user_id")

		record, err := a.storage.Users.Show(
			ctx,
			id,
		)

		if err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				a.RenderNotify(w, r, Notification{
					Message: ToPtr("Failed to find user"),
					Status:  ToPtr(http.StatusNotFound),
				})

				return
			}

			log.Error().
				Err(err).
				Str("action", "UserToContext").
				Str("user", id).
				Msg("Failed to load user")

			a.RenderNotify(w, r, Notification{
				Message: ToPtr("Failed to load user"),
				Status:  ToPtr(http.StatusInternalServerError),
			})

			return
		}

		ctx = context.WithValue(
			ctx,
			userContext,
			record,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserFromContext is used to get the requested user from the context.
func (a *API) UserFromContext(ctx context.Context) *model.User {
	record, ok := ctx.Value(userContext).(*model.User)

	if !ok {
		return nil
	}

	return record
}
