package slackcmd

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	regCommand     = map[string]Commander{}
	regMutex       = &sync.Mutex{}
	errCmdNotFound = errors.New("command not found")
)

// Commander interface define function that a command must have;
// If your commander also implement PayloadValidator interface,
// then Handler will use it instead of default one (token validator).
type Commander interface {
	// GetCommand return list of command can be serve by commander
	GetCommand() []string
	// Execute will receive Slack command payload when Slack send payload to server
	// when an error is returned, error message will be sent to user who type command
	Execute(payload *Payload, w http.ResponseWriter) error
}

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

// registerByName a command with commander which will process command
func registerByName(name string, cmd Commander) {
	regMutex.Lock()
	defer regMutex.Unlock()

	if _, found := regCommand[name]; found {
		log.Fatalf("Command %q is registered already!", name)
	}

	regCommand[name] = cmd
}

// getCommander return registered commander by name, if not found then return error.
func getCommander(name string) (cmd Commander, err error) {
	regMutex.Lock()
	defer regMutex.Unlock()

	if cmd, found := regCommand[name]; found {
		return cmd, nil
	}

	return nil, errCmdNotFound
}
