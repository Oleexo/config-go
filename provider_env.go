package config

import (
	"os"
	"strconv"
)

type envProvider struct {
	prefix string
}

func (p envProvider) Precedence() int {
	return 200
}

func (p envProvider) GetEntry(key string) Entry {
	lookupKey := key
	if p.prefix != "" {
		lookupKey = p.prefix + key
	}

	value, exists := os.LookupEnv(lookupKey)
	if !exists {
		return Empty()
	}

	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return NewInt(intValue)
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return NewFloat(floatValue)
	}
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return NewBool(boolValue)
	}
	return NewString(value)
}

func WithEnvironmentVariables() Option {
	return WithEnvPrefix("")
}

func WithEnvPrefix(prefix string) Option {
	return func(c *optionSet) error {
		c.addProvider(newEnvProvider(prefix))
		return nil
	}
}

func newEnvProvider(prefix string) *envProvider {
	return &envProvider{prefix: prefix}
}
