package config

import (
	"fmt"
	"sort"
)

type Configuration interface {
	String(key string) (string, error)
	StringDefault(key string, defaultValue string) string
	MustString(key string) string

	Int(key string) (int64, error)
	IntDefault(key string, defaultValue int64) int64
	MustInt(key string) int64

	Float(key string) (float64, error)
	FloatDefault(key string, defaultValue float64) float64
	MustFloat(key string) float64

	Bool(key string) (bool, error)
	BoolDefault(key string, defaultValue bool) bool
	MustBool(key string) bool
}

type configuration struct {
	providers []Provider
}

type Scalar interface {
	string | int64 | float64 | bool
}

// New creates a new configuration instance.
func New(options ...Option) (Configuration, error) {
	var opt optionSet
	for _, applyOption := range options {
		if err := applyOption(&opt); err != nil {
			return nil, err
		}
	}

	providers := opt.providersList()
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Precedence() < providers[j].Precedence()
	})

	return &configuration{
		providers: providers,
	}, nil
}

func (c *configuration) lookupEntry(key string) (Entry, error) {
	for _, provider := range c.providers {
		entry := provider.GetEntry(key)
		if entry.Exists() {
			return entry, nil
		}
	}

	return Empty(), fmt.Errorf("config key %q: %w", key, ErrKeyNotFound)
}

func (c *configuration) String(key string) (string, error) {
	entry, err := c.lookupEntry(key)
	if err != nil {
		return "", err
	}

	v, err := entry.String()
	if err != nil {
		return "", fmt.Errorf("config key %q: %w", key, err)
	}

	return v, nil
}

func (c *configuration) StringDefault(key, defaultValue string) string {
	v, err := c.String(key)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c *configuration) MustString(key string) string {
	v, err := c.String(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *configuration) Int(key string) (int64, error) {
	entry, err := c.lookupEntry(key)
	if err != nil {
		return 0, err
	}

	v, err := entry.Int()
	if err != nil {
		return 0, fmt.Errorf("config key %q: %w", key, err)
	}

	return v, nil
}

func (c *configuration) IntDefault(key string, defaultValue int64) int64 {
	v, err := c.Int(key)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c *configuration) MustInt(key string) int64 {
	v, err := c.Int(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *configuration) Float(key string) (float64, error) {
	entry, err := c.lookupEntry(key)
	if err != nil {
		return 0, err
	}

	v, err := entry.Float()
	if err != nil {
		return 0, fmt.Errorf("config key %q: %w", key, err)
	}

	return v, nil
}

func (c *configuration) FloatDefault(key string, defaultValue float64) float64 {
	v, err := c.Float(key)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c *configuration) MustFloat(key string) float64 {
	v, err := c.Float(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *configuration) Bool(key string) (bool, error) {
	entry, err := c.lookupEntry(key)
	if err != nil {
		return false, err
	}

	v, err := entry.Bool()
	if err != nil {
		return false, fmt.Errorf("config key %q: %w", key, err)
	}

	return v, nil
}

func (c *configuration) BoolDefault(key string, defaultValue bool) bool {
	v, err := c.Bool(key)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c *configuration) MustBool(key string) bool {
	v, err := c.Bool(key)
	if err != nil {
		panic(err)
	}
	return v
}

func Get[T Scalar](configuration Configuration, key string) (T, error) {
	var zero T

	switch any(zero).(type) {
	case string:
		v, err := configuration.String(key)
		if err != nil {
			return zero, err
		}
		typed, ok := any(v).(T)
		if !ok {
			return zero, ErrInvalidEntry
		}
		return typed, nil
	case int64:
		v, err := configuration.Int(key)
		if err != nil {
			return zero, err
		}
		typed, ok := any(v).(T)
		if !ok {
			return zero, ErrInvalidEntry
		}
		return typed, nil
	case float64:
		v, err := configuration.Float(key)
		if err != nil {
			return zero, err
		}
		typed, ok := any(v).(T)
		if !ok {
			return zero, ErrInvalidEntry
		}
		return typed, nil
	case bool:
		v, err := configuration.Bool(key)
		if err != nil {
			return zero, err
		}
		typed, ok := any(v).(T)
		if !ok {
			return zero, ErrInvalidEntry
		}
		return typed, nil
	default:
		return zero, ErrInvalidEntry
	}
}

func MustGet[T Scalar](configuration Configuration, key string) T {
	v, err := Get[T](configuration, key)
	if err != nil {
		panic(err)
	}
	return v
}
