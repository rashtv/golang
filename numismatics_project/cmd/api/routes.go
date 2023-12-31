package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/coins", app.requirePermission("coins:read", app.listCoinsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/coins", app.requirePermission("coins:write", app.createCoinHandler))
	router.HandlerFunc(http.MethodGet, "/v1/coins/:id", app.requirePermission("coins:read", app.showCoinHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/coins/:id", app.requirePermission("coins:write", app.updateCoinHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/coins/:id", app.requirePermission("coins:write", app.deleteCoinHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
