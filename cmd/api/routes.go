package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
    // initialize a new instance of httprouter

    router := httprouter.New()

    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHander)
    router.HandlerFunc(http.MethodPost, "/v1/movies", app.MovieHander)
    router.HandlerFunc(http.MethodGet, "/v1/movies/:id", appshowMovieHander)

    return router
}
