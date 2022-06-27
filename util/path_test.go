package util

import (
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	dir := WorkingPath()
	if strings.Contains(dir, "T/go-build") {
		t.Fatal("WorkingPath getError")
	}
}
