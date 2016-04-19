package main

import (
	"log"
	"time"
	"net/http"
)

func logRequest(r *http.Request, t time.Duration) {
	log.Printf(
		"%-9s\t%-9s\t%s",
		r.Method,
		r.RequestURI,
		t,
	)
}