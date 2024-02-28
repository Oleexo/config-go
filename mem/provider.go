package mem

import "github.com/Oleexo/config-go"

type Provider struct {
	data map[string]config.Entry
}

func (p Provider) Priority() int {
	return 100
}

func (p Provider) GetEntry(key string) config.Entry {
	if value, exists := p.data[key]; exists {
		return value
	}
	return config.NewEntryEmpty()
}

func NewProvider(data map[string]config.Entry) *Provider {
	return &Provider{
		data: data,
	}
}

func WithMemory(data map[string]config.Entry) config.ConfigurationOptionFunc {
	return func(c *config.ConfigurationOptions) {
		c.AddProvider(NewProvider(data))
	}
}
