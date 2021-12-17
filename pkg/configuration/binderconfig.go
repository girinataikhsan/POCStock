package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type ServicesApp struct {
	EnvVariable string
	Config      ConfigurationApp
	Path        string
}

func (s *ServicesApp) Load() {
	if err := envconfig.Process(s.EnvVariable, &s.Config); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("Loaded configs: ", s.Config)
	log.Info("Root Path: ", s.Path)
}

func (s *ServicesApp) GetRoot() string {
	return s.Config.RootURL
}

func (s *ServicesApp) GetVersion() string {
	return s.Config.Version
}

func (s *ServicesApp) GetAppName() string {
	return s.Config.AppName
}
