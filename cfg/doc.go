// Package cfg provides interfaces and utilities for application configuration management.
// Allows each package to define its own config struct and registration logic while still
// calling them all from a central place without needing to know the details of each struct.
//
// The main interfaces are:
// - Registerable: for config structs that can register themselves to a FlagSet.
// - Validatable: for config structs that can validate themselves after flags are parsed.
package cfg
