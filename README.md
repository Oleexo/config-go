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
)

func main() {
 config := config.NewConfiguration(
  config.WithMemory(map[string]config.Entry{
   "key": config.NewEntryString("value"),
  }),
  config.WithEnvironmentVariables(),
  config.WithDotenvFiles(".env", ".env.local"),
 )

 fmt.Println(config.GetString("key"))
}
```
