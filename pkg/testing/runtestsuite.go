package testing

import (
	"fmt"

	"github.com/jayhelton/webauthncf/pkg/test_cases"
	attestationoptions "github.com/jayhelton/webauthncf/pkg/test_cases/attestionoptions"
)

func RunAttestationTestSuites(data test_cases.AttestationRequestAndResponse) []*test_cases.Result {
	results := []*test_cases.Result{}
	testSuites := []test_cases.ConformanceTestSuite{
		attestationoptions.P1TestSuite(),
	}

	for _, suite := range testSuites {
		fmt.Println("Running...")
		fmt.Printf("Name : %s", suite.Name)
		fmt.Printf("\nDescription : %s", suite.Description)
		result := suite.Execute(data)
		results = append(results, result...)
	}

	return results
}
