package system

import "os"

func DebugMode() bool {
	return os.Getenv("DEBUG") != ""
}
