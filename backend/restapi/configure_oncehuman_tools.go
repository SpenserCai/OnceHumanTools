/*
 * @Author: SpenserCai
 * @Date: 2025-06-23 17:50:55
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2025-06-23 18:00:09
 * @Description: file content
 */
// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/SpenserCai/OnceHumanTools/backend/internal/handlers"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/mod"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/system"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/tools"
	// 导入我们的处理器
)

//go:generate swagger generate server --target ../../backend --name OncehumanTools --spec ../api/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.OncehumanToolsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OncehumanToolsAPI) http.Handler {
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

	// 创建处理器实例
	systemHandler := handlers.NewSystemHandler()
	toolsHandler := handlers.NewToolsHandler()
	modHandler := handlers.NewModHandler()

	// 连接模组相关处理器
	api.ModCalculateAffixProbabilityHandler = mod.CalculateAffixProbabilityHandlerFunc(modHandler.CalculateAffixProbability)
	api.ModCalculateStrengthenProbabilityHandler = mod.CalculateStrengthenProbabilityHandlerFunc(modHandler.CalculateStrengthenProbability)
	api.ModListAffixesHandler = mod.ListAffixesHandlerFunc(modHandler.ListAffixes)

	// 连接系统处理器
	api.SystemHealthCheckHandler = system.HealthCheckHandlerFunc(systemHandler.HealthCheck)

	// 连接工具处理器
	api.ToolsListToolsHandler = tools.ListToolsHandlerFunc(toolsHandler.ListTools)

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
