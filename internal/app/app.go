package app

import (
	"fmt"
	"github.com/weazyexe/passkeys/internal/router"
	"github.com/weazyexe/passkeys/pkg/fido2"
	"log"

	"github.com/gin-gonic/gin"
)

type Config struct {
}

func Run(configPath string) {
	config, err := ReadConfig(configPath)
	if err != nil {
		log.Fatalf("error while reading the config: %s", err)
	}

	fido2.Initialize(config.Webauthn.RPDisplayName, config.Webauthn.RPID, config.Webauthn.RPOrigins)

	r := gin.New()
	router.Setup(r)

	if err := r.Run(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		log.Fatalf("error while trying to listen the port %s: %s", config.Server.Port, err)
	}
}
