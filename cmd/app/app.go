package main

import (
	"github.com/weazyexe/passkeys/internal/app"
	"path/filepath"
)

var configPath = filepath.Join("config", "app.yaml")

func main() {
	app.Run(configPath)
}
