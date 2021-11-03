package system

import "os"

func GO111MODULE() bool {
	return os.Getenv("GO111MODULE") != ""
}
