package log

import (
	"errors"
	"flag"
	"fmt"
)

// Config adds log-specific configuration fields to the
// common cfg.Registerable and cfg.Validatable interfaces
type Config struct {
	Level             string // log level (debug|info|warn|error)
	JsonFormat        bool   // use JSON log format (vs logfmt)
	StacktraceLevel   string // min level for stack traces (debug|info|warn|error)
	IncludeErrorLinks bool   // include source links in error logs
	MaxErrorLinks     int    // max error chain depth for links (1..64)
}

// RegisterFlags binds Config fields to the given FlagSet with defaults inline
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Level, "log-level", "info", "log level (debug|info|warn|error)")
	fs.BoolVar(&c.JsonFormat, "log-json", true, "JSON log format")
	fs.StringVar(&c.StacktraceLevel, "stacktrace-level", "error", "min level for stack traces (debug|info|warn|error)")
	fs.BoolVar(&c.IncludeErrorLinks, "include-error-links", true, "include source links in error logs")
	fs.IntVar(&c.MaxErrorLinks, "max-error-links", 5, "max error chain depth for links (1..64)")
}

func (c *Config) Validate() error {
	var errs []error

	// Log level
	if _, err := ParseLevel(c.Level); err != nil {
		errs = append(errs, fmt.Errorf("invalid LOG_LEVEL %q: %w", c.Level, err))
	}

	// Stacktrace level
	if _, err := ParseLevel(c.StacktraceLevel); err != nil {
		errs = append(errs, fmt.Errorf("invalid STACKTRACE_LEVEL %q: %w", c.StacktraceLevel, err))
	}

	// Error link limits
	if c.IncludeErrorLinks {
		if c.MaxErrorLinks < 1 || c.MaxErrorLinks > 64 {
			errs = append(errs, fmt.Errorf("MAX_ERROR_LINKS must be 1..64 (got %d)", c.MaxErrorLinks))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
