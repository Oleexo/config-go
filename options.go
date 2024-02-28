package config

type ConfigurationOptionFunc func(*ConfigurationOptions)

type ConfigurationOptions struct {
	providers Providers
}

func (c *ConfigurationOptions) AddProvider(provider Provider) {
	c.providers = append(c.providers, provider)
}

func (c *ConfigurationOptions) GetProviders() Providers {
	return c.providers
}
