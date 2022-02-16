package system

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Exit(exitCode ...int) {
	fmt.Printf("exit stack: %+v \n", string(debug.Stack()))
	os.Exit(exitCode[0])
	return
}

func ExIt(i ...int) {
	stacks := debug.Stack()
	fmt.Printf("exit stack: %s \n", string(stacks))
	if GO111MODULE() {
		return
	}
	os.Exit(i[0])
	return
}
