package types

import "time"

type RetryConfig struct {
	Attempts int
	Delay    time.Duration
}
