package httpmw

import (
	"errors"
	"flag"
	"fmt"
)

// Config adds httpmw-specific configuration fields to the
// common cfg.Registerable and cfg.Validatable interfaces
type Config struct {
	TrustedProxyHops int
}

// RegisterFlags binds Config fields to the given FlagSet with defaults inline
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.IntVar(&c.TrustedProxyHops, "trusted-proxy-hops", 1, "number of trusted reverse proxies (0=direct, 1=ALB, 2=CDN+ALB, etc.)")
}

func (c *Config) Validate() error {
	var errs []error

	// Trusted proxy hops
	if c.TrustedProxyHops < 0 || c.TrustedProxyHops > 10 {
		errs = append(errs, fmt.Errorf("invalid TRUSTED_PROXY_HOPS %d (must be 0..10)", c.TrustedProxyHops))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
