package db

import "github.com/weazyexe/passkeys/internal/models"

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
