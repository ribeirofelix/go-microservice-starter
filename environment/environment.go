package environment

import "github.com/urfave/negroni"

// Environment -
type Environment interface {
	GetHandler() (*negroni.Negroni, error)
}
