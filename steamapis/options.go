package steamapis

import net "github.com/cyrillemad/csmt/internal/httpclient"

type Option func(config *Config)

func WithHTTPOption(option net.Option) Option {
	return func(config *Config) {
		config.HTTPOptions = append(
			config.HTTPOptions,
			option)
	}
}
