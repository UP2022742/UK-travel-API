package main

import (
	"flag"
	"web-app-template/worker"
)

func main() {

	// Declare the API port.
	apiPort := flag.Int("api-port", 8080, "The port for the API.")
	apiIP := flag.String("api-ip", "localhost", "The IP is required to start the server.")

	// Make a bool to tell the thread when to stop.
	stop := make(chan bool)

	// Declare the new structure.
	c := worker.NewWorkerStructure(*apiPort, *apiIP)

	// Create the API server, declare endpoints etc.
	c.CreateAPIServer()

	// Start listening for API requests.
	c.ListenAPIServer(stop)

	<-stop

	// Stop the API.
	c.APIShutdown()
}
