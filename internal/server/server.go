package server

import (
	"fmt"
	"net/http"
)

func Start(port string) {
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/authentication", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
