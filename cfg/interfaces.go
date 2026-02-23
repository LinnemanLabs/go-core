package cfg

import "flag"

// Registerable is for config structs that can register themselves to a FlagSet.
type Registerable interface {
	RegisterFlags(fs *flag.FlagSet)
}

// Validatable is for config structs that can validate themselves after flags are parsed.
type Validatable interface {
	Validate() error
}
