package config

type Option func(*optionSet) error

type optionSet struct {
	providers Providers
}

func (c *optionSet) addProvider(provider Provider) {
	c.providers = append(c.providers, provider)
}

func (c *optionSet) providersList() Providers {
	return c.providers
}
