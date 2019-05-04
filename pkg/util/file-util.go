package util

import (
	"os"
	"path/filepath"
	"github.com/qkgo/scaff/pkg/cfg"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		cfg.Log.Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
