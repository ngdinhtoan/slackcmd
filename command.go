package slackcmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Commander interface define function that a command must have
type Commander interface {
	// GetCommand return list of command can be serve by commander
	GetCommand() []string
	// Execute will receive Slack command payload when Slack send payload to server
	// when an error is returned, error message will be sent to user who type command
	Execute(payload *Payload, w http.ResponseWriter) error
	// ValidateToken will check if given token (sent by Slack) is same as registered token for command
	ValidateToken(token string) bool
}

var (
	regCommand   = map[string]Commander{}
	registerLock = sync.Mutex{}
)

// Register a commander which will process command
func Register(cmd Commander) {
	cmdNames := cmd.GetCommand()
	if len(cmdNames) == 0 {
		return
	}

	for _, cmdName := range cmdNames {
		if !strings.HasPrefix(cmdName, "/") {
			cmdName = "/" + cmdName
		}

		registerByName(cmdName, cmd)
	}
}

// RegisterByName a command with commander which will process command
func registerByName(name string, cmd Commander) {
	registerLock.Lock()
	defer registerLock.Unlock()

	if _, found := regCommand[name]; found {
		log.Fatalf("Command %q is registered already!", name)
	}

	regCommand[name] = cmd
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := newPayloadByForm(r.Form)

	cmd, found := regCommand[payload.Command]
	if !found {
		io.WriteString(w, fmt.Sprintf("Command %q was not found: %+v", payload.Command, payload))
		return
	}

	if cmd.ValidateToken(payload.Token) == false {
		io.WriteString(w, "Your command is invalid! Please contact your administrator for advice.")
		return
	}

	if err := cmd.Execute(payload, w); err != nil {
		io.WriteString(w, err.Error())
	}
}
