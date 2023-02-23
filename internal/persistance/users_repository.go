package persistance

import (
	"github.com/weazyexe/passkeys/internal/db"
	"github.com/weazyexe/passkeys/internal/models"
)

type UsersRepository struct {
}

var usersRepository *UsersRepository

func GetUsersRepository() *UsersRepository {
	if usersRepository == nil {
		usersRepository = &UsersRepository{}
	}
	return usersRepository
}

func (r *UsersRepository) IsUserExists(username string) bool {
	return db.IsUserExist(username)
}

func (r *UsersRepository) CreateUser(user models.User) {
	db.CreateUser(user)
}
