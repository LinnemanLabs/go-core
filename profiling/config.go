package profiling

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
)

// Config adds prof-specific configuration fields to the
// common cfg.Registerable and cfg.Validatable interfaces
type Config struct {
	EnablePyroscope bool
	PyroServer      string
	PyroTenantID    string
}

// RegisterFlags binds Config fields to the given FlagSet with defaults inline
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.EnablePyroscope, "enable-pyroscope", false, "Enable pushing Pyroscope data to server set in -pyro-server")
	fs.StringVar(&c.PyroServer, "pyro-server", "", "pyroscope server url to push to")
	fs.StringVar(&c.PyroTenantID, "pyro-tenant", "", "tenant (x-scope-orgid) to use for pyro-server")
}

func (c *Config) Validate() error {
	var errs []error

	// Pyroscope (URL and scheme)
	if c.EnablePyroscope {
		if c.PyroServer == "" {
			errs = append(errs, fmt.Errorf("PYRO_SERVER required when ENABLE_PYROSCOPE=true"))
		} else if u, err := url.Parse(c.PyroServer); err != nil || u.Scheme == "" || u.Host == "" {
			errs = append(errs, fmt.Errorf("PYRO_SERVER must be a URL (got %q)", c.PyroServer))
		}
	}

	// Pyroscope tenant
	if c.EnablePyroscope && c.PyroTenantID == "" {
		errs = append(errs, fmt.Errorf("PYRO_TENANT required when ENABLE_PYROSCOPE=true"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
