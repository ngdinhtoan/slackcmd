package slackcmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Handler is a http.Handler function that can be added to http server to handle Slack Commands
func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		contactAdminMsg(w)
		return
	}

	payload := newPayloadByForm(r.Form)
	if !payload.IsValid() {
		contactAdminMsg(w)
		return
	}

	cmd, err := getCommander(payload.Command)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Command %q was not found", payload.Command))
		return
	}

	var validator PayloadValidator
	if _, ok := cmd.(PayloadValidator); ok {
		validator = cmd.(PayloadValidator)
	} else {
		validator = NewTokenValidator()
	}

	if err := validator.Validate(payload); err != nil {
		contactAdminMsg(w)
		return
	}

	if err := cmd.Execute(payload, w); err != nil {
		io.WriteString(w, err.Error())
	}
}

func contactAdminMsg(w io.Writer) {
	io.WriteString(w, "Something goes wrong, please contact your administrator.")
}

// StartServer start a server which will receive and process Slack command
func StartServer(host, port, path string) {
	if strings.TrimSpace(path) == "" {
		path = "/"
	}

	http.HandleFunc(path, Handler)

	address := host + ":" + port

	log.Printf("Server is listening at %s%s", address, path)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
