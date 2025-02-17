package handlers

import (
	"flag"
)

// NewFlagSet creates and returns a new flag set
func NewFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("", flag.ExitOnError)
}
