package httpclient

import (
	"errors"
	"io"
	"net"
)

func shouldRetryStatus(code int) bool {
	switch code {
	case 429, 502, 503, 504:
		return true
	default:
		return false
	}
}

func shouldRetryError(err error) bool {
	if errors.Is(err, io.EOF) {
		return true
	}
	if netErr, ok := errors.AsType[net.Error](err); ok {
		return netErr.Timeout()
	}

	return false
}
