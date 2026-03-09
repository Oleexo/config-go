// Package config provides layered configuration loading for Go applications.
//
// Providers are resolved by precedence (lower value wins): memory, environment,
// then dotenv by default. Accessors are error-first and expose Must* helpers for
// startup-time fail-fast usage.
//
// Construction uses options:
//
//	cfg, err := config.New(
//		config.WithMemory(map[string]config.Entry{"KEY": config.NewString("v")}),
//		config.WithEnvPrefix("APP_"),
//		config.WithDotenvFiles(".env", ".env.local"),
//	)
//
// Dotenv sources can be optional (WithDotenvFiles) or strict
// (WithDotenvFilesStrict).
package config
