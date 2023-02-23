package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/weazyexe/passkeys/internal/router"
	"github.com/weazyexe/passkeys/pkg/fido2"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
}

func Run(configPath string) {
	config, err := ReadConfig(configPath)
	if err != nil {
		log.Fatalf("error while reading the config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading .env file")
	}

	fido2.Initialize(config.Webauthn.RPDisplayName, config.Webauthn.RPID, config.Webauthn.RPOrigins)

	r := gin.New()
	router.Setup(r)

	port := os.Getenv("PORT")
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error while trying to listen the port %s: %s", port, err)
	}
}
