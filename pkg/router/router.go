package router

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	oamw "github.com/go-openapi/runtime/middleware"
	v1 "github.com/gopad/gopad-api/pkg/api/v1"
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/metrics"
	"github.com/gopad/gopad-api/pkg/middleware/current"
	"github.com/gopad/gopad-api/pkg/middleware/header"
	"github.com/gopad/gopad-api/pkg/model"
	"github.com/gopad/gopad-api/pkg/respond"
	"github.com/gopad/gopad-api/pkg/service/members"
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/gopad/gopad-api/pkg/session"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/gopad/gopad-api/pkg/token"
	"github.com/gopad/gopad-api/pkg/upload"
	cgmw "github.com/oapi-codegen/nethttp-middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	registry *metrics.Metrics,
	sess *session.Session,
	uploads upload.Upload,
	storage store.Store,
	teamsService teams.Service,
	usersService users.Service,
	membersService members.Service,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.RemoteAddrHandler("ip"))
	mux.Use(hlog.URLHandler("path"))
	mux.Use(hlog.MethodHandler("method"))
	mux.Use(hlog.RequestIDHandler("request_id", "Request-Id"))

	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("request")
	}))

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)
	mux.Use(sess.Middleware)
	mux.Use(current.Middleware)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		root.Get("/", func(w http.ResponseWriter, r *http.Request) {

			respond.JSON(
				w,
				r,
				[]string{
					sess.Get(
						r.Context(),
						"user",
					),
				},
			)

		})

		root.Route("/api/v1", func(rapi chi.Router) {
			swagger, err := v1.GetSwagger()

			if err != nil {
				log.Error().
					Err(err).
					Str("version", "v1").
					Msg("Failed to load openapi spec")
			}

			swagger.Servers = openapi3.Servers{
				{
					URL: cfg.Server.Host + path.Join(
						cfg.Server.Root,
						"api",
						"v1",
					),
				},
			}

			rapi.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
				respond.JSON(
					w,
					r,
					swagger,
				)
			})

			rapi.Handle("/docs", oamw.SwaggerUI(oamw.SwaggerUIOpts{
				Path: path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"docs",
				),
				SpecURL: cfg.Server.Host + path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"swagger",
				),
			}, nil))

			rapi.With(cgmw.OapiRequestValidatorWithOptions(
				swagger,
				&cgmw.Options{
					SilenceServersWarning: true,
					Options: openapi3filter.Options{
						AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
							authenticating := &model.User{}
							scheme := input.SecuritySchemeName
							operation := input.RequestValidationInput.Route.Operation.OperationID

							logger := log.With().
								Str("scheme", scheme).
								Str("operation", operation).
								Logger()

							switch scheme {
							case "Cookie":
								userID := sess.Get(
									input.RequestValidationInput.Request.Context(),
									"user",
								)

								if userID == "" {
									return fmt.Errorf("no session cookie present")
								}

								user, err := usersService.AuthByID(
									ctx,
									userID,
								)

								if err != nil {
									logger.Error().
										Err(err).
										Str("user", userID).
										Msg("failed to find user")

									return fmt.Errorf("failed to find user")
								}

								logger.Trace().
									Str("user", userID).
									Msg("authentication")

								authenticating = user

							case "Header":
								header := input.RequestValidationInput.Request.Header.Get(
									input.SecurityScheme.Name,
								)

								if header == "" {
									return fmt.Errorf("missing authorization header")
								}

								t, err := token.Parse(
									strings.TrimSpace(
										header,
									),
									cfg.Session.Secret,
								)

								if err != nil {
									return fmt.Errorf("failed to parse auth token")
								}

								user, err := usersService.AuthByID(
									ctx,
									t.Text,
								)

								if err != nil {
									logger.Error().
										Err(err).
										Str("user", t.Text).
										Msg("failed to find user")

									return fmt.Errorf("failed to find user")
								}

								logger.Trace().
									Str("user", t.Text).
									Msg("authentication")

								authenticating = user

							case "Bearer":
								header := input.RequestValidationInput.Request.Header.Get(
									"Authorization",
								)

								if header == "" {
									return fmt.Errorf("missing authorization header")
								}

								t, err := token.Parse(
									strings.TrimSpace(
										strings.Replace(
											header,
											"Bearer",
											"",
											1,
										),
									),
									cfg.Session.Secret,
								)

								if err != nil {
									return fmt.Errorf("failed to parse auth token")
								}

								user, err := usersService.AuthByID(
									ctx,
									t.Text,
								)

								if err != nil {
									logger.Error().
										Err(err).
										Str("user", t.Text).
										Msg("failed to find user")

									return fmt.Errorf("failed to find user")
								}

								logger.Trace().
									Str("user", t.Text).
									Msg("authentication")

								authenticating = user

							case "Basic":
								username, password, ok := input.RequestValidationInput.Request.BasicAuth()

								if !ok {
									return fmt.Errorf("missing basic credentials")
								}

								user, err := usersService.AuthByCreds(
									ctx,
									username,
									password,
								)

								if err != nil {
									logger.Error().
										Err(err).
										Str("user", username).
										Msg("wrong credentials")

									return fmt.Errorf("wrong credentials")
								}

								logger.Trace().
									Str("user", username).
									Msg("authentication")

								authenticating = user

							default:
								return fmt.Errorf("unknown security scheme: %s", scheme)
							}

							log.Trace().
								Str("username", authenticating.Username).
								Str("operation", operation).
								Msg("authenticated")

							current.SetUser(
								input.RequestValidationInput.Request.Context(),
								authenticating,
							)

							return nil
						},
					},
				},
			)).Mount("/", v1.Handler(
				v1.NewStrictHandler(
					v1.New(
						cfg,
						registry,
						sess,
						uploads,
						storage,
						teamsService,
						usersService,
						membersService,
					),
					make([]v1.StrictMiddlewareFunc, 0),
				),
			))

			rapi.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"storage",
				),
			))
		})
	})

	return mux
}

// Metrics initializes the routing of metrics and health.
func Metrics(
	cfg *config.Config,
	registry *metrics.Metrics,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)

	mux.Route("/", func(root chi.Router) {
		root.Get("/metrics", registry.Handler())

		if cfg.Metrics.Pprof {
			root.Mount("/debug", middleware.Profiler())
		}

		root.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			_, _ = io.WriteString(w, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)

			_, _ = io.WriteString(w, http.StatusText(http.StatusOK))
		})
	})

	return mux
}
