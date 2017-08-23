package types

import "errors"

// Environment represents a Sensu environment in RBAC
type Environment struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Validate returns an error if the environment does not pass validation tests
func (e *Environment) Validate() error {
	if err := ValidateName(e.Name); err != nil {
		return errors.New("environment name " + err.Error())
	}

	return nil
}

// FixtureEnvironment returns a mocked environment
func FixtureEnvironment(name string) *Environment {
	return &Environment{
		Name: name,
	}
}