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
	certFile  string
	keyFile   string
	cert      tls.Certificate
}

// NewWorkerStructure Defines an instance of the worker structure.
func NewWorkerStructure(apiPort int, apiIP string, certFile string, keyFile string) *Worker {
	route := &Worker{
		logger:    log.New("module", "Worker"),
		apiServer: nil,
		apiPort:   apiPort,
		apiIP:     apiIP,
		certFile:  certFile,
		keyFile:   keyFile,
	}
	route.logger.SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat()))
	return route
}
