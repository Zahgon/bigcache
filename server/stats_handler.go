package main

import (
	"net/http"
)

func statsIndexHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func getCacheStatsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
