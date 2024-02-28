package configfx

import (
	"github.com/Oleexo/config-go"
	"go.uber.org/fx"
)

func BuildConfigModule(f ...config.ConfigurationOptionFunc) fx.Option {
	return fx.Module("configfx",
		fx.Provide(
			func() config.Configuration {
				return config.NewConfiguration(f...)
			},
		),
	)
}
