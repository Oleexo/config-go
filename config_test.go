package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigurationGetters(t *testing.T) {
	cfg, err := New(WithMemory(map[string]Entry{
		"STRING_KEY": NewString("value"),
		"INT_KEY":    NewInt(12),
		"FLOAT_KEY":  NewFloat(1.5),
		"BOOL_KEY":   NewBool(true),
	}))
	require.NoError(t, err)

	s, err := cfg.String("STRING_KEY")
	require.NoError(t, err)
	assert.Equal(t, "value", s)

	i, err := cfg.Int("INT_KEY")
	require.NoError(t, err)
	assert.Equal(t, int64(12), i)

	f, err := cfg.Float("FLOAT_KEY")
	require.NoError(t, err)
	assert.Equal(t, 1.5, f)

	b, err := cfg.Bool("BOOL_KEY")
	require.NoError(t, err)
	assert.True(t, b)
}

func TestConfigurationErrors(t *testing.T) {
	cfg, err := New(WithMemory(map[string]Entry{"BOOL_KEY": NewBool(true)}))
	require.NoError(t, err)

	_, err = cfg.String("BOOL_KEY")
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrTypeMismatch))

	_, err = cfg.String("MISSING_KEY")
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrKeyNotFound))
}

func TestGenericGet(t *testing.T) {
	cfg, err := New(WithMemory(map[string]Entry{"INT_KEY": NewInt(42)}))
	require.NoError(t, err)

	v, err := Get[int64](cfg, "INT_KEY")
	require.NoError(t, err)
	assert.Equal(t, int64(42), v)
}

func TestWithEnvPrefix(t *testing.T) {
	t.Setenv("APP_PORT", "8080")

	cfg, err := New(WithEnvPrefix("APP_"))
	require.NoError(t, err)

	port, err := cfg.Int("PORT")
	require.NoError(t, err)
	assert.Equal(t, int64(8080), port)
}

func TestWithDotenvFilesStrict(t *testing.T) {
	dir := t.TempDir()
	present := filepath.Join(dir, ".env")
	missing := filepath.Join(dir, ".env.local")

	require.NoError(t, os.WriteFile(present, []byte("A=1\n"), 0o600))

	_, err := New(WithDotenvFilesStrict(present, missing))
	require.Error(t, err)
	assert.True(t, errors.Is(err, ErrDotenvFileNotFound))
}
