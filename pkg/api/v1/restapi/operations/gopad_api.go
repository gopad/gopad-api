// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/profile"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/team"
	"github.com/gopad/gopad-api/pkg/api/v1/restapi/operations/user"
)

// NewGopadAPI creates a new Gopad instance
func NewGopadAPI(spec *loads.Document) *GopadAPI {
	return &GopadAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AuthAuthLoginHandler: auth.AuthLoginHandlerFunc(func(params auth.AuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthAuthLogin has not yet been implemented")
		}),
		AuthAuthRefreshHandler: auth.AuthRefreshHandlerFunc(func(params auth.AuthRefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthAuthRefresh has not yet been implemented")
		}),
		AuthAuthVerifyHandler: auth.AuthVerifyHandlerFunc(func(params auth.AuthVerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthAuthVerify has not yet been implemented")
		}),
		ProfileProfileShowHandler: profile.ProfileShowHandlerFunc(func(params profile.ProfileShowParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileProfileShow has not yet been implemented")
		}),
		ProfileProfileTokenHandler: profile.ProfileTokenHandlerFunc(func(params profile.ProfileTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileProfileToken has not yet been implemented")
		}),
		ProfileProfileUpdateHandler: profile.ProfileUpdateHandlerFunc(func(params profile.ProfileUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileProfileUpdate has not yet been implemented")
		}),
		TeamTeamCreateHandler: team.TeamCreateHandlerFunc(func(params team.TeamCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamCreate has not yet been implemented")
		}),
		TeamTeamDeleteHandler: team.TeamDeleteHandlerFunc(func(params team.TeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamDelete has not yet been implemented")
		}),
		TeamTeamIndexHandler: team.TeamIndexHandlerFunc(func(params team.TeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamIndex has not yet been implemented")
		}),
		TeamTeamShowHandler: team.TeamShowHandlerFunc(func(params team.TeamShowParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamShow has not yet been implemented")
		}),
		TeamTeamUpdateHandler: team.TeamUpdateHandlerFunc(func(params team.TeamUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamUpdate has not yet been implemented")
		}),
		TeamTeamUserAppendHandler: team.TeamUserAppendHandlerFunc(func(params team.TeamUserAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamUserAppend has not yet been implemented")
		}),
		TeamTeamUserDeleteHandler: team.TeamUserDeleteHandlerFunc(func(params team.TeamUserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamUserDelete has not yet been implemented")
		}),
		TeamTeamUserIndexHandler: team.TeamUserIndexHandlerFunc(func(params team.TeamUserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamUserIndex has not yet been implemented")
		}),
		TeamTeamUserPermHandler: team.TeamUserPermHandlerFunc(func(params team.TeamUserPermParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamTeamUserPerm has not yet been implemented")
		}),
		UserUserCreateHandler: user.UserCreateHandlerFunc(func(params user.UserCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserCreate has not yet been implemented")
		}),
		UserUserDeleteHandler: user.UserDeleteHandlerFunc(func(params user.UserDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserDelete has not yet been implemented")
		}),
		UserUserIndexHandler: user.UserIndexHandlerFunc(func(params user.UserIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserIndex has not yet been implemented")
		}),
		UserUserShowHandler: user.UserShowHandlerFunc(func(params user.UserShowParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserShow has not yet been implemented")
		}),
		UserUserTeamAppendHandler: user.UserTeamAppendHandlerFunc(func(params user.UserTeamAppendParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserTeamAppend has not yet been implemented")
		}),
		UserUserTeamDeleteHandler: user.UserTeamDeleteHandlerFunc(func(params user.UserTeamDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserTeamDelete has not yet been implemented")
		}),
		UserUserTeamIndexHandler: user.UserTeamIndexHandlerFunc(func(params user.UserTeamIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserTeamIndex has not yet been implemented")
		}),
		UserUserTeamPermHandler: user.UserTeamPermHandlerFunc(func(params user.UserTeamPermParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserTeamPerm has not yet been implemented")
		}),
		UserUserUpdateHandler: user.UserUpdateHandlerFunc(func(params user.UserUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUserUpdate has not yet been implemented")
		}),
	}
}

