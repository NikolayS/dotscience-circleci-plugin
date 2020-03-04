package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/dotmesh-io/dotscience-circleci-plugin/pkg/config"
	"github.com/dotmesh-io/dotscience-circleci-plugin/pkg/logger"
)

func main() {

	log := logger.GetLoggerInstance(zap.InfoLevel)

	conf := config.MustLoad()

	if conf.Token == "" {
		log.Fatal("TOKEN not set, cannot continue")
		os.Exit(1)
	}
}
