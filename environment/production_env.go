package environment

import (
	raven "github.com/getsentry/raven-go"
	"github.com/hellomd/go-sdk/config"
	loggerMw "github.com/hellomd/go-sdk/logger"
	newrelicMw "github.com/hellomd/go-sdk/newrelic"
	"github.com/hellomd/go-sdk/recovery"
	"github.com/hellomd/go-sdk/requestid"
	logmatic "github.com/marcelloma/logmatic-go"
	newrelic "github.com/newrelic/go-agent"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type productionEnv struct{}

// NewProductionEnv -
func NewProductionEnv() Environment {
	return &productionEnv{}
}

// GetHandler -
func (env *productionEnv) GetHandler() (*negroni.Negroni, error) {
	stack := negroni.New()

	stack.Use(requestid.NewMiddleware())

	logger := logrus.New()
	logger.Formatter = &logmatic.JSONFormatter{}
	hook := logmatic.NewLogmaticHook(config.Get(LogmaticAPIKeyCfgKey))
	logger.Hooks.Add(hook)
	stack.Use(loggerMw.NewMiddleware(logger))

	ravenCli, err := raven.New(config.Get(SentryDSNCfgKey))
	if err != nil {
		return nil, err
	}
	stack.Use(recovery.NewMiddleware(ravenCli, logger))

	newRelicApp, err := newrelic.NewApplication(newrelic.NewConfig(config.Get(AppNameCfgKey), config.Get(NewRelicLicenseKeyCfgKey)))
	if err != nil {
		return nil, err
	}
	stack.Use(newrelicMw.NewMiddleware(newRelicApp))

	return stack, nil
}
