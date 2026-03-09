package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	t.Run("should parse line with STRING value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=value")
		assert.NoError(t, err)
		assert.Equal(t, STRING, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, "value", entry.String())
	})

	t.Run("should parse line with INT value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=4321")
		assert.NoError(t, err)
		assert.Equal(t, INT, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, int64(4321), entry.Int())
	})

	t.Run("should parse line with FLOAT value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=1234.4321")
		assert.NoError(t, err)
		assert.Equal(t, FLOAT, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, 1234.4321, entry.Float())
	})

	t.Run("should parse line with BOOL value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=true")
		assert.NoError(t, err)
		assert.Equal(t, BOOL, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, true, entry.Bool())
	})

	t.Run("should parse line with ${VALUE} as environment variable", func(t *testing.T) {
		t.Setenv("TEST_VALUE", "12345")
		t.Setenv("OTHER_VALUE", "67890")

		key, entry, err := parseLineToEntry("key=lorem ${TEST_VALUE} :: ${OTHER_VALUE} ipsum")
		assert.NoError(t, err)
		assert.Equal(t, STRING, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, "lorem 12345 :: 67890 ipsum", entry.String())
	})
}

func TestWithDotenvFiles(t *testing.T) {
	dir := t.TempDir()

	baseFile := filepath.Join(dir, ".env")
	localFile := filepath.Join(dir, ".env.local")

	assert.NoError(t, os.WriteFile(baseFile, []byte("SHARED=from-base\nBASE_ONLY=base\nINT_KEY=42\n"), 0o600))
	assert.NoError(t, os.WriteFile(localFile, []byte("SHARED=from-local\nLOCAL_ONLY=local\n"), 0o600))

	c := NewConfiguration(WithDotenvFiles(baseFile, localFile))

	assert.Equal(t, "from-local", c.GetString("SHARED"))
	assert.Equal(t, "base", c.GetString("BASE_ONLY"))
	assert.Equal(t, "local", c.GetString("LOCAL_ONLY"))
	assert.Equal(t, int64(42), c.GetInt("INT_KEY"))
}

func TestWithDotenvFilesSkipsMissing(t *testing.T) {
	dir := t.TempDir()
	existingFile := filepath.Join(dir, ".env")
	missingFile := filepath.Join(dir, ".env.local")

	assert.NoError(t, os.WriteFile(existingFile, []byte("VALUE=present\n"), 0o600))

	c := NewConfiguration(WithDotenvFiles(missingFile, existingFile))

	assert.Equal(t, "present", c.GetString("VALUE"))
}
