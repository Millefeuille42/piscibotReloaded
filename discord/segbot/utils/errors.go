package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

// LogError Prints error + StackTrace to stderr if error
func LogError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err, string(debug.Stack()))
	}
}

// CheckError Panic if error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
