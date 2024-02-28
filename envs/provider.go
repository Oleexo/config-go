package envs

import (
	"github.com/Oleexo/config-go"
	"os"
	"strconv"
)

type Provider struct {
}

func (p Provider) Priority() int {
	return 200
}

func (p Provider) GetEntry(key string) config.Entry {
	value, exists := os.LookupEnv(key)
	if !exists {
		return config.NewEntryEmpty()
	}

	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return config.NewEntryInt(intValue)
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return config.NewEntryFloat(floatValue)
	}
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return config.NewEntryBool(boolValue)
	}
	return config.NewEntryString(value)
}

func WithEnvironmentVariables() config.ConfigurationOptionFunc {
	return func(c *config.ConfigurationOptions) {
		c.AddProvider(NewProvider())
	}
}

func NewProvider() *Provider {
	return &Provider{}
}
