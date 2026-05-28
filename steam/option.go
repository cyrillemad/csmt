package steam

import (
	"time"

	net "github.com/cyrillemad/csmt/internal/httpclient"
	"github.com/cyrillemad/csmt/types"
)

type Option func(config *Config)

func WithCountry(country string) Option {
	if country != "" &&
		len(country) == 2 {
		return func(config *Config) {
			config.Country = country
		}
	}
	return func(config *Config) {}
}

func WithCurrency(currency types.Currency) Option {
	return func(config *Config) {
		config.Currency = currency
	}
}

func WithTimeout(timeout time.Duration) Option {
	if timeout > 0 {
		return func(config *Config) {
			config.Timeout = timeout
		}
	}
	return func(config *Config) {}
}

func WithHTTPOption(option net.Option) Option {
	return func(config *Config) {
		config.HTTPOptions = append(
			config.HTTPOptions,
			option)
	}
}
