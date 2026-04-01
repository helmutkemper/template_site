// Package devices — EchoServer is the configuration device for the
// Echo Hello World template.
//
// Place this device on the canvas.
// Wire a ConstInt block to the port connector to set the HTTP listen port.
// Wire a ConstString block to the message connector to set the greeting text.
// Set Module Name in the Inspect panel.
// Click Generate ZIP to download a ready-to-run Go project.
package devices

import (
	"fmt"
)

// EchoServer holds the runtime configuration for the Echo Hello World server.
//
// icon:server. label:Echo Server.
type EchoServer struct {
	// port is the HTTP listen port, received via wire from a ConstInt block.
	// Not a prop — the value comes exclusively from the wire connection.
	port int

	// message is the greeting text, received via wire from a ConstString block.
	// Not a prop — the value comes exclusively from the wire connection.
	message string

	// ModuleName is the Go module name written into go.mod.
	// Set in the Inspect panel; no wire needed.
	ModuleName string `prop:"Module Name" default:"github.com/example/hello"`
}

// Init receives the server configuration from wires and validates the port.
//
// executionOrder:1. icon:server. label:Init.
//
// Params
//
//	port:    HTTP listen port — wire from a ConstInt block.     connection:mandatory.  unit:port.  range:1..65535.
//	message: greeting shown on the index page — wire from a ConstString block.  connection:mandatory.
//
// Returns
//
//	err: non-nil when port is outside 1–65535.  connection:optional.
func (s *EchoServer) Init(port int, message string) (err error) {
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port %d: must be between 1 and 65535", port)
	}
	s.port = port
	s.message = message
	return nil
}
