package util

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joycastle/casual-server-lib/log"
)

//util.WorkingPath
func WorkingPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Get("error").Fatal("util.WorkingPath: ", err)
	}

	if strings.Contains(dir, os.TempDir()) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			return path.Dir(filename)
		}
	}
	return dir
}
