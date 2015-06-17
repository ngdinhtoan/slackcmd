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

	http.HandleFunc(path, commandHandler)

	address := host + ":" + port
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
