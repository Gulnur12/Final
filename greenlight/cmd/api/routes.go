package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/books", app.requirePermission("books:read", app.listBooksHandler))
	router.HandlerFunc(http.MethodPost, "/v1/books", app.requirePermission("books:write", app.createBookHandler))
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.requirePermission("books:read", app.showBookHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/books/:id", app.requirePermission("books:write", app.updateBookHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/books/:id", app.requirePermission("books:write", app.deleteBookHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
