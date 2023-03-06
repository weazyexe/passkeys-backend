package db

import (
	"errors"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/weazyexe/passkeys/internal/models"
)

var usersDb []models.User

func IsUserExist(username string) bool {
	for i := range usersDb {
		if usersDb[i].Username == username {
			return true
		}
	}
	return false
}

func CreateUser(user models.User) {
	usersDb = append(usersDb, user)
}

func GetUserByID(userId string) (*models.User, error) {
	for _, user := range usersDb {
		if user.ID == userId {
			return &user, nil
		}
	}
	return nil, errors.New("user does not exist")
}

func GetUserByUsername(username string) (*models.User, error) {
	for _, user := range usersDb {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user does not exist")
}

func SaveCredentialForUser(userId string, credential webauthn.Credential) {
	for _, user := range usersDb {
		if user.ID == userId {
			user.Credentials = append(user.Credentials, credential)
		}
	}
}
