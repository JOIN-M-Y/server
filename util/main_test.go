package util_test

import (
	"testing"

	"github.com/JOIN-M-Y/server/util"
)

// TestInitialize test InitizlizeUtil method
func TestInitializeUtil(t *testing.T) {
	instance := util.Initialize()
	if instance == nil {
		t.Error("Can not create util instance")
	}
}
