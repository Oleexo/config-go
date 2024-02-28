package config

import (
	"fmt"
	"sort"
)

type Configuration interface {
	// GetString returns the value associated with the key as a string.
	// If the key does not exist, panic.
	GetString(key string) string
	// GetStringOrDefault returns the value associated with the key as a string.
	// If the key does not exist, return defaultValue.
	GetStringOrDefault(key string, defaultValue string) string
	// GetInt returns the value associated with the key as an int.
	// If the key does not exist, panic.
	GetInt(key string) int64
	// GetIntOrDefault returns the value associated with the key as an int.
	// If the key does not exist, return defaultValue.
	GetIntOrDefault(key string, defaultValue int64) int64
	// GetBool returns the value associated with the key as a bool.
	// If the key does not exist, panic.
	GetBool(key string) bool
	// GetBoolOrDefault returns the value associated with the key as a bool.
	// If the key does not exist, return defaultValue.
	GetBoolOrDefault(key string, defaultValue bool) bool
}

type configuration struct {
	providers []Provider
}

// NewConfiguration creates a new configuration.
func NewConfiguration(f ...ConfigurationOptionFunc) Configuration {
	var opt ConfigurationOptions
	for _, v := range f {
		v(&opt)
	}
	providers := opt.GetProviders()
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Priority() < providers[j].Priority()
	})
	return &configuration{
		providers: providers,
	}
}

func (c *configuration) GetString(key string) string {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == STRING {
			return entry.String()
		}
	}
	panic(fmt.Sprintf("key %s not found", key))
}

func (c *configuration) GetStringOrDefault(key string, defaultValue string) string {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == STRING {
			return entry.String()
		}
	}
	return defaultValue
}

func (c *configuration) GetInt(key string) int64 {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == INT {
			return entry.Int()
		}
	}
	panic(fmt.Sprintf("key %s not found", key))
}

func (c *configuration) GetIntOrDefault(key string, defaultValue int64) int64 {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == INT {
			return entry.Int()
		}
	}
	return defaultValue
}

func (c *configuration) GetBool(key string) bool {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == BOOL {
			return entry.Bool()
		}
	}
	panic(fmt.Sprintf("key %s not found", key))
}

func (c *configuration) GetBoolOrDefault(key string, defaultValue bool) bool {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() && entry.ValueType() == BOOL {
			return entry.Bool()
		}
	}
	return defaultValue
}
