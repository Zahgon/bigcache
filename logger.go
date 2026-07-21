package bigcache

import (
	"log"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

var _ Logger = &log.Logger{}

func DefaultLogger() *log.Logger { _ = "STUB: not implemented"; return nil }

func newLogger(custom Logger) Logger { _ = "STUB: not implemented"; return *new(Logger) }
