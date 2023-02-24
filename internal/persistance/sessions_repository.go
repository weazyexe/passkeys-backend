package persistance

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/weazyexe/passkeys/internal/db"
)

type SessionsRepository struct {
}

var sessionsRepository *SessionsRepository

func GetSessionsRepository() *SessionsRepository {
	if sessionsRepository == nil {
		sessionsRepository = &SessionsRepository{}
	}
	return sessionsRepository
}

func (r *SessionsRepository) StoreSession(session webauthn.SessionData) {
	db.StoreSession(session)
}

func (r *SessionsRepository) RemoveSession(session webauthn.SessionData) {
	db.RemoveSession(session)
}

func (r *SessionsRepository) GetSessionByUserID(id string) (*webauthn.SessionData, error) {
	return db.GetSessionByUserID(id)
}

func (r *SessionsRepository) GetSessionByChallenge(challenge string) (*webauthn.SessionData, error) {
	return db.GetSessionByChallenge(challenge)
}
