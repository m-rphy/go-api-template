package main

import (
	"fmt"
	"net/http"
)


func (app *application) healthCheckHander(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s \n", app.config.env)
    fmt.Fprintf(w, "version: %s\n", VERSION)
}
