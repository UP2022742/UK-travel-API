package routing

import (
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
}

// NewRouterStructure Defines an instance of the worker structure.
func NewRouterStructure(title string, webPort int, webIP string, apiPort int, apiIP string) *Router {
	route := &Router{logger: log.New("module", "Router"), apiServer: nil, projectTitle: title, webPort: webPort, webIP: webIP, apiPort: apiPort, apiIP: apiIP}
	route.logger.SetHandler(log.StreamHandler(os.Stderr, log.TerminalFormat()))
	return route
}
