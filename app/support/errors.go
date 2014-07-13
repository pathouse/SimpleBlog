package support

import (
	"github.com/golang/glog"
	"log"
	"runtime"
)

func LogStacktrace(err error) {
	if err != nil {
		var stack [4096]byte
		// Stack formats a stack trace of the calling goroutine into buf and returns the number of bytes written to buf.
		// If all is true, Stack formats stack traces of all other goroutines into buf after the trace for the current goroutine.
		runtime.Stack(stack[:], false)
		// %q = double quoted string safely escaped with Go syntax
		log.Printf("%q\n%s\n", err, stack[:])
		// conditional logging depending on verbosity - see https://google-glog.googlecode.com/svn/trunk/doc/glog.html
		// V reports whether verbosity at the call site is at least the requested level.
		glog.V(3).Infoln("%q\n%s\n", err)
	}
}
