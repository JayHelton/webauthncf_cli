package test_cases

type AuthenticatorSelection struct {
	RequireResidentKey bool
	UserVerification   string
}

type AttestationOptionsResponse struct {
	Username               string
	DisplayName            string
	AuthenticatorSelection AuthenticatorSelection
	Status                 string `json:"status,omitempty"`
}

type AttestationRequestAndResponse struct {
	Response AttestationOptionsResponse
}
