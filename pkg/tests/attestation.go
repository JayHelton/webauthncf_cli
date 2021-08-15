package tests

type AuthenticatorSelection struct {
	requireResidentKey bool
	userVerification   string
}

type AttestationOptionsRequest struct {
	username               string
	displayName            string
	authenticatorSelection AuthenticatorSelection
	attestation            string
}
