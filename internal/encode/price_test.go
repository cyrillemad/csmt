package encode

import (
	"fmt"
	"testing"
)

func TestParsePrice(t *testing.T) {
	var data = []string{"126.34 RUB", "$1.44", "164,43"}
	var responses = []float64{126.34, 1.44, 164.43}

	for index, example := range data {
		result, err := ParsePrice(example)
		if err != nil {
			t.Error(err)
		}
		if result != responses[index] {
			t.Error(
				fmt.Errorf(
					"expected %v, got %v",
					responses[index], result))
		}
	}
}
