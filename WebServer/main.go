package main

import (
	flag "flag"
	"web-api-template/routing"
)

// main func
func main() {
	// Declare the API port.

	projectTitle := flag.String("title", "TemplateApplication", "This will contain the title of the client.")
	webPort := flag.Int("web-port", 443, "The Port is required to start the server.")
	webIP := flag.String("web-ip", "localhost", "The IP is required to start the server.")
	apiPort := flag.Int("api-port", 8080, "The port for the API.")
	apiIP := flag.String("api-ip", "localhost", "The IP is required to start the server.")
	certificate := flag.String("cert", "certs/certificate.crt", "Location of the certificate.")
	key := flag.String("key", "certs/private.key", "Location of th key.")

	// Make a bool to tell the thread when to stop.
	stop := make(chan bool)

	// Declare the new structure.
	c := routing.NewRouterStructure(*projectTitle, *webPort, *webIP, *apiPort, *apiIP, *certificate, *key)

	// Create the API server, declare endpoints etc.
	c.CreateWebServer()

	// Start listening for API requests.
	c.ListenWebServer(stop)

	<-stop

	// Stop the API.
	c.WebShutdown()
}
