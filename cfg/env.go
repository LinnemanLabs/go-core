package cfg

import (
	"flag"
	"os"
	"strings"
)

// FillFromEnv sets any flag not explicitly passed on the CLI from
// environment variables. Flag "foo-bar" maps to PREFIX_FOO_BAR.
// Precedence: cli flag > env var > default.
func FillFromEnv(fs *flag.FlagSet, prefix string, logf func(string, ...any)) {
	explicit := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { explicit[f.Name] = true })

	fs.VisitAll(func(f *flag.Flag) {
		key := prefix + strings.ReplaceAll(strings.ToUpper(f.Name), "-", "_")
		envVal, envSet := os.LookupEnv(key)
		if !envSet {
			return
		}
		if explicit[f.Name] {
			if logf != nil {
				logf("flag -%s: cli value %q overrides env %s=%q", f.Name, f.Value.String(), key, envVal)
			}
			return
		}
		prev := f.Value.String()
		if err := fs.Set(f.Name, envVal); err != nil {
			_ = fs.Set(f.Name, prev) // restore; prev was already valid
			if logf != nil {
				logf("flag -%s: ignoring invalid env %s=%q: %v", f.Name, key, envVal, err)
			}
		}
	})
}
