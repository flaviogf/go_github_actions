package math

import "testing"

func TestSum(t *testing.T) {
	tableTest := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 10, y: 10, expected: 20},
		{x: 20, y: 20, expected: 40},
		{x: 30, y: 30, expected: 60},
	}

	for _, test := range tableTest {
		got := Sum(test.x, test.y)

		if test.expected != got {
			t.Errorf("expected: %d, got: %d", test.expected, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	tableTest := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 10, y: 10, expected: 100},
		{x: 20, y: 20, expected: 400},
		{x: 30, y: 30, expected: 900},
	}

	for _, test := range tableTest {
		got := Multiply(test.x, test.y)

		if test.expected != got {
			t.Errorf("expected: %d, got: %d", test.expected, got)
		}
	}
}
