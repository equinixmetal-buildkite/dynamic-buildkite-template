package types

import (
	"strings"
)

// ArrayFlags can be used when you want to accept array of strings via command line flag
type ArrayFlags []string

// String method returns a comma concatenated string from the ArrayFlags
func (i *ArrayFlags) String() string {
	return strings.Join(*i, ",")
}

// Set appends value to the ArrayFlags
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
