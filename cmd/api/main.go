package main

import (
	"flag"

	"github.com/glats/go-ms/pkg/api"
	"github.com/glats/go-ms/pkg/config"
)

func main() {
	cfgPath := flag.String("p", "./cmd/api/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
