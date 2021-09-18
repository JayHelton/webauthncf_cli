package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jayhelton/webauthncf/pkg/test_cases"
	"github.com/jayhelton/webauthncf/pkg/testing"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func registration(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/registration" || req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(req.Body)

	var a test_cases.AttestationRequestAndResponse
	err := decoder.Decode(&a)

	if err != nil {
		panic(err)
	}

	result := []*test_cases.Result{}
	testResponses := testing.RunAttestationTestSuites(a)

	verbose := req.URL.Query().Get("verbose")

	if verbose != "true" {
		for _, test := range testResponses {
			if !test.Passed {
				result = append(result, test)
			}
		}
	} else {
		result = testResponses
	}

	json, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
