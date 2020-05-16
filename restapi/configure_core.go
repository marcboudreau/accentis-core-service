// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/marcboudreau/accentis-core-service/handlers"
	"github.com/marcboudreau/accentis-core-service/health"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/marcboudreau/accentis-core-service/restapi/operations"
	"github.com/marcboudreau/accentis-core-service/restapi/operations/healthcheck"
)

//go:generate swagger generate server --target ../../core-service --name Core --spec ../swagger/swagger.yaml

func configureFlags(api *operations.CoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HealthcheckHealthCheckHandler = healthcheck.HealthCheckHandlerFunc(func(params healthcheck.HealthCheckParams) middleware.Responder {
		return handlers.HandleHealthCheck(params)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	health.Set(true)

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
	//http.Handle("/metrics", promhttp.Handler())

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			promhttp.Handler().ServeHTTP(rw, r)
			return
		}

		handler.ServeHTTP(rw, r)
	})
}
