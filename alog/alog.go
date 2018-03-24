package alog

import (
	"log"
	"os"
)

// AccessLog is wrapper above log.Logger
type AccessLog struct {
	Logger *log.Logger
}

// New returns new AccessLogger without date
func New() *AccessLog {
	ac := new(AccessLog)
	ac.Logger = log.New(os.Stderr, "", log.Ltime)
	return ac
}
