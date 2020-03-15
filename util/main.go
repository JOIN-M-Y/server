package util

import (
	"github.com/JOIN-M-Y/server/util/error"
)

// Interface Utile interface
type Interface interface {
}

// Util provide utilities
type Util struct {
	Error *error.Error
}

// Initialize initialize utilities
func Initialize() *Util {
	error := error.New()
	return &Util{Error: error}
}
