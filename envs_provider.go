package config

import (
	"os"
	"strconv"
)

type envsProvider struct{}

func (p envsProvider) Priority() int {
	return 200
}

func (p envsProvider) GetEntry(key string) Entry {
	value, exists := os.LookupEnv(key)
	if !exists {
		return NewEntryEmpty()
	}

	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return NewEntryInt(intValue)
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return NewEntryFloat(floatValue)
	}
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return NewEntryBool(boolValue)
	}
	return NewEntryString(value)
}

func WithEnvironmentVariables() ConfigurationOptionFunc {
	return func(c *ConfigurationOptions) {
		c.AddProvider(newEnvsProvider())
	}
}

func newEnvsProvider() *envsProvider {
	return &envsProvider{}
}
