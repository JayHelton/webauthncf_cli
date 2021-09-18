package attestationoptions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jayhelton/webauthncf/pkg/test_cases"
)

func TestP1AMustBeOkError(t *testing.T) {
	conformanceTest := p1A()
	str := `{"options": {}, "response":{"username": "test",  "displayName": "test", "authenticatorSelection": {}, "attestation": "" } }`
	res := test_cases.AttestationRequestAndResponse{}
	json.Unmarshal([]byte(str), &res)
	result := conformanceTest.Execute(res)
	if assert.NotNil(t, result) {
		assert.Equal(t, result.Message, "Response.status MUST be set to 'ok'", "Response message is incorrect")
		assert.False(t, result.Passed, "assertions passed but should have failed")
	}
}
