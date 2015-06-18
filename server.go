package slackcmd

import (
	"log"
	"net/http"
	"strings"
)

// StartServer start server which will receive and process Slack command
func StartServer(host, port, path string) {
	if strings.TrimSpace(path) == "" {
		path = "/"
	}

	http.HandleFunc(path, Handler)

	address := host + ":" + port

	log.Printf("Server is listening at %s%s...", address, path)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
