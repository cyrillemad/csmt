package steam

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
