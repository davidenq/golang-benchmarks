package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.ListenAndServe ("127.0.0.1:6060", nil) 
}