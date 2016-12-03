package config

import (
	"log"
	"reflect"
	"strings"

	"github.com/caarlos0/env"
)

var config *Config

// This gets the config that we've already loaded.
func GetConfig() *Config {
	return config
}

type Config struct {
	Port          int    `env:"PORT"`
	Auth0ClientID string `env:"AUTH0_CLIENT_ID"`
}

func ParseConfig() (*Config, error) {
	// Parse the config or return the error.
	c := new(Config)
	err := env.Parse(c)
	if err != nil {
		return nil, err
	}

	// If we encountered no errors, still make sure that every variable was
	// defined.
	missingVars := c.getMissingVars()
	if len(missingVars) > 0 {
		log.Fatalf("Missing the following environment variables: %s", strings.Join(missingVars, ", "))
	}

	// At this point, we should have a healthy, complete config! Store it for
	// other accessors to use.
	config = c

	return c, nil
}

func (c *Config) getMissingVars() []string {
	missingVars := []string{}
	configReflectValue := reflect.Indirect(reflect.ValueOf(c))
	numFields := configReflectValue.NumField()
	for i := 0; i < numFields; i++ {
		// If this entry is `nil` or the zero-value of the type, log an error.
		if isZeroOfUnderlyingType(configReflectValue.Field(i).Interface()) {
			envVarName := reflect.TypeOf(c).Elem().Field(i).Tag.Get("env")
			missingVars = append(missingVars, envVarName)
		}
	}
	return missingVars
}

// Is this value nil or the zero-value of the concrete type?
// http://stackoverflow.com/a/13906031/1473205
func isZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
