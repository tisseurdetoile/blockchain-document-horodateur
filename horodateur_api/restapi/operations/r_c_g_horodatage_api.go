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
)

// NewRCGHorodatageAPI creates a new RCGHorodatage instance
func NewRCGHorodatageAPI(spec *loads.Document) *RCGHorodatageAPI {
	return &RCGHorodatageAPI{
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
		BinProducer:         runtime.ByteStreamProducer(),
		ConfigureSAMLHandler: ConfigureSAMLHandlerFunc(func(params ConfigureSAMLParams) middleware.Responder {
			return middleware.NotImplemented("operation ConfigureSAML has not yet been implemented")
		}),
		DelreceiptsHandler: DelreceiptsHandlerFunc(func(params DelreceiptsParams) middleware.Responder {
			return middleware.NotImplemented("operation Delreceipts has not yet been implemented")
		}),
		GetreceiptHandler: GetreceiptHandlerFunc(func(params GetreceiptParams) middleware.Responder {
			return middleware.NotImplemented("operation Getreceipt has not yet been implemented")
		}),
		ListtimestampedHandler: ListtimestampedHandlerFunc(func(params ListtimestampedParams) middleware.Responder {
			return middleware.NotImplemented("operation Listtimestamped has not yet been implemented")
		}),
		MonitoringHandler: MonitoringHandlerFunc(func(params MonitoringParams) middleware.Responder {
			return middleware.NotImplemented("operation Monitoring has not yet been implemented")
		}),
	}
}

/*RCGHorodatageAPI RCG horodatage est un service qui permet l'horodatage numérique via
sur la blockchain Ethereum.
Le principe est d'envoyer des fichiers qui sont ensuite passer dans
une fonction hachage SHA3-256. Les « hash » sont ensuite intégrés
dans un arbre de Merkle dont la racine est inséré dans une
transaction blockchain, l'(es) adresse(s) signant la transaction
identifie le Registre du Commerce, c'est une information qui doit
être publique.
*/
type RCGHorodatageAPI struct {
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
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// BinProducer registers a producer for a "application/octet-stream" mime type
	BinProducer runtime.Producer

	// ConfigureSAMLHandler sets the operation handler for the configure s a m l operation
	ConfigureSAMLHandler ConfigureSAMLHandler
	// DelreceiptsHandler sets the operation handler for the delreceipts operation
	DelreceiptsHandler DelreceiptsHandler
	// GetreceiptHandler sets the operation handler for the getreceipt operation
	GetreceiptHandler GetreceiptHandler
	// ListtimestampedHandler sets the operation handler for the listtimestamped operation
	ListtimestampedHandler ListtimestampedHandler
	// MonitoringHandler sets the operation handler for the monitoring operation
	MonitoringHandler MonitoringHandler

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
func (o *RCGHorodatageAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *RCGHorodatageAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *RCGHorodatageAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *RCGHorodatageAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *RCGHorodatageAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *RCGHorodatageAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *RCGHorodatageAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the RCGHorodatageAPI
func (o *RCGHorodatageAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}

	if o.ConfigureSAMLHandler == nil {
		unregistered = append(unregistered, "ConfigureSAMLHandler")
	}

	if o.DelreceiptsHandler == nil {
		unregistered = append(unregistered, "DelreceiptsHandler")
	}

	if o.GetreceiptHandler == nil {
		unregistered = append(unregistered, "GetreceiptHandler")
	}

	if o.ListtimestampedHandler == nil {
		unregistered = append(unregistered, "ListtimestampedHandler")
	}

	if o.MonitoringHandler == nil {
		unregistered = append(unregistered, "MonitoringHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *RCGHorodatageAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *RCGHorodatageAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *RCGHorodatageAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *RCGHorodatageAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

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
func (o *RCGHorodatageAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *RCGHorodatageAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the r c g horodatage API
func (o *RCGHorodatageAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *RCGHorodatageAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/saml"] = NewConfigureSAML(o.context, o.ConfigureSAMLHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/recu"] = NewDelreceipts(o.context, o.DelreceiptsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/recu"] = NewGetreceipt(o.context, o.GetreceiptHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/horodatage"] = NewListtimestamped(o.context, o.ListtimestampedHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sonde"] = NewMonitoring(o.context, o.MonitoringHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *RCGHorodatageAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *RCGHorodatageAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *RCGHorodatageAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *RCGHorodatageAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
