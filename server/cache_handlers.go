package main

import (
	"net/http"
)

func cacheIndexHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func cacheClearHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func clearCache(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func getCacheHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func putCacheHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func deleteCacheHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
