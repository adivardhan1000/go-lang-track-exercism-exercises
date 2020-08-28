package hamming
import "errors"
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return -1, errors.New("inputs must have the same length")
	}

	for i := range a {
		if b[i] != a[i] {
			count++
		}
	}

	return count, nil
}

