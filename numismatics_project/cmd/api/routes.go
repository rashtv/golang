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
	router.HandlerFunc(http.MethodGet, "/v1/coins", app.listCoinsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/coins", app.createCoinHandler)
	router.HandlerFunc(http.MethodGet, "/v1/coins/:id", app.showCoinHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/coins/:id", app.updateCoinHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/coins/:id", app.deleteCoinHandler)

	return app.recoverPanic(router)
}
