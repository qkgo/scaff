package errs

import (
	"fmt"
	"runtime/debug"
)

func HandleError(err error) {
	fmt.Printf("%v from %s", err, string(debug.Stack()))
}

func Handle(err error) {
	fmt.Printf("%v from %s", err, string(debug.Stack()))
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("%v from %s", err, string(debug.Stack()))
	}
}
