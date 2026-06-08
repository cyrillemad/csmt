package convert

import "strconv"

func Int(value string) (int, error) {
	return strconv.Atoi(value)
}
