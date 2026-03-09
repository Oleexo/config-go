package config

type memProvider struct {
	data map[string]Entry
}

func (p memProvider) Priority() int {
	return 100
}

func (p memProvider) GetEntry(key string) Entry {
	if value, exists := p.data[key]; exists {
		return value
	}
	return NewEntryEmpty()
}

func newMemProvider(data map[string]Entry) *memProvider {
	return &memProvider{data: data}
}

func WithMemory(data map[string]Entry) ConfigurationOptionFunc {
	return func(c *ConfigurationOptions) {
		c.AddProvider(newMemProvider(data))
	}
}
