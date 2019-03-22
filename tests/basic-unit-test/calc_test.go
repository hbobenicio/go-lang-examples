package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{0, 0, 9},
		{2, 3, 9},
		{-1, 1, 9},
	}

	for _, tc := range testCases {
		actual := Add(tc.x, tc.y)
		if actual != tc.expected {
			t.Errorf("Add(%d, %d) returned %d instead of %d", tc.x, tc.y, actual, tc.expected)
		}
	}
}
