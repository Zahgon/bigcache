package main

import (
	"log"
	"net/http"
)

type service func(http.Handler) http.Handler

func serviceLoader(h http.Handler, svcs ...service) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func requestMetrics(l *log.Logger) service { _ = "STUB: not implemented"; return *new(service) }
