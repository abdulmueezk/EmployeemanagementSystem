// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"employeeManagementSystem/gen/restapi/operations"
	"employeeManagementSystem/gen/restapi/operations/admin"
	"employeeManagementSystem/gen/restapi/operations/employee"
	"employeeManagementSystem/gen/restapi/operations/team_lead"
)

//go:generate swagger generate server --target ../../gen --name EmployeeManagemntSystem --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.EmployeeManagemntSystemAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EmployeeManagemntSystemAPI) http.Handler {
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

	if api.AdminDeleteAdminDeleteCnicHandler == nil {
		api.AdminDeleteAdminDeleteCnicHandler = admin.DeleteAdminDeleteCnicHandlerFunc(func(params admin.DeleteAdminDeleteCnicParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.DeleteAdminDeleteCnic has not yet been implemented")
		})
	}
	if api.EmployeeGetEmployeeGetrecoardCnicHandler == nil {
		api.EmployeeGetEmployeeGetrecoardCnicHandler = employee.GetEmployeeGetrecoardCnicHandlerFunc(func(params employee.GetEmployeeGetrecoardCnicParams) middleware.Responder {
			return middleware.NotImplemented("operation employee.GetEmployeeGetrecoardCnic has not yet been implemented")
		})
	}
	if api.TeamLeadGetTeamLeadGetrecoardCnicHandler == nil {
		api.TeamLeadGetTeamLeadGetrecoardCnicHandler = team_lead.GetTeamLeadGetrecoardCnicHandlerFunc(func(params team_lead.GetTeamLeadGetrecoardCnicParams) middleware.Responder {
			return middleware.NotImplemented("operation team_lead.GetTeamLeadGetrecoardCnic has not yet been implemented")
		})
	}
	if api.TeamLeadGetTeamLeadGetteamrecoardCnicHandler == nil {
		api.TeamLeadGetTeamLeadGetteamrecoardCnicHandler = team_lead.GetTeamLeadGetteamrecoardCnicHandlerFunc(func(params team_lead.GetTeamLeadGetteamrecoardCnicParams) middleware.Responder {
			return middleware.NotImplemented("operation team_lead.GetTeamLeadGetteamrecoardCnic has not yet been implemented")
		})
	}
	if api.AdminPostAdminAddHandler == nil {
		api.AdminPostAdminAddHandler = admin.PostAdminAddHandlerFunc(func(params admin.PostAdminAddParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.PostAdminAdd has not yet been implemented")
		})
	}
	if api.AdminPostAdminUpdateCnicHandler == nil {
		api.AdminPostAdminUpdateCnicHandler = admin.PostAdminUpdateCnicHandlerFunc(func(params admin.PostAdminUpdateCnicParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.PostAdminUpdateCnic has not yet been implemented")
		})
	}
	if api.AdminGetemployeeRecoardHandler == nil {
		api.AdminGetemployeeRecoardHandler = admin.GetemployeeRecoardHandlerFunc(func(params admin.GetemployeeRecoardParams) middleware.Responder {
			return middleware.NotImplemented("operation admin.GetemployeeRecoard has not yet been implemented")
		})
	}

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
