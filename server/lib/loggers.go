package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var (
	// InfoLog ...info logger
	InfoLog *log.Logger

	// ErrLog ...err logger
	ErrLog *log.Logger
)

func init() {
	InfoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	ErrLog = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

// ServerError ...for internal server errors
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ErrLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError ...for client errors such as 400 errors
func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound ...404 errors
func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
