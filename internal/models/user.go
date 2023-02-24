package models

import "github.com/go-webauthn/webauthn/webauthn"

type User struct {
	webauthn.User
	ID          string
	Username    string
	Credentials []webauthn.Credential
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u *User) WebAuthnName() string {
	return u.Username
}

func (u *User) WebAuthnDisplayName() string {
	return u.Username
}

func (u *User) WebAuthnIcon() string {
	return ""
}
