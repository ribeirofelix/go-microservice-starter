package environment

import (
	loggerMw "github.com/hellomd/go-sdk/logger"
	"github.com/hellomd/go-sdk/requestid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type developmentEnv struct{}

// NewDevelopmentEnv -
func NewDevelopmentEnv() Environment {
	return &developmentEnv{}
}

// GetHandler -
func (env *developmentEnv) GetHandler() (*negroni.Negroni, error) {
	stack := negroni.New()

	stack.Use(negroni.NewRecovery())
	stack.Use(requestid.NewMiddleware())

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}
	stack.Use(loggerMw.NewMiddleware(logger))

	return stack, nil
}
