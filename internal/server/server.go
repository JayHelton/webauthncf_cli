package server

import (
	"fmt"
	"net/http"
)

func Start(port string) {
	http.HandleFunc("/attestation", attestation)
	http.HandleFunc("/assertion", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
