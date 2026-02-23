package httpserver

import (
	"errors"
	"flag"
	"fmt"
)

// Config adds httpserver-specific configuration fields to the
// common cfg.Registerable and cfg.Validatable interfaces
type Config struct {
	Port int
}

// RegisterFlags binds Config fields to the given FlagSet with defaults inline
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.IntVar(&c.Port, "http-port", 8080, "listen TCP port (1..65535)")
}

func (c *Config) Validate() error {
	var errs []error

	// Ports
	if c.Port < 1 || c.Port > 65535 {
		errs = append(errs, fmt.Errorf("invalid HTTP_PORT %d (must be 1..65535)", c.Port))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
