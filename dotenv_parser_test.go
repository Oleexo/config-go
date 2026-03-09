package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLine(t *testing.T) {
	t.Run("should parse line with STRING value", func(t *testing.T) {
		key, entry := parseLineToEntry("key=value")
		assert.Equal(t, KindString, entry.Kind())
		assert.Equal(t, "key", key)

		value, stringErr := entry.String()
		require.NoError(t, stringErr)
		assert.Equal(t, "value", value)
	})

	t.Run("should parse line with INT value", func(t *testing.T) {
		key, entry := parseLineToEntry("key=4321")
		assert.Equal(t, KindInt, entry.Kind())
		assert.Equal(t, "key", key)

		value, intErr := entry.Int()
		require.NoError(t, intErr)
		assert.Equal(t, int64(4321), value)
	})

	t.Run("should parse line with FLOAT value", func(t *testing.T) {
		key, entry := parseLineToEntry("key=1234.4321")
		assert.Equal(t, KindFloat, entry.Kind())
		assert.Equal(t, "key", key)

		value, floatErr := entry.Float()
		require.NoError(t, floatErr)
		assert.Equal(t, 1234.4321, value)
	})

	t.Run("should parse line with BOOL value", func(t *testing.T) {
		key, entry := parseLineToEntry("key=true")
		assert.Equal(t, KindBool, entry.Kind())
		assert.Equal(t, "key", key)

		value, boolErr := entry.Bool()
		require.NoError(t, boolErr)
		assert.Equal(t, true, value)
	})

	t.Run("should parse line with ${VALUE} as environment variable", func(t *testing.T) {
		t.Setenv("TEST_VALUE", "12345")
		t.Setenv("OTHER_VALUE", "67890")

		key, entry := parseLineToEntry("key=lorem ${TEST_VALUE} :: ${OTHER_VALUE} ipsum")
		assert.Equal(t, KindString, entry.Kind())
		assert.Equal(t, "key", key)

		value, stringErr := entry.String()
		require.NoError(t, stringErr)
		assert.Equal(t, "lorem 12345 :: 67890 ipsum", value)
	})
}

func TestWithDotenvFiles(t *testing.T) {
	dir := t.TempDir()

	baseFile := filepath.Join(dir, ".env")
	localFile := filepath.Join(dir, ".env.local")

	assert.NoError(t, os.WriteFile(baseFile, []byte("SHARED=from-base\nBASE_ONLY=base\nINT_KEY=42\n"), 0o600))
	assert.NoError(t, os.WriteFile(localFile, []byte("SHARED=from-local\nLOCAL_ONLY=local\n"), 0o600))

	c, err := New(WithDotenvFiles(baseFile, localFile))
	require.NoError(t, err)

	assert.Equal(t, "from-local", c.MustString("SHARED"))
	assert.Equal(t, "base", c.MustString("BASE_ONLY"))
	assert.Equal(t, "local", c.MustString("LOCAL_ONLY"))
	assert.Equal(t, int64(42), c.MustInt("INT_KEY"))
}

func TestWithDotenvFilesSkipsMissing(t *testing.T) {
	dir := t.TempDir()
	existingFile := filepath.Join(dir, ".env")
	missingFile := filepath.Join(dir, ".env.local")

	assert.NoError(t, os.WriteFile(existingFile, []byte("VALUE=present\n"), 0o600))

	c, err := New(WithDotenvFiles(missingFile, existingFile))
	require.NoError(t, err)

	assert.Equal(t, "present", c.MustString("VALUE"))
}
