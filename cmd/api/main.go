package main

import (
	"flag"
	"os"

	"github.com/veep-provider/pkg/api"
	"github.com/veep-provider/pkg/utl/config"
)

func main() {
	os.Setenv("ENVIRONMENT_NAME", "dev")
	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
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
