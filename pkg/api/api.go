package api

import (
	"fmt"

	"github.com/glats/go-ms/pkg/config"
	"github.com/glats/go-ms/pkg/postgres"
)

type ContextKey string

func (c ContextKey) String() string {
	return "server " + string(c)
}

// Start method to start the api server
func Start(cfg *config.Configuration) error {
	_, err := postgres.New(cfg.DB.Psn, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}


	//log := log.New()

	fmt.Println(cfg.Server)
	fmt.Println(cfg.Test)
	return nil
}

