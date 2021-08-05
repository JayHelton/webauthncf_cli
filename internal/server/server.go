package server

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func Start(port string) {
	http.HandleFunc("/attestation", hello)
	http.HandleFunc("/assertion", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
