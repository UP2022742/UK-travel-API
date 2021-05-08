package worker

import (
	"context"
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
	}
}

func (route *Worker) ListenAPIServer(stop chan bool) {
	route.logger.Info(
		"Starting Web Server",
		"address", route.apiServer.Addr,
	)

	go func() {
		err := route.apiServer.ListenAndServe()
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
