package handlers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/google/uuid"
	"github.com/weazyexe/passkeys/internal/models"
	"github.com/weazyexe/passkeys/internal/persistance"
	"github.com/weazyexe/passkeys/pkg/fido2"
	"log"
	"net/http"
)

type requestRegistrationRequestBody struct {
	Username string `json:"username"`
}

func RequestRegistrationHandler(c *gin.Context) {
	usersRepository := persistance.GetUsersRepository()
	sessionsRepository := persistance.GetSessionsRepository()
	webauthn := fido2.GetWebAuthn()

	var body requestRegistrationRequestBody
	if err := c.BindJSON(&body); err != nil {
		errorMessage := "error while request json parsing"
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	user, err := usersRepository.GetUser(body.Username)
	if err != nil {
		user = &models.User{
			ID:       uuid.New().String(),
			Username: body.Username,
		}
		usersRepository.CreateUser(*user)
	}

	options, session, err := webauthn.BeginRegistration(user)
	if err != nil {
		errorMessage := fmt.Sprintf("error while generating webauthn credentials: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	sessionsRepository.StoreSession(*session)
	c.IndentedJSON(http.StatusOK, options.Response)
}

func ResponseRegistrationHandler(c *gin.Context) {
	usersRepository := persistance.GetUsersRepository()
	sessionsRepository := persistance.GetSessionsRepository()
	webauthn := fido2.GetWebAuthn()

	response, err := protocol.ParseCredentialCreationResponse(c.Request)
	if err != nil {
		errorMessage := "error while request json parsing"
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	session, err := sessionsRepository.GetSessionByChallenge(response.Response.CollectedClientData.Challenge)
	if err != nil {
		errorMessage := fmt.Sprintf("error while finding the registration session: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	decodedId, err := base64.URLEncoding.DecodeString(string(session.UserID))
	if err != nil {
		errorMessage := fmt.Sprintf("can't decode user ID from base64 in session: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	user, err := usersRepository.GetUser(string(decodedId))
	if err != nil {
		errorMessage := fmt.Sprintf("error while finding the user by session: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	credential, err := webauthn.CreateCredential(user, *session, response)
	if err != nil {
		errorMessage := fmt.Sprintf("error while creating credentials: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	usersRepository.SaveCredentialForUser(user.ID, *credential)

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}
