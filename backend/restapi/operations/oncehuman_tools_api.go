// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/mod"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/system"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/tools"
)

// NewOncehumanToolsAPI creates a new OncehumanTools instance
func NewOncehumanToolsAPI(spec *loads.Document) *OncehumanToolsAPI {
	return &OncehumanToolsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ModCalculateAffixProbabilityHandler: mod.CalculateAffixProbabilityHandlerFunc(func(params mod.CalculateAffixProbabilityParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.CalculateAffixProbability has not yet been implemented")
		}),
		ModCalculateStrengthenProbabilityHandler: mod.CalculateStrengthenProbabilityHandlerFunc(func(params mod.CalculateStrengthenProbabilityParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.CalculateStrengthenProbability has not yet been implemented")
		}),
		SystemHealthCheckHandler: system.HealthCheckHandlerFunc(func(params system.HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation system.HealthCheck has not yet been implemented")
		}),
		ModListAffixesHandler: mod.ListAffixesHandlerFunc(func(params mod.ListAffixesParams) middleware.Responder {
			return middleware.NotImplemented("operation mod.ListAffixes has not yet been implemented")
		}),
		ToolsListToolsHandler: tools.ListToolsHandlerFunc(func(params tools.ListToolsParams) middleware.Responder {
			return middleware.NotImplemented("operation tools.ListTools has not yet been implemented")
		}),
	}
}

/*OncehumanToolsAPI 提供OnceHuman游戏相关的计算工具API */
type OncehumanToolsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// ModCalculateAffixProbabilityHandler sets the operation handler for the calculate affix probability operation
	ModCalculateAffixProbabilityHandler mod.CalculateAffixProbabilityHandler
	// ModCalculateStrengthenProbabilityHandler sets the operation handler for the calculate strengthen probability operation
	ModCalculateStrengthenProbabilityHandler mod.CalculateStrengthenProbabilityHandler
	// SystemHealthCheckHandler sets the operation handler for the health check operation
	SystemHealthCheckHandler system.HealthCheckHandler
	// ModListAffixesHandler sets the operation handler for the list affixes operation
	ModListAffixesHandler mod.ListAffixesHandler
	// ToolsListToolsHandler sets the operation handler for the list tools operation
	ToolsListToolsHandler tools.ListToolsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *OncehumanToolsAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *OncehumanToolsAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *OncehumanToolsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *OncehumanToolsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *OncehumanToolsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *OncehumanToolsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *OncehumanToolsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *OncehumanToolsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *OncehumanToolsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the OncehumanToolsAPI
func (o *OncehumanToolsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ModCalculateAffixProbabilityHandler == nil {
		unregistered = append(unregistered, "mod.CalculateAffixProbabilityHandler")
	}
	if o.ModCalculateStrengthenProbabilityHandler == nil {
		unregistered = append(unregistered, "mod.CalculateStrengthenProbabilityHandler")
	}
	if o.SystemHealthCheckHandler == nil {
		unregistered = append(unregistered, "system.HealthCheckHandler")
	}
	if o.ModListAffixesHandler == nil {
		unregistered = append(unregistered, "mod.ListAffixesHandler")
	}
	if o.ToolsListToolsHandler == nil {
		unregistered = append(unregistered, "tools.ListToolsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *OncehumanToolsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *OncehumanToolsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *OncehumanToolsAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *OncehumanToolsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *OncehumanToolsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
func (o *OncehumanToolsAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the oncehuman tools API
func (o *OncehumanToolsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *OncehumanToolsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mod/affix/probability"] = mod.NewCalculateAffixProbability(o.context, o.ModCalculateAffixProbabilityHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mod/strengthen/probability"] = mod.NewCalculateStrengthenProbability(o.context, o.ModCalculateStrengthenProbabilityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = system.NewHealthCheck(o.context, o.SystemHealthCheckHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mod/affix/list"] = mod.NewListAffixes(o.context, o.ModListAffixesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tools"] = tools.NewListTools(o.context, o.ToolsListToolsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *OncehumanToolsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *OncehumanToolsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *OncehumanToolsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *OncehumanToolsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *OncehumanToolsAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
