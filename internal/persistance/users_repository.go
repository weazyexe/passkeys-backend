package persistance

import (
	"github.com/go-webauthn/webauthn/webauthn"
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

func (r *UsersRepository) GetUserByID(userId string) (*models.User, error) {
	return db.GetUserByID(userId)
}

func (r *UsersRepository) GetUserByUsername(username string) (*models.User, error) {
	return db.GetUserByUsername(username)
}

func (r *UsersRepository) SaveCredentialForUser(userId string, credential webauthn.Credential) {
	db.SaveCredentialForUser(userId, credential)
}
