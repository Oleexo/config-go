# Config go

Config go is a simple go package that allows you to read configuration from a file or environment variables.

## Installation

```bash
go get github.com/Oleexo/config-go
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Oleexo/config-go"
	"github.com/Oleexo/config-go/dotenv"
	"github.com/Oleexo/config-go/envs"
	"github.com/Oleexo/config-go/mem"
)

func main() {
	config := config.NewConfiguration(
		mem.WithMemory(map[string]string{
			"key": "value",
		}),
		envs.WithEnvironmentVariables(),
		dotenv.WithDotenv(),
	)

	fmt.Println(config.Get("key"))
}
```