package routing

import (
	"crypto/tls"
	"net/http"
	"os"

	log "github.com/inconshreveable/log15"
)

// Router is the struct that organises the Clinic Service.
type Router struct {
	logger       log.Logger
	apiServer    *http.Server
	projectTitle string
	webPort      int
	webIP        string
	apiPort      int
	apiIP        string
	cert         tls.Certificate
}

// NewRouterStructure Defines an instance of the worker structure.
func NewRouterStructure(title string, webPort int, webIP string, apiPort int, apiIP string, certFile string, keyFile string) *Router {
	cert, _ := tls.LoadX509KeyPair("certs/certificate.crt", "certs/private.key")
	route := &Router{
		logger:       log.New("module", "Router"),
		apiServer:    nil,
		projectTitle: title,
		webPort:      webPort,
		webIP:        webIP,
		apiPort:      apiPort,
		apiIP:        apiIP,
		cert:         cert,
	}
	route.logger.SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat()))
	return route
}
