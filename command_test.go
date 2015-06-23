package slackcmd

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRegister
func TestRegister(t *testing.T) {
	assert := assert.New(t)

	helloCmd := &hello{}
	Register(helloCmd)

	cmd, found := regCommand["/command1"]
	assert.True(found)
	assert.Implements((*Commander)(nil), cmd)

	cmd, found = regCommand["/command2"]
	assert.True(found)
	assert.Implements((*Commander)(nil), cmd)
}

type hello struct{}

func (h *hello) GetCommand() []string { return []string{"/command1", "command2"} }
func (h *hello) Execute(_ *Payload, _ http.ResponseWriter) error {
	return errors.New("command executed")
}
func (h *hello) ValidateToken(token string) bool { return token == "token" }

// TestCommandHandler
func TestCommandHandler(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)

	var (
		req      *http.Request
		w        = &writer{}
		adminMsg = "Something goes wrong, please contact your administrator."
	)

	// case ParseForm() getErr
	req = &http.Request{}
	req.Method = "POST"

	Handler(w, req)
	assert.Equal(adminMsg, w.data)

	// case payload invalid
	var payload = url.Values{}
	req.PostForm = payload
	req.Form = payload

	Handler(w, req)
	assert.Equal(adminMsg, w.data)

	// case command not found
	payload.Set("command", "/abc")
	payload.Set("token", "token")
	payload.Set("channel_name", "general")
	payload.Set("channel_id", "T2029")

	Handler(w, req)
	assert.Equal(`Command "/abc" was not found`, w.data)

	payload.Set("command", "/command1")
	Handler(w, req)
	assert.Equal("command executed", w.data)

	// case wrong tocken
	payload.Set("token", "102932083213")
	Handler(w, req)
	assert.Equal(adminMsg, w.data)
}

type writer struct {
	data string
}

var _ http.ResponseWriter = (*writer)(nil)

func (w *writer) Header() http.Header { return http.Header{} }
func (w *writer) Write(data []byte) (int, error) {
	w.data = string(data)
	return 1, nil
}
func (w *writer) WriteHeader(_ int) {}