/*GopadAPI API definition for Gopad */
type GopadAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AuthAuthLoginHandler sets the operation handler for the auth login operation
	AuthAuthLoginHandler auth.AuthLoginHandler
	// AuthAuthRefreshHandler sets the operation handler for the auth refresh operation
	AuthAuthRefreshHandler auth.AuthRefreshHandler
	// AuthAuthVerifyHandler sets the operation handler for the auth verify operation
	AuthAuthVerifyHandler auth.AuthVerifyHandler
	// ProfileProfileShowHandler sets the operation handler for the profile show operation
	ProfileProfileShowHandler profile.ProfileShowHandler
	// ProfileProfileTokenHandler sets the operation handler for the profile token operation
	ProfileProfileTokenHandler profile.ProfileTokenHandler
	// ProfileProfileUpdateHandler sets the operation handler for the profile update operation
	ProfileProfileUpdateHandler profile.ProfileUpdateHandler
	// TeamTeamCreateHandler sets the operation handler for the team create operation
	TeamTeamCreateHandler team.TeamCreateHandler
	// TeamTeamDeleteHandler sets the operation handler for the team delete operation
	TeamTeamDeleteHandler team.TeamDeleteHandler
	// TeamTeamIndexHandler sets the operation handler for the team index operation
	TeamTeamIndexHandler team.TeamIndexHandler
	// TeamTeamShowHandler sets the operation handler for the team show operation
	TeamTeamShowHandler team.TeamShowHandler
	// TeamTeamUpdateHandler sets the operation handler for the team update operation
	TeamTeamUpdateHandler team.TeamUpdateHandler
	// TeamTeamUserAppendHandler sets the operation handler for the team user append operation
	TeamTeamUserAppendHandler team.TeamUserAppendHandler
	// TeamTeamUserDeleteHandler sets the operation handler for the team user delete operation
	TeamTeamUserDeleteHandler team.TeamUserDeleteHandler
	// TeamTeamUserIndexHandler sets the operation handler for the team user index operation
	TeamTeamUserIndexHandler team.TeamUserIndexHandler
	// TeamTeamUserPermHandler sets the operation handler for the team user perm operation
	TeamTeamUserPermHandler team.TeamUserPermHandler
	// UserUserCreateHandler sets the operation handler for the user create operation
	UserUserCreateHandler user.UserCreateHandler
	// UserUserDeleteHandler sets the operation handler for the user delete operation
	UserUserDeleteHandler user.UserDeleteHandler
	// UserUserIndexHandler sets the operation handler for the user index operation
	UserUserIndexHandler user.UserIndexHandler
	// UserUserShowHandler sets the operation handler for the user show operation
	UserUserShowHandler user.UserShowHandler
	// UserUserTeamAppendHandler sets the operation handler for the user team append operation
	UserUserTeamAppendHandler user.UserTeamAppendHandler
	// UserUserTeamDeleteHandler sets the operation handler for the user team delete operation
	UserUserTeamDeleteHandler user.UserTeamDeleteHandler
	// UserUserTeamIndexHandler sets the operation handler for the user team index operation
	UserUserTeamIndexHandler user.UserTeamIndexHandler
	// UserUserTeamPermHandler sets the operation handler for the user team perm operation
	UserUserTeamPermHandler user.UserTeamPermHandler
	// UserUserUpdateHandler sets the operation handler for the user update operation
	UserUserUpdateHandler user.UserUpdateHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *GopadAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GopadAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GopadAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GopadAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GopadAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GopadAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GopadAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GopadAPI
