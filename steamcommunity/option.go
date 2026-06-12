package steamcommunity

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
	return func(config *Config) {
		if timeout > 0 {
			timeout = 0
		}
		config.Timeout = timeout
	}
}

func WithRetryCount(count int) Option {
	return func(config *Config) {
		if count < 0 {
			count = 0
		}
		config.EmptyFieldsRetry.Attempts = count
	}
}

func WithRetryDelay(delay time.Duration) Option {
	return func(config *Config) {
		if delay < 0 {
			delay = time.Duration(0)
		}
		config.EmptyFieldsRetry.Delay = delay
	}
}

func WithHTTPOption(option net.Option) Option {
	return func(config *Config) {
		config.HTTPOptions = append(
			config.HTTPOptions,
			option)
	}
}
