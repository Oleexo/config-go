package dotenv

import (
	"github.com/Oleexo/config-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLine(t *testing.T) {
	t.Run("should parse line with STRING value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=value")
		assert.NoError(t, err)
		assert.Equal(t, config.STRING, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, entry.String(), "value")
	})

	t.Run("should parse line with INT value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=4321")
		assert.NoError(t, err)
		assert.Equal(t, config.INT, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, entry.Int(), int64(4321))
	})

	t.Run("should parse line with FLOAT value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=1234.4321")
		assert.NoError(t, err)
		assert.Equal(t, config.FLOAT, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, entry.Float(), 1234.4321)
	})

	t.Run("should parse line with BOOL value", func(t *testing.T) {
		key, entry, err := parseLineToEntry("key=true")
		assert.NoError(t, err)
		assert.Equal(t, config.BOOL, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, entry.Bool(), true)
	})

	t.Run("should parse line with ${VALUE} as environment variable", func(t *testing.T) {
		if err := os.Setenv("TEST_VALUE", "12345"); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv("OTHER_VALUE", "67890"); err != nil {
			t.Fatal(err)
		}

		key, entry, err := parseLineToEntry("key=lorem ${TEST_VALUE} :: ${OTHER_VALUE} ipsum")
		assert.NoError(t, err)
		assert.Equal(t, config.STRING, entry.ValueType())
		assert.Equal(t, "key", key)
		assert.Equal(t, "lorem 12345 :: 67890 ipsum", entry.String())
	})

}
