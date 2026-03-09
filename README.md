# Config Go

Config Go is a small library for layered configuration loading with an error-first API.

## Installation

```bash
go get github.com/Oleexo/config-go
```

## Usage

```go
package main

import (
    "errors"
    "fmt"

 "github.com/Oleexo/config-go"
)

func main() {
    cfg, err := config.New(config.WithEnvPrefix("APP_"), config.WithDotenvFiles(".env", ".env.local"))
    if err != nil {
        if errors.Is(err, config.ErrDotenvFileNotFound) {
            fmt.Println("missing dotenv file")
            return
        }
        panic(err)
    }

    v, err := cfg.String("DB_HOST")
    if err != nil {
        if errors.Is(err, config.ErrKeyNotFound) {
            fmt.Println("DB_HOST not configured")
            return
        }
        panic(err)
    }

    fmt.Println(v)
}
```

## Provider precedence

Providers are resolved in precedence order (lower value wins):

1. `WithMemory(...)`
2. `WithEnvironmentVariables()` or `WithEnvPrefix(...)`
3. `WithDotenv()` or `WithDotenvFiles(...)`

The first provider that returns an existing key is used.

### Strict dotenv loading

Use `WithDotenvFilesStrict` when missing files should fail initialization.

```go
cfg, err := config.New(config.WithDotenvFilesStrict(".env", ".env.local"))
if err != nil {
    // handle ErrDotenvFileNotFound or parse errors
}

_ = cfg
```

## Error handling

Lookup methods return typed errors that can be matched with `errors.Is`.

```go
value, err := cfg.Int("PORT")
if err != nil {
    if errors.Is(err, config.ErrKeyNotFound) {
        value = 8080
    } else if errors.Is(err, config.ErrTypeMismatch) {
        panic("PORT must be an integer")
    } else {
        panic(err)
    }
}
```

## Migration guide (old API to new API)

| Old API | New API |
| --- | --- |
| `NewConfiguration(...)` | `New(...) (Configuration, error)` |
| `GetString(key)` | `String(key) (string, error)` |
| `GetInt(key)` | `Int(key) (int64, error)` |
| `GetBool(key)` | `Bool(key) (bool, error)` |
| `GetStringOrDefault(key, def)` | `StringDefault(key, def)` |
| `GetIntOrDefault(key, def)` | `IntDefault(key, def)` |
| `GetBoolOrDefault(key, def)` | `BoolDefault(key, def)` |
| `NewEntryString(v)` | `NewString(v)` |
| `NewEntryInt(v)` | `NewInt(v)` |
| `NewEntryFloat(v)` | `NewFloat(v)` |
| `NewEntryBool(v)` | `NewBool(v)` |

Must-style helpers are available for fail-fast startup paths:

1. `MustString(key)`
2. `MustInt(key)`
3. `MustFloat(key)`
4. `MustBool(key)`
5. `MustGet[T](cfg, key)`
