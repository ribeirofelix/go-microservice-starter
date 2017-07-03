package environment

import (
	"github.com/hellomd/go-sdk/requestid"
	"github.com/urfave/negroni"
)

type testEnv struct{}

// NewTestEnv returns a pointer to the test environment
func NewTestEnv() Environment {
	return &testEnv{}
}

// GetHandler bootstraps a web middleware with necovery, logging and DB session
func (env *testEnv) GetHandler() (*negroni.Negroni, error) {
	stack := negroni.New()

	stack.Use(negroni.NewRecovery())
	stack.Use(requestid.NewMiddleware())

	return stack, nil
}
