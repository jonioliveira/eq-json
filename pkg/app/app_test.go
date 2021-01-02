package app

import (
	"testing"
)

func TestApp(t *testing.T) {
	if Start() != nil {
		t.Error("Start failed")
	}
}
