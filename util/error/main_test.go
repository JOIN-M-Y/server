package error_test

import (
	"testing"

	"github.com/JOIN-M-Y/server/util/error"
)

// TestNew test new method in error package
func TestNew(t *testing.T) {
	instance := &error.Error{}
	if instance == nil {
		t.Error("Can not create error instance")
	}
}
