package config

import (
	"os"
	"strings"
)

type dotenvProvider struct {
	entries map[string]Entry
}

func WithDotenv() ConfigurationOptionFunc {
	return WithDotenvFiles(resolveDefaultDotenvFiles()...)
}

// WithDotenvFiles loads dotenv files in order.
// Later files override values from earlier files.
func WithDotenvFiles(files ...string) ConfigurationOptionFunc {
	return func(c *ConfigurationOptions) {
		p := newDotenvProvider(files...)
		if p == nil {
			return
		}
		c.AddProvider(p)
	}
}

func resolveDefaultDotenvFiles() []string {
	file, exists := os.LookupEnv(dotenvEnvironmentKey)
	if !exists {
		return []string{defaultDotenvFileName}
	}
	parts := strings.Split(file, ",")
	files := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			files = append(files, trimmed)
		}
	}
	if len(files) == 0 {
		return []string{defaultDotenvFileName}
	}
	return files
}

func newDotenvProvider(files ...string) *dotenvProvider {
	if len(files) == 0 {
		files = []string{defaultDotenvFileName}
	}

	entries := make(map[string]Entry)
	loaded := false
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}

		fileEntries, err := readDotenvFile(file)
		if err != nil {
			panic(err)
		}
		for key, value := range fileEntries {
			entries[key] = value
		}
		loaded = true
	}

	if !loaded {
		return nil
	}

	return &dotenvProvider{entries: entries}
}

func (p dotenvProvider) Priority() int {
	return 1000
}

func (p dotenvProvider) GetEntry(key string) Entry {
	entry, ok := p.entries[key]
	if !ok {
		return NewEntryEmpty()
	}
	return entry
}
