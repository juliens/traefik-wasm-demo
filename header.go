package main

import (
	"encoding/json"
	"fmt"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

type Middleware struct {
	Config *Config
}

var mw = &Middleware{}

func init() {
	err := json.Unmarshal(handler.Host.GetConfig(), &mw.Config)
	if err != nil {
		handler.Host.Log(api.LogLevelError, fmt.Sprintf("Could not load config %v", err)) // TODO: how to panic? api does not have that?
	}
}

func main() {
	handler.HandleRequestFn = mw.handleRequest
	handler.HandleResponseFn = mw.handleResponse
}

// handleRequest implements a simple request middleware.
func (mw *Middleware) handleRequest(req api.Request, _ api.Response) (next bool, reqCtx uint32) {
	handler.Host.Log(api.LogLevelInfo, "hello from handleRequest")
	for k, v := range mw.Config.Headers {
		req.Headers().Add(k, v)
	}
	// proceed to the next handler on the host.
	next = true
	return
}

// handleResponse implements a simple response middleware.
func (mw *Middleware) handleResponse(reqCtx uint32, req api.Request, resp api.Response, _ bool) {
	handler.Host.Log(api.LogLevelInfo, "hello from handleResponse")
}
