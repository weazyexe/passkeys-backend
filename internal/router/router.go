package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weazyexe/passkeys/internal/handlers"
)

func Setup(r *gin.Engine) {
	r.StaticFile("/.well-known/assetlinks.json", "./assets/assetlinks.json")

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register/request", handlers.RequestRegistrationHandler)
		authRouter.POST("/register/response", handlers.ResponseRegistrationHandler)
	}

}
