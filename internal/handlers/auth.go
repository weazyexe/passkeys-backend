package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/weazyexe/passkeys/internal/models"
	"github.com/weazyexe/passkeys/internal/persistance"
	"github.com/weazyexe/passkeys/pkg/fido2"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestRegistrationRequestBody struct {
	Username string `json:"username"`
}

func RequestRegistrationHandler(c *gin.Context) {
	repository := persistance.GetUsersRepository()
	webauthn := fido2.GetWebAuthn()

	var body requestRegistrationRequestBody
	if err := c.BindJSON(&body); err != nil {
		errorMessage := "error while json parsing"
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	if repository.IsUserExists(body.Username) {
		errorMessage := fmt.Sprintf("user %s is already exist", body.Username)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	user := models.User{
		ID:       uuid.New().String(),
		Username: body.Username,
	}

	repository.CreateUser(user)
	options, _, err := webauthn.BeginRegistration(&user)
	if err != nil {
		errorMessage := fmt.Sprintf("error while generating webauthn credentials: %s", err)
		log.Print(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, options.Response)
}

func ResponseRegistrationHandler(c *gin.Context) {

}
