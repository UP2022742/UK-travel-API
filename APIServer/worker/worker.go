package worker

import (
	"crypto/tls"
	"net/http"
	"os"

	log "github.com/inconshreveable/log15"
)

// Worker is the struct that organises the Clinic Service.
type Worker struct {
	logger    log.Logger
	apiServer *http.Server
	apiPort   int
	apiIP     string
	cert      tls.Certificate
}

// NewWorkerStructure Defines an instance of the worker structure.
func NewWorkerStructure(apiPort int, apiIP string, certFile string, keyFile string) *Worker {
	cert, _ := tls.LoadX509KeyPair("certs/certificate.crt", "certs/private.key")
	route := &Worker{
		logger:    log.New("module", "Worker"),
		apiServer: nil,
		apiPort:   apiPort,
		apiIP:     apiIP,
		cert:      cert,
	}
	route.logger.SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat()))
	return route
}
