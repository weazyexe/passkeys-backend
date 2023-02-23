package fido2

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"log"
)

var webAuthn *webauthn.WebAuthn

func Initialize(rpDisplayName, rpId string, rpOrigins []string) {
	webauthnConfig := &webauthn.Config{
		RPDisplayName: rpDisplayName,
		RPID:          rpId,
		RPOrigins:     rpOrigins,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			UserVerification:   protocol.VerificationRequired,
			ResidentKey:        protocol.ResidentKeyRequirementRequired,
			RequireResidentKey: protocol.ResidentKeyRequired(),
		},
	}

	w, err := webauthn.New(webauthnConfig)
	if err != nil {
		log.Fatal("error while initializing webauthn entity")
	}
	webAuthn = w
}

func GetWebAuthn() *webauthn.WebAuthn {
	return webAuthn
}
