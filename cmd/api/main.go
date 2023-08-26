package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


const VERSION = "1.0.0"

type config struct {
    port int
    env string
}

type application struct {
    config config 
    logger *log.Logger
}


func main() {
    var cfg config 

    // Read the port and env command-line flags
    // default to 4000 and "development"
    flag.IntVar(&cfg.port, "port", 4000, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    // Initialize a new logger which writes message to stdout
    // prefix with date and time
    logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

    // Declare an instance of applications

    app := &application{
        config: cfg,
        logger: logger,
    }

    // Declare a new servemux and add /v1/healthcheck route
    // this route dispatched the healthCheckHandler method

    mux := http.NewServeMux()
    mux.HandleFunc("/v1/healthCheck", app.healthCheckHander)

    srv := &http.Server{
        Addr: fmt.Sprintf(":%d", cfg.port),
        Handler: mux,
        IdleTimeout: time.Minute,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    // Start server
    logger.Printf("Beep Boop: Starting %s server on port: %s", cfg.env, srv.Addr)
    err := srv.ListenAndServe()
    logger.Fatal(err)
}
