package db

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
)

var sessionsDb []webauthn.SessionData

func StoreSession(session webauthn.SessionData) {
	sessionsDb = append(sessionsDb, session)
}

func RemoveSession(session webauthn.SessionData) {
	for i, s := range sessionsDb {
		if bytes.Compare(s.UserID, session.UserID) == 0 {
			sessionsDb = append(sessionsDb[:i], sessionsDb[i+1:]...)
		}
	}
}

func GetSessionByUserID(id string) (*webauthn.SessionData, error) {
	for _, session := range sessionsDb {
		if string(session.UserID) == id {
			return &session, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("there is no session for user %s", id))
}

func GetSessionByChallenge(challenge string) (*webauthn.SessionData, error) {
	for _, session := range sessionsDb {
		if challenge == session.Challenge {
			return &session, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("there is no session for challenge %s", challenge))
}
