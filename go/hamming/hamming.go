package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("asdf")
	}
	h_dist := 0
	for i, _ := range a {
		if a[i] != b[i] {
			h_dist++
		}
	}
	return h_dist, nil
}
