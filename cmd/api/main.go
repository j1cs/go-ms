package main

import (
	"path/filepath"
	"runtime"

	"github.com/glats/go-ms/pkg/api"
	"github.com/glats/go-ms/pkg/config"
)

//TODO check my other projects on the asus.
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	cfg, err := config.Load(basepath, "config", "yaml")
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
