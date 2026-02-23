package otelx

import (
	"errors"
	"flag"
	"fmt"
	"net"
)

// Config adds otelx-specific configuration fields to the
// common cfg.Registerable and cfg.Validatable interfaces
type Config struct {
	EnableTracing bool
	OTLPEndpoint  string
	TraceSample   float64
}

// RegisterFlags binds Config fields to the given FlagSet with defaults inline
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.EnableTracing, "enable-tracing", false, "Enable OTLP tracing and push to otlp-endpoint")
	fs.Float64Var(&c.TraceSample, "trace-sample", 0.0, "trace sampling ratio (0..1)")
	fs.StringVar(&c.OTLPEndpoint, "otlp-endpoint", "", "OTLP endpoint to push to (gRPC) (host:port)")
}

func (c *Config) Validate() error {
	var errs []error

	// Tracing sample
	if c.TraceSample < 0 || c.TraceSample > 1 {
		errs = append(errs, fmt.Errorf("invalid TRACE_SAMPLE %.3f (must be 0..1)", c.TraceSample))
	}

	// OTLP tracing (grpc exporter wants host:port, no scheme)
	if c.EnableTracing {
		if c.OTLPEndpoint == "" {
			errs = append(errs, fmt.Errorf("OTLP_ENDPOINT required when ENABLE_TRACING=true"))
		} else if _, _, err := net.SplitHostPort(c.OTLPEndpoint); err != nil {
			errs = append(errs, fmt.Errorf("OTLP_ENDPOINT must be host:port (got %q): %w", c.OTLPEndpoint, err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
