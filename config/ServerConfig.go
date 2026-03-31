// Package config — ServerConfig is the configuration device for the Echo Hello World template.
//
// Place this device on the canvas, set Port, Message and Module Name
// in the Inspect panel, then click Templates → Echo Hello World → Generate ZIP.
// The downloaded ZIP is a ready-to-run Go project.
package config

import (
	"fmt"
)

// ServerConfig holds the top-level configuration for the Hello World server.
//
// icon:server. label:Server Config.
type ServerConfig struct {
	// Port is the HTTP listen port.
	Port string `prop:"Port" default:"8081"`

	// Message is the greeting shown on the index page.
	Message string `prop:"Message" default:"Hello, World!"`

	// ModuleName is the Go module name written into go.mod.
	ModuleName string `prop:"Module Name" default:"github.com/example/hello"`
}

// Init validates and stores the port.
//
// icon:plug. label:Init.
//
// Params
//
//	port: HTTP listen port, receives wire from a ConstInt device. connection:mandatory. unit:port.
//
// Returns
//
//	err: non-nil if port is out of valid range. connection:optional.
func (s *ServerConfig) Init(port int) (err error) {
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port: %d", port)
	}
	s.Port = fmt.Sprintf("%d", port)
	return nil
}