func (o *GopadAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthAuthLoginHandler == nil {
		unregistered = append(unregistered, "auth.AuthLoginHandler")
	}

	if o.AuthAuthRefreshHandler == nil {
		unregistered = append(unregistered, "auth.AuthRefreshHandler")
	}

	if o.AuthAuthVerifyHandler == nil {
		unregistered = append(unregistered, "auth.AuthVerifyHandler")
	}

	if o.ProfileProfileShowHandler == nil {
		unregistered = append(unregistered, "profile.ProfileShowHandler")
	}

	if o.ProfileProfileTokenHandler == nil {
		unregistered = append(unregistered, "profile.ProfileTokenHandler")
	}

	if o.ProfileProfileUpdateHandler == nil {
		unregistered = append(unregistered, "profile.ProfileUpdateHandler")
	}

	if o.TeamTeamCreateHandler == nil {
		unregistered = append(unregistered, "team.TeamCreateHandler")
	}

	if o.TeamTeamDeleteHandler == nil {
		unregistered = append(unregistered, "team.TeamDeleteHandler")
	}

	if o.TeamTeamIndexHandler == nil {
		unregistered = append(unregistered, "team.TeamIndexHandler")
	}

	if o.TeamTeamShowHandler == nil {
		unregistered = append(unregistered, "team.TeamShowHandler")
	}

	if o.TeamTeamUpdateHandler == nil {
		unregistered = append(unregistered, "team.TeamUpdateHandler")
	}

	if o.TeamTeamUserAppendHandler == nil {
		unregistered = append(unregistered, "team.TeamUserAppendHandler")
	}

	if o.TeamTeamUserDeleteHandler == nil {
		unregistered = append(unregistered, "team.TeamUserDeleteHandler")
	}

	if o.TeamTeamUserIndexHandler == nil {
		unregistered = append(unregistered, "team.TeamUserIndexHandler")
	}

	if o.TeamTeamUserPermHandler == nil {
		unregistered = append(unregistered, "team.TeamUserPermHandler")
	}

	if o.UserUserCreateHandler == nil {
		unregistered = append(unregistered, "user.UserCreateHandler")
	}

	if o.UserUserDeleteHandler == nil {
		unregistered = append(unregistered, "user.UserDeleteHandler")
	}

	if o.UserUserIndexHandler == nil {
		unregistered = append(unregistered, "user.UserIndexHandler")
	}

	if o.UserUserShowHandler == nil {
		unregistered = append(unregistered, "user.UserShowHandler")
	}

	if o.UserUserTeamAppendHandler == nil {
		unregistered = append(unregistered, "user.UserTeamAppendHandler")
	}

	if o.UserUserTeamDeleteHandler == nil {
		unregistered = append(unregistered, "user.UserTeamDeleteHandler")
	}

	if o.UserUserTeamIndexHandler == nil {
		unregistered = append(unregistered, "user.UserTeamIndexHandler")
	}

	if o.UserUserTeamPermHandler == nil {
		unregistered = append(unregistered, "user.UserTeamPermHandler")
	}

	if o.UserUserUpdateHandler == nil {
		unregistered = append(unregistered, "user.UserUpdateHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GopadAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GopadAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *GopadAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *GopadAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *GopadAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GopadAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the gopad API
func (o *GopadAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GopadAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = auth.NewAuthLogin(o.context, o.AuthAuthLoginHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/refresh"] = auth.NewAuthRefresh(o.context, o.AuthAuthRefreshHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/verify/{token}"] = auth.NewAuthVerify(o.context, o.AuthAuthVerifyHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/self"] = profile.NewProfileShow(o.context, o.ProfileProfileShowHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/token"] = profile.NewProfileToken(o.context, o.ProfileProfileTokenHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/profile/self"] = profile.NewProfileUpdate(o.context, o.ProfileProfileUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams"] = team.NewTeamCreate(o.context, o.TeamTeamCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{teamID}"] = team.NewTeamDelete(o.context, o.TeamTeamDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams"] = team.NewTeamIndex(o.context, o.TeamTeamIndexHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{teamID}"] = team.NewTeamShow(o.context, o.TeamTeamShowHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{teamID}"] = team.NewTeamUpdate(o.context, o.TeamTeamUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams/{teamID}/users"] = team.NewTeamUserAppend(o.context, o.TeamTeamUserAppendHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{teamID}/users"] = team.NewTeamUserDelete(o.context, o.TeamTeamUserDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{teamID}/users"] = team.NewTeamUserIndex(o.context, o.TeamTeamUserIndexHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{teamID}/users"] = team.NewTeamUserPerm(o.context, o.TeamTeamUserPermHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewUserCreate(o.context, o.UserUserCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userID}"] = user.NewUserDelete(o.context, o.UserUserDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewUserIndex(o.context, o.UserUserIndexHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userID}"] = user.NewUserShow(o.context, o.UserUserShowHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{userID}/teams"] = user.NewUserTeamAppend(o.context, o.UserUserTeamAppendHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userID}/teams"] = user.NewUserTeamDelete(o.context, o.UserUserTeamDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userID}/teams"] = user.NewUserTeamIndex(o.context, o.UserUserTeamIndexHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{userID}/teams"] = user.NewUserTeamPerm(o.context, o.UserUserTeamPermHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{userID}"] = user.NewUserUpdate(o.context, o.UserUserUpdateHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GopadAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GopadAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GopadAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GopadAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
