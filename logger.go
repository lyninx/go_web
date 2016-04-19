package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
)
const CLR_R = "\x1b[31;1m"
const CLR_W = "\x1b[37;1m"

func logRequest(r *http.Request, t time.Duration) {
	log.Printf(
		"%-9s\t%-9s\t%s",
		r.Method,
		r.RequestURI,
		t,
	)
}
func logError(r *http.Request, err error) {
	fmt.Print(CLR_R)
	log.Printf(
		"%-9s\t%-9s",
		r.Method,
		r.RequestURI,
	)
	fmt.Println(CLR_W, "[", CLR_R, "ERROR", CLR_W, "]",err)
	fmt.Print(CLR_W)
}