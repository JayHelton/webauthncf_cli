package server

import (
	"encoding/json"
	"fmt"
	"github.com/jayhelton/webauthncf/pkg/testing"
	"github.com/jayhelton/webauthncf/pkg/tests"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

type attestationBody struct {
	options  tests.AttestationOptionsRequest
	response interface{}
}

func attestation(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/attestation" || req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(req.Body)

	var a attestationBody
	err := decoder.Decode(&a)

	if err != nil {
		panic(err)
	}

	result := testing.RunAttestationTestSuites(a)

	json, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
