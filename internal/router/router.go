package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weazyexe/passkeys/internal/handlers"
)

func Setup(r *gin.Engine) {
	r.POST("/auth/register/request", handlers.RequestRegistrationHandler)
	r.POST("/auth/register/response", handlers.ResponseRegistrationHandler)
}
