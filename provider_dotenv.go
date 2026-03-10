package config

import (
	"fmt"
	"os"
	"strings"
)

type dotenvProvider struct {
	entries map[string]Entry
}

func WithDotenv() Option {
	return WithDotenvFiles(resolveDefaultDotenvFiles()...)
}

// WithDotenvFiles loads dotenv files in order.
// Later files override values from earlier files.
func WithDotenvFiles(files ...string) Option {
	return withDotenvFiles(false, files...)
}

// WithDotenvFilesStrict loads dotenv files in order and fails if any file is missing.
func WithDotenvFilesStrict(files ...string) Option {
	return withDotenvFiles(true, files...)
}

func withDotenvFiles(strict bool, files ...string) Option {
	return func(c *optionSet) error {
		p, err := newDotenvProvider(strict, files...)
		if err != nil {
			return err
		}
		if p == nil {
			return nil
		}
		c.addProvider(p)
		return nil
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

func newDotenvProvider(strict bool, files ...string) (*dotenvProvider, error) {
	if len(files) == 0 {
		files = []string{defaultDotenvFileName}
	}

	entries := make(map[string]Entry)
	loaded := false
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			if strict {
				return nil, fmt.Errorf("dotenv file %q: %w", file, ErrDotenvFileNotFound)
			}
			continue
		} else if err != nil {
			return nil, err
		}

		fileEntries, err := readDotenvFile(file)
		if err != nil {
			return nil, err
		}
		for key, value := range fileEntries {
			entries[key] = value
		}
		loaded = true
	}

	if !loaded {
		if strict {
			return nil, ErrDotenvFileNotFound
		}
		return nil, nil
	}

	return &dotenvProvider{entries: entries}, nil
}

func (p dotenvProvider) Precedence() int {
	return 1000
}

func (p dotenvProvider) GetEntry(key string) Entry {
	entry, ok := p.entries[key]
	if !ok {
		return Empty()
	}
	return entry
}
