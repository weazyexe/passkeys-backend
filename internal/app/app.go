package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/weazyexe/passkeys/internal/router"
	"github.com/weazyexe/passkeys/pkg/fido2"
	"log"
	"os"
	"strconv"

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
		log.Print("error while loading .env file")
	}

	fido2.Initialize(config.Webauthn.RPDisplayName, config.Webauthn.RPID, config.Webauthn.RPOrigins)

	r := gin.New()
	router.Setup(r)

	port := os.Getenv("PORT")
	shouldUseTls, err := strconv.ParseBool(os.Getenv("SELF_SIGNED_CERTIFICATES"))
	if err != nil {
		log.Print("there is no environment variable SELF_SIGNED_CERTIFICATES")
	}

	if shouldUseTls {
		if err := r.RunTLS(fmt.Sprintf(":%s", port), config.Ssl.TlsCertificatePath, config.Ssl.TlsKeyPath); err != nil {
			log.Fatalf("error while trying to listen the port %s: %s", port, err)
		}
	} else {
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("error while trying to listen the port %s: %s", port, err)
		}
	}
}
