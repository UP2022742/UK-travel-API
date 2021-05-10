package worker

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (route *Worker) CreateAPIServer() {
	r := mux.NewRouter()
	r.HandleFunc("/green", route.GreenList).Methods("GET")
	r.HandleFunc("/amber", route.AmberList).Methods("GET")
	r.HandleFunc("/red", route.RedList).Methods("GET")

	route.apiServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", route.apiPort),
		Handler: handlers.CORS()(r),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{route.cert},
		},
	}
}

func (route *Worker) ListenAPIServer(stop chan bool) {
	route.logger.Info(
		"Starting API Server",
		"address", route.apiServer.Addr,
	)

	go func() {

		// Shouldn't need to justify the certificate and key again but it has
		// problems. Look into this later.
		err := route.apiServer.ListenAndServeTLS(route.certFile, route.keyFile)
		if err != nil {
			route.logger.Crit(err.Error())
			stop <- true
		}
	}()
}

func (route *Worker) APIShutdown() error {
	route.logger.Info("Web Server Shutting down")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	return route.apiServer.Shutdown(ctx)
}
