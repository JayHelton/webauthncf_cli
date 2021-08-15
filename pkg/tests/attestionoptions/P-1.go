package attestationoptions

import (
	"fmt"
	"github.com/jayhelton/webauthncf/pkg/tests"
)

func P1TestSuite() tests.ConformanceTestSuite {
	return tests.ConformanceTestSuite{
		Name: "P-1",
		Description: `
/**
 * P-1
 * Get ServerPublicKeyCredentialCreationOptionsResponse:
 */
		`,
		Run: func(data interface{}) []tests.Result {
			results := []tests.Result{}
			tests := []tests.ConformanceTest{
				p1A(),
				p1B(),
				p1C(),
				p1D(),
				p1E(),
				p1F(),
				p1G(),
				p1H(),
			}
			for _, test := range tests {
				fmt.Println("Running...")
				fmt.Printf("Name : %s", test.Name)
				fmt.Printf("Description : %s", test.Description)
				result := test.Run(data)
				fmt.Printf("Result : %+v", result)
				results = append(results, result)
			}
			return results
		},
	}
}

func p1A() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.a",
		Description: `
* (a) response MUST contain "status" field, and it MUST be of type DOMString and set to "ok"
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1B() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.b",
		Description: `
 *     (b) response MUST contain "errorMessage" field, and it MUST be of type DOMString and set to
 *         an empty string
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1C() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.c",
		Description: `
 *     (c) response contains "user" field, of type Object and:
 *         (1) check that user.name is not missing, and is of type DOMString
 *         (2) check that user.displayName is not missing, and is of type DOMString
 *         (3) check that user.id is not missing, and is of type DOMString, and is not empty. It
 *             MUST be base64url encoded byte sequence, and is not longer than 64 bytes.
 *         (4) If user.icon is presented, check that it's is of type DOMString
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1D() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.d",
		Description: `
 *     (d) response contains "rp" field, of type Object and:
 *         (1) check that rp.name is not missing, and is of type DOMString
 *         (2) check that rp.id is not missing, and is of type DOMString.
 *         (3) If rp.icon is presented, check that it's is of type DOMString
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1E() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.e",
		Description: `
 *     (e) response contains "challenge" field, of type String, base64url encoded and not less than
 *         16 bytes.
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1F() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.f",
		Description: `
 *     (f) response contains "pubKeyCredParams" field, of type Array and:
 *         (1) each member MUST be of type Object
 *         (2) each member MUST contain "type" field of type DOMString
 *         (3) each member MUST contain "alg" field of type Number
 *         (4) MUST contain one member with type set to "public-key" and alg set to an algorithm
 *             that is supported by the authenticator
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}
func p1G() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.g",
		Description: `
 *     (g) If response contains "timeout" field, check that it's of type Number and is bigger than 0
		`, Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}

func p1H() tests.ConformanceTest {
	return tests.ConformanceTest{
		Name: "P-1.h",
		Description: `
 *     (h) response contains "extensions" field, with "example.extension" key presented
		`,
		Run: func(interface{}) tests.Result {
			return tests.Result{}
		},
	}
}
