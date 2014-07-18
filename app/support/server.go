// once again thanks to git-go-websiteskeletong
// which was in turn inspired by https://gist.github.com/cespare/3985516
// this one has a lot more in common with the gist
package support

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ServerLogger struct {
	mux http.Handler
	out io.Writer
}

type serverEvent struct {
	// composed of http.ResponseWriter so that we can
	// intercept Write and WriteHeader for logging purposes
	http.ResponseWriter

	ip                    string
	time                  time.Time
	method, uri, protocol string
	status                int
	responseBytes         int64
	elapsedTime           time.Duration
}

func NewServerLogger(mux http.Handler, out io.Writer) http.Handler {
	return &ServerLogger{mux: mux, out: out}
}

func (s *ServerLogger) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. The HTTP server in this package
	// sets RemoteAddr to an "IP:port"
	clientIP := req.RemoteAddr
	if port := strings.LastIndex(clientIP, ":"); port != -1 {
		clientIP = clientIP[:port]
	}

	event := &serverEvent{
		ResponseWriter: resp,
		ip:             clientIP,
		time:           time.Time{},
		method:         req.Method,
		uri:            req.RequestURI,
		protocol:       req.Proto,
		status:         http.StatusOK,
		elapsedTime:    time.Duration(0),
	}

	startTime := time.Now()
	s.mux.ServeHTTP(event, req)
	finishTime := time.Now()

	event.time = finishTime.UTC()
	event.elapsedTime = finishTime.Sub(startTime)

	event.Log(s.out)
}

func (e *serverEvent) Write(p []byte) (int, error) {
	written, err := e.ResponseWriter.Write(p)
	// we interrupt this regularly scheduled response
	// to take note of how many bytes were written
	e.responseBytes += int64(written)
	return written, err
}

func (e *serverEvent) WriteHeader(status int) {
	// we interrupt this regularly scheduled header
	// to take note of the status and overwrite the
	// http.StatusOK placeholder we initialized the event with
	e.status = status
	e.ResponseWriter.WriteHeader(status)
}

func (e *serverEvent) Log(out io.Writer) {
	timestamp := e.time.Format("Jan _2 2006 03:04:05")
	requestLine := strings.Join([]string{e.method, e.uri, e.protocol}, " ")
	fmt.Fprintf(out, "%s - - [%s] \"%s\" %d %d (load time: %.4f)\n", e.ip, timestamp,
		requestLine, e.status, e.responseBytes, e.elapsedTime.Seconds())
}
