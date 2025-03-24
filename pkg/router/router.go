package router

import (
	"encoding/json"
	"net/http"
	"path"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	oamw "github.com/go-openapi/runtime/middleware"
	v1 "github.com/gopad/gopad-api/pkg/api/v1"
	"github.com/gopad/gopad-api/pkg/authn"
	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/handler"
	"github.com/gopad/gopad-api/pkg/metrics"
	"github.com/gopad/gopad-api/pkg/middleware/current"
	"github.com/gopad/gopad-api/pkg/middleware/header"
	"github.com/gopad/gopad-api/pkg/scim"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/gopad/gopad-api/pkg/upload"
	cgmw "github.com/oapi-codegen/nethttp-middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// Server initializes the routing of the server.
func Server(
	cfg *config.Config,
	registry *metrics.Metrics,
	identity *authn.Authn,
	uploads upload.Upload,
	storage *store.Store,
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
			Msg("Accesslog")
	}))

	mux.Use(render.SetContentType(render.ContentTypeJSON))
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(header.Version)
	mux.Use(header.Cache)
	mux.Use(header.Secure)
	mux.Use(header.Options)
	mux.Use(current.Middleware)

	mux.Route(cfg.Server.Root, func(root chi.Router) {
		if cfg.Scim.Enabled {
			srv, err := scim.New(
				scim.WithRoot(
					path.Join(
						cfg.Server.Root,
						"api",
						"scim",
						"v2",
					),
				),
				scim.WithStore(
					storage.Handle(),
				),
				scim.WithConfig(
					cfg.Scim,
				),
			).Server()

			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to linitialize scim server")
			}

			root.Mount("/api/scim/v2", srv)
		}

		root.Route("/api/v1", func(r chi.Router) {
			swagger, err := v1.GetSwagger()

			if err != nil {
				log.Error().
					Err(err).
					Str("version", "v1").
					Msg("Failed to load openapi spec")
			}

			swagger.Servers = openapi3.Servers{
				{
					URL: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
					),
				},
			}

			if cfg.Server.Docs {
				r.Get("/spec", func(w http.ResponseWriter, r *http.Request) {
					render.Status(r, http.StatusOK)
					render.JSON(w, r, swagger)
				})

				r.Handle("/docs", oamw.SwaggerUI(oamw.SwaggerUIOpts{
					Path: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
						"docs",
					),
					SpecURL: path.Join(
						cfg.Server.Root,
						"api",
						"v1",
						"spec",
					),
				}, nil))
			}

			apiv1 := v1.New(
				cfg,
				registry,
				identity,
				uploads,
				storage,
			)

			wrapper := v1.ServerInterfaceWrapper{
				Handler: apiv1,
				ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
					apiv1.RenderNotify(w, r, v1.Notification{
						Message: v1.ToPtr(err.Error()),
						Status:  v1.ToPtr(http.StatusBadRequest),
					})
				},
			}

			r.With(cgmw.OapiRequestValidatorWithOptions(
				swagger,
				&cgmw.Options{
					SilenceServersWarning: true,
					Options: openapi3filter.Options{
						AuthenticationFunc: apiv1.Authentication,
					},
					ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(statusCode)

						_ = json.NewEncoder(w).Encode(v1.Notification{
							Message: v1.ToPtr(message),
							Status:  v1.ToPtr(statusCode),
						})
					},
				},
			)).Route("/", func(r chi.Router) {
				r.Route("/auth", func(r chi.Router) {
					r.Group(func(r chi.Router) {
						r.Post("/redirect", wrapper.RedirectAuth)
						r.Post("/login", wrapper.LoginAuth)
						r.Get("/refresh", wrapper.RefreshAuth)
						r.Get("/verify", wrapper.VerifyAuth)
					})

					r.Group(func(r chi.Router) {
						r.Get("/providers", wrapper.ListProviders)

						r.Route("/{provider}", func(r chi.Router) {
							r.Use(render.SetContentType(render.ContentTypeHTML))

							r.Get("/callback", wrapper.CallbackProvider)
							r.Get("/request", wrapper.RequestProvider)
						})
					})
				})

				r.Route("/profile", func(r chi.Router) {
					r.Get("/self", wrapper.ShowProfile)
					r.Put("/self", wrapper.UpdateProfile)
					r.Get("/token", wrapper.TokenProfile)
				})

				r.Route("/groups", func(r chi.Router) {
					r.Get("/", wrapper.ListGroups)
					r.With(apiv1.AllowAdminAccessOnly).Post("/", wrapper.CreateGroup)

					r.Route("/{group_id}", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.GroupToContext)

						r.Get("/", wrapper.ShowGroup)
						r.Delete("/", wrapper.DeleteGroup)
						r.Put("/", wrapper.UpdateGroup)

						r.Route("/users", func(r chi.Router) {
							r.Get("/", wrapper.ListGroupUsers)
							r.Delete("/", wrapper.DeleteGroupFromUser)
							r.Post("/", wrapper.AttachGroupToUser)
							r.Put("/", wrapper.PermitGroupUser)
						})
					})
				})

				r.Route("/users", func(r chi.Router) {
					r.Get("/", wrapper.ListUsers)
					r.With(apiv1.AllowAdminAccessOnly).Post("/", wrapper.CreateUser)

					r.Route("/{user_id}", func(r chi.Router) {
						r.Use(apiv1.AllowAdminAccessOnly)
						r.Use(apiv1.UserToContext)

						r.Get("/", wrapper.ShowUser)
						r.Delete("/", wrapper.DeleteUser)
						r.Put("/", wrapper.UpdateUser)

						r.Route("/groups", func(r chi.Router) {
							r.Get("/", wrapper.ListUserGroups)
							r.Delete("/", wrapper.DeleteUserFromGroup)
							r.Post("/", wrapper.AttachUserToGroup)
							r.Put("/", wrapper.PermitUserGroup)
						})
					})
				})
			})

			r.Handle("/storage/*", uploads.Handler(
				path.Join(
					cfg.Server.Root,
					"api",
					"v1",
					"storage",
				),
			))
		})

		handlers := handler.New(cfg)
		root.Get("/", handlers.Index())
		root.Get("/favicon.svg", handlers.Favicon())
		root.Get("/config.json", handlers.Config())
		root.Get("/manifest.json", handlers.Manifest())
		root.Handle("/assets/*", handlers.Assets())
		root.NotFound(handlers.Index())
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

		root.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
		})

		root.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
		})
	})

	return mux
}
