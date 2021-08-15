package testing

import (
	"fmt"
	"github.com/jayhelton/webauthncf/pkg/tests"
	"github.com/jayhelton/webauthncf/pkg/tests/attestionoptions"
)

func RunAttestationTestSuites(data interface{}) []tests.Result {
	results := []tests.Result{}
	testSuites := []tests.ConformanceTestSuite{
		attestationoptions.P1TestSuite(),
	}

	for _, suite := range testSuites {
		fmt.Println("Running...")
		fmt.Printf("Name : %s", suite.Name)
		fmt.Printf("Description : %s", suite.Description)
		result := suite.Run(data)
		results = append(results, result...)
	}

	return results
}
