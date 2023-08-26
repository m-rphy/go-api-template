package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
    // When httprouter is parsing a request, and interpolated URL parameters will be
    // stored in the request context. We can use the ParamsFromContext() func
    // to retrieve a slice containing these parameter names and values

    params := httprouter.ParamsFromContext(r.Context())

    // ByName() will always return a string. We must type cast it
    // to a base 10 int64 numnber. Furthermore the number will always
    // be postive
    id , err := strconv.ParseInt(params.ByName("id"), 10, 64)
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "show the detail of the movie %d\n", id)
}
