package api

import (
	"fmt"

	"github.com/glats/go-ms/pkg/config"
)

// Start method to start the api server
func Start(cfg *config.Configuration) error {
	fmt.Println(cfg.Server)
	fmt.Println(cfg.Test)
	return nil
}
