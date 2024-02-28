package dotenv

import (
	"github.com/Oleexo/config-go"
	"os"
)

type Provider struct {
	entries map[string]config.Entry
}

func NewProvider() *Provider {
	file, exists := os.LookupEnv(environmentKey)
	if !exists {
		file = defaultFileName
	}
	// Check if exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil
	}

	entries, err := readFile(file)
	if err != nil {
		panic(err)
	}
	return &Provider{
		entries: entries,
	}
}

func WithDotenv() config.ConfigurationOptionFunc {
	return func(c *config.ConfigurationOptions) {
		p := NewProvider()
		if p == nil {
			return
		}
		c.AddProvider(p)
	}
}

func (p Provider) Priority() int {
	return 1000
}

func (p Provider) GetEntry(key string) config.Entry {
	entry, ok := p.entries[key]
	if !ok {
		return config.NewEntryEmpty()
	}
	return entry
}
