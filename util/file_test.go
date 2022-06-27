package util

import (
	"testing"
)

func Test_IsDir(t *testing.T) {
	path := "/Users/mac2022/joycastle/code/casual-server-lib/util"
	if !IsDir(path) {
		t.Fatal("not a dir")
	}
	path = "./file.go"
	if IsDir(path) {
		t.Fatal("is a dir")
	}
}
