package support

import (
	"log"
	"runtime"
)

func LogStacktrace(err error) {
	var stack [4096]byte
	// Stack formats a stack trace of the calling goroutine into buf and returns the number of bytes written to buf.
	// If all is true, Stack formats stack traces of all other goroutines into buf after the trace for the current goroutine.
	runtime.Stack(stack[:], false)
	// %q = double quoted string safely escaped with Go syntax
	log.Printf("%q\n%s\n", err, stack[:])
}
