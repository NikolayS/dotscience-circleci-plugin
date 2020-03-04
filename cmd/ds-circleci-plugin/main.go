package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/dotmesh-io/dotscience-circleci-plugin/pkg/client"
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

	circleCIClient := client.New(log, conf, http.DefaultClient)

	log.With(
		zap.String("host", conf.Host),
		zap.String("username", conf.Username),
		zap.String("project", conf.Project),
		zap.String("vcs_type", conf.VCSType),
		zap.String("revision", conf.Revision),
		zap.String("tag", conf.Tag),
		zap.String("branch", conf.Branch),
	).Info("client initialized, triggering job")

	err := circleCIClient.TriggerNewJob()
	if err != nil {
		log.With(zap.Error(err)).Fatal("failed to trigger CircleCI job")
		os.Exit(1)
	}

	log.Info("job triggered, exiting.")
	// done
}
